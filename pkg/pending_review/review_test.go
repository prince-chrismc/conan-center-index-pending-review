package pending_review

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func parseReviewJSON(t *testing.T, str string) []*PullRequestReview {
	var files []*PullRequestReview

	if err := json.Unmarshal([]byte(str), &files); err != nil {
		t.Fatal()
	}

	return files
}

func TestKnowCase6144(t *testing.T) {
	reviews := parseReviewJSON(t, `[
		{
		  "id": 698590899,
		  "node_id": "MDE3OlB1bGxSZXF1ZXN0UmV2aWV3Njk4NTkwODk5",
		  "user": {
			"login": "prince-chrismc",
			"id": 16867443,
			"node_id": "MDQ6VXNlcjE2ODY3NDQz",
			"avatar_url": "https://avatars.githubusercontent.com/u/16867443?u=410263f66886d2d12cdb8da43e7da02d5423380a&v=4",
			"gravatar_id": "",
			"url": "https://api.github.com/users/prince-chrismc",
			"html_url": "https://github.com/prince-chrismc",
			"followers_url": "https://api.github.com/users/prince-chrismc/followers",
			"following_url": "https://api.github.com/users/prince-chrismc/following{/other_user}",
			"gists_url": "https://api.github.com/users/prince-chrismc/gists{/gist_id}",
			"starred_url": "https://api.github.com/users/prince-chrismc/starred{/owner}{/repo}",
			"subscriptions_url": "https://api.github.com/users/prince-chrismc/subscriptions",
			"organizations_url": "https://api.github.com/users/prince-chrismc/orgs",
			"repos_url": "https://api.github.com/users/prince-chrismc/repos",
			"events_url": "https://api.github.com/users/prince-chrismc/events{/privacy}",
			"received_events_url": "https://api.github.com/users/prince-chrismc/received_events",
			"type": "User",
			"site_admin": false
		  },
		  "body": "",
		  "state": "APPROVED",
		  "html_url": "https://github.com/conan-io/conan-center-index/pull/6144#pullrequestreview-698590899",
		  "pull_request_url": "https://api.github.com/repos/conan-io/conan-center-index/pulls/6144",
		  "author_association": "CONTRIBUTOR",
		  "_links": {
			"html": {
			  "href": "https://github.com/conan-io/conan-center-index/pull/6144#pullrequestreview-698590899"
			},
			"pull_request": {
			  "href": "https://api.github.com/repos/conan-io/conan-center-index/pulls/6144"
			}
		  },
		  "submitted_at": "2021-07-03T12:32:33Z",
		  "commit_id": "3093bad9162e288d55eeddec288b0481d964518e"
		},
		{
		  "id": 698857675,
		  "node_id": "MDE3OlB1bGxSZXF1ZXN0UmV2aWV3Njk4ODU3Njc1",
		  "user": {
			"login": "SSE4",
			"id": 870236,
			"node_id": "MDQ6VXNlcjg3MDIzNg==",
			"avatar_url": "https://avatars.githubusercontent.com/u/870236?v=4",
			"gravatar_id": "",
			"url": "https://api.github.com/users/SSE4",
			"html_url": "https://github.com/SSE4",
			"followers_url": "https://api.github.com/users/SSE4/followers",
			"following_url": "https://api.github.com/users/SSE4/following{/other_user}",
			"gists_url": "https://api.github.com/users/SSE4/gists{/gist_id}",
			"starred_url": "https://api.github.com/users/SSE4/starred{/owner}{/repo}",
			"subscriptions_url": "https://api.github.com/users/SSE4/subscriptions",
			"organizations_url": "https://api.github.com/users/SSE4/orgs",
			"repos_url": "https://api.github.com/users/SSE4/repos",
			"events_url": "https://api.github.com/users/SSE4/events{/privacy}",
			"received_events_url": "https://api.github.com/users/SSE4/received_events",
			"type": "User",
			"site_admin": false
		  },
		  "body": "",
		  "state": "APPROVED",
		  "html_url": "https://github.com/conan-io/conan-center-index/pull/6144#pullrequestreview-698857675",
		  "pull_request_url": "https://api.github.com/repos/conan-io/conan-center-index/pulls/6144",
		  "author_association": "COLLABORATOR",
		  "_links": {
			"html": {
			  "href": "https://github.com/conan-io/conan-center-index/pull/6144#pullrequestreview-698857675"
			},
			"pull_request": {
			  "href": "https://api.github.com/repos/conan-io/conan-center-index/pulls/6144"
			}
		  },
		  "submitted_at": "2021-07-05T07:07:00Z",
		  "commit_id": "3093bad9162e288d55eeddec288b0481d964518e"
		},
		{
		  "id": 698894036,
		  "node_id": "MDE3OlB1bGxSZXF1ZXN0UmV2aWV3Njk4ODk0MDM2",
		  "user": {
			"login": "jgsogo",
			"id": 1406456,
			"node_id": "MDQ6VXNlcjE0MDY0NTY=",
			"avatar_url": "https://avatars.githubusercontent.com/u/1406456?u=b056762d4b8488fb294022c204d8b79389debe76&v=4",
			"gravatar_id": "",
			"url": "https://api.github.com/users/jgsogo",
			"html_url": "https://github.com/jgsogo",
			"followers_url": "https://api.github.com/users/jgsogo/followers",
			"following_url": "https://api.github.com/users/jgsogo/following{/other_user}",
			"gists_url": "https://api.github.com/users/jgsogo/gists{/gist_id}",
			"starred_url": "https://api.github.com/users/jgsogo/starred{/owner}{/repo}",
			"subscriptions_url": "https://api.github.com/users/jgsogo/subscriptions",
			"organizations_url": "https://api.github.com/users/jgsogo/orgs",
			"repos_url": "https://api.github.com/users/jgsogo/repos",
			"events_url": "https://api.github.com/users/jgsogo/events{/privacy}",
			"received_events_url": "https://api.github.com/users/jgsogo/received_events",
			"type": "User",
			"site_admin": false
		  },
		  "body": "",
		  "state": "APPROVED",
		  "html_url": "https://github.com/conan-io/conan-center-index/pull/6144#pullrequestreview-698894036",
		  "pull_request_url": "https://api.github.com/repos/conan-io/conan-center-index/pulls/6144",
		  "author_association": "MEMBER",
		  "_links": {
			"html": {
			  "href": "https://github.com/conan-io/conan-center-index/pull/6144#pullrequestreview-698894036"
			},
			"pull_request": {
			  "href": "https://api.github.com/repos/conan-io/conan-center-index/pulls/6144"
			}
		  },
		  "submitted_at": "2021-07-05T07:53:50Z",
		  "commit_id": "3093bad9162e288d55eeddec288b0481d964518e"
		},
		{
		  "id": 698926654,
		  "node_id": "MDE3OlB1bGxSZXF1ZXN0UmV2aWV3Njk4OTI2NjU0",
		  "user": {
			"login": "AndreyMlashkin",
			"id": 3842441,
			"node_id": "MDQ6VXNlcjM4NDI0NDE=",
			"avatar_url": "https://avatars.githubusercontent.com/u/3842441?v=4",
			"gravatar_id": "",
			"url": "https://api.github.com/users/AndreyMlashkin",
			"html_url": "https://github.com/AndreyMlashkin",
			"followers_url": "https://api.github.com/users/AndreyMlashkin/followers",
			"following_url": "https://api.github.com/users/AndreyMlashkin/following{/other_user}",
			"gists_url": "https://api.github.com/users/AndreyMlashkin/gists{/gist_id}",
			"starred_url": "https://api.github.com/users/AndreyMlashkin/starred{/owner}{/repo}",
			"subscriptions_url": "https://api.github.com/users/AndreyMlashkin/subscriptions",
			"organizations_url": "https://api.github.com/users/AndreyMlashkin/orgs",
			"repos_url": "https://api.github.com/users/AndreyMlashkin/repos",
			"events_url": "https://api.github.com/users/AndreyMlashkin/events{/privacy}",
			"received_events_url": "https://api.github.com/users/AndreyMlashkin/received_events",
			"type": "User",
			"site_admin": false
		  },
		  "body": "",
		  "state": "APPROVED",
		  "html_url": "https://github.com/conan-io/conan-center-index/pull/6144#pullrequestreview-698926654",
		  "pull_request_url": "https://api.github.com/repos/conan-io/conan-center-index/pulls/6144",
		  "author_association": "CONTRIBUTOR",
		  "_links": {
			"html": {
			  "href": "https://github.com/conan-io/conan-center-index/pull/6144#pullrequestreview-698926654"
			},
			"pull_request": {
			  "href": "https://api.github.com/repos/conan-io/conan-center-index/pulls/6144"
			}
		  },
		  "submitted_at": "2021-07-05T08:30:09Z",
		  "commit_id": "3093bad9162e288d55eeddec288b0481d964518e"
		}
	  ]`)

	result := ProcessReviewComments(reviews, "3093bad9162e288d55eeddec288b0481d964518e")
	assert.Equal(t, Reviews{
		Count: 4, ValidApprovals: 3, TeamApproval: true,
		Approvals: []string{"prince-chrismc", "SSE4", "jgsogo", "AndreyMlashkin"},
		Blockers:  nil, LastReview: &Review{
			ReviewerName: reviews[len(reviews)-1].GetUser().GetLogin(),
			SubmittedAt:  reviews[len(reviews)-1].GetSubmittedAt(),
			HTMLURL:      reviews[len(reviews)-1].GetHTMLURL(),
		},
	}, result)
}
