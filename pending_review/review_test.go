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

func TestProcessReviewComments6144(t *testing.T) {
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
			"login": "AndreyMlashkin_",
			"id": 3842441,
			"node_id": "MDQ6VXNlcjM4NDI0NDE=",
			"avatar_url": "https://avatars.githubusercontent.com/u/3842441?v=4",
			"gravatar_id": "",
			"url": "https://api.github.com/users/AndreyMlashkin_",
			"html_url": "https://github.com/AndreyMlashkin_",
			"followers_url": "https://api.github.com/users/AndreyMlashkin_/followers",
			"following_url": "https://api.github.com/users/AndreyMlashkin_/following{/other_user}",
			"gists_url": "https://api.github.com/users/AndreyMlashkin_/gists{/gist_id}",
			"starred_url": "https://api.github.com/users/AndreyMlashkin_/starred{/owner}{/repo}",
			"subscriptions_url": "https://api.github.com/users/AndreyMlashkin_/subscriptions",
			"organizations_url": "https://api.github.com/users/AndreyMlashkin_/orgs",
			"repos_url": "https://api.github.com/users/AndreyMlashkin_/repos",
			"events_url": "https://api.github.com/users/AndreyMlashkin_/events{/privacy}",
			"received_events_url": "https://api.github.com/users/AndreyMlashkin_/received_events",
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
	reviewers := ConanCenterReviewers{Reviewers: []Reviewer{{User: "danimtb", Type: "team", Requested: true}, {User: "lasote", Type: "team", Requested: false}, {User: "jgsogo", Type: "team", Requested: true}, {User: "czoido", Type: "team", Requested: false}, {User: "memsharded", Type: "team", Requested: false}, {User: "SSE4", Type: "team", Requested: true}, {User: "uilianries", Type: "team", Requested: true}, {User: "madebr", Type: "community", Requested: false}, {User: "SpaceIm", Type: "community", Requested: false}, {User: "ericLemanissier", Type: "community", Requested: false}, {User: "prince-chrismc", Type: "team", Requested: false}, {User: "Croydon", Type: "community", Requested: false}, {User: "intelligide", Type: "community", Requested: false}, {User: "theirix", Type: "community", Requested: false}, {User: "gocarlos", Type: "community", Requested: false}, {User: "mathbunnyru", Type: "community", Requested: false}, {User: "ericriff", Type: "community", Requested: false}, {User: "toge", Type: "community", Requested: false}, {User: "AndreyMlashkin", Type: "community", Requested: false}, {User: "MartinDelille", Type: "community", Requested: false}, {User: "dmn-star", Type: "community", Requested: false}}}
	result := ProcessReviewComments(&reviewers, reviews, "3093bad9162e288d55eeddec288b0481d964518e")
	assert.Equal(t, Reviews{
		Count: 4, ValidApprovals: 3, TeamApproval: true,
		Approvals: []Approver{
			{Name: "prince-chrismc", Tier: Team},
			{Name: "SSE4", Tier: Team},
			{Name: "jgsogo", Tier: Team},
			{Name: "AndreyMlashkin_", Tier: Unofficial}},
		Blockers: nil, LastReview: &Review{
			ReviewerName: reviews[len(reviews)-1].GetUser().GetLogin(),
			SubmittedAt:  reviews[len(reviews)-1].GetSubmittedAt(),
			HTMLURL:      reviews[len(reviews)-1].GetHTMLURL(),
		},
	}, result)
}

func TestProcessReviewComments16144(t *testing.T) {
	reviews := parseReviewJSON(t, `[
		{
		  "id": 1311404502,
		  "node_id": "PRR_kwDODDMJAM5OKnHW",
		  "user": {
			"login": "toge",
			"id": 465629,
			"node_id": "MDQ6VXNlcjQ2NTYyOQ==",
			"avatar_url": "https://avatars.githubusercontent.com/u/465629?u=fc95d16a396044be5625091463f1f89711bdc05e&v=4",
			"gravatar_id": "",
			"url": "https://api.github.com/users/toge",
			"html_url": "https://github.com/toge",
			"followers_url": "https://api.github.com/users/toge/followers",
			"following_url": "https://api.github.com/users/toge/following{/other_user}",
			"gists_url": "https://api.github.com/users/toge/gists{/gist_id}",
			"starred_url": "https://api.github.com/users/toge/starred{/owner}{/repo}",
			"subscriptions_url": "https://api.github.com/users/toge/subscriptions",
			"organizations_url": "https://api.github.com/users/toge/orgs",
			"repos_url": "https://api.github.com/users/toge/repos",
			"events_url": "https://api.github.com/users/toge/events{/privacy}",
			"received_events_url": "https://api.github.com/users/toge/received_events",
			"type": "User",
			"site_admin": false
		  },
		  "body": "",
		  "state": "APPROVED",
		  "html_url": "https://github.com/conan-io/conan-center-index/pull/16144#pullrequestreview-1311404502",
		  "pull_request_url": "https://api.github.com/repos/conan-io/conan-center-index/pulls/16144",
		  "author_association": "CONTRIBUTOR",
		  "_links": {
			"html": {
			  "href": "https://github.com/conan-io/conan-center-index/pull/16144#pullrequestreview-1311404502"
			},
			"pull_request": {
			  "href": "https://api.github.com/repos/conan-io/conan-center-index/pulls/16144"
			}
		  },
		  "submitted_at": "2023-02-23T14:15:27Z",
		  "commit_id": "e2aa65c961d48d688dd5450811229eb1d62649ba"
		},
		{
		  "id": 1335829632,
		  "node_id": "PRR_kwDODDMJAM5PnySA",
		  "user": {
			"login": "prince-chrismc",
			"id": 16867443,
			"node_id": "MDQ6VXNlcjE2ODY3NDQz",
			"avatar_url": "https://avatars.githubusercontent.com/u/16867443?u=d7c5b45b864fe5d26e44d38645a22f8af18a2a16&v=4",
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
		  "html_url": "https://github.com/conan-io/conan-center-index/pull/16144#pullrequestreview-1335829632",
		  "pull_request_url": "https://api.github.com/repos/conan-io/conan-center-index/pulls/16144",
		  "author_association": "CONTRIBUTOR",
		  "_links": {
			"html": {
			  "href": "https://github.com/conan-io/conan-center-index/pull/16144#pullrequestreview-1335829632"
			},
			"pull_request": {
			  "href": "https://api.github.com/repos/conan-io/conan-center-index/pulls/16144"
			}
		  },
		  "submitted_at": "2023-03-11T06:46:57Z",
		  "commit_id": "e2aa65c961d48d688dd5450811229eb1d62649ba"
		}
	  ]`)
	reviewers := ConanCenterReviewers{Reviewers: []Reviewer{
		{User: "czoido", Type: "team", Requested: false},
		{User: "memsharded", Type: "team", Requested: false},
		{User: "uilianries", Type: "team", Requested: true},
		{User: "SpaceIm", Type: "community", Requested: false},
		{User: "ericLemanissier", Type: "community", Requested: false},
		{User: "prince-chrismc", Type: "team", Requested: false},
		{User: "Croydon", Type: "community", Requested: false},
		{User: "toge", Type: "community", Requested: false},
	}}
	result := ProcessReviewComments(&reviewers, reviews, "e2aa65c961d48d688dd5450811229eb1d62649ba")
	assert.Equal(t, Reviews{
		Count: 2, ValidApprovals: 2, TeamApproval: true,
		Approvals: []Approver{{Name: "toge", Tier: Community}, {Name: "prince-chrismc", Tier: Team}},
		Blockers:  nil, LastReview: &Review{
			ReviewerName: reviews[len(reviews)-1].GetUser().GetLogin(),
			SubmittedAt:  reviews[len(reviews)-1].GetSubmittedAt(),
			HTMLURL:      reviews[len(reviews)-1].GetHTMLURL(),
		},
	}, result)
}

func TestProcessReviewComments16187_1(t *testing.T) {
	reviews := parseReviewJSON(t, `[
		{
		  "id": 1323937342,
		  "node_id": "PRR_kwDODDMJAM5O6a4-",
		  "user": {
			"login": "jwillikers",
			"id": 19399197,
			"node_id": "MDQ6VXNlcjE5Mzk5MTk3",
			"avatar_url": "https://avatars.githubusercontent.com/u/19399197?u=dac900db43788310516c636d63cd9245b37b51a9&v=4",
			"gravatar_id": "",
			"url": "https://api.github.com/users/jwillikers",
			"html_url": "https://github.com/jwillikers",
			"followers_url": "https://api.github.com/users/jwillikers/followers",
			"following_url": "https://api.github.com/users/jwillikers/following{/other_user}",
			"gists_url": "https://api.github.com/users/jwillikers/gists{/gist_id}",
			"starred_url": "https://api.github.com/users/jwillikers/starred{/owner}{/repo}",
			"subscriptions_url": "https://api.github.com/users/jwillikers/subscriptions",
			"organizations_url": "https://api.github.com/users/jwillikers/orgs",
			"repos_url": "https://api.github.com/users/jwillikers/repos",
			"events_url": "https://api.github.com/users/jwillikers/events{/privacy}",
			"received_events_url": "https://api.github.com/users/jwillikers/received_events",
			"type": "User",
			"site_admin": false
		  },
		  "body": "",
		  "state": "DISMISSED",
		  "html_url": "https://github.com/conan-io/conan-center-index/pull/16187#pullrequestreview-1323937342",
		  "pull_request_url": "https://api.github.com/repos/conan-io/conan-center-index/pulls/16187",
		  "author_association": "CONTRIBUTOR",
		  "_links": {
			"html": {
			  "href": "https://github.com/conan-io/conan-center-index/pull/16187#pullrequestreview-1323937342"
			},
			"pull_request": {
			  "href": "https://api.github.com/repos/conan-io/conan-center-index/pulls/16187"
			}
		  },
		  "submitted_at": "2023-03-03T14:36:58Z",
		  "commit_id": "d242bc7a59e64a83e45f9cf84b520de97c48619f"
		}
	  ]`)
	reviewers := ConanCenterReviewers{Reviewers: []Reviewer{
		{User: "czoido", Type: "team", Requested: false},
		{User: "memsharded", Type: "team", Requested: false},
		{User: "uilianries", Type: "team", Requested: true},
		{User: "SpaceIm", Type: "community", Requested: false},
		{User: "ericLemanissier", Type: "community", Requested: false},
		{User: "prince-chrismc", Type: "team", Requested: false},
		{User: "Croydon", Type: "community", Requested: false},
		{User: "jwillikers", Type: "community", Requested: false},
	}}
	result := ProcessReviewComments(&reviewers, reviews, "d242bc7a59e64a83e45f9cf84b520de97c48619f")
	assert.Equal(t, Reviews{
		Count: 1, ValidApprovals: 0, TeamApproval: false,
		Approvals: nil,
		Blockers:  nil, LastReview: &Review{
			ReviewerName: reviews[len(reviews)-1].GetUser().GetLogin(),
			SubmittedAt:  reviews[len(reviews)-1].GetSubmittedAt(),
			HTMLURL:      reviews[len(reviews)-1].GetHTMLURL(),
		},
	}, result)
}

func TestProcessReviewComments16187_2(t *testing.T) {
	reviews := parseReviewJSON(t, `[
		{
		  "id": 1323937342,
		  "node_id": "PRR_kwDODDMJAM5O6a4-",
		  "user": {
			"login": "jwillikers",
			"id": 19399197,
			"node_id": "MDQ6VXNlcjE5Mzk5MTk3",
			"avatar_url": "https://avatars.githubusercontent.com/u/19399197?u=dac900db43788310516c636d63cd9245b37b51a9&v=4",
			"gravatar_id": "",
			"url": "https://api.github.com/users/jwillikers",
			"html_url": "https://github.com/jwillikers",
			"followers_url": "https://api.github.com/users/jwillikers/followers",
			"following_url": "https://api.github.com/users/jwillikers/following{/other_user}",
			"gists_url": "https://api.github.com/users/jwillikers/gists{/gist_id}",
			"starred_url": "https://api.github.com/users/jwillikers/starred{/owner}{/repo}",
			"subscriptions_url": "https://api.github.com/users/jwillikers/subscriptions",
			"organizations_url": "https://api.github.com/users/jwillikers/orgs",
			"repos_url": "https://api.github.com/users/jwillikers/repos",
			"events_url": "https://api.github.com/users/jwillikers/events{/privacy}",
			"received_events_url": "https://api.github.com/users/jwillikers/received_events",
			"type": "User",
			"site_admin": false
		  },
		  "body": "",
		  "state": "DISMISSED",
		  "html_url": "https://github.com/conan-io/conan-center-index/pull/16187#pullrequestreview-1323937342",
		  "pull_request_url": "https://api.github.com/repos/conan-io/conan-center-index/pulls/16187",
		  "author_association": "CONTRIBUTOR",
		  "_links": {
			"html": {
			  "href": "https://github.com/conan-io/conan-center-index/pull/16187#pullrequestreview-1323937342"
			},
			"pull_request": {
			  "href": "https://api.github.com/repos/conan-io/conan-center-index/pulls/16187"
			}
		  },
		  "submitted_at": "2023-03-03T14:36:58Z",
		  "commit_id": "d242bc7a59e64a83e45f9cf84b520de97c48619f"
		},
		{
		  "id": 1374813738,
		  "node_id": "PRR_kwDODDMJAM5R8f4q",
		  "user": {
			"login": "jwillikers",
			"id": 19399197,
			"node_id": "MDQ6VXNlcjE5Mzk5MTk3",
			"avatar_url": "https://avatars.githubusercontent.com/u/19399197?u=dac900db43788310516c636d63cd9245b37b51a9&v=4",
			"gravatar_id": "",
			"url": "https://api.github.com/users/jwillikers",
			"html_url": "https://github.com/jwillikers",
			"followers_url": "https://api.github.com/users/jwillikers/followers",
			"following_url": "https://api.github.com/users/jwillikers/following{/other_user}",
			"gists_url": "https://api.github.com/users/jwillikers/gists{/gist_id}",
			"starred_url": "https://api.github.com/users/jwillikers/starred{/owner}{/repo}",
			"subscriptions_url": "https://api.github.com/users/jwillikers/subscriptions",
			"organizations_url": "https://api.github.com/users/jwillikers/orgs",
			"repos_url": "https://api.github.com/users/jwillikers/repos",
			"events_url": "https://api.github.com/users/jwillikers/events{/privacy}",
			"received_events_url": "https://api.github.com/users/jwillikers/received_events",
			"type": "User",
			"site_admin": false
		  },
		  "body": "",
		  "state": "APPROVED",
		  "html_url": "https://github.com/conan-io/conan-center-index/pull/16187#pullrequestreview-1374813738",
		  "pull_request_url": "https://api.github.com/repos/conan-io/conan-center-index/pulls/16187",
		  "author_association": "CONTRIBUTOR",
		  "_links": {
			"html": {
			  "href": "https://github.com/conan-io/conan-center-index/pull/16187#pullrequestreview-1374813738"
			},
			"pull_request": {
			  "href": "https://api.github.com/repos/conan-io/conan-center-index/pulls/16187"
			}
		  },
		  "submitted_at": "2023-04-06T12:57:34Z",
		  "commit_id": "47efa240d39a5e74b7b02cbe15702ea85215145d"
		}
	  ]`)
	reviewers := ConanCenterReviewers{Reviewers: []Reviewer{
		{User: "czoido", Type: "team", Requested: false},
		{User: "memsharded", Type: "team", Requested: false},
		{User: "uilianries", Type: "team", Requested: true},
		{User: "SpaceIm", Type: "community", Requested: false},
		{User: "ericLemanissier", Type: "community", Requested: false},
		{User: "prince-chrismc", Type: "team", Requested: false},
		{User: "Croydon", Type: "community", Requested: false},
		{User: "jwillikers", Type: "community", Requested: false},
	}}
	result := ProcessReviewComments(&reviewers, reviews, "47efa240d39a5e74b7b02cbe15702ea85215145d")
	assert.Equal(t, Reviews{
		Count: 2, ValidApprovals: 1, TeamApproval: false,
		Approvals: []Approver{{Name: "jwillikers", Tier: Community}},
		Blockers:  nil, LastReview: &Review{
			ReviewerName: reviews[len(reviews)-1].GetUser().GetLogin(),
			SubmittedAt:  reviews[len(reviews)-1].GetSubmittedAt(),
			HTMLURL:      reviews[len(reviews)-1].GetHTMLURL(),
		},
	}, result)
}

func TestProcessReviewComments17752(t *testing.T) {
	reviews := parseReviewJSON(t, `[
		{
		  "id": 1457612113,
		  "node_id": "PRR_kwDODDMJAM5W4WVR",
		  "user": {
			"login": "RubenRBS",
			"id": 5364255,
			"node_id": "MDQ6VXNlcjUzNjQyNTU=",
			"avatar_url": "https://avatars.githubusercontent.com/u/5364255?u=1b2a8263db33c7906c65cd447567ab0639712be0&v=4",
			"gravatar_id": "",
			"url": "https://api.github.com/users/RubenRBS",
			"html_url": "https://github.com/RubenRBS",
			"followers_url": "https://api.github.com/users/RubenRBS/followers",
			"following_url": "https://api.github.com/users/RubenRBS/following{/other_user}",
			"gists_url": "https://api.github.com/users/RubenRBS/gists{/gist_id}",
			"starred_url": "https://api.github.com/users/RubenRBS/starred{/owner}{/repo}",
			"subscriptions_url": "https://api.github.com/users/RubenRBS/subscriptions",
			"organizations_url": "https://api.github.com/users/RubenRBS/orgs",
			"repos_url": "https://api.github.com/users/RubenRBS/repos",
			"events_url": "https://api.github.com/users/RubenRBS/events{/privacy}",
			"received_events_url": "https://api.github.com/users/RubenRBS/received_events",
			"type": "User",
			"site_admin": false
		  },
		  "body": "Hi! Thanks a lot for your contribution. I only have one minor question before approving :)",
		  "state": "COMMENTED",
		  "html_url": "https://github.com/conan-io/conan-center-index/pull/17752#pullrequestreview-1457612113",
		  "pull_request_url": "https://api.github.com/repos/conan-io/conan-center-index/pulls/17752",
		  "author_association": "MEMBER",
		  "_links": {
			"html": {
			  "href": "https://github.com/conan-io/conan-center-index/pull/17752#pullrequestreview-1457612113"
			},
			"pull_request": {
			  "href": "https://api.github.com/repos/conan-io/conan-center-index/pulls/17752"
			}
		  },
		  "submitted_at": "2023-06-02T14:20:16Z",
		  "commit_id": "49161ef083ed0ce6988d43f72944e2f0e513d5c0"
		},
		{
		  "id": 1460602759,
		  "node_id": "PRR_kwDODDMJAM5XDweH",
		  "user": {
			"login": "0xFireWolf",
			"id": 10460478,
			"node_id": "MDQ6VXNlcjEwNDYwNDc4",
			"avatar_url": "https://avatars.githubusercontent.com/u/10460478?v=4",
			"gravatar_id": "",
			"url": "https://api.github.com/users/0xFireWolf",
			"html_url": "https://github.com/0xFireWolf",
			"followers_url": "https://api.github.com/users/0xFireWolf/followers",
			"following_url": "https://api.github.com/users/0xFireWolf/following{/other_user}",
			"gists_url": "https://api.github.com/users/0xFireWolf/gists{/gist_id}",
			"starred_url": "https://api.github.com/users/0xFireWolf/starred{/owner}{/repo}",
			"subscriptions_url": "https://api.github.com/users/0xFireWolf/subscriptions",
			"organizations_url": "https://api.github.com/users/0xFireWolf/orgs",
			"repos_url": "https://api.github.com/users/0xFireWolf/repos",
			"events_url": "https://api.github.com/users/0xFireWolf/events{/privacy}",
			"received_events_url": "https://api.github.com/users/0xFireWolf/received_events",
			"type": "User",
			"site_admin": false
		  },
		  "body": "",
		  "state": "COMMENTED",
		  "html_url": "https://github.com/conan-io/conan-center-index/pull/17752#pullrequestreview-1460602759",
		  "pull_request_url": "https://api.github.com/repos/conan-io/conan-center-index/pulls/17752",
		  "author_association": "CONTRIBUTOR",
		  "_links": {
			"html": {
			  "href": "https://github.com/conan-io/conan-center-index/pull/17752#pullrequestreview-1460602759"
			},
			"pull_request": {
			  "href": "https://api.github.com/repos/conan-io/conan-center-index/pulls/17752"
			}
		  },
		  "submitted_at": "2023-06-04T01:03:08Z",
		  "commit_id": "49161ef083ed0ce6988d43f72944e2f0e513d5c0"
		},
		{
		  "id": 1485955540,
		  "node_id": "PRR_kwDODDMJAM5YkeHU",
		  "user": {
			"login": "ericLemanissier",
			"id": 1926390,
			"node_id": "MDQ6VXNlcjE5MjYzOTA=",
			"avatar_url": "https://avatars.githubusercontent.com/u/1926390?v=4",
			"gravatar_id": "",
			"url": "https://api.github.com/users/ericLemanissier",
			"html_url": "https://github.com/ericLemanissier",
			"followers_url": "https://api.github.com/users/ericLemanissier/followers",
			"following_url": "https://api.github.com/users/ericLemanissier/following{/other_user}",
			"gists_url": "https://api.github.com/users/ericLemanissier/gists{/gist_id}",
			"starred_url": "https://api.github.com/users/ericLemanissier/starred{/owner}{/repo}",
			"subscriptions_url": "https://api.github.com/users/ericLemanissier/subscriptions",
			"organizations_url": "https://api.github.com/users/ericLemanissier/orgs",
			"repos_url": "https://api.github.com/users/ericLemanissier/repos",
			"events_url": "https://api.github.com/users/ericLemanissier/events{/privacy}",
			"received_events_url": "https://api.github.com/users/ericLemanissier/received_events",
			"type": "User",
			"site_admin": false
		  },
		  "body": "",
		  "state": "APPROVED",
		  "html_url": "https://github.com/conan-io/conan-center-index/pull/17752#pullrequestreview-1485955540",
		  "pull_request_url": "https://api.github.com/repos/conan-io/conan-center-index/pulls/17752",
		  "author_association": "CONTRIBUTOR",
		  "_links": {
			"html": {
			  "href": "https://github.com/conan-io/conan-center-index/pull/17752#pullrequestreview-1485955540"
			},
			"pull_request": {
			  "href": "https://api.github.com/repos/conan-io/conan-center-index/pulls/17752"
			}
		  },
		  "submitted_at": "2023-06-19T10:54:02Z",
		  "commit_id": "e5d87da51c7d00bad28b958c29383d709b5f0592"
		},
		{
		  "id": 1499553254,
		  "node_id": "PRR_kwDODDMJAM5ZYV3m",
		  "user": {
			"login": "prince-chrismc",
			"id": 16867443,
			"node_id": "MDQ6VXNlcjE2ODY3NDQz",
			"avatar_url": "https://avatars.githubusercontent.com/u/16867443?u=d7c5b45b864fe5d26e44d38645a22f8af18a2a16&v=4",
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
		  "body": "Dont mind me just trying to test some code 🧑‍💻 ",
		  "state": "DISMISSED",
		  "html_url": "https://github.com/conan-io/conan-center-index/pull/17752#pullrequestreview-1499553254",
		  "pull_request_url": "https://api.github.com/repos/conan-io/conan-center-index/pulls/17752",
		  "author_association": "CONTRIBUTOR",
		  "_links": {
			"html": {
			  "href": "https://github.com/conan-io/conan-center-index/pull/17752#pullrequestreview-1499553254"
			},
			"pull_request": {
			  "href": "https://api.github.com/repos/conan-io/conan-center-index/pulls/17752"
			}
		  },
		  "submitted_at": "2023-06-26T21:55:37Z",
		  "commit_id": "e5d87da51c7d00bad28b958c29383d709b5f0592"
		},
		{
		  "id": 1499560222,
		  "node_id": "PRR_kwDODDMJAM5ZYXke",
		  "user": {
			"login": "prince-chrismc",
			"id": 16867443,
			"node_id": "MDQ6VXNlcjE2ODY3NDQz",
			"avatar_url": "https://avatars.githubusercontent.com/u/16867443?u=d7c5b45b864fe5d26e44d38645a22f8af18a2a16&v=4",
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
		  "html_url": "https://github.com/conan-io/conan-center-index/pull/17752#pullrequestreview-1499560222",
		  "pull_request_url": "https://api.github.com/repos/conan-io/conan-center-index/pulls/17752",
		  "author_association": "CONTRIBUTOR",
		  "_links": {
			"html": {
			  "href": "https://github.com/conan-io/conan-center-index/pull/17752#pullrequestreview-1499560222"
			},
			"pull_request": {
			  "href": "https://api.github.com/repos/conan-io/conan-center-index/pulls/17752"
			}
		  },
		  "submitted_at": "2023-06-26T22:00:02Z",
		  "commit_id": "fce3c24a57563620b8915a307171f9b2b8b2686e"
		}
	  ]`)
	reviewers := ConanCenterReviewers{Reviewers: []Reviewer{
		{User: "RubenRBS", Type: Team, Requested: false},
		{User: "ericLemanissier", Type: Community, Requested: false},
		{User: "prince-chrismc", Type: Team, Requested: false},
	}}
	result := ProcessReviewComments(&reviewers, reviews, "fce3c24a57563620b8915a307171f9b2b8b2686e")
	assert.Equal(t, Reviews{
		Count: 5, ValidApprovals: 1, TeamApproval: true,
		Approvals: []Approver{{Name: "prince-chrismc", Tier: Team}},
		Blockers:  nil, LastReview: &Review{
			ReviewerName: reviews[len(reviews)-1].GetUser().GetLogin(),
			SubmittedAt:  reviews[len(reviews)-1].GetSubmittedAt(),
			HTMLURL:      reviews[len(reviews)-1].GetHTMLURL(),
		},
	}, result)
}

func TestIsApproved_Community(t *testing.T) {
	var reviews = Reviews{
		Count: 1, ValidApprovals: 1, TeamApproval: false, Approvals: []Approver{{Name: "foo", Tier: Community}},
	}
	reviews.IsBump = true
	assert.Equal(t, false, reviews.IsApproved())

	reviews.IsBump = false
	assert.Equal(t, false, reviews.IsApproved())
}

func TestIsApproved_Team(t *testing.T) {
	var reviews = Reviews{
		Count: 1, ValidApprovals: 1, TeamApproval: true, Approvals: []Approver{{Name: "foo", Tier: Team}},
	}
	reviews.IsBump = true
	assert.Equal(t, true, reviews.IsApproved())

	reviews.IsBump = false
	assert.Equal(t, false, reviews.IsApproved())
}
