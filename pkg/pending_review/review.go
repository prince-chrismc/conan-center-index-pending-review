package pending_review

import "time"

// Review contains teh essentials of a submission
type Review struct {
	ReviewerName string
	SubmittedAt  time.Time
	HTMLURL      string
}

// Reviews summarizes all the reviews of a given pull request
type Reviews struct {
	Count          int  // Total number of comments, requested changes, and approvals
	ValidApprovals int  // Counted by head commit approvals from official community reviewers and the Conan team
	TeamApproval   bool // At least one approval from the Conan team

	Approvals []string // List of users who have approved the pull request on the head commit
	Blockers  []string // List of Conan team members who have requested changes on the head commit

	LastReview *Review `json:",omitempty"` // Snapshot of the last review
}

// IsApproved when the conditions for merging are meet as per https://github.com/conan-io/conan-center-index/blob/master/docs/review_process.md
func (r *Reviews) IsApproved() bool {
	return r.TeamApproval && r.ValidApprovals >= 3 && len(r.Blockers) == 0
}

// ProcessReviewComments interprets the all the reviews to extract a summary based on the requirements of CCI
func ProcessReviewComments(reviews []*PullRequestReview, sha string) Reviews {
	summary := Reviews{
		Count:        len(reviews),
		TeamApproval: false,
	}

	if len := len(reviews); len > 0 {
		lastReview := reviews[len-1]
		summary.LastReview = &Review{
			ReviewerName: lastReview.GetUser().GetLogin(),
			SubmittedAt:  lastReview.GetSubmittedAt(),
			HTMLURL:      lastReview.GetHTMLURL(),
		}
	}

	for _, review := range reviews {
		onBranchHead := sha == review.GetCommitID()

		reviewerName := review.GetUser().GetLogin()
		isTeamMember := isTeamMember(reviewerName)
		isMember := isTeamMember || isCommunityMember(reviewerName)

		switch review.GetState() { // Either as indicated by the reviewer or by updates from the GitHub API
		case "CHANGES_REQUESTED":
			if isTeamMember {
				summary.Blockers, _ = appendUnique(summary.Blockers, reviewerName)
			}

			removed := false
			summary.Approvals, removed = removeUnique(summary.Approvals, reviewerName)
			if removed && isMember {
				// If a reviewer retracted their reivew, the count needs to be adjusted
				summary.ValidApprovals = summary.ValidApprovals - 1
			}

		case "APPROVED":
			summary.Blockers, _ = removeUnique(summary.Blockers, reviewerName)

			if !onBranchHead {
				break // If the approval is on anything other than the HEAD it's not counted by @conan-center-bot
			}

			new := false
			summary.Approvals, new = appendUnique(summary.Approvals, reviewerName)
			if !new { // Duplicate review (usually an accident)
				break
			}

			if isTeamMember {
				summary.TeamApproval = true
			}

			if isMember {
				summary.ValidApprovals = summary.ValidApprovals + 1
			}

		case "DISMISSED":
			summary.Blockers, _ = removeUnique(summary.Blockers, reviewerName)

		case "COMMENTED":
			// Out-dated Approvals are transformed into comments https://github.com/conan-io/conan-center-index/pull/3855#issuecomment-770120073
			// TODO: Figure out how GitHub knows what they were!
		default:
		}
	}

	return summary
}

func isTeamMember(reviewerName string) bool {
	switch reviewerName {
	// As defined by https://github.com/conan-io/conan-center-index/blob/master/docs/review_process.md#official-reviewers
	case "memsharded", "lasote", "danimtb", "jgsogo", "czoido", "SSE4", "uilianries":
		return true
	default:
		return false
	}
}

func isCommunityMember(reviewerName string) bool {
	switch reviewerName {
	// As defined by https://github.com/conan-io/conan-center-index/issues/2857
	// and https://github.com/conan-io/conan-center-index/blob/master/docs/review_process.md#community-reviewers
	case "madebr", "SpaceIm", "ericLemanissier", "prince-chrismc", "Croydon", "intelligide", "theirix", "gocarlos", "mathbunnyru", "ericriff", "toge":
		return true
	default:
		return false
	}
}

func appendUnique(slice []string, elem string) ([]string, bool) {
	for _, e := range slice {
		if e == elem {
			return slice, false
		}
	}

	return append(slice, elem), true
}

func removeUnique(slice []string, elem string) ([]string, bool) {
	for i, e := range slice {
		if e == elem {
			return append(slice[:i], slice[i+1:]...), true
		}
	}

	return slice, false
}
