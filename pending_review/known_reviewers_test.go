package pending_review

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	gock "gopkg.in/h2non/gock.v1"
)

func TestIsReviwers(t *testing.T) {
	reviewers := ConanCenterReviewers{
		[]Reviewer{
			{User: "danimtb", Type: Team, Requested: true},
			{User: "lasote", Type: Team, Requested: false},
			{User: "jgsogo", Type: Team, Requested: true},
			{User: "czoido", Type: Team, Requested: false},
			{User: "memsharded", Type: Team, Requested: false},
			{User: "SSE4", Type: Team, Requested: true},
			{User: "uilianries", Type: Team, Requested: true},
			{User: "madebr", Type: Community, Requested: false},
			{User: "SpaceIm", Type: Community, Requested: false},
			{User: "ericLemanissier", Type: Community, Requested: false},
		},
	}

	assert.Equal(t, true, reviewers.IsTeamMember("danimtb"))
	assert.Equal(t, false, reviewers.IsCommunityMember("danimtb"))
	assert.Equal(t, true, reviewers.IsTeamMember("czoido"))
	assert.Equal(t, false, reviewers.IsCommunityMember("czoido"))

	assert.Equal(t, false, reviewers.IsTeamMember("madebr"))
	assert.Equal(t, true, reviewers.IsCommunityMember("madebr"))
	assert.Equal(t, false, reviewers.IsTeamMember("ericLemanissier"))
	assert.Equal(t, true, reviewers.IsCommunityMember("ericLemanissier"))
}

func TestParseReviewers(t *testing.T) {
	reviewers, err := parseReviewers(`reviewers:
  # List with users whose review is taken into account so that a pull-request is merged.
  #   - <user>: Name of the github user
  #   - <type>: Either 'community' for community reviewers or 'team' for Conan reviewers.
  #   - <request_reviews>: Make the bot proactively request the user's review of pull-requests ready for review.
  - user: "danimtb"
    type: "team"
    request_reviews: true
  - user: "lasote"
    type: "team"
    request_reviews: false
  - user: "jgsogo"
    type: "team"
    request_reviews: true
  - user: "czoido"
    type: "team"
    request_reviews: false
  - user: "memsharded"
    type: "team"
    request_reviews: false
  - user: "SSE4"
    type: "team"
    request_reviews: true
  - user: "uilianries"
    type: "team"
    request_reviews: true
  - user: "madebr"
    type: "community"
    request_reviews: false
  - user: "SpaceIm"
    type: "community"
    request_reviews: false
  - user: "ericLemanissier"
    type: "community"
    request_reviews: false
  `)

	expected := ConanCenterReviewers{
		[]Reviewer{
			{User: "danimtb", Type: Team, Requested: true},
			{User: "lasote", Type: Team, Requested: false},
			{User: "jgsogo", Type: Team, Requested: true},
			{User: "czoido", Type: Team, Requested: false},
			{User: "memsharded", Type: Team, Requested: false},
			{User: "SSE4", Type: Team, Requested: true},
			{User: "uilianries", Type: Team, Requested: true},
			{User: "madebr", Type: Community, Requested: false},
			{User: "SpaceIm", Type: Community, Requested: false},
			{User: "ericLemanissier", Type: Community, Requested: false},
		},
	}

	assert.Equal(t, nil, err)
	assert.Equal(t, &expected, reviewers)
}

func TestDownloadKnownReviewersList(t *testing.T) {
	defer gock.Off()

	gock.New("https://api.github.com").
		Get("/repos/conan-io/conan-center-index/contents/.c3i/reviewers.yml").
		MatchParam("ref", "master").
		Reply(200).
		BodyString(`{
			"name": "reviewers.yml",
			"path": ".c3i/reviewers.yml",
			"sha": "0f6506f08eb81cf7ce0a50db7564786bf2b25cf2",
			"size": 1783,
			"url": "https://api.github.com/repos/conan-io/conan-center-index/contents/.c3i/reviewers.yml?ref=master",
			"html_url": "https://github.com/conan-io/conan-center-index/blob/master/.c3i/reviewers.yml",
			"git_url": "https://api.github.com/repos/conan-io/conan-center-index/git/blobs/0f6506f08eb81cf7ce0a50db7564786bf2b25cf2",
			"download_url": "https://raw.githubusercontent.com/conan-io/conan-center-index/master/.c3i/reviewers.yml",
			"type": "file",
			"content": "cmV2aWV3ZXJzOgogICMgTGlzdCB3aXRoIHVzZXJzIHdob3NlIHJldmlldyBp\ncyB0YWtlbiBpbnRvIGFjY291bnQgc28gdGhhdCBhIHB1bGwtcmVxdWVzdCBp\ncyBtZXJnZWQuCiAgIyAgIC0gPHVzZXI+OiBOYW1lIG9mIHRoZSBnaXRodWIg\ndXNlcgogICMgICAtIDx0eXBlPjogRWl0aGVyICdjb21tdW5pdHknIGZvciBj\nb21tdW5pdHkgcmV2aWV3ZXJzIG9yICd0ZWFtJyBmb3IgQ29uYW4gcmV2aWV3\nZXJzLgogICMgICAtIDxyZXF1ZXN0X3Jldmlld3M+OiBNYWtlIHRoZSBib3Qg\ncHJvYWN0aXZlbHkgcmVxdWVzdCB0aGUgdXNlcidzIHJldmlldyBvZiBwdWxs\nLXJlcXVlc3RzIHJlYWR5IGZvciByZXZpZXcuCiAgLSB1c2VyOiAiZGFuaW10\nYiIKICAgIHR5cGU6ICJ0ZWFtIgogICAgcmVxdWVzdF9yZXZpZXdzOiB0cnVl\nCiAgLSB1c2VyOiAibGFzb3RlIgogICAgdHlwZTogInRlYW0iCiAgICByZXF1\nZXN0X3Jldmlld3M6IGZhbHNlCiAgLSB1c2VyOiAiamdzb2dvIgogICAgdHlw\nZTogInRlYW0iCiAgICByZXF1ZXN0X3Jldmlld3M6IHRydWUKICAtIHVzZXI6\nICJjem9pZG8iCiAgICB0eXBlOiAidGVhbSIKICAgIHJlcXVlc3RfcmV2aWV3\nczogZmFsc2UKICAtIHVzZXI6ICJtZW1zaGFyZGVkIgogICAgdHlwZTogInRl\nYW0iCiAgICByZXF1ZXN0X3Jldmlld3M6IGZhbHNlCiAgLSB1c2VyOiAiU1NF\nNCIKICAgIHR5cGU6ICJ0ZWFtIgogICAgcmVxdWVzdF9yZXZpZXdzOiB0cnVl\nCiAgLSB1c2VyOiAidWlsaWFucmllcyIKICAgIHR5cGU6ICJ0ZWFtIgogICAg\ncmVxdWVzdF9yZXZpZXdzOiB0cnVlCiAgLSB1c2VyOiAibWFkZWJyIgogICAg\ndHlwZTogImNvbW11bml0eSIKICAgIHJlcXVlc3RfcmV2aWV3czogZmFsc2UK\nICAtIHVzZXI6ICJTcGFjZUltIgogICAgdHlwZTogImNvbW11bml0eSIKICAg\nIHJlcXVlc3RfcmV2aWV3czogZmFsc2UKICAtIHVzZXI6ICJlcmljTGVtYW5p\nc3NpZXIiCiAgICB0eXBlOiAiY29tbXVuaXR5IgogICAgcmVxdWVzdF9yZXZp\nZXdzOiBmYWxzZQogIC0gdXNlcjogInByaW5jZS1jaHJpc21jIgogICAgdHlw\nZTogInRlYW0iCiAgICByZXF1ZXN0X3Jldmlld3M6IGZhbHNlCiAgLSB1c2Vy\nOiAiQ3JveWRvbiIKICAgIHR5cGU6ICJjb21tdW5pdHkiCiAgICByZXF1ZXN0\nX3Jldmlld3M6IGZhbHNlCiAgLSB1c2VyOiAiaW50ZWxsaWdpZGUiCiAgICB0\neXBlOiAiY29tbXVuaXR5IgogICAgcmVxdWVzdF9yZXZpZXdzOiBmYWxzZQog\nIC0gdXNlcjogInRoZWlyaXgiCiAgICB0eXBlOiAiY29tbXVuaXR5IgogICAg\ncmVxdWVzdF9yZXZpZXdzOiBmYWxzZQogIC0gdXNlcjogImdvY2FybG9zIgog\nICAgdHlwZTogImNvbW11bml0eSIKICAgIHJlcXVlc3RfcmV2aWV3czogZmFs\nc2UKICAtIHVzZXI6ICJtYXRoYnVubnlydSIKICAgIHR5cGU6ICJjb21tdW5p\ndHkiCiAgICByZXF1ZXN0X3Jldmlld3M6IGZhbHNlCiAgLSB1c2VyOiAiZXJp\nY3JpZmYiCiAgICB0eXBlOiAiY29tbXVuaXR5IgogICAgcmVxdWVzdF9yZXZp\nZXdzOiBmYWxzZQogIC0gdXNlcjogInRvZ2UiCiAgICB0eXBlOiAiY29tbXVu\naXR5IgogICAgcmVxdWVzdF9yZXZpZXdzOiBmYWxzZQogIC0gdXNlcjogIkFu\nZHJleU1sYXNoa2luIgogICAgdHlwZTogImNvbW11bml0eSIKICAgIHJlcXVl\nc3RfcmV2aWV3czogZmFsc2UKICAtIHVzZXI6ICJNYXJ0aW5EZWxpbGxlIgog\nICAgdHlwZTogImNvbW11bml0eSIKICAgIHJlcXVlc3RfcmV2aWV3czogZmFs\nc2UKICAtIHVzZXI6ICJkbW4tc3RhciIKICAgIHR5cGU6ICJjb21tdW5pdHki\nCiAgICByZXF1ZXN0X3Jldmlld3M6IGZhbHNlCg==\n",
			"encoding": "base64",
			"_links": {
			  "self": "https://api.github.com/repos/conan-io/conan-center-index/contents/.c3i/reviewers.yml?ref=master",
			  "git": "https://api.github.com/repos/conan-io/conan-center-index/git/blobs/0f6506f08eb81cf7ce0a50db7564786bf2b25cf2",
			  "html": "https://github.com/conan-io/conan-center-index/blob/master/.c3i/reviewers.yml"
			}
		  }`)

	expected := ConanCenterReviewers{Reviewers: []Reviewer{{User: "danimtb", Type: "team", Requested: true}, {User: "lasote", Type: "team", Requested: false}, {User: "jgsogo", Type: "team", Requested: true}, {User: "czoido", Type: "team", Requested: false}, {User: "memsharded", Type: "team", Requested: false}, {User: "SSE4", Type: "team", Requested: true}, {User: "uilianries", Type: "team", Requested: true}, {User: "madebr", Type: "community", Requested: false}, {User: "SpaceIm", Type: "community", Requested: false}, {User: "ericLemanissier", Type: "community", Requested: false}, {User: "prince-chrismc", Type: "team", Requested: false}, {User: "Croydon", Type: "community", Requested: false}, {User: "intelligide", Type: "community", Requested: false}, {User: "theirix", Type: "community", Requested: false}, {User: "gocarlos", Type: "community", Requested: false}, {User: "mathbunnyru", Type: "community", Requested: false}, {User: "ericriff", Type: "community", Requested: false}, {User: "toge", Type: "community", Requested: false}, {User: "AndreyMlashkin", Type: "community", Requested: false}, {User: "MartinDelille", Type: "community", Requested: false}, {User: "dmn-star", Type: "community", Requested: false}}}

	reviewers, err := DownloadKnownReviewersList(context.Background(), NewClient(&http.Client{}, WorkingRepository{}))
	assert.Equal(t, nil, err)
	assert.Equal(t, &expected, reviewers)

	assert.Equal(t, true, gock.IsDone())
}
