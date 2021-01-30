package pending_review

type reviewState string

// State of the review, either as indicated by the reviewer or by updates from the GitHub API
const (
	CHANGE    reviewState = "CHANGES_REQUESTED"
	APPRVD    reviewState = "APPROVED"
	DISMISSED reviewState = "DISMISSED"
)

func (rs reviewState) String() string {
	return string(rs)
}

type authorAssociation string

// Reviewer's author associations to the repository
const (
	COLLABORATOR authorAssociation = "COLLABORATOR"
	CONTRIBUTOR  authorAssociation = "CONTRIBUTOR"
	MEMBER       authorAssociation = "MEMBER"
)

func (as authorAssociation) String() string {
	return string(as)
}

// ReviewsSummary digested representation of all the reviews of a given pull request
type ReviewsSummary struct {
	Count            int  // Total number of comments, requested changes, and approvals
	ValidApprovals   int  // Counted by head commit approvals from official community reviewers and c3i team
	PriorityApproval bool // At least one approval from a c31 team member

	HeadCommitApprovals []string // List of users who have approved the pull request
	HeadCommitBlockers  []string // List of c3i team members who have requested changes on the head commit
}

// ProcessReviewComments interprets the all the reviews to extract a summary based on the requirements of CCI
func ProcessReviewComments(reviews []*PullRequestReview, sha string) ReviewsSummary {
	summary := ReviewsSummary{
		Count:            len(reviews),
		PriorityApproval: false,
	}

	for _, review := range reviews {
		onBranchHead := sha == review.GetCommitID()

		reviewerName := review.GetUser().GetLogin()
		reviewerAssociation := review.GetAuthorAssociation()

		isC3iTeam := reviewerAssociation == MEMBER.String() || reviewerAssociation == COLLABORATOR.String()

		switch review.GetState() {
		case CHANGE.String():
			if isC3iTeam {
				summary.HeadCommitBlockers = appendUnique(summary.HeadCommitBlockers, reviewerName)
			}

			summary.HeadCommitApprovals = removeUnique(summary.HeadCommitApprovals, reviewerName)

		case APPRVD.String():
			if onBranchHead {
				summary.HeadCommitApprovals = appendUnique(summary.HeadCommitApprovals, reviewerName)
				if isC3iTeam {
					summary.PriorityApproval = true
					summary.ValidApprovals = summary.ValidApprovals + 1
				}
			}

			// We always count the community reviewer approvals (for any commit) because this matches closer with the GitHub UI
			switch reviewerName {
			case "madebr", "SpaceIm", "ericLemanissier", "prince-chrismc", "Croydon", "intelligide", "theirix", "gocarlos":
				summary.ValidApprovals = summary.ValidApprovals + 1
			default:
			}

			summary.HeadCommitBlockers = removeUnique(summary.HeadCommitBlockers, reviewerName)

		case DISMISSED.String():
			// Out-dated Approvals are transformed into comments https://github.com/conan-io/conan-center-index/pull/3855#issuecomment-770120073
			summary.HeadCommitBlockers = removeUnique(summary.HeadCommitBlockers, reviewerName)
		default:
		}
	}

	return summary
}

func appendUnique(slice []string, elem string) []string {
	for _, e := range slice {
		if e == elem {
			return slice
		}
	}

	return append(slice, elem)
}

func removeUnique(slice []string, elem string) []string {
	for i, e := range slice {
		if e == elem {
			return append(slice[:i], slice[i+1:]...)
		}
	}

	return slice
}
