package pending_review

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	gock "gopkg.in/h2non/gock.v1"
)

func TestGetDataFile(t *testing.T) {
	defer gock.Off()

	gock.New("https://api.github.com").
		Get("/repos/conan-io/conan-center/contents/.c3i/reviewers.yml").
		MatchParam("ref", "master").
		Reply(200).
		BodyString(`{
			"name": ".c3i/reviewers.yml",
			"path": ".c3i/reviewers.yml",
			"sha": "c958c6f2fae4c522d1bc690d2dc4bfc1d3101474",
			"size": 16,
			"url": "https://api.github.com/repos/conan-io/conan-center/contents/.c3i/reviewers.yml?ref=master",
			"html_url": "https://github.com/conan-io/conan-center/blob/master/.c3i/reviewers.yml",
			"git_url": "https://api.github.com/repos/conan-io/conan-center/git/blobs/c958c6f2fae4c522d1bc690d2dc4bfc1d3101474",
			"download_url": "https://raw.githubusercontent.com/conan-io/conan-center/master/.c3i/reviewers.yml",
			"type": "file",
			"content": "SGVsbG8gV29ybGQh",
			"encoding": "base64",
			"_links": {
			  "self": "https://api.github.com/repos/conan-io/conan-center/contents/.c3i/reviewers.yml?ref=master",
			  "git": "https://api.github.com/repos/conan-io/conan-center/git/blobs/c958c6f2fae4c522d1bc690d2dc4bfc1d3101474",
			  "html": "https://github.com/conan-io/conan-center/blob/master/.c3i/reviewers.yml"
			}
		  }`)

	reviewers, err := DownloadKnownReviewersList(context.Background(), NewClient(&http.Client{}))
	assert.Equal(t, nil, err)

	assert.Equal(t, nil, err)
	assert.Equal(t, "Hello World!", reviewers)

	assert.Equal(t, true, gock.IsDone())
}
