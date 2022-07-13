package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	git "github.com/go-git/go-git/v5/plumbing"
	"github.com/google/go-github/v42/github"
	"github.com/prince-chrismc/conan-center-index-pending-review/v3/pending_review"
)

// GetCommits returns a list of of number from the raw-data branch
func GetCommits(context context.Context, client *pending_review.Client, file string, count int, owner string, repo string) ([]*github.RepositoryCommit, error) {
	commits, _, err := client.Repositories.ListCommits(context, owner, repo,
		&github.CommitsListOptions{SHA: "raw-data", Path: file, ListOptions: github.ListOptions{PerPage: count}})
	if err != nil {
		return nil, err
	}

	return commits, nil
}

// GetFileAtRef returns the content of file from the root directory from a commit sha
func GetFileAtRef(context context.Context, client *pending_review.Client, file string, sha string, owner string, repo string) (*github.RepositoryContent, error) {
	fileContent, _, _, err := client.Repositories.GetContents(context, owner, repo, file,
		&github.RepositoryContentGetOptions{Ref: sha})
	if err != nil {
		return nil, err
	}

	return fileContent, nil
}

// GetDataFile returns the content of file from the root directory of the raw-data branch
func GetDataFile(context context.Context, client *pending_review.Client, file string, owner string, repo string) (*github.RepositoryContent, error) {
	fileContent, _, _, err := client.Repositories.GetContents(context, owner, repo, file,
		&github.RepositoryContentGetOptions{Ref: "raw-data"})
	if err != nil {
		return nil, err
	}

	return fileContent, nil
}

// GetJSONFile returns the JSON structure of file from the root directory of the raw-data branch
func GetJSONFile(context context.Context, client *pending_review.Client, file string, content interface{}, owner string, repo string) error {
	fileContent, err := GetDataFile(context, client, file, owner, repo)
	if err != nil {
		return err
	}

	str, err := fileContent.GetContent()
	if err != nil {
		return err
	}

	if err := json.Unmarshal([]byte(str), content); err != nil {
		return err
	}

	return nil
}

// UpdateDataFileAtRef commits the new content if it's different. It returns if the modification took place and any error encountered.
func UpdateFileAtRef(context context.Context, client *pending_review.Client, file string, branch string, content []byte, owner string, repo string) (bool, error) {
	fileContent, err := GetFileAtRef(context, client, file, branch, owner, repo)
	if err != nil {
		return false, err
	}

	newSha := git.ComputeHash(git.BlobObject, content).String()
	if newSha == fileContent.GetSHA() {
		fmt.Printf("Content for '%s' was the same\n", file)
		return false, nil
	}

	opts := &github.RepositoryContentFileOptions{
		SHA:     fileContent.SHA, // Required to edit the file
		Message: github.String(file + ": New data - " + time.Now().Format(time.RFC3339)),
		Content: content,
		Branch:  github.String(branch),
		Committer: &github.CommitAuthor{Name: github.String("github-actions[bot]"),
			Email: github.String("github-actions[bot]@users.noreply.github.com")},
	}
	_, _, err = client.Repositories.UpdateFile(context, owner, repo, file, opts)
	if err != nil {
		return false, err
	}

	return true, nil
}

// UpdateDataFile commits the new content to `raw-data` if it's different. It returns if the modification took place and any error encountered.
func UpdateDataFile(context context.Context, client *pending_review.Client, file string, content []byte, owner string, repo string) (bool, error) {
	return UpdateFileAtRef(context, client, file, "raw-data", content, owner, repo)
}

// UpdateJSONFile commits the new content if it's different. It returns if the modification took place and any error encountered.
func UpdateJSONFile(context context.Context, client *pending_review.Client, file string, content interface{}, owner string, repo string) (bool, error) {
	data, err := json.MarshalIndent(content, "", "   ")
	if err != nil {
		return false, err
	}

	updated, err := UpdateDataFile(context, client, file, data, owner, repo)
	if err != nil {
		return false, err
	}

	return updated, nil
}
