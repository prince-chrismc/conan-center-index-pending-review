package main

import (
	"context"
	"image"
	"image/png"
	"strings"
	"time"

	"github.com/prince-chrismc/conan-center-index-pending-review/v2/internal"
	"github.com/prince-chrismc/conan-center-index-pending-review/v2/internal/duration"
	"github.com/prince-chrismc/conan-center-index-pending-review/v2/pkg/pending_review"
)

func GetOvmPngFromThisWeek(context context.Context, client *pending_review.Client) ([]image.Image, error) {
	window := time.Now().Truncate(duration.WEEK)

	commits, err := internal.GetCommitsSince(context, client, "open-versus-merged.png", window)
	if err != nil {
		return nil, err
	}

	snapshots := make([]image.Image, 0)

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
