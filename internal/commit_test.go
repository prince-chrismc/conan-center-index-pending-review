package internal

import (
	"context"
	"net/http"
	"testing"

	"github.com/prince-chrismc/conan-center-index-pending-review/v4/pending_review"
	"github.com/stretchr/testify/assert"
	gock "gopkg.in/h2non/gock.v1"
)

func TestGetDataFile(t *testing.T) {
	defer gock.Off()

	gock.New("https://api.github.com").
		Get("/repos/prince-chrismc/conan-center-index-pending-review/contents/data.txt").
		MatchParam("ref", "raw-data").
		Reply(200).
		BodyString(`{
			"name": "data.txt",
			"path": "data.txt",
			"sha": "c958c6f2fae4c522d1bc690d2dc4bfc1d3101474",
			"size": 16,
			"url": "https://api.github.com/repos/prince-chrismc/conan-center-index-pending-review/contents/data.txt?ref=raw-data",
			"html_url": "https://github.com/prince-chrismc/conan-center-index-pending-review/blob/raw-data/data.txt",
			"git_url": "https://api.github.com/repos/prince-chrismc/conan-center-index-pending-review/git/blobs/c958c6f2fae4c522d1bc690d2dc4bfc1d3101474",
			"download_url": "https://raw.githubusercontent.com/prince-chrismc/conan-center-index-pending-review/raw-data/data.txt",
			"type": "file",
			"content": "SGVsbG8gV29ybGQh",
			"encoding": "base64",
			"_links": {
			  "self": "https://api.github.com/repos/prince-chrismc/conan-center-index-pending-review/contents/data.txt?ref=raw-data",
			  "git": "https://api.github.com/repos/prince-chrismc/conan-center-index-pending-review/git/blobs/c958c6f2fae4c522d1bc690d2dc4bfc1d3101474",
			  "html": "https://github.com/prince-chrismc/conan-center-index-pending-review/blob/raw-data/data.txt"
			}
		  }`)

	fileContent, err := GetDataFile(context.Background(), pending_review.NewClient(&http.Client{}, pending_review.WorkingRepository{Owner: "prince-chrismc", Name: "conan-center-index-pending-review"}), "data.txt")
	assert.Equal(t, nil, err)

	data, err := fileContent.GetContent()
	assert.Equal(t, nil, err)
	assert.Equal(t, "Hello World!", data)

	assert.Equal(t, true, gock.IsDone())
}

func TestGetJSONFile(t *testing.T) {
	defer gock.Off()

	gock.New("https://api.github.com").
		Get("/repos/prince-chrismc/conan-center-index-pending-review/contents/data.json").
		MatchParam("ref", "raw-data").
		Reply(200).
		BodyString(`{
			"name": "data.json",
			"path": "data.json",
			"sha": "c958c6f2fae4c522d1bc690d2dc4bfc1d3101474",
			"size": 20,
			"url": "https://api.github.com/repos/prince-chrismc/conan-center-index-pending-review/contents/data.json?ref=raw-data",
			"html_url": "https://github.com/prince-chrismc/conan-center-index-pending-review/blob/raw-data/data.json",
			"git_url": "https://api.github.com/repos/prince-chrismc/conan-center-index-pending-review/git/blobs/c958c6f2fae4c522d1bc690d2dc4bfc1d3101474",
			"download_url": "https://raw.githubusercontent.com/prince-chrismc/conan-center-index-pending-review/raw-data/data.json",
			"type": "file",
			"content": "eyJmb28iOiJiYXIifQ==",
			"encoding": "base64",
			"_links": {
			  "self": "https://api.github.com/repos/prince-chrismc/conan-center-index-pending-review/contents/data.json?ref=raw-data",
			  "git": "https://api.github.com/repos/prince-chrismc/conan-center-index-pending-review/git/blobs/c958c6f2fae4c522d1bc690d2dc4bfc1d3101474",
			  "html": "https://github.com/prince-chrismc/conan-center-index-pending-review/blob/raw-data/data.json"
			}
		  }`)

	contents := map[string]string{}
	err := GetJSONFile(context.Background(), pending_review.NewClient(&http.Client{}, pending_review.WorkingRepository{
		Owner: "prince-chrismc", Name: "conan-center-index-pending-review",
	}), "data.json", &contents)
	assert.Equal(t, nil, err)

	assert.Equal(t, nil, err)
	assert.Equal(t, map[string]string{"foo": "bar"}, contents)

	assert.Equal(t, true, gock.IsDone())
}
