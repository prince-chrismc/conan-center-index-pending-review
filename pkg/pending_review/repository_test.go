package pending_review

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	gock "gopkg.in/h2non/gock.v1"
)

func TestGetRepositorySummary(t *testing.T) {
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

	assert.Equal(t, nil, err)
	assert.Equal(t, 200, res.StatusCode)

	assert.Equal(t, &RepositorySumarry{Name: "conan-center-index", Owner: "conan-io", FullName: "conan-io/conan-center-index", Description: "Recipes for the ConanCenter repository", StarsCount: 304, ForksCount: 440, OpenIssuesCount: 549}, repository)

	assert.Equal(t, true, gock.IsDone())
}

func TestGetStatus(t *testing.T) {
	defer gock.Off()

	gock.New("https://api.github.com").
		Get("/commits/08f356aabf77ff55d96ae43de3e3bfdfb67f6018/status").
		Reply(200).
		BodyString(`{
			"state": "pending",
			"statuses": [
			  {
				"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/08f356aabf77ff55d96ae43de3e3bfdfb67f6018",
				"avatar_url": "https://avatars.githubusercontent.com/oa/160438?v=4",
				"id": 14713069337,
				"node_id": "SC_kwDODDMJAM8AAAADbPefGQ",
				"state": "success",
				"description": "Contributor License Agreement is signed.",
				"target_url": "https://cla-assistant.io/conan-io/conan-center-index?pullRequest=7543",
				"context": "license/cla",
				"created_at": "2021-10-05T14:58:24Z",
				"updated_at": "2021-10-05T14:58:24Z"
			  },
			  {
				"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/08f356aabf77ff55d96ae43de3e3bfdfb67f6018",
				"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
				"id": 14713079872,
				"node_id": "SC_kwDODDMJAM8AAAADbPfIQA",
				"state": "pending",
				"description": "This commit is being built",
				"target_url": "https://ci-conan-prod.jfrog.team/job/cci/job/PR-7543/2/display/redirect",
				"context": "continuous-integration/jenkins/pr-merge",
				"created_at": "2021-10-05T14:58:57Z",
				"updated_at": "2021-10-05T14:58:57Z"
			  },
			  {
				"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/08f356aabf77ff55d96ae43de3e3bfdfb67f6018",
				"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
				"id": 14713093821,
				"node_id": "SC_kwDODDMJAM8AAAADbPf-vQ",
				"state": "pending",
				"description": "running...",
				"target_url": "https://c3i.jfrog.io/c3i/misc/summary.html?json=https://c3i.jfrog.io/c3i/misc/logs/pr/7543/2-configs/windows-visual_studio/aws-c-http/0.6.7//summary.json",
				"context": "[required] aws-c-http/0.6.7@ Windows, Visual Studio",
				"created_at": "2021-10-05T14:59:43Z",
				"updated_at": "2021-10-05T14:59:43Z"
			  },
			  {
				"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/08f356aabf77ff55d96ae43de3e3bfdfb67f6018",
				"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
				"id": 14713239118,
				"node_id": "SC_kwDODDMJAM8AAAADbPo2Tg",
				"state": "success",
				"description": "All green! (24)",
				"target_url": "https://c3i.jfrog.io/c3i/misc/summary.html?json=https://c3i.jfrog.io/c3i/misc/logs/pr/7543/2-configs/linux-gcc/aws-c-http/0.6.5//summary.json",
				"context": "[required] aws-c-http/0.6.5@ Linux, GCC",
				"created_at": "2021-10-05T15:07:32Z",
				"updated_at": "2021-10-05T15:07:32Z"
			  },
			  {
				"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/08f356aabf77ff55d96ae43de3e3bfdfb67f6018",
				"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
				"id": 14713253044,
				"node_id": "SC_kwDODDMJAM8AAAADbPpstA",
				"state": "success",
				"description": "All green! (24)",
				"target_url": "https://c3i.jfrog.io/c3i/misc/summary.html?json=https://c3i.jfrog.io/c3i/misc/logs/pr/7543/2-configs/linux-gcc/aws-c-http/0.6.7//summary.json",
				"context": "[required] aws-c-http/0.6.7@ Linux, GCC",
				"created_at": "2021-10-05T15:08:22Z",
				"updated_at": "2021-10-05T15:08:22Z"
			  },
			  {
				"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/08f356aabf77ff55d96ae43de3e3bfdfb67f6018",
				"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
				"id": 14713288509,
				"node_id": "SC_kwDODDMJAM8AAAADbPr3PQ",
				"state": "success",
				"description": "All green! (8)",
				"target_url": "https://c3i.jfrog.io/c3i/misc/summary.html?json=https://c3i.jfrog.io/c3i/misc/logs/pr/7543/2-configs/linux-clang/aws-c-http/0.6.5//summary.json",
				"context": "[required] aws-c-http/0.6.5@ Linux, Clang",
				"created_at": "2021-10-05T15:10:22Z",
				"updated_at": "2021-10-05T15:10:22Z"
			  },
			  {
				"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/08f356aabf77ff55d96ae43de3e3bfdfb67f6018",
				"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
				"id": 14713311226,
				"node_id": "SC_kwDODDMJAM8AAAADbPtP-g",
				"state": "success",
				"description": "All green! (8)",
				"target_url": "https://c3i.jfrog.io/c3i/misc/summary.html?json=https://c3i.jfrog.io/c3i/misc/logs/pr/7543/2-configs/linux-clang/aws-c-http/0.6.7//summary.json",
				"context": "[required] aws-c-http/0.6.7@ Linux, Clang",
				"created_at": "2021-10-05T15:11:36Z",
				"updated_at": "2021-10-05T15:11:36Z"
			  },
			  {
				"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/08f356aabf77ff55d96ae43de3e3bfdfb67f6018",
				"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
				"id": 14713712378,
				"node_id": "SC_kwDODDMJAM8AAAADbQFu-g",
				"state": "success",
				"description": "All green! (8)",
				"target_url": "https://c3i.jfrog.io/c3i/misc/summary.html?json=https://c3i.jfrog.io/c3i/misc/logs/pr/7543/2-configs/macos-clang/aws-c-http/0.6.5//summary.json",
				"context": "[required] aws-c-http/0.6.5@ macOS, Clang",
				"created_at": "2021-10-05T15:33:54Z",
				"updated_at": "2021-10-05T15:33:54Z"
			  },
			  {
				"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/08f356aabf77ff55d96ae43de3e3bfdfb67f6018",
				"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
				"id": 14713735124,
				"node_id": "SC_kwDODDMJAM8AAAADbQHH1A",
				"state": "success",
				"description": "All green! (8)",
				"target_url": "https://c3i.jfrog.io/c3i/misc/summary.html?json=https://c3i.jfrog.io/c3i/misc/logs/pr/7543/2-configs/macos-clang/aws-c-http/0.6.7//summary.json",
				"context": "[required] aws-c-http/0.6.7@ macOS, Clang",
				"created_at": "2021-10-05T15:35:13Z",
				"updated_at": "2021-10-05T15:35:13Z"
			  },
			  {
				"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/08f356aabf77ff55d96ae43de3e3bfdfb67f6018",
				"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
				"id": 14713761894,
				"node_id": "SC_kwDODDMJAM8AAAADbQIwZg",
				"state": "success",
				"description": "All green! (4)",
				"target_url": "https://c3i.jfrog.io/c3i/misc/summary.html?json=https://c3i.jfrog.io/c3i/misc/logs/pr/7543/2-configs/macos-m1-clang/aws-c-http/0.6.5//summary.json",
				"context": "[required] aws-c-http/0.6.5@ macOS, Clang (M1/arm64)",
				"created_at": "2021-10-05T15:36:44Z",
				"updated_at": "2021-10-05T15:36:44Z"
			  },
			  {
				"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/08f356aabf77ff55d96ae43de3e3bfdfb67f6018",
				"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
				"id": 14713793899,
				"node_id": "SC_kwDODDMJAM8AAAADbQKtaw",
				"state": "success",
				"description": "All green! (4)",
				"target_url": "https://c3i.jfrog.io/c3i/misc/summary.html?json=https://c3i.jfrog.io/c3i/misc/logs/pr/7543/2-configs/macos-m1-clang/aws-c-http/0.6.7//summary.json",
				"context": "[required] aws-c-http/0.6.7@ macOS, Clang (M1/arm64)",
				"created_at": "2021-10-05T15:38:35Z",
				"updated_at": "2021-10-05T15:38:35Z"
			  },
			  {
				"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/08f356aabf77ff55d96ae43de3e3bfdfb67f6018",
				"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
				"id": 14714044849,
				"node_id": "SC_kwDODDMJAM8AAAADbQaBsQ",
				"state": "success",
				"description": "All green! (16)",
				"target_url": "https://c3i.jfrog.io/c3i/misc/summary.html?json=https://c3i.jfrog.io/c3i/misc/logs/pr/7543/2-configs/windows-visual_studio/aws-c-http/0.6.5//summary.json",
				"context": "[required] aws-c-http/0.6.5@ Windows, Visual Studio",
				"created_at": "2021-10-05T15:52:47Z",
				"updated_at": "2021-10-05T15:52:47Z"
			  }
			],
			"sha": "08f356aabf77ff55d96ae43de3e3bfdfb67f6018",
			"total_count": 12,
			"repository": {
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
			  "deployments_url": "https://api.github.com/repos/conan-io/conan-center-index/deployments"
			},
			"commit_url": "https://api.github.com/repos/conan-io/conan-center-index/commits/08f356aabf77ff55d96ae43de3e3bfdfb67f6018",
			"url": "https://api.github.com/repos/conan-io/conan-center-index/commits/08f356aabf77ff55d96ae43de3e3bfdfb67f6018/status"
		  }`)

	status, res, err := NewClient(&http.Client{}).Repository.GetStatus(context.Background(), "conan-io", "conan-center-index", "08f356aabf77ff55d96ae43de3e3bfdfb67f6018")

	assert.Equal(t, nil, err)
	assert.Equal(t, 200, res.StatusCode)

	assert.Equal(t, "pending", status.GetState())
	assert.Equal(t, 12, status.GetTotalCount())

	assert.Equal(t, true, gock.IsDone())
}

func TestGetStatusTwo(t *testing.T) {
	defer gock.Off()

	gock.New("https://api.github.com").
		Get("/commits/fa60b0dec5ce1ef3f902b2b110321d05f7730a4d/status").
		Reply(200).
		BodyString(`{
			"state": "pending",
			"statuses": [
			  {
				"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/fa60b0dec5ce1ef3f902b2b110321d05f7730a4d",
				"avatar_url": "https://avatars.githubusercontent.com/oa/160438?v=4",
				"id": 14714563456,
				"node_id": "SC_kwDODDMJAM8AAAADbQ5rgA",
				"state": "success",
				"description": "Contributor License Agreement is signed.",
				"target_url": "https://cla-assistant.io/conan-io/conan-center-index?pullRequest=7547",
				"context": "license/cla",
				"created_at": "2021-10-05T16:21:58Z",
				"updated_at": "2021-10-05T16:21:58Z"
			  },
			  {
				"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/fa60b0dec5ce1ef3f902b2b110321d05f7730a4d",
				"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
				"id": 14714573353,
				"node_id": "SC_kwDODDMJAM8AAAADbQ6SKQ",
				"state": "pending",
				"description": "This commit is being built",
				"target_url": "https://ci-conan-prod.jfrog.team/job/cci/job/PR-7547/1/display/redirect",
				"context": "continuous-integration/jenkins/pr-merge",
				"created_at": "2021-10-05T16:22:31Z",
				"updated_at": "2021-10-05T16:22:31Z"
			  },
			  {
				"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/fa60b0dec5ce1ef3f902b2b110321d05f7730a4d",
				"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
				"id": 14714586395,
				"node_id": "SC_kwDODDMJAM8AAAADbQ7FGw",
				"state": "pending",
				"description": "running...",
				"target_url": "https://c3i.jfrog.io/c3i/misc/summary.html?json=https://c3i.jfrog.io/c3i/misc/logs/pr/7547/1-configs/macos-m1-clang/sqlite3/3.30.1//summary.json",
				"context": "[required] sqlite3/3.30.1@ macOS, Clang (M1/arm64)",
				"created_at": "2021-10-05T16:23:18Z",
				"updated_at": "2021-10-05T16:23:18Z"
			  },
			  {
				"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/fa60b0dec5ce1ef3f902b2b110321d05f7730a4d",
				"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
				"id": 14714586661,
				"node_id": "SC_kwDODDMJAM8AAAADbQ7GJQ",
				"state": "pending",
				"description": "running...",
				"target_url": "https://c3i.jfrog.io/c3i/misc/summary.html?json=https://c3i.jfrog.io/c3i/misc/logs/pr/7547/1-configs/windows-visual_studio/sqlite3/3.30.1//summary.json",
				"context": "[required] sqlite3/3.30.1@ Windows, Visual Studio",
				"created_at": "2021-10-05T16:23:19Z",
				"updated_at": "2021-10-05T16:23:19Z"
			  },
			  {
				"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/fa60b0dec5ce1ef3f902b2b110321d05f7730a4d",
				"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
				"id": 14714587309,
				"node_id": "SC_kwDODDMJAM8AAAADbQ7IrQ",
				"state": "pending",
				"description": "running...",
				"target_url": "https://c3i.jfrog.io/c3i/misc/summary.html?json=https://c3i.jfrog.io/c3i/misc/logs/pr/7547/1-configs/macos-clang/sqlite3/3.29.0//summary.json",
				"context": "[required] sqlite3/3.29.0@ macOS, Clang",
				"created_at": "2021-10-05T16:23:21Z",
				"updated_at": "2021-10-05T16:23:21Z"
			  },
			  {
				"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/fa60b0dec5ce1ef3f902b2b110321d05f7730a4d",
				"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
				"id": 14714587582,
				"node_id": "SC_kwDODDMJAM8AAAADbQ7Jvg",
				"state": "pending",
				"description": "running...",
				"target_url": "https://c3i.jfrog.io/c3i/misc/summary.html?json=https://c3i.jfrog.io/c3i/misc/logs/pr/7547/1-configs/macos-m1-clang/sqlite3/3.29.0//summary.json",
				"context": "[required] sqlite3/3.29.0@ macOS, Clang (M1/arm64)",
				"created_at": "2021-10-05T16:23:22Z",
				"updated_at": "2021-10-05T16:23:22Z"
			  },
			  {
				"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/fa60b0dec5ce1ef3f902b2b110321d05f7730a4d",
				"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
				"id": 14714587876,
				"node_id": "SC_kwDODDMJAM8AAAADbQ7K5A",
				"state": "pending",
				"description": "running...",
				"target_url": "https://c3i.jfrog.io/c3i/misc/summary.html?json=https://c3i.jfrog.io/c3i/misc/logs/pr/7547/1-configs/windows-visual_studio/sqlite3/3.29.0//summary.json",
				"context": "[required] sqlite3/3.29.0@ Windows, Visual Studio",
				"created_at": "2021-10-05T16:23:24Z",
				"updated_at": "2021-10-05T16:23:24Z"
			  },
			  {
				"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/fa60b0dec5ce1ef3f902b2b110321d05f7730a4d",
				"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
				"id": 14714588496,
				"node_id": "SC_kwDODDMJAM8AAAADbQ7NUA",
				"state": "pending",
				"description": "running...",
				"target_url": "https://c3i.jfrog.io/c3i/misc/summary.html?json=https://c3i.jfrog.io/c3i/misc/logs/pr/7547/1-configs/macos-clang/sqlite3/3.31.0//summary.json",
				"context": "[required] sqlite3/3.31.0@ macOS, Clang",
				"created_at": "2021-10-05T16:23:26Z",
				"updated_at": "2021-10-05T16:23:26Z"
			  },
			  {
				"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/fa60b0dec5ce1ef3f902b2b110321d05f7730a4d",
				"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
				"id": 14714588710,
				"node_id": "SC_kwDODDMJAM8AAAADbQ7OJg",
				"state": "pending",
				"description": "running...",
				"target_url": "https://c3i.jfrog.io/c3i/misc/summary.html?json=https://c3i.jfrog.io/c3i/misc/logs/pr/7547/1-configs/macos-m1-clang/sqlite3/3.31.0//summary.json",
				"context": "[required] sqlite3/3.31.0@ macOS, Clang (M1/arm64)",
				"created_at": "2021-10-05T16:23:27Z",
				"updated_at": "2021-10-05T16:23:27Z"
			  },
			  {
				"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/fa60b0dec5ce1ef3f902b2b110321d05f7730a4d",
				"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
				"id": 14714588953,
				"node_id": "SC_kwDODDMJAM8AAAADbQ7PGQ",
				"state": "pending",
				"description": "running...",
				"target_url": "https://c3i.jfrog.io/c3i/misc/summary.html?json=https://c3i.jfrog.io/c3i/misc/logs/pr/7547/1-configs/windows-visual_studio/sqlite3/3.31.0//summary.json",
				"context": "[required] sqlite3/3.31.0@ Windows, Visual Studio",
				"created_at": "2021-10-05T16:23:28Z",
				"updated_at": "2021-10-05T16:23:28Z"
			  },
			  {
				"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/fa60b0dec5ce1ef3f902b2b110321d05f7730a4d",
				"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
				"id": 14714589665,
				"node_id": "SC_kwDODDMJAM8AAAADbQ7R4Q",
				"state": "pending",
				"description": "running...",
				"target_url": "https://c3i.jfrog.io/c3i/misc/summary.html?json=https://c3i.jfrog.io/c3i/misc/logs/pr/7547/1-configs/macos-clang/sqlite3/3.31.1//summary.json",
				"context": "[required] sqlite3/3.31.1@ macOS, Clang",
				"created_at": "2021-10-05T16:23:30Z",
				"updated_at": "2021-10-05T16:23:30Z"
			  },
			  {
				"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/fa60b0dec5ce1ef3f902b2b110321d05f7730a4d",
				"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
				"id": 14714589924,
				"node_id": "SC_kwDODDMJAM8AAAADbQ7S5A",
				"state": "pending",
				"description": "running...",
				"target_url": "https://c3i.jfrog.io/c3i/misc/summary.html?json=https://c3i.jfrog.io/c3i/misc/logs/pr/7547/1-configs/macos-m1-clang/sqlite3/3.31.1//summary.json",
				"context": "[required] sqlite3/3.31.1@ macOS, Clang (M1/arm64)",
				"created_at": "2021-10-05T16:23:31Z",
				"updated_at": "2021-10-05T16:23:31Z"
			  },
			  {
				"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/fa60b0dec5ce1ef3f902b2b110321d05f7730a4d",
				"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
				"id": 14714590152,
				"node_id": "SC_kwDODDMJAM8AAAADbQ7TyA",
				"state": "pending",
				"description": "running...",
				"target_url": "https://c3i.jfrog.io/c3i/misc/summary.html?json=https://c3i.jfrog.io/c3i/misc/logs/pr/7547/1-configs/windows-visual_studio/sqlite3/3.31.1//summary.json",
				"context": "[required] sqlite3/3.31.1@ Windows, Visual Studio",
				"created_at": "2021-10-05T16:23:32Z",
				"updated_at": "2021-10-05T16:23:32Z"
			  },
			  {
				"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/fa60b0dec5ce1ef3f902b2b110321d05f7730a4d",
				"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
				"id": 14714590429,
				"node_id": "SC_kwDODDMJAM8AAAADbQ7U3Q",
				"state": "pending",
				"description": "running...",
				"target_url": "https://c3i.jfrog.io/c3i/misc/summary.html?json=https://c3i.jfrog.io/c3i/misc/logs/pr/7547/1-configs/linux-gcc/sqlite3/3.32.2//summary.json",
				"context": "[required] sqlite3/3.32.2@ Linux, GCC",
				"created_at": "2021-10-05T16:23:33Z",
				"updated_at": "2021-10-05T16:23:33Z"
			  },
			  {
				"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/fa60b0dec5ce1ef3f902b2b110321d05f7730a4d",
				"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
				"id": 14714590665,
				"node_id": "SC_kwDODDMJAM8AAAADbQ7VyQ",
				"state": "pending",
				"description": "running...",
				"target_url": "https://c3i.jfrog.io/c3i/misc/summary.html?json=https://c3i.jfrog.io/c3i/misc/logs/pr/7547/1-configs/linux-clang/sqlite3/3.32.2//summary.json",
				"context": "[required] sqlite3/3.32.2@ Linux, Clang",
				"created_at": "2021-10-05T16:23:33Z",
				"updated_at": "2021-10-05T16:23:33Z"
			  },
			  {
				"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/fa60b0dec5ce1ef3f902b2b110321d05f7730a4d",
				"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
				"id": 14714590938,
				"node_id": "SC_kwDODDMJAM8AAAADbQ7W2g",
				"state": "pending",
				"description": "running...",
				"target_url": "https://c3i.jfrog.io/c3i/misc/summary.html?json=https://c3i.jfrog.io/c3i/misc/logs/pr/7547/1-configs/macos-clang/sqlite3/3.32.2//summary.json",
				"context": "[required] sqlite3/3.32.2@ macOS, Clang",
				"created_at": "2021-10-05T16:23:34Z",
				"updated_at": "2021-10-05T16:23:34Z"
			  },
			  {
				"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/fa60b0dec5ce1ef3f902b2b110321d05f7730a4d",
				"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
				"id": 14714591171,
				"node_id": "SC_kwDODDMJAM8AAAADbQ7Xww",
				"state": "pending",
				"description": "running...",
				"target_url": "https://c3i.jfrog.io/c3i/misc/summary.html?json=https://c3i.jfrog.io/c3i/misc/logs/pr/7547/1-configs/macos-m1-clang/sqlite3/3.32.2//summary.json",
				"context": "[required] sqlite3/3.32.2@ macOS, Clang (M1/arm64)",
				"created_at": "2021-10-05T16:23:35Z",
				"updated_at": "2021-10-05T16:23:35Z"
			  },
			  {
				"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/fa60b0dec5ce1ef3f902b2b110321d05f7730a4d",
				"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
				"id": 14714591394,
				"node_id": "SC_kwDODDMJAM8AAAADbQ7Yog",
				"state": "pending",
				"description": "running...",
				"target_url": "https://c3i.jfrog.io/c3i/misc/summary.html?json=https://c3i.jfrog.io/c3i/misc/logs/pr/7547/1-configs/windows-visual_studio/sqlite3/3.32.2//summary.json",
				"context": "[required] sqlite3/3.32.2@ Windows, Visual Studio",
				"created_at": "2021-10-05T16:23:36Z",
				"updated_at": "2021-10-05T16:23:36Z"
			  },
			  {
				"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/fa60b0dec5ce1ef3f902b2b110321d05f7730a4d",
				"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
				"id": 14714591629,
				"node_id": "SC_kwDODDMJAM8AAAADbQ7ZjQ",
				"state": "pending",
				"description": "running...",
				"target_url": "https://c3i.jfrog.io/c3i/misc/summary.html?json=https://c3i.jfrog.io/c3i/misc/logs/pr/7547/1-configs/linux-gcc/sqlite3/3.35.5//summary.json",
				"context": "[required] sqlite3/3.35.5@ Linux, GCC",
				"created_at": "2021-10-05T16:23:37Z",
				"updated_at": "2021-10-05T16:23:37Z"
			  },
			  {
				"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/fa60b0dec5ce1ef3f902b2b110321d05f7730a4d",
				"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
				"id": 14714591861,
				"node_id": "SC_kwDODDMJAM8AAAADbQ7adQ",
				"state": "pending",
				"description": "running...",
				"target_url": "https://c3i.jfrog.io/c3i/misc/summary.html?json=https://c3i.jfrog.io/c3i/misc/logs/pr/7547/1-configs/linux-clang/sqlite3/3.35.5//summary.json",
				"context": "[required] sqlite3/3.35.5@ Linux, Clang",
				"created_at": "2021-10-05T16:23:37Z",
				"updated_at": "2021-10-05T16:23:37Z"
			  },
			  {
				"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/fa60b0dec5ce1ef3f902b2b110321d05f7730a4d",
				"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
				"id": 14714592186,
				"node_id": "SC_kwDODDMJAM8AAAADbQ7bug",
				"state": "pending",
				"description": "running...",
				"target_url": "https://c3i.jfrog.io/c3i/misc/summary.html?json=https://c3i.jfrog.io/c3i/misc/logs/pr/7547/1-configs/macos-clang/sqlite3/3.35.5//summary.json",
				"context": "[required] sqlite3/3.35.5@ macOS, Clang",
				"created_at": "2021-10-05T16:23:38Z",
				"updated_at": "2021-10-05T16:23:38Z"
			  },
			  {
				"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/fa60b0dec5ce1ef3f902b2b110321d05f7730a4d",
				"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
				"id": 14714592404,
				"node_id": "SC_kwDODDMJAM8AAAADbQ7clA",
				"state": "pending",
				"description": "running...",
				"target_url": "https://c3i.jfrog.io/c3i/misc/summary.html?json=https://c3i.jfrog.io/c3i/misc/logs/pr/7547/1-configs/macos-m1-clang/sqlite3/3.35.5//summary.json",
				"context": "[required] sqlite3/3.35.5@ macOS, Clang (M1/arm64)",
				"created_at": "2021-10-05T16:23:39Z",
				"updated_at": "2021-10-05T16:23:39Z"
			  },
			  {
				"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/fa60b0dec5ce1ef3f902b2b110321d05f7730a4d",
				"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
				"id": 14714592641,
				"node_id": "SC_kwDODDMJAM8AAAADbQ7dgQ",
				"state": "pending",
				"description": "running...",
				"target_url": "https://c3i.jfrog.io/c3i/misc/summary.html?json=https://c3i.jfrog.io/c3i/misc/logs/pr/7547/1-configs/windows-visual_studio/sqlite3/3.35.5//summary.json",
				"context": "[required] sqlite3/3.35.5@ Windows, Visual Studio",
				"created_at": "2021-10-05T16:23:40Z",
				"updated_at": "2021-10-05T16:23:40Z"
			  },
			  {
				"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/fa60b0dec5ce1ef3f902b2b110321d05f7730a4d",
				"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
				"id": 14714592855,
				"node_id": "SC_kwDODDMJAM8AAAADbQ7eVw",
				"state": "pending",
				"description": "running...",
				"target_url": "https://c3i.jfrog.io/c3i/misc/summary.html?json=https://c3i.jfrog.io/c3i/misc/logs/pr/7547/1-configs/linux-gcc/sqlite3/3.34.0//summary.json",
				"context": "[required] sqlite3/3.34.0@ Linux, GCC",
				"created_at": "2021-10-05T16:23:41Z",
				"updated_at": "2021-10-05T16:23:41Z"
			  },
			  {
				"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/fa60b0dec5ce1ef3f902b2b110321d05f7730a4d",
				"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
				"id": 14714593109,
				"node_id": "SC_kwDODDMJAM8AAAADbQ7fVQ",
				"state": "pending",
				"description": "running...",
				"target_url": "https://c3i.jfrog.io/c3i/misc/summary.html?json=https://c3i.jfrog.io/c3i/misc/logs/pr/7547/1-configs/linux-clang/sqlite3/3.34.0//summary.json",
				"context": "[required] sqlite3/3.34.0@ Linux, Clang",
				"created_at": "2021-10-05T16:23:42Z",
				"updated_at": "2021-10-05T16:23:42Z"
			  },
			  {
				"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/fa60b0dec5ce1ef3f902b2b110321d05f7730a4d",
				"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
				"id": 14714593343,
				"node_id": "SC_kwDODDMJAM8AAAADbQ7gPw",
				"state": "pending",
				"description": "running...",
				"target_url": "https://c3i.jfrog.io/c3i/misc/summary.html?json=https://c3i.jfrog.io/c3i/misc/logs/pr/7547/1-configs/macos-clang/sqlite3/3.34.0//summary.json",
				"context": "[required] sqlite3/3.34.0@ macOS, Clang",
				"created_at": "2021-10-05T16:23:42Z",
				"updated_at": "2021-10-05T16:23:42Z"
			  },
			  {
				"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/fa60b0dec5ce1ef3f902b2b110321d05f7730a4d",
				"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
				"id": 14714593593,
				"node_id": "SC_kwDODDMJAM8AAAADbQ7hOQ",
				"state": "pending",
				"description": "running...",
				"target_url": "https://c3i.jfrog.io/c3i/misc/summary.html?json=https://c3i.jfrog.io/c3i/misc/logs/pr/7547/1-configs/macos-m1-clang/sqlite3/3.34.0//summary.json",
				"context": "[required] sqlite3/3.34.0@ macOS, Clang (M1/arm64)",
				"created_at": "2021-10-05T16:23:43Z",
				"updated_at": "2021-10-05T16:23:43Z"
			  },
			  {
				"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/fa60b0dec5ce1ef3f902b2b110321d05f7730a4d",
				"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
				"id": 14714593818,
				"node_id": "SC_kwDODDMJAM8AAAADbQ7iGg",
				"state": "pending",
				"description": "running...",
				"target_url": "https://c3i.jfrog.io/c3i/misc/summary.html?json=https://c3i.jfrog.io/c3i/misc/logs/pr/7547/1-configs/windows-visual_studio/sqlite3/3.34.0//summary.json",
				"context": "[required] sqlite3/3.34.0@ Windows, Visual Studio",
				"created_at": "2021-10-05T16:23:44Z",
				"updated_at": "2021-10-05T16:23:44Z"
			  },
			  {
				"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/fa60b0dec5ce1ef3f902b2b110321d05f7730a4d",
				"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
				"id": 14714594029,
				"node_id": "SC_kwDODDMJAM8AAAADbQ7i7Q",
				"state": "pending",
				"description": "running...",
				"target_url": "https://c3i.jfrog.io/c3i/misc/summary.html?json=https://c3i.jfrog.io/c3i/misc/logs/pr/7547/1-configs/linux-gcc/sqlite3/3.35.4//summary.json",
				"context": "[required] sqlite3/3.35.4@ Linux, GCC",
				"created_at": "2021-10-05T16:23:45Z",
				"updated_at": "2021-10-05T16:23:45Z"
			  },
			  {
				"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/fa60b0dec5ce1ef3f902b2b110321d05f7730a4d",
				"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
				"id": 14714594241,
				"node_id": "SC_kwDODDMJAM8AAAADbQ7jwQ",
				"state": "pending",
				"description": "running...",
				"target_url": "https://c3i.jfrog.io/c3i/misc/summary.html?json=https://c3i.jfrog.io/c3i/misc/logs/pr/7547/1-configs/linux-clang/sqlite3/3.35.4//summary.json",
				"context": "[required] sqlite3/3.35.4@ Linux, Clang",
				"created_at": "2021-10-05T16:23:46Z",
				"updated_at": "2021-10-05T16:23:46Z"
			  }
			],
			"sha": "fa60b0dec5ce1ef3f902b2b110321d05f7730a4d",
			"total_count": 82,
			"repository": {
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
			  "deployments_url": "https://api.github.com/repos/conan-io/conan-center-index/deployments"
			},
			"commit_url": "https://api.github.com/repos/conan-io/conan-center-index/commits/fa60b0dec5ce1ef3f902b2b110321d05f7730a4d",
			"url": "https://api.github.com/repos/conan-io/conan-center-index/commits/fa60b0dec5ce1ef3f902b2b110321d05f7730a4d/status"
		  }`)

	status, res, err := NewClient(&http.Client{}).Repository.GetStatus(context.Background(), "conan-io", "conan-center-index", "fa60b0dec5ce1ef3f902b2b110321d05f7730a4d")

	assert.Equal(t, nil, err)
	assert.Equal(t, 200, res.StatusCode)

	assert.Equal(t, "pending", status.GetState())
	assert.Equal(t, 82, status.GetTotalCount())

	assert.Equal(t, true, gock.IsDone())
}

func TestGetStatusNoStatus(t *testing.T) {
	defer gock.Off()

	gock.New("https://api.github.com").
		Get("/commits/08f356aabf77ff55d96ae43de3e3bfdfb67f6018/status").
		Reply(200).
		BodyString(`{
			"state": "pending",
			"statuses": [
			],
			"sha": "08f356aabf77ff55d96ae43de3e3bfdfb67f6018",
			"total_count": 0,
			"repository": {
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
			  "deployments_url": "https://api.github.com/repos/conan-io/conan-center-index/deployments"
			},
			"commit_url": "https://api.github.com/repos/conan-io/conan-center-index/commits/08f356aabf77ff55d96ae43de3e3bfdfb67f6018",
			"url": "https://api.github.com/repos/conan-io/conan-center-index/commits/08f356aabf77ff55d96ae43de3e3bfdfb67f6018/status"
		  }`)

	status, res, err := NewClient(&http.Client{}).Repository.GetStatus(context.Background(), "conan-io", "conan-center-index", "08f356aabf77ff55d96ae43de3e3bfdfb67f6018")

	assert.Equal(t, ErrNoCommitStatus, err)
	assert.Equal(t, 200, res.StatusCode)

	assert.Nil(t, status)

	assert.Equal(t, true, gock.IsDone())
}

func TestGetStatusDocs(t *testing.T) {
	defer gock.Off()

	gock.New("https://api.github.com").
		Get("/commits/96474105a96d24ab7b735b6afc126614cb8158f9/status").
		Reply(200).
		BodyString(`{
			"state": "success",
			"statuses": [
			  {
				"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/96474105a96d24ab7b735b6afc126614cb8158f9",
				"avatar_url": "https://avatars.githubusercontent.com/oa/160438?v=4",
				"id": 14778350863,
				"node_id": "SC_kwDODDMJAM8AAAADcNu9Dw",
				"state": "success",
				"description": "Contributor License Agreement is signed.",
				"target_url": "https://cla-assistant.io/conan-io/conan-center-index?pullRequest=7643",
				"context": "license/cla",
				"created_at": "2021-10-11T08:27:56Z",
				"updated_at": "2021-10-11T08:27:56Z"
			  },
			  {
				"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/96474105a96d24ab7b735b6afc126614cb8158f9",
				"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
				"id": 14778377152,
				"node_id": "SC_kwDODDMJAM8AAAADcNwjwA",
				"state": "success",
				"description": "This commit looks good",
				"target_url": "https://ci-conan-prod.jfrog.team/job/cci/job/PR-7643/1/display/redirect",
				"context": "continuous-integration/jenkins/pr-merge",
				"created_at": "2021-10-11T08:30:03Z",
				"updated_at": "2021-10-11T08:30:03Z"
			  }
			],
			"sha": "96474105a96d24ab7b735b6afc126614cb8158f9",
			"total_count": 2,
			"repository": {
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
			  "deployments_url": "https://api.github.com/repos/conan-io/conan-center-index/deployments"
			},
			"commit_url": "https://api.github.com/repos/conan-io/conan-center-index/commits/96474105a96d24ab7b735b6afc126614cb8158f9",
			"url": "https://api.github.com/repos/conan-io/conan-center-index/commits/96474105a96d24ab7b735b6afc126614cb8158f9/status"
		  }`)

	status, res, err := NewClient(&http.Client{}).Repository.GetStatus(context.Background(), "conan-io", "conan-center-index", "96474105a96d24ab7b735b6afc126614cb8158f9")

	assert.Equal(t, nil, err)
	assert.Equal(t, 200, res.StatusCode)

	assert.Equal(t, "success", status.GetState())
	assert.Equal(t, 2, status.GetTotalCount())

	assert.Equal(t, true, gock.IsDone())
}
