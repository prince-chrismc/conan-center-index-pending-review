package pending_review

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	gock "gopkg.in/h2non/gock.v1"
)

func TestSomething(t *testing.T) {

	// assert equality
	assert.Equal(t, 123, 123, "they should be equal")

	defer gock.Off()

	gock.New("https://api.github.com").
		Get("/repos/conan-io/conan-center-index").
		Reply(200).
		BodyString(`{
			"id": 204671232,
			"node_id": "MDEwOlJlcG9zaXRvcnkyMDQ2NzEyMzI=",
			"name": "conan-center-index",
			"full_name": "conan-io/conan-center-index",
			"private": false,
			"owner": {
			  "login": "conan-io",
			  "id": 15212165,
			  "node_id": "MDEyOk9yZ2FuaXphdGlvbjE1MjEyMTY1",
			  "avatar_url": "https://avatars.githubusercontent.com/u/15212165?v=4",
			  "gravatar_id": "",
			  "url": "https://api.github.com/users/conan-io",
			  "html_url": "https://github.com/conan-io",
			  "followers_url": "https://api.github.com/users/conan-io/followers",
			  "following_url": "https://api.github.com/users/conan-io/following{/other_user}",
			  "gists_url": "https://api.github.com/users/conan-io/gists{/gist_id}",
			  "starred_url": "https://api.github.com/users/conan-io/starred{/owner}{/repo}",
			  "subscriptions_url": "https://api.github.com/users/conan-io/subscriptions",
			  "organizations_url": "https://api.github.com/users/conan-io/orgs",
			  "repos_url": "https://api.github.com/users/conan-io/repos",
			  "events_url": "https://api.github.com/users/conan-io/events{/privacy}",
			  "received_events_url": "https://api.github.com/users/conan-io/received_events",
			  "type": "Organization",
			  "site_admin": false
			},
			"html_url": "https://github.com/conan-io/conan-center-index",
			"description": "Recipes for the ConanCenter repository",
			"fork": false,
			"url": "https://api.github.com/repos/conan-io/conan-center-index",
			"forks_url": "https://api.github.com/repos/conan-io/conan-center-index/forks",
			"keys_url": "https://api.github.com/repos/conan-io/conan-center-index/keys{/key_id}",
			"collaborators_url": "https://api.github.com/repos/conan-io/conan-center-index/collaborators{/collaborator}",
			"teams_url": "https://api.github.com/repos/conan-io/conan-center-index/teams",
			"hooks_url": "https://api.github.com/repos/conan-io/conan-center-index/hooks",
			"issue_events_url": "https://api.github.com/repos/conan-io/conan-center-index/issues/events{/number}",
			"events_url": "https://api.github.com/repos/conan-io/conan-center-index/events",
			"assignees_url": "https://api.github.com/repos/conan-io/conan-center-index/assignees{/user}",
			"branches_url": "https://api.github.com/repos/conan-io/conan-center-index/branches{/branch}",
			"tags_url": "https://api.github.com/repos/conan-io/conan-center-index/tags",
			"blobs_url": "https://api.github.com/repos/conan-io/conan-center-index/git/blobs{/sha}",
			"git_tags_url": "https://api.github.com/repos/conan-io/conan-center-index/git/tags{/sha}",
			"git_refs_url": "https://api.github.com/repos/conan-io/conan-center-index/git/refs{/sha}",
			"trees_url": "https://api.github.com/repos/conan-io/conan-center-index/git/trees{/sha}",
			"statuses_url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/{sha}",
			"languages_url": "https://api.github.com/repos/conan-io/conan-center-index/languages",
			"stargazers_url": "https://api.github.com/repos/conan-io/conan-center-index/stargazers",
			"contributors_url": "https://api.github.com/repos/conan-io/conan-center-index/contributors",
			"subscribers_url": "https://api.github.com/repos/conan-io/conan-center-index/subscribers",
			"subscription_url": "https://api.github.com/repos/conan-io/conan-center-index/subscription",
			"commits_url": "https://api.github.com/repos/conan-io/conan-center-index/commits{/sha}",
			"git_commits_url": "https://api.github.com/repos/conan-io/conan-center-index/git/commits{/sha}",
			"comments_url": "https://api.github.com/repos/conan-io/conan-center-index/comments{/number}",
			"issue_comment_url": "https://api.github.com/repos/conan-io/conan-center-index/issues/comments{/number}",
			"contents_url": "https://api.github.com/repos/conan-io/conan-center-index/contents/{+path}",
			"compare_url": "https://api.github.com/repos/conan-io/conan-center-index/compare/{base}...{head}",
			"merges_url": "https://api.github.com/repos/conan-io/conan-center-index/merges",
			"archive_url": "https://api.github.com/repos/conan-io/conan-center-index/{archive_format}{/ref}",
			"downloads_url": "https://api.github.com/repos/conan-io/conan-center-index/downloads",
			"issues_url": "https://api.github.com/repos/conan-io/conan-center-index/issues{/number}",
			"pulls_url": "https://api.github.com/repos/conan-io/conan-center-index/pulls{/number}",
			"milestones_url": "https://api.github.com/repos/conan-io/conan-center-index/milestones{/number}",
			"notifications_url": "https://api.github.com/repos/conan-io/conan-center-index/notifications{?since,all,participating}",
			"labels_url": "https://api.github.com/repos/conan-io/conan-center-index/labels{/name}",
			"releases_url": "https://api.github.com/repos/conan-io/conan-center-index/releases{/id}",
			"deployments_url": "https://api.github.com/repos/conan-io/conan-center-index/deployments",
			"created_at": "2019-08-27T09:43:58Z",
			"updated_at": "2021-03-10T22:01:19Z",
			"pushed_at": "2021-03-11T00:16:37Z",
			"git_url": "git://github.com/conan-io/conan-center-index.git",
			"ssh_url": "git@github.com:conan-io/conan-center-index.git",
			"clone_url": "https://github.com/conan-io/conan-center-index.git",
			"svn_url": "https://github.com/conan-io/conan-center-index",
			"homepage": "https://conan.io/center",
			"size": 20649,
			"stargazers_count": 304,
			"watchers_count": 304,
			"language": "Python",
			"has_issues": true,
			"has_projects": true,
			"has_downloads": true,
			"has_wiki": true,
			"has_pages": false,
			"forks_count": 440,
			"mirror_url": null,
			"archived": false,
			"disabled": false,
			"open_issues_count": 549,
			"license": {
			  "key": "mit",
			  "name": "MIT License",
			  "spdx_id": "MIT",
			  "url": "https://api.github.com/licenses/mit",
			  "node_id": "MDc6TGljZW5zZTEz"
			},
			"forks": 440,
			"open_issues": 549,
			"watchers": 304,
			"default_branch": "master",
			"temp_clone_token": null,
			"organization": {
			  "login": "conan-io",
			  "id": 15212165,
			  "node_id": "MDEyOk9yZ2FuaXphdGlvbjE1MjEyMTY1",
			  "avatar_url": "https://avatars.githubusercontent.com/u/15212165?v=4",
			  "gravatar_id": "",
			  "url": "https://api.github.com/users/conan-io",
			  "html_url": "https://github.com/conan-io",
			  "followers_url": "https://api.github.com/users/conan-io/followers",
			  "following_url": "https://api.github.com/users/conan-io/following{/other_user}",
			  "gists_url": "https://api.github.com/users/conan-io/gists{/gist_id}",
			  "starred_url": "https://api.github.com/users/conan-io/starred{/owner}{/repo}",
			  "subscriptions_url": "https://api.github.com/users/conan-io/subscriptions",
			  "organizations_url": "https://api.github.com/users/conan-io/orgs",
			  "repos_url": "https://api.github.com/users/conan-io/repos",
			  "events_url": "https://api.github.com/users/conan-io/events{/privacy}",
			  "received_events_url": "https://api.github.com/users/conan-io/received_events",
			  "type": "Organization",
			  "site_admin": false
			},
			"network_count": 440,
			"subscribers_count": 15
		  }`)

	repository, res, err := NewClient(&http.Client{}).Repository.GetSummary(context.Background(), "conan-io", "conan-center-index")

	assert.Equal(t, err, nil)
	assert.Equal(t, res.StatusCode, 200)

	assert.Equal(t, repository, &RepositorySumarry{Name: "conan-center-index", Owner: "conan-io", FullName: "conan-io/conan-center-index", Description: "Recipes for the ConanCenter repository", StarsCount: 304, ForksCount: 440, OpenIssuesCount: 549})

	assert.Equal(t, gock.IsDone(), true)
}
