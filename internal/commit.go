package internal

import (
	"context"
	"fmt"
	"time"

	git "github.com/go-git/go-git/v5/plumbing"
	"github.com/google/go-github/v33/github"
	"github.com/prince-chrismc/conan-center-index-pending-review/v2/pkg/pending_review"
)

func UpdateDataFile(context context.Context, client *pending_review.Client, file string, content []byte) error {
	fileContent, _, _, err := client.Repositories.GetContents(context, "prince-chrismc", "conan-center-index-pending-review", file, &github.RepositoryContentGetOptions{
		Ref: "raw-data",
	})
	if err != nil {
		return err
	}

	newSha := git.ComputeHash(git.BlobObject, content).String()
	if newSha != fileContent.GetSHA() {
		opts := &github.RepositoryContentFileOptions{
			SHA:       fileContent.SHA, // Required to edit the file
			Message:   github.String(file + ": New data - " + time.Now().Format(time.RFC3339)),
			Content:   content,
			Branch:    github.String("raw-data"),
			Committer: &github.CommitAuthor{Name: github.String("github-actions[bot]"), Email: github.String("github-actions[bot]@users.noreply.github.com")},
		}
		_, _, err = client.Repositories.UpdateFile(context, "prince-chrismc", "conan-center-index-pending-review", file, opts)
		if err != nil {
			return err
		}
	} else {
		fmt.Printf("Content for '%s' was the same\n", file)
	}

	return nil
}
