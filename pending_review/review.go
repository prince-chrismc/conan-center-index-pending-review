package pending_review

import "time"

// Approver captures the user login or display name along with the tier of the reviewer
type Approver struct {
	Name string
	Tier ReviewerType
}

// Review contains the essentials of a submission
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

	Approvals []Approver // List of users who have approved the pull request on the head commit
	Blockers  []string   // List of Conan team members who have requested changes on the head commit

	LastReview *Review `json:",omitempty"` // Snapshot of the last review

	IsBump			bool // PR is a bump of version or dependencies
}

// IsApproved when the conditions for merging are meet as per https://github.com/conan-io/conan-center-index/blob/master/docs/review_process.md
func (r *Reviews) IsApproved() bool {
	var nRequiredApprovals = 2
	if r.IsBump {
		nRequiredApprovals = 1
	}
	return r.TeamApproval && r.ValidApprovals >= nRequiredApprovals && len(r.Blockers) == 0
}

// ProcessReviewComments interprets the all the reviews to extract a summary based on the requirements of CCI
func ProcessReviewComments(reviewers *ConanCenterReviewers, reviews []*PullRequestReview, sha string) Reviews {
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
		isTeamMember := reviewers.IsTeamMember(reviewerName)
		isMember := isTeamMember || reviewers.IsCommunityMember(reviewerName)

		reviewer := Approver{Name: reviewerName, Tier: Unofficial}
		if isMember {
			reviewer.Tier = Community
		}
		if isTeamMember {
			reviewer.Tier = Team
		}

		switch review.GetState() { // Either as indicated by the reviewer or by updates from the GitHub API
		case "CHANGES_REQUESTED":
			if isTeamMember {
				summary.Blockers, _ = appendUnique(summary.Blockers, reviewerName)
			}

			removed := false
			summary.Approvals, removed = removeUnique(summary.Approvals, reviewer)
			if removed && isMember {
				// If a reviewer retracted their review, the count needs to be adjusted
				summary.ValidApprovals = summary.ValidApprovals - 1
			}

		case "APPROVED":
			summary.Blockers, _ = removeUnique(summary.Blockers, reviewerName)

			if !onBranchHead {
				break // If the approval is on anything other than the HEAD it's not counted by @conan-center-bot
			}

			new := false
			summary.Approvals, new = appendUnique(summary.Approvals, reviewer)
			if !new {
				// Duplicate review (usually an accident) or might be after an empty or merge commit
				// https://github.com/conan-io/conan-center-index/pull/16475#pullrequestreview-1376616442
				break
			}

			if isTeamMember {
				summary.TeamApproval = true
			}

			if isMember {
				summary.ValidApprovals = summary.ValidApprovals + 1
			}

		case "DISMISSED":
			// This is applied to both "Approvals" (15202) in addition to "Requested Changes" (16034)
			// The previous state is not accessible (2023-04-07 at least) so *NO* modifications are
			// required to the existing lists

			// https://api.github.com/repos/conan-io/conan-center-index/pulls/15202/reviews
			// https://github.com/conan-io/conan-center-index/pull/15202#pullrequestreview-1244441259

			// https://api.github.com/repos/conan-io/conan-center-index/pulls/16034/reviews
			// https://github.com/conan-io/conan-center-index/pull/16034#pullrequestreview-1330280912

		case "COMMENTED":
			// TODO: Figure out if there is something useful to extract from here
		default:
		}
	}

	return summary
}

func appendUnique[K comparable](slice []K, elem K) ([]K, bool) {
	for _, e := range slice {
		if e == elem {
			return slice, false
		}
	}

	return append(slice, elem), true
}

func removeUnique[K comparable](slice []K, elem K) ([]K, bool) {
	for i, e := range slice {
		if e == elem {
			return append(slice[:i], slice[i+1:]...), true
		}
	}

	return slice, false
}

// FilterAuthor removes any reviews by a given user's login
func FilterAuthor(reviews []*PullRequestReview, author string) []*PullRequestReview {
	filtered := make([]*PullRequestReview, 0, len(reviews))
	for _, r := range reviews {
		if r.GetUser().GetLogin() != author {
			filtered = append(filtered, r)
		}
	}
	return filtered
}
