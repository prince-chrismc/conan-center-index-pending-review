package main

import (
	"context"
	"image"
	"image/png"
	"strings"

	"github.com/prince-chrismc/conan-center-index-pending-review/v2/internal"
	"github.com/prince-chrismc/conan-center-index-pending-review/v2/pkg/pending_review"
)

// GetOvmPngFromThisWeek returns the last sevent PNGs that have been uploaded.
// Note: This is under the basis that the deployment does this every ~24hrs.
func GetOvmPngFromThisWeek(context context.Context, client *pending_review.Client) ([]image.Image, error) {
	commits, err := internal.GetCommits(context, client, "open-versus-merged.png", 7)
	if err != nil {
		return nil, err
	}

	snapshots := make([]image.Image, 0, 7)

	for _, commit := range commits {
		fileContent, err := internal.GetDataFileAtRef(context, client, "open-versus-merged.png", commit.GetSHA())
		if err != nil {
			return snapshots, err
		}

		str, err := fileContent.GetContent()
		if err != nil {
			return snapshots, err
		}

		img, err := png.Decode(strings.NewReader(str))
		if err != nil {
			return snapshots, err
		}

		snapshots = append(snapshots, img)
	}

	return snapshots, nil
}
