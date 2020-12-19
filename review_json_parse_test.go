package main

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/google/go-github/github"
)

func TestMain_UnmarshallReviews(t *testing.T) {
	sample := `[
		{
		  "id": 555740828,
		  "node_id": "MDE3OlB1bGxSZXF1ZXN0UmV2aWV3NTU1NzQwODI4",
		  "user": {
			"login": "Croydon",
			"id": 1593194
		  },
		  "body": "",
		  "state": "APPROVED",
		  "html_url": "https://github.com/conan-io/conan-center-index/pull/3953#pullrequestreview-555740828",
		  "pull_request_url": "https://api.github.com/repos/conan-io/conan-center-index/pulls/3953",
		  "author_association": "CONTRIBUTOR",
		  "_links": {
			"html": {
			  "href": "https://github.com/conan-io/conan-center-index/pull/3953#pullrequestreview-555740828"
			},
			"pull_request": {
			  "href": "https://api.github.com/repos/conan-io/conan-center-index/pulls/3953"
			}
		  },
		  "submitted_at": "2020-12-18T19:02:20Z",
		  "commit_id": "50e1ef0062f5961effbaacb8caa0c3590b7e61bf"
		},
		{
		  "id": 555908227,
		  "node_id": "MDE3OlB1bGxSZXF1ZXN0UmV2aWV3NTU1OTA4MjI3",
		  "user": {
			"login": "madebr",
			"id": 4138939
		  },
		  "body": "",
		  "state": "COMMENTED",
		  "html_url": "https://github.com/conan-io/conan-center-index/pull/3953#pullrequestreview-555908227",
		  "pull_request_url": "https://api.github.com/repos/conan-io/conan-center-index/pulls/3953",
		  "author_association": "CONTRIBUTOR",
		  "_links": {
			"html": {
			  "href": "https://github.com/conan-io/conan-center-index/pull/3953#pullrequestreview-555908227"
			},
			"pull_request": {
			  "href": "https://api.github.com/repos/conan-io/conan-center-index/pulls/3953"
			}
		  },
		  "submitted_at": "2020-12-19T02:26:45Z",
		  "commit_id": "50e1ef0062f5961effbaacb8caa0c3590b7e61bf"
		}
	  ]`

	reviews := new(ReviewsResponse)
	json.Unmarshal([]byte(sample), reviews)

	if want := testReviews; !reflect.DeepEqual(reviews, want) {
		t.Errorf("json.Unmarshal_Reviews returned %+v, \n\n want %+v", reviews, want)
	}
}

var testReviews = &ReviewsResponse{
	&Review{
		ID:     github.Int(555740828),
		NodeID: github.String("MDE3OlB1bGxSZXF1ZXN0UmV2aWV3NTU1NzQwODI4"),
		User: &github.User{
			Login: github.String("Croydon"),
			ID:    github.Int64(1593194),
		},
		State:             github.String("APPROVED"),
		AuthorAssociation: github.String("CONTRIBUTOR"),
		CommitID:          github.String("50e1ef0062f5961effbaacb8caa0c3590b7e61bf"),
	},
	&Review{
		ID:     github.Int(555908227),
		NodeID: github.String("MDE3OlB1bGxSZXF1ZXN0UmV2aWV3NTU1OTA4MjI3"),
		User: &github.User{
			Login: github.String("madebr"),
			ID:    github.Int64(4138939),
		},
		State:             github.String("COMMENTED"),
		AuthorAssociation: github.String("CONTRIBUTOR"),
		CommitID:          github.String("50e1ef0062f5961effbaacb8caa0c3590b7e61bf"),
	},
}
