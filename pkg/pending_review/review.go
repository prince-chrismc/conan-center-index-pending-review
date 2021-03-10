package pending_review

type authorAssociation string

// Reviews summary of all the reviews of a given pull request
type Reviews struct {
	Count            int  // Total number of comments, requested changes, and approvals
	ValidApprovals   int  // Counted by head commit approvals from official community reviewers and the Conan team
	PriorityApproval bool // At least one approval from the Conan team

	HeadCommitApprovals []string // List of users who have approved the pull request
	HeadCommitBlockers  []string // List of Conan team members who have requested changes on the head commit
}

// ProcessReviewComments interprets the all the reviews to extract a summary based on the requirements of CCI
func ProcessReviewComments(reviews []*PullRequestReview, sha string) Reviews {
	summary := Reviews{
		Count:            len(reviews),
		PriorityApproval: false,
	}

	for _, review := range reviews {
		onBranchHead := sha == review.GetCommitID()

		reviewerName := review.GetUser().GetLogin()
		reviewerAssociation := review.GetAuthorAssociation()
		isC3iTeam := reviewerAssociation == "MEMBER" || reviewerAssociation == "COLLABORATOR"

		switch review.GetState() { // Either as indicated by the reviewer or by updates from the GitHub API
		case "CHANGES_REQUESTED":
			if isC3iTeam {
				summary.HeadCommitBlockers, _ = appendUnique(summary.HeadCommitBlockers, reviewerName)
			}

			removed := false
			summary.HeadCommitApprovals, removed = removeUnique(summary.HeadCommitApprovals, reviewerName)
			if removed {
				// If a reviewer retracted their reivew, the count needs to be adjusted

				if isC3iTeam {
					summary.ValidApprovals = summary.ValidApprovals - 1
				}

				switch reviewerName {
				case "madebr", "SpaceIm", "ericLemanissier", "prince-chrismc", "Croydon", "intelligide", "theirix", "gocarlos":
					// If a community reviewer retracted their reivew, the count needs to be adjusted
					summary.ValidApprovals = summary.ValidApprovals - 1
				default:
				}
			}

		case "APPROVED":
			if onBranchHead {
				approvals, new := appendUnique(summary.HeadCommitApprovals, reviewerName)
				if !new { // Duplicate review (usually an accident)
					break
				}

				summary.HeadCommitApprovals = approvals
				if isC3iTeam {
					summary.PriorityApproval = true
					summary.ValidApprovals = summary.ValidApprovals + 1
				}
			}

			// We always count the community reviewer approvals (for any commit) because this matches closer with the GitHub UI
			switch reviewerName {
			case "madebr", "SpaceIm", "ericLemanissier", "prince-chrismc", "Croydon", "intelligide", "theirix", "gocarlos", "mathbunnyru":
				summary.ValidApprovals = summary.ValidApprovals + 1
			default:
			}

			summary.HeadCommitBlockers, _ = removeUnique(summary.HeadCommitBlockers, reviewerName)

		case "DISMISSED":
			// Out-dated Approvals are transformed into comments https://github.com/conan-io/conan-center-index/pull/3855#issuecomment-770120073
			summary.HeadCommitBlockers, _ = removeUnique(summary.HeadCommitBlockers, reviewerName)
		default:
		}
	}

	return summary
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
