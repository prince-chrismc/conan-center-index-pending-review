package pending_review

import "fmt"

func evaluateSummary(p *PullRequestSummary) error {
	if p.Change == DOCS { // Always save documentation pull requests
		return nil
	}

	if p.Change == CONFIG && p.CciBotPassed { // Always save infrastructure pull requests that are passing
		return nil
	}

	if p.Summary.Count < 1 { // Has not been looked at...
		return nil // let's save it! So it can get some attention
	}

	if len(p.Summary.Approvals) > 0 { // It's been approved!
		return nil
	}

	if p.LastCommitAt.After(p.Summary.LastReview.SubmittedAt) { // OP has presumably applied review comments
		return nil // Let's save it so it gets another pass
	}

	// Pull request has comment which have not been corrected by a commit.
	// Assuming more work is required.

	return fmt.Errorf("%w", ErrWorkRequired)
}
