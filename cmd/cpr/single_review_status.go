package main

import (
	"context"
	"fmt"

	"github.com/prince-chrismc/conan-center-index-pending-review/v4/internal"
	"github.com/prince-chrismc/conan-center-index-pending-review/v4/pending_review"
)

// SingleReviewStatus analysis a pull requests
func SingleReviewStatus(token string, pr uint, owner string, repo string) error {
	context := context.Background()
	client, err := internal.MakeClient(context, token, pending_review.WorkingRepository{Owner: owner, Name: repo})
	if err != nil {
		return fmt.Errorf("problem making client %w", err)
	}

	defer fmt.Println("::endgroup") // Always print when we return

	fmt.Println("::group::👤 Initializing list of known reviewers")
	reviewers, err := pending_review.DownloadKnownReviewersList(context, client)
	if err != nil {
		return fmt.Errorf("problem getting list of known reviewers from CCI %w", err)
	}
	fmt.Printf("%+v\n", reviewers)
	fmt.Println("::endgroup")

	fmt.Println("::group::🔎 Gathering data on all Pull Requests")

	pull, _, err := client.PullRequests.Get(context, "conan-io", "conan-center-index", int(pr))
	if err != nil {
		return fmt.Errorf("problem getting pull request list %w", err)
	}

	pulls := []*pending_review.PullRequest{pull}
	_, _, err = gatherReviewStatus(context, client, reviewers, pulls)
	if err != nil {
		return fmt.Errorf("problem getting review status %w", err)
	}

	fmt.Println("::endgroup")

	fmt.Println("::group::🖊️ Rendering data and saving results!")

	fmt.Println("::endgroup")

	return nil
}
