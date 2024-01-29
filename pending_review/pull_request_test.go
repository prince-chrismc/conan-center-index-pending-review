package pending_review

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	gock "gopkg.in/h2non/gock.v1"
)

func TestGetReviewSummary16144(t *testing.T) {
	defer gock.Off()

	gock.New("https://api.github.com").
		Get("/repos/conan-io/conan-center-index/pulls/16144/reviews").
		Reply(200).
		BodyString(`[
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
	gock.New("https://api.github.com").
		Get("/repos/conan-io/conan-center-index/pulls/16144/files").
		Reply(200).
		BodyString(`[
			{
			  "sha": "c8ad0901b061ac1977cb8f354a1c4d124474a548",
			  "filename": "recipes/re2/all/conanfile.py",
			  "status": "modified",
			  "additions": 5,
			  "deletions": 7,
			  "changes": 12,
			  "blob_url": "https://github.com/conan-io/conan-center-index/blob/e2aa65c961d48d688dd5450811229eb1d62649ba/recipes%2Fre2%2Fall%2Fconanfile.py",
			  "raw_url": "https://github.com/conan-io/conan-center-index/raw/e2aa65c961d48d688dd5450811229eb1d62649ba/recipes%2Fre2%2Fall%2Fconanfile.py",
			  "contents_url": "https://api.github.com/repos/conan-io/conan-center-index/contents/recipes%2Fre2%2Fall%2Fconanfile.py?ref=e2aa65c961d48d688dd5450811229eb1d62649ba",
			  "patch": "@@ -4,17 +4,18 @@\n from conan.tools.files import copy, get, rmdir\n import os\n \n-required_conan_version = \">=1.53.0\"\n+required_conan_version = \">=1.54.0\"\n \n \n class Re2Conan(ConanFile):\n     name = \"re2\"\n     description = \"Fast, safe, thread-friendly regular expression library\"\n-    topics = (\"regex\")\n+    topics = (\"regex\",)\n     url = \"https://github.com/conan-io/conan-center-index\"\n     homepage = \"https://github.com/google/re2\"\n     license = \"BSD-3-Clause\"\n \n+    package_type = \"library\"\n     settings = \"os\", \"arch\", \"compiler\", \"build_type\"\n     options = {\n         \"shared\": [True, False],\n@@ -37,18 +38,15 @@ def layout(self):\n         cmake_layout(self, src_folder=\"src\")\n \n     def validate(self):\n-        if self.info.settings.compiler.get_safe(\"cppstd\"):\n+        if self.settings.compiler.get_safe(\"cppstd\"):\n             check_min_cppstd(self, 11)\n \n     def source(self):\n-        get(self, **self.conan_data[\"sources\"][self.version],\n-            destination=self.source_folder, strip_root=True)\n+        get(self, **self.conan_data[\"sources\"][self.version], strip_root=True)\n \n     def generate(self):\n         tc = CMakeToolchain(self)\n         tc.variables[\"RE2_BUILD_TESTING\"] = False\n-        # Honor BUILD_SHARED_LIBS from conan_toolchain (see https://github.com/conan-io/conan/issues/11840)\n-        tc.cache_variables[\"CMAKE_POLICY_DEFAULT_CMP0077\"] = \"NEW\"\n         tc.generate()\n \n     def build(self):"
			}
		  ]`)
	gock.New("https://api.github.com").
		Get("/repos/SpaceIm/conan-center-index/commits/e2aa65c961d48d688dd5450811229eb1d62649ba").
		Reply(200).
		BodyString(`{
				"sha": "e2aa65c961d48d688dd5450811229eb1d62649ba",
				"node_id": "C_kwDODapHUtoAKGUyYWE2NWM5NjFkNDhkNjg4ZGQ1NDUwODExMjI5ZWIxZDYyNjQ5YmE",
				"commit": {
				  "author": {
					"name": "SpaceIm",
					"email": "30052553+SpaceIm@users.noreply.github.com",
					"date": "2023-02-19T15:10:08Z"
				  },
				  "committer": {
					"name": "SpaceIm",
					"email": "30052553+SpaceIm@users.noreply.github.com",
					"date": "2023-02-19T15:10:08Z"
				  },
				  "message": "fix topics",
				  "tree": {
					"sha": "98b53c8eba48a5c11a4dd763ad2974f17428f0f4",
					"url": "https://api.github.com/repos/SpaceIm/conan-center-index/git/trees/98b53c8eba48a5c11a4dd763ad2974f17428f0f4"
				  },
				  "url": "https://api.github.com/repos/SpaceIm/conan-center-index/git/commits/e2aa65c961d48d688dd5450811229eb1d62649ba",
				  "comment_count": 0,
				  "verification": {
					"verified": false,
					"reason": "unsigned",
					"signature": null,
					"payload": null
				  }
				},
				"url": "https://api.github.com/repos/SpaceIm/conan-center-index/commits/e2aa65c961d48d688dd5450811229eb1d62649ba",
				"html_url": "https://github.com/SpaceIm/conan-center-index/commit/e2aa65c961d48d688dd5450811229eb1d62649ba",
				"comments_url": "https://api.github.com/repos/SpaceIm/conan-center-index/commits/e2aa65c961d48d688dd5450811229eb1d62649ba/comments",
				"author": {
				  "login": "SpaceIm",
				  "id": 30052553,
				  "node_id": "MDQ6VXNlcjMwMDUyNTUz",
				  "avatar_url": "https://avatars.githubusercontent.com/u/30052553?v=4",
				  "gravatar_id": "",
				  "url": "https://api.github.com/users/SpaceIm",
				  "html_url": "https://github.com/SpaceIm",
				  "followers_url": "https://api.github.com/users/SpaceIm/followers",
				  "following_url": "https://api.github.com/users/SpaceIm/following{/other_user}",
				  "gists_url": "https://api.github.com/users/SpaceIm/gists{/gist_id}",
				  "starred_url": "https://api.github.com/users/SpaceIm/starred{/owner}{/repo}",
				  "subscriptions_url": "https://api.github.com/users/SpaceIm/subscriptions",
				  "organizations_url": "https://api.github.com/users/SpaceIm/orgs",
				  "repos_url": "https://api.github.com/users/SpaceIm/repos",
				  "events_url": "https://api.github.com/users/SpaceIm/events{/privacy}",
				  "received_events_url": "https://api.github.com/users/SpaceIm/received_events",
				  "type": "User",
				  "site_admin": false
				},
				"committer": {
				  "login": "SpaceIm",
				  "id": 30052553,
				  "node_id": "MDQ6VXNlcjMwMDUyNTUz",
				  "avatar_url": "https://avatars.githubusercontent.com/u/30052553?v=4",
				  "gravatar_id": "",
				  "url": "https://api.github.com/users/SpaceIm",
				  "html_url": "https://github.com/SpaceIm",
				  "followers_url": "https://api.github.com/users/SpaceIm/followers",
				  "following_url": "https://api.github.com/users/SpaceIm/following{/other_user}",
				  "gists_url": "https://api.github.com/users/SpaceIm/gists{/gist_id}",
				  "starred_url": "https://api.github.com/users/SpaceIm/starred{/owner}{/repo}",
				  "subscriptions_url": "https://api.github.com/users/SpaceIm/subscriptions",
				  "organizations_url": "https://api.github.com/users/SpaceIm/orgs",
				  "repos_url": "https://api.github.com/users/SpaceIm/repos",
				  "events_url": "https://api.github.com/users/SpaceIm/events{/privacy}",
				  "received_events_url": "https://api.github.com/users/SpaceIm/received_events",
				  "type": "User",
				  "site_admin": false
				},
				"parents": [
				  {
					"sha": "ee78bc3c07f6885cb6d2c0b3943bb21c52c584d4",
					"url": "https://api.github.com/repos/SpaceIm/conan-center-index/commits/ee78bc3c07f6885cb6d2c0b3943bb21c52c584d4",
					"html_url": "https://github.com/SpaceIm/conan-center-index/commit/ee78bc3c07f6885cb6d2c0b3943bb21c52c584d4"
				  }
				],
				"stats": {
				  "total": 2,
				  "additions": 1,
				  "deletions": 1
				},
				"files": [
				  {
					"sha": "c8ad0901b061ac1977cb8f354a1c4d124474a548",
					"filename": "recipes/re2/all/conanfile.py",
					"status": "modified",
					"additions": 1,
					"deletions": 1,
					"changes": 2,
					"blob_url": "https://github.com/SpaceIm/conan-center-index/blob/e2aa65c961d48d688dd5450811229eb1d62649ba/recipes%2Fre2%2Fall%2Fconanfile.py",
					"raw_url": "https://github.com/SpaceIm/conan-center-index/raw/e2aa65c961d48d688dd5450811229eb1d62649ba/recipes%2Fre2%2Fall%2Fconanfile.py",
					"contents_url": "https://api.github.com/repos/SpaceIm/conan-center-index/contents/recipes%2Fre2%2Fall%2Fconanfile.py?ref=e2aa65c961d48d688dd5450811229eb1d62649ba",
					"patch": "@@ -10,7 +10,7 @@\n class Re2Conan(ConanFile):\n     name = \"re2\"\n     description = \"Fast, safe, thread-friendly regular expression library\"\n-    topics = (\"regex\")\n+    topics = (\"regex\",)\n     url = \"https://github.com/conan-io/conan-center-index\"\n     homepage = \"https://github.com/google/re2\"\n     license = \"BSD-3-Clause\""
				  }
				]
			  }`)
	gock.New("https://api.github.com").
		Get("/repos/conan-io/conan-center-index/commits/e2aa65c961d48d688dd5450811229eb1d62649ba/status").
		Reply(200).
		BodyString(`{
					"state": "failure",
					"statuses": [
					  {
						"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/e2aa65c961d48d688dd5450811229eb1d62649ba",
						"avatar_url": "https://avatars.githubusercontent.com/oa/160438?v=4",
						"id": 21637263246,
						"node_id": "SC_kwDODDMJAM8AAAAFCa5vjg",
						"state": "success",
						"description": "Contributor License Agreement is signed.",
						"target_url": "https://cla-assistant.io/conan-io/conan-center-index?pullRequest=16144",
						"context": "license/cla",
						"created_at": "2023-02-19T15:10:42Z",
						"updated_at": "2023-02-19T15:10:42Z"
					  },
					  {
						"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/e2aa65c961d48d688dd5450811229eb1d62649ba",
						"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
						"id": 21637332718,
						"node_id": "SC_kwDODDMJAM8AAAAFCa9-7g",
						"state": "success",
						"description": "All green! (48)",
						"target_url": "https://c3i.jfrog.io/c3i/misc/summary.html?json=https://c3i.jfrog.io/c3i/misc/logs/pr/16144/1-linux-gcc/re2/20221201//summary.json",
						"context": "[required] re2/20221201@ Linux, GCC",
						"created_at": "2023-02-19T15:34:05Z",
						"updated_at": "2023-02-19T15:34:05Z"
					  },
					  {
						"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/e2aa65c961d48d688dd5450811229eb1d62649ba",
						"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
						"id": 21637336168,
						"node_id": "SC_kwDODDMJAM8AAAAFCa-MaA",
						"state": "success",
						"description": "All green! (48)",
						"target_url": "https://c3i.jfrog.io/c3i/misc/summary.html?json=https://c3i.jfrog.io/c3i/misc/logs/pr/16144/1-linux-gcc/re2/20220601//summary.json",
						"context": "[required] re2/20220601@ Linux, GCC",
						"created_at": "2023-02-19T15:35:12Z",
						"updated_at": "2023-02-19T15:35:12Z"
					  },
					  {
						"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/e2aa65c961d48d688dd5450811229eb1d62649ba",
						"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
						"id": 21637336169,
						"node_id": "SC_kwDODDMJAM8AAAAFCa-MaQ",
						"state": "success",
						"description": "All green! (48)",
						"target_url": "https://c3i.jfrog.io/c3i/misc/summary.html?json=https://c3i.jfrog.io/c3i/misc/logs/pr/16144/1-linux-gcc/re2/20211101//summary.json",
						"context": "[required] re2/20211101@ Linux, GCC",
						"created_at": "2023-02-19T15:35:12Z",
						"updated_at": "2023-02-19T15:35:12Z"
					  },
					  {
						"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/e2aa65c961d48d688dd5450811229eb1d62649ba",
						"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
						"id": 21637336172,
						"node_id": "SC_kwDODDMJAM8AAAAFCa-MbA",
						"state": "success",
						"description": "All green! (48)",
						"target_url": "https://c3i.jfrog.io/c3i/misc/summary.html?json=https://c3i.jfrog.io/c3i/misc/logs/pr/16144/1-linux-gcc/re2/20230201//summary.json",
						"context": "[required] re2/20230201@ Linux, GCC",
						"created_at": "2023-02-19T15:35:12Z",
						"updated_at": "2023-02-19T15:35:12Z"
					  },
					  {
						"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/e2aa65c961d48d688dd5450811229eb1d62649ba",
						"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
						"id": 21637367600,
						"node_id": "SC_kwDODDMJAM8AAAAFCbAHMA",
						"state": "success",
						"description": "All green! (24)",
						"target_url": "https://c3i.jfrog.io/c3i/misc/summary.html?json=https://c3i.jfrog.io/c3i/misc/logs/pr/16144/1-linux-clang/re2/20221201//summary.json",
						"context": "[required] re2/20221201@ Linux, Clang",
						"created_at": "2023-02-19T15:45:38Z",
						"updated_at": "2023-02-19T15:45:38Z"
					  },
					  {
						"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/e2aa65c961d48d688dd5450811229eb1d62649ba",
						"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
						"id": 21637371034,
						"node_id": "SC_kwDODDMJAM8AAAAFCbAUmg",
						"state": "success",
						"description": "All green! (24)",
						"target_url": "https://c3i.jfrog.io/c3i/misc/summary.html?json=https://c3i.jfrog.io/c3i/misc/logs/pr/16144/1-linux-clang/re2/20220601//summary.json",
						"context": "[required] re2/20220601@ Linux, Clang",
						"created_at": "2023-02-19T15:46:53Z",
						"updated_at": "2023-02-19T15:46:53Z"
					  },
					  {
						"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/e2aa65c961d48d688dd5450811229eb1d62649ba",
						"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
						"id": 21637372924,
						"node_id": "SC_kwDODDMJAM8AAAAFCbAb_A",
						"state": "success",
						"description": "All green! (24)",
						"target_url": "https://c3i.jfrog.io/c3i/misc/summary.html?json=https://c3i.jfrog.io/c3i/misc/logs/pr/16144/1-linux-clang/re2/20211101//summary.json",
						"context": "[required] re2/20211101@ Linux, Clang",
						"created_at": "2023-02-19T15:47:31Z",
						"updated_at": "2023-02-19T15:47:31Z"
					  },
					  {
						"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/e2aa65c961d48d688dd5450811229eb1d62649ba",
						"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
						"id": 21637373457,
						"node_id": "SC_kwDODDMJAM8AAAAFCbAeEQ",
						"state": "success",
						"description": "All green! (24)",
						"target_url": "https://c3i.jfrog.io/c3i/misc/summary.html?json=https://c3i.jfrog.io/c3i/misc/logs/pr/16144/1-linux-clang/re2/20230201//summary.json",
						"context": "[required] re2/20230201@ Linux, Clang",
						"created_at": "2023-02-19T15:47:43Z",
						"updated_at": "2023-02-19T15:47:43Z"
					  },
					  {
						"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/e2aa65c961d48d688dd5450811229eb1d62649ba",
						"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
						"id": 21637585129,
						"node_id": "SC_kwDODDMJAM8AAAAFCbNY6Q",
						"state": "success",
						"description": "All green! (12)",
						"target_url": "https://c3i.jfrog.io/c3i/misc/summary.html?json=https://c3i.jfrog.io/c3i/misc/logs/pr/16144/1-configs/macos-clang/re2/20221201//summary.json",
						"context": "[required] re2/20221201@ macOS, Clang",
						"created_at": "2023-02-19T16:59:50Z",
						"updated_at": "2023-02-19T16:59:50Z"
					  },
					  {
						"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/e2aa65c961d48d688dd5450811229eb1d62649ba",
						"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
						"id": 21637632179,
						"node_id": "SC_kwDODDMJAM8AAAAFCbQQsw",
						"state": "success",
						"description": "All green! (12)",
						"target_url": "https://c3i.jfrog.io/c3i/misc/summary.html?json=https://c3i.jfrog.io/c3i/misc/logs/pr/16144/1-configs/macos-clang/re2/20220601//summary.json",
						"context": "[required] re2/20220601@ macOS, Clang",
						"created_at": "2023-02-19T17:15:16Z",
						"updated_at": "2023-02-19T17:15:16Z"
					  },
					  {
						"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/e2aa65c961d48d688dd5450811229eb1d62649ba",
						"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
						"id": 21637632184,
						"node_id": "SC_kwDODDMJAM8AAAAFCbQQuA",
						"state": "success",
						"description": "All green! (12)",
						"target_url": "https://c3i.jfrog.io/c3i/misc/summary.html?json=https://c3i.jfrog.io/c3i/misc/logs/pr/16144/1-configs/macos-clang/re2/20211101//summary.json",
						"context": "[required] re2/20211101@ macOS, Clang",
						"created_at": "2023-02-19T17:15:16Z",
						"updated_at": "2023-02-19T17:15:16Z"
					  },
					  {
						"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/e2aa65c961d48d688dd5450811229eb1d62649ba",
						"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
						"id": 21637634466,
						"node_id": "SC_kwDODDMJAM8AAAAFCbQZog",
						"state": "success",
						"description": "All green! (12)",
						"target_url": "https://c3i.jfrog.io/c3i/misc/summary.html?json=https://c3i.jfrog.io/c3i/misc/logs/pr/16144/1-configs/macos-clang/re2/20230201//summary.json",
						"context": "[required] re2/20230201@ macOS, Clang",
						"created_at": "2023-02-19T17:15:53Z",
						"updated_at": "2023-02-19T17:15:53Z"
					  },
					  {
						"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/e2aa65c961d48d688dd5450811229eb1d62649ba",
						"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
						"id": 21637721461,
						"node_id": "SC_kwDODDMJAM8AAAAFCbVtdQ",
						"state": "success",
						"description": "All green! (8)",
						"target_url": "https://c3i.jfrog.io/c3i/misc/summary.html?json=https://c3i.jfrog.io/c3i/misc/logs/pr/16144/1-configs/macos-m1-clang/re2/20221201//summary.json",
						"context": "[required] re2/20221201@ macOS, Clang (M1/arm64)",
						"created_at": "2023-02-19T17:51:01Z",
						"updated_at": "2023-02-19T17:51:01Z"
					  },
					  {
						"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/e2aa65c961d48d688dd5450811229eb1d62649ba",
						"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
						"id": 21637740310,
						"node_id": "SC_kwDODDMJAM8AAAAFCbW3Fg",
						"state": "success",
						"description": "All green! (8)",
						"target_url": "https://c3i.jfrog.io/c3i/misc/summary.html?json=https://c3i.jfrog.io/c3i/misc/logs/pr/16144/1-configs/macos-m1-clang/re2/20220601//summary.json",
						"context": "[required] re2/20220601@ macOS, Clang (M1/arm64)",
						"created_at": "2023-02-19T17:58:40Z",
						"updated_at": "2023-02-19T17:58:40Z"
					  },
					  {
						"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/e2aa65c961d48d688dd5450811229eb1d62649ba",
						"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
						"id": 21637745648,
						"node_id": "SC_kwDODDMJAM8AAAAFCbXL8A",
						"state": "success",
						"description": "All green! (8)",
						"target_url": "https://c3i.jfrog.io/c3i/misc/summary.html?json=https://c3i.jfrog.io/c3i/misc/logs/pr/16144/1-configs/macos-m1-clang/re2/20211101//summary.json",
						"context": "[required] re2/20211101@ macOS, Clang (M1/arm64)",
						"created_at": "2023-02-19T18:00:30Z",
						"updated_at": "2023-02-19T18:00:30Z"
					  },
					  {
						"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/e2aa65c961d48d688dd5450811229eb1d62649ba",
						"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
						"id": 21637751933,
						"node_id": "SC_kwDODDMJAM8AAAAFCbXkfQ",
						"state": "success",
						"description": "All green! (8)",
						"target_url": "https://c3i.jfrog.io/c3i/misc/summary.html?json=https://c3i.jfrog.io/c3i/misc/logs/pr/16144/1-configs/macos-m1-clang/re2/20230201//summary.json",
						"context": "[required] re2/20230201@ macOS, Clang (M1/arm64)",
						"created_at": "2023-02-19T18:02:46Z",
						"updated_at": "2023-02-19T18:02:46Z"
					  },
					  {
						"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/e2aa65c961d48d688dd5450811229eb1d62649ba",
						"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
						"id": 21637976079,
						"node_id": "SC_kwDODDMJAM8AAAAFCblQDw",
						"state": "success",
						"description": "All green! (12)",
						"target_url": "https://c3i.jfrog.io/c3i/misc/summary.html?json=https://c3i.jfrog.io/c3i/misc/logs/pr/16144/1-configs/windows-visual_studio/re2/20211101//summary.json",
						"context": "[required] re2/20211101@ Windows, Visual Studio",
						"created_at": "2023-02-19T19:32:14Z",
						"updated_at": "2023-02-19T19:32:14Z"
					  },
					  {
						"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/e2aa65c961d48d688dd5450811229eb1d62649ba",
						"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
						"id": 21637977400,
						"node_id": "SC_kwDODDMJAM8AAAAFCblVOA",
						"state": "success",
						"description": "All green! (12)",
						"target_url": "https://c3i.jfrog.io/c3i/misc/summary.html?json=https://c3i.jfrog.io/c3i/misc/logs/pr/16144/1-configs/windows-visual_studio/re2/20221201//summary.json",
						"context": "[required] re2/20221201@ Windows, Visual Studio",
						"created_at": "2023-02-19T19:32:46Z",
						"updated_at": "2023-02-19T19:32:46Z"
					  },
					  {
						"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/e2aa65c961d48d688dd5450811229eb1d62649ba",
						"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
						"id": 21637978937,
						"node_id": "SC_kwDODDMJAM8AAAAFCblbOQ",
						"state": "success",
						"description": "All green! (12)",
						"target_url": "https://c3i.jfrog.io/c3i/misc/summary.html?json=https://c3i.jfrog.io/c3i/misc/logs/pr/16144/1-configs/windows-visual_studio/re2/20220601//summary.json",
						"context": "[required] re2/20220601@ Windows, Visual Studio",
						"created_at": "2023-02-19T19:33:33Z",
						"updated_at": "2023-02-19T19:33:33Z"
					  },
					  {
						"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/e2aa65c961d48d688dd5450811229eb1d62649ba",
						"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
						"id": 21637978938,
						"node_id": "SC_kwDODDMJAM8AAAAFCblbOg",
						"state": "success",
						"description": "All green! (12)",
						"target_url": "https://c3i.jfrog.io/c3i/misc/summary.html?json=https://c3i.jfrog.io/c3i/misc/logs/pr/16144/1-configs/windows-visual_studio/re2/20230201//summary.json",
						"context": "[required] re2/20230201@ Windows, Visual Studio",
						"created_at": "2023-02-19T19:33:33Z",
						"updated_at": "2023-02-19T19:33:33Z"
					  },
					  {
						"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/e2aa65c961d48d688dd5450811229eb1d62649ba",
						"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
						"id": 21638024417,
						"node_id": "SC_kwDODDMJAM8AAAAFCboM4Q",
						"state": "success",
						"description": "All green! (48)",
						"target_url": "https://c3i.jfrog.io/c3i/misc/summary.html?json=https://c3i.jfrog.io/c3i/misc/logs/pr/16144/1-linux-gcc/re2/20210601//summary.json",
						"context": "[required] re2/20210601@ Linux, GCC",
						"created_at": "2023-02-19T19:53:35Z",
						"updated_at": "2023-02-19T19:53:35Z"
					  },
					  {
						"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/e2aa65c961d48d688dd5450811229eb1d62649ba",
						"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
						"id": 21638026009,
						"node_id": "SC_kwDODDMJAM8AAAAFCboTGQ",
						"state": "success",
						"description": "All green! (48)",
						"target_url": "https://c3i.jfrog.io/c3i/misc/summary.html?json=https://c3i.jfrog.io/c3i/misc/logs/pr/16144/1-linux-gcc/re2/20220201//summary.json",
						"context": "[required] re2/20220201@ Linux, GCC",
						"created_at": "2023-02-19T19:54:11Z",
						"updated_at": "2023-02-19T19:54:11Z"
					  },
					  {
						"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/e2aa65c961d48d688dd5450811229eb1d62649ba",
						"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
						"id": 21638026010,
						"node_id": "SC_kwDODDMJAM8AAAAFCboTGg",
						"state": "success",
						"description": "All green! (48)",
						"target_url": "https://c3i.jfrog.io/c3i/misc/summary.html?json=https://c3i.jfrog.io/c3i/misc/logs/pr/16144/1-linux-gcc/re2/20210401//summary.json",
						"context": "[required] re2/20210401@ Linux, GCC",
						"created_at": "2023-02-19T19:54:11Z",
						"updated_at": "2023-02-19T19:54:11Z"
					  },
					  {
						"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/e2aa65c961d48d688dd5450811229eb1d62649ba",
						"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
						"id": 21638029481,
						"node_id": "SC_kwDODDMJAM8AAAAFCbogqQ",
						"state": "success",
						"description": "All green! (48)",
						"target_url": "https://c3i.jfrog.io/c3i/misc/summary.html?json=https://c3i.jfrog.io/c3i/misc/logs/pr/16144/1-linux-gcc/re2/20210202//summary.json",
						"context": "[required] re2/20210202@ Linux, GCC",
						"created_at": "2023-02-19T19:56:00Z",
						"updated_at": "2023-02-19T19:56:00Z"
					  },
					  {
						"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/e2aa65c961d48d688dd5450811229eb1d62649ba",
						"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
						"id": 21638053441,
						"node_id": "SC_kwDODDMJAM8AAAAFCbp-QQ",
						"state": "success",
						"description": "All green! (24)",
						"target_url": "https://c3i.jfrog.io/c3i/misc/summary.html?json=https://c3i.jfrog.io/c3i/misc/logs/pr/16144/1-linux-clang/re2/20210601//summary.json",
						"context": "[required] re2/20210601@ Linux, Clang",
						"created_at": "2023-02-19T20:04:57Z",
						"updated_at": "2023-02-19T20:04:57Z"
					  },
					  {
						"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/e2aa65c961d48d688dd5450811229eb1d62649ba",
						"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
						"id": 21638055735,
						"node_id": "SC_kwDODDMJAM8AAAAFCbqHNw",
						"state": "success",
						"description": "All green! (24)",
						"target_url": "https://c3i.jfrog.io/c3i/misc/summary.html?json=https://c3i.jfrog.io/c3i/misc/logs/pr/16144/1-linux-clang/re2/20220201//summary.json",
						"context": "[required] re2/20220201@ Linux, Clang",
						"created_at": "2023-02-19T20:05:41Z",
						"updated_at": "2023-02-19T20:05:41Z"
					  },
					  {
						"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/e2aa65c961d48d688dd5450811229eb1d62649ba",
						"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
						"id": 21638055845,
						"node_id": "SC_kwDODDMJAM8AAAAFCbqHpQ",
						"state": "success",
						"description": "All green! (24)",
						"target_url": "https://c3i.jfrog.io/c3i/misc/summary.html?json=https://c3i.jfrog.io/c3i/misc/logs/pr/16144/1-linux-clang/re2/20210401//summary.json",
						"context": "[required] re2/20210401@ Linux, Clang",
						"created_at": "2023-02-19T20:05:43Z",
						"updated_at": "2023-02-19T20:05:43Z"
					  },
					  {
						"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/e2aa65c961d48d688dd5450811229eb1d62649ba",
						"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
						"id": 21638061816,
						"node_id": "SC_kwDODDMJAM8AAAAFCbqe-A",
						"state": "success",
						"description": "All green! (24)",
						"target_url": "https://c3i.jfrog.io/c3i/misc/summary.html?json=https://c3i.jfrog.io/c3i/misc/logs/pr/16144/1-linux-clang/re2/20210202//summary.json",
						"context": "[required] re2/20210202@ Linux, Clang",
						"created_at": "2023-02-19T20:07:27Z",
						"updated_at": "2023-02-19T20:07:27Z"
					  },
					  {
						"url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/e2aa65c961d48d688dd5450811229eb1d62649ba",
						"avatar_url": "https://avatars.githubusercontent.com/u/54393557?v=4",
						"id": 21638092734,
						"node_id": "SC_kwDODDMJAM8AAAAFCbsXvg",
						"state": "success",
						"description": "All green! (12)",
						"target_url": "https://c3i.jfrog.io/c3i/misc/summary.html?json=https://c3i.jfrog.io/c3i/misc/logs/pr/16144/1-configs/macos-clang/re2/20210601//summary.json",
						"context": "[required] re2/20210601@ macOS, Clang",
						"created_at": "2023-02-19T20:16:47Z",
						"updated_at": "2023-02-19T20:16:47Z"
					  }
					],
					"sha": "e2aa65c961d48d688dd5450811229eb1d62649ba",
					"total_count": 53,
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
					"commit_url": "https://api.github.com/repos/conan-io/conan-center-index/commits/e2aa65c961d48d688dd5450811229eb1d62649ba",
					"url": "https://api.github.com/repos/conan-io/conan-center-index/commits/e2aa65c961d48d688dd5450811229eb1d62649ba/status"
				  }`)

	client := NewClient(&http.Client{}, WorkingRepository{
		Owner: "prince-chrismc", Name: "conan-center-index-pending-review",
	})
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
	pr := parsePullRequestJSON(t, `{
		"url": "https://api.github.com/repos/conan-io/conan-center-index/pulls/16144",
		"id": 1246324241,
		"node_id": "PR_kwDODDMJAM5KSWYR",
		"html_url": "https://github.com/conan-io/conan-center-index/pull/16144",
		"diff_url": "https://github.com/conan-io/conan-center-index/pull/16144.diff",
		"patch_url": "https://github.com/conan-io/conan-center-index/pull/16144.patch",
		"issue_url": "https://api.github.com/repos/conan-io/conan-center-index/issues/16144",
		"number": 16144,
		"state": "open",
		"locked": false,
		"title": "re2: modernize more for conan v2",
		"user": {
		  "login": "SpaceIm",
		  "id": 30052553,
		  "node_id": "MDQ6VXNlcjMwMDUyNTUz",
		  "avatar_url": "https://avatars.githubusercontent.com/u/30052553?v=4",
		  "gravatar_id": "",
		  "url": "https://api.github.com/users/SpaceIm",
		  "html_url": "https://github.com/SpaceIm",
		  "followers_url": "https://api.github.com/users/SpaceIm/followers",
		  "following_url": "https://api.github.com/users/SpaceIm/following{/other_user}",
		  "gists_url": "https://api.github.com/users/SpaceIm/gists{/gist_id}",
		  "starred_url": "https://api.github.com/users/SpaceIm/starred{/owner}{/repo}",
		  "subscriptions_url": "https://api.github.com/users/SpaceIm/subscriptions",
		  "organizations_url": "https://api.github.com/users/SpaceIm/orgs",
		  "repos_url": "https://api.github.com/users/SpaceIm/repos",
		  "events_url": "https://api.github.com/users/SpaceIm/events{/privacy}",
		  "received_events_url": "https://api.github.com/users/SpaceIm/received_events",
		  "type": "User",
		  "site_admin": false
		},
		"body": "Specify library name and version:  **lib/1.0**\r\n\r\n<!-- This is also a good place to share with all of us **why you are submitting this PR** (specially if it is a new addition to ConanCenter): is it a dependency of other libraries you want to package? Are you the author of the library? Thanks! -->\r\n\r\n\r\n---\r\n\r\n- [ ] I've read the [contributing guidelines](https://github.com/conan-io/conan-center-index/blob/master/CONTRIBUTING.md).\r\n- [ ] I've used a [recent](https://github.com/conan-io/conan/releases/latest) Conan client version close to the [currently deployed](https://github.com/conan-io/conan-center-index/blob/master/.c3i/config_v1.yml#L6).\r\n- [ ] I've tried at least one configuration locally with the [conan-center hook](https://github.com/conan-io/hooks.git) activated.\r\n",
		"created_at": "2023-02-19T15:10:36Z",
		"updated_at": "2023-03-14T21:19:06Z",
		"closed_at": null,
		"merged_at": null,
		"merge_commit_sha": "6009cd08f7f9639d0882e41d4e9c69c31a843689",
		"assignee": null,
		"assignees": [
	  
		],
		"requested_reviewers": [
		  {
			"login": "danimtb",
			"id": 10808592,
			"node_id": "MDQ6VXNlcjEwODA4NTky",
			"avatar_url": "https://avatars.githubusercontent.com/u/10808592?v=4",
			"gravatar_id": "",
			"url": "https://api.github.com/users/danimtb",
			"html_url": "https://github.com/danimtb",
			"followers_url": "https://api.github.com/users/danimtb/followers",
			"following_url": "https://api.github.com/users/danimtb/following{/other_user}",
			"gists_url": "https://api.github.com/users/danimtb/gists{/gist_id}",
			"starred_url": "https://api.github.com/users/danimtb/starred{/owner}{/repo}",
			"subscriptions_url": "https://api.github.com/users/danimtb/subscriptions",
			"organizations_url": "https://api.github.com/users/danimtb/orgs",
			"repos_url": "https://api.github.com/users/danimtb/repos",
			"events_url": "https://api.github.com/users/danimtb/events{/privacy}",
			"received_events_url": "https://api.github.com/users/danimtb/received_events",
			"type": "User",
			"site_admin": false
		  }
		],
		"requested_teams": [
	  
		],
		"labels": [
	  
		],
		"milestone": null,
		"draft": false,
		"commits_url": "https://api.github.com/repos/conan-io/conan-center-index/pulls/16144/commits",
		"review_comments_url": "https://api.github.com/repos/conan-io/conan-center-index/pulls/16144/comments",
		"review_comment_url": "https://api.github.com/repos/conan-io/conan-center-index/pulls/comments{/number}",
		"comments_url": "https://api.github.com/repos/conan-io/conan-center-index/issues/16144/comments",
		"statuses_url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/e2aa65c961d48d688dd5450811229eb1d62649ba",
		"head": {
		  "label": "SpaceIm:re2-modernize-more",
		  "ref": "re2-modernize-more",
		  "sha": "e2aa65c961d48d688dd5450811229eb1d62649ba",
		  "user": {
			"login": "SpaceIm",
			"id": 30052553,
			"node_id": "MDQ6VXNlcjMwMDUyNTUz",
			"avatar_url": "https://avatars.githubusercontent.com/u/30052553?v=4",
			"gravatar_id": "",
			"url": "https://api.github.com/users/SpaceIm",
			"html_url": "https://github.com/SpaceIm",
			"followers_url": "https://api.github.com/users/SpaceIm/followers",
			"following_url": "https://api.github.com/users/SpaceIm/following{/other_user}",
			"gists_url": "https://api.github.com/users/SpaceIm/gists{/gist_id}",
			"starred_url": "https://api.github.com/users/SpaceIm/starred{/owner}{/repo}",
			"subscriptions_url": "https://api.github.com/users/SpaceIm/subscriptions",
			"organizations_url": "https://api.github.com/users/SpaceIm/orgs",
			"repos_url": "https://api.github.com/users/SpaceIm/repos",
			"events_url": "https://api.github.com/users/SpaceIm/events{/privacy}",
			"received_events_url": "https://api.github.com/users/SpaceIm/received_events",
			"type": "User",
			"site_admin": false
		  },
		  "repo": {
			"id": 229263186,
			"node_id": "MDEwOlJlcG9zaXRvcnkyMjkyNjMxODY=",
			"name": "conan-center-index",
			"full_name": "SpaceIm/conan-center-index",
			"private": false,
			"owner": {
			  "login": "SpaceIm",
			  "id": 30052553,
			  "node_id": "MDQ6VXNlcjMwMDUyNTUz",
			  "avatar_url": "https://avatars.githubusercontent.com/u/30052553?v=4",
			  "gravatar_id": "",
			  "url": "https://api.github.com/users/SpaceIm",
			  "html_url": "https://github.com/SpaceIm",
			  "followers_url": "https://api.github.com/users/SpaceIm/followers",
			  "following_url": "https://api.github.com/users/SpaceIm/following{/other_user}",
			  "gists_url": "https://api.github.com/users/SpaceIm/gists{/gist_id}",
			  "starred_url": "https://api.github.com/users/SpaceIm/starred{/owner}{/repo}",
			  "subscriptions_url": "https://api.github.com/users/SpaceIm/subscriptions",
			  "organizations_url": "https://api.github.com/users/SpaceIm/orgs",
			  "repos_url": "https://api.github.com/users/SpaceIm/repos",
			  "events_url": "https://api.github.com/users/SpaceIm/events{/privacy}",
			  "received_events_url": "https://api.github.com/users/SpaceIm/received_events",
			  "type": "User",
			  "site_admin": false
			},
			"html_url": "https://github.com/SpaceIm/conan-center-index",
			"description": "Recipes for the conan-center repository",
			"fork": true,
			"url": "https://api.github.com/repos/SpaceIm/conan-center-index",
			"forks_url": "https://api.github.com/repos/SpaceIm/conan-center-index/forks",
			"keys_url": "https://api.github.com/repos/SpaceIm/conan-center-index/keys{/key_id}",
			"collaborators_url": "https://api.github.com/repos/SpaceIm/conan-center-index/collaborators{/collaborator}",
			"teams_url": "https://api.github.com/repos/SpaceIm/conan-center-index/teams",
			"hooks_url": "https://api.github.com/repos/SpaceIm/conan-center-index/hooks",
			"issue_events_url": "https://api.github.com/repos/SpaceIm/conan-center-index/issues/events{/number}",
			"events_url": "https://api.github.com/repos/SpaceIm/conan-center-index/events",
			"assignees_url": "https://api.github.com/repos/SpaceIm/conan-center-index/assignees{/user}",
			"branches_url": "https://api.github.com/repos/SpaceIm/conan-center-index/branches{/branch}",
			"tags_url": "https://api.github.com/repos/SpaceIm/conan-center-index/tags",
			"blobs_url": "https://api.github.com/repos/SpaceIm/conan-center-index/git/blobs{/sha}",
			"git_tags_url": "https://api.github.com/repos/SpaceIm/conan-center-index/git/tags{/sha}",
			"git_refs_url": "https://api.github.com/repos/SpaceIm/conan-center-index/git/refs{/sha}",
			"trees_url": "https://api.github.com/repos/SpaceIm/conan-center-index/git/trees{/sha}",
			"statuses_url": "https://api.github.com/repos/SpaceIm/conan-center-index/statuses/{sha}",
			"languages_url": "https://api.github.com/repos/SpaceIm/conan-center-index/languages",
			"stargazers_url": "https://api.github.com/repos/SpaceIm/conan-center-index/stargazers",
			"contributors_url": "https://api.github.com/repos/SpaceIm/conan-center-index/contributors",
			"subscribers_url": "https://api.github.com/repos/SpaceIm/conan-center-index/subscribers",
			"subscription_url": "https://api.github.com/repos/SpaceIm/conan-center-index/subscription",
			"commits_url": "https://api.github.com/repos/SpaceIm/conan-center-index/commits{/sha}",
			"git_commits_url": "https://api.github.com/repos/SpaceIm/conan-center-index/git/commits{/sha}",
			"comments_url": "https://api.github.com/repos/SpaceIm/conan-center-index/comments{/number}",
			"issue_comment_url": "https://api.github.com/repos/SpaceIm/conan-center-index/issues/comments{/number}",
			"contents_url": "https://api.github.com/repos/SpaceIm/conan-center-index/contents/{+path}",
			"compare_url": "https://api.github.com/repos/SpaceIm/conan-center-index/compare/{base}...{head}",
			"merges_url": "https://api.github.com/repos/SpaceIm/conan-center-index/merges",
			"archive_url": "https://api.github.com/repos/SpaceIm/conan-center-index/{archive_format}{/ref}",
			"downloads_url": "https://api.github.com/repos/SpaceIm/conan-center-index/downloads",
			"issues_url": "https://api.github.com/repos/SpaceIm/conan-center-index/issues{/number}",
			"pulls_url": "https://api.github.com/repos/SpaceIm/conan-center-index/pulls{/number}",
			"milestones_url": "https://api.github.com/repos/SpaceIm/conan-center-index/milestones{/number}",
			"notifications_url": "https://api.github.com/repos/SpaceIm/conan-center-index/notifications{?since,all,participating}",
			"labels_url": "https://api.github.com/repos/SpaceIm/conan-center-index/labels{/name}",
			"releases_url": "https://api.github.com/repos/SpaceIm/conan-center-index/releases{/id}",
			"deployments_url": "https://api.github.com/repos/SpaceIm/conan-center-index/deployments",
			"created_at": "2019-12-20T12:43:32Z",
			"updated_at": "2023-01-31T17:47:31Z",
			"pushed_at": "2023-03-14T20:11:54Z",
			"git_url": "git://github.com/SpaceIm/conan-center-index.git",
			"ssh_url": "git@github.com:SpaceIm/conan-center-index.git",
			"clone_url": "https://github.com/SpaceIm/conan-center-index.git",
			"svn_url": "https://github.com/SpaceIm/conan-center-index",
			"homepage": "https://bintray.com/conan/conan-center",
			"size": 45518,
			"stargazers_count": 0,
			"watchers_count": 0,
			"language": "Python",
			"has_issues": false,
			"has_projects": true,
			"has_downloads": true,
			"has_wiki": true,
			"has_pages": false,
			"has_discussions": false,
			"forks_count": 0,
			"mirror_url": null,
			"archived": false,
			"disabled": false,
			"open_issues_count": 0,
			"license": {
			  "key": "mit",
			  "name": "MIT License",
			  "spdx_id": "MIT",
			  "url": "https://api.github.com/licenses/mit",
			  "node_id": "MDc6TGljZW5zZTEz"
			},
			"allow_forking": true,
			"is_template": false,
			"web_commit_signoff_required": false,
			"topics": [
	  
			],
			"visibility": "public",
			"forks": 0,
			"open_issues": 0,
			"watchers": 0,
			"default_branch": "master"
		  }
		},
		"base": {
		  "label": "conan-io:master",
		  "ref": "master",
		  "sha": "c37a2f18f60b8087ddc0948e7234be47e8c4f100",
		  "user": {
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
		  "repo": {
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
			"updated_at": "2023-03-11T20:13:08Z",
			"pushed_at": "2023-03-14T21:20:53Z",
			"git_url": "git://github.com/conan-io/conan-center-index.git",
			"ssh_url": "git@github.com:conan-io/conan-center-index.git",
			"clone_url": "https://github.com/conan-io/conan-center-index.git",
			"svn_url": "https://github.com/conan-io/conan-center-index",
			"homepage": "https://conan.io/center",
			"size": 44651,
			"stargazers_count": 704,
			"watchers_count": 704,
			"language": "Python",
			"has_issues": true,
			"has_projects": true,
			"has_downloads": true,
			"has_wiki": true,
			"has_pages": false,
			"has_discussions": true,
			"forks_count": 1234,
			"mirror_url": null,
			"archived": false,
			"disabled": false,
			"open_issues_count": 1622,
			"license": {
			  "key": "mit",
			  "name": "MIT License",
			  "spdx_id": "MIT",
			  "url": "https://api.github.com/licenses/mit",
			  "node_id": "MDc6TGljZW5zZTEz"
			},
			"allow_forking": true,
			"is_template": false,
			"web_commit_signoff_required": false,
			"topics": [
			  "conan",
			  "conan-center",
			  "conan-index",
			  "conan-packages",
			  "conan-recipe",
			  "cpp",
			  "cpp-library",
			  "dependencies",
			  "hacktoberfest",
			  "package-management",
			  "package-manager"
			],
			"visibility": "public",
			"forks": 1234,
			"open_issues": 1622,
			"watchers": 704,
			"default_branch": "master"
		  }
		},
		"_links": {
		  "self": {
			"href": "https://api.github.com/repos/conan-io/conan-center-index/pulls/16144"
		  },
		  "html": {
			"href": "https://github.com/conan-io/conan-center-index/pull/16144"
		  },
		  "issue": {
			"href": "https://api.github.com/repos/conan-io/conan-center-index/issues/16144"
		  },
		  "comments": {
			"href": "https://api.github.com/repos/conan-io/conan-center-index/issues/16144/comments"
		  },
		  "review_comments": {
			"href": "https://api.github.com/repos/conan-io/conan-center-index/pulls/16144/comments"
		  },
		  "review_comment": {
			"href": "https://api.github.com/repos/conan-io/conan-center-index/pulls/comments{/number}"
		  },
		  "commits": {
			"href": "https://api.github.com/repos/conan-io/conan-center-index/pulls/16144/commits"
		  },
		  "statuses": {
			"href": "https://api.github.com/repos/conan-io/conan-center-index/statuses/e2aa65c961d48d688dd5450811229eb1d62649ba"
		  }
		},
		"author_association": "CONTRIBUTOR",
		"auto_merge": null,
		"active_lock_reason": null,
		"merged": false,
		"mergeable": true,
		"rebaseable": true,
		"mergeable_state": "unstable",
		"merged_by": null,
		"comments": 4,
		"review_comments": 0,
		"maintainer_can_modify": true,
		"commits": 2,
		"additions": 5,
		"deletions": 7,
		"changed_files": 1
	  }`)

	review, _, err := client.PullRequest.GetReviewSummary(context.Background(), "conan-io", "conan-center-index", &reviewers, pr)
	assert.Equal(t, nil, err)

	const layout = "2006-01-02 15:04:05 -0700 MST"
	createdAt, err := time.Parse(layout, "2023-02-19 15:10:36 +0000 UTC") // This is the debug time from `%+v` formatter
	assert.Equal(t, nil, err)
	lastCommitAt, err := time.Parse(layout, "2023-02-19 15:10:08 +0000 UTC")
	assert.Equal(t, nil, err)
	submittedAt, err := time.Parse(layout, "2023-03-11 06:46:57 +0000 UTC")
	assert.Equal(t, nil, err)

	assert.Equal(t, &PullRequestSummary{
		Number:   16144,
		OpenedBy: "SpaceIm", CreatedAt: createdAt, Recipe: "re2", Change: EDIT, Weight: SMALL,
		ReviewURL:     "https://github.com/conan-io/conan-center-index/pull/16144",
		LastCommitSHA: "e2aa65c961d48d688dd5450811229eb1d62649ba", LastCommitAt: lastCommitAt, CciBotPassed: false,
		Summary: Reviews{Count: 2, ValidApprovals: 2, TeamApproval: true, Approvals: []Approver{{Name: "toge", Tier: Community}, {Name: "prince-chrismc", Tier: Team}},
			Blockers: nil, LastReview: &Review{ReviewerName: "prince-chrismc", SubmittedAt: submittedAt,
				HTMLURL: "https://github.com/conan-io/conan-center-index/pull/16144#pullrequestreview-1335829632",
			},
			IsBump: false,
		},
	}, review)

	assert.Equal(t, true, gock.IsDone())
}

func TestGetReviewSummary22576(t *testing.T) {
	client := NewClient(&http.Client{}, WorkingRepository{
		Owner: "prince-chrismc", Name: "conan-center-index-pending-review",
	})

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
	
	pr := parsePullRequestJSON(t, `{
		"url": "https://api.github.com/repos/conan-io/conan-center-index/pulls/22576",
		"id": 1699009947,
		"node_id": "PR_kwDODDMJAM5lRNWb",
		"html_url": "https://github.com/conan-io/conan-center-index/pull/22576",
		"diff_url": "https://github.com/conan-io/conan-center-index/pull/22576.diff",
		"patch_url": "https://github.com/conan-io/conan-center-index/pull/22576.patch",
		"issue_url": "https://api.github.com/repos/conan-io/conan-center-index/issues/22576",
		"number": 22576,
		"state": "open",
		"locked": false,
		"title": "fast_float: add version 6.1.0",
		"user": {
		  "login": "toge",
		  "id": 465629,
		  "node_id": "MDQ6VXNlcjQ2NTYyOQ==",
		  "avatar_url": "https://avatars.githubusercontent.com/u/465629?v=4",
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
		"body": "Specify library name and version:  **fast_float/6.1.0**\r\n\r\n---\r\n\r\n- [x] I've read the [contributing guidelines](https://github.com/conan-io/conan-center-index/blob/master/CONTRIBUTING.md).\r\n- [x] I've used a [recent](https://github.com/conan-io/conan/releases/latest) Conan client version close to the [currently deployed](https://github.com/conan-io/conan-center-index/blob/master/.c3i/config_v1.yml#L6).\r\n- [x] I've tried at least one configuration locally with the [conan-center hook](https://github.com/conan-io/hooks.git) activated.\r\n",
		"created_at": "2024-01-29T00:31:19Z",
		"updated_at": "2024-01-29T00:50:40Z",
		"closed_at": null,
		"merged_at": null,
		"merge_commit_sha": "e75e30de4d180523bc6c37b74036e5afb2a22401",
		"assignee": null,
		"assignees": [
	  
		],
		"requested_reviewers": [
	  
		],
		"requested_teams": [
	  
		],
		"labels": [
		  {
			"id": 1983649076,
			"node_id": "MDU6TGFiZWwxOTgzNjQ5MDc2",
			"url": "https://api.github.com/repos/conan-io/conan-center-index/labels/Bump%20version",
			"name": "Bump version",
			"color": "e5fc3a",
			"default": false,
			"description": "PR bumping version without recipe modifications"
		  }
		],
		"milestone": null,
		"draft": false,
		"commits_url": "https://api.github.com/repos/conan-io/conan-center-index/pulls/22576/commits",
		"review_comments_url": "https://api.github.com/repos/conan-io/conan-center-index/pulls/22576/comments",
		"review_comment_url": "https://api.github.com/repos/conan-io/conan-center-index/pulls/comments{/number}",
		"comments_url": "https://api.github.com/repos/conan-io/conan-center-index/issues/22576/comments",
		"statuses_url": "https://api.github.com/repos/conan-io/conan-center-index/statuses/bb02be24067706b1696de3c62d5b1b40b1d87567",
		"head": {
		  "label": "toge:fast_float-6.1.0",
		  "ref": "fast_float-6.1.0",
		  "sha": "bb02be24067706b1696de3c62d5b1b40b1d87567",
		  "user": {
			"login": "toge",
			"id": 465629,
			"node_id": "MDQ6VXNlcjQ2NTYyOQ==",
			"avatar_url": "https://avatars.githubusercontent.com/u/465629?v=4",
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
		  "repo": {
			"id": 341992172,
			"node_id": "MDEwOlJlcG9zaXRvcnkzNDE5OTIxNzI=",
			"name": "conan-center-index",
			"full_name": "toge/conan-center-index",
			"private": false,
			"owner": {
			  "login": "toge",
			  "id": 465629,
			  "node_id": "MDQ6VXNlcjQ2NTYyOQ==",
			  "avatar_url": "https://avatars.githubusercontent.com/u/465629?v=4",
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
			"html_url": "https://github.com/toge/conan-center-index",
			"description": "Recipes for the ConanCenter repository",
			"fork": true,
			"url": "https://api.github.com/repos/toge/conan-center-index",
			"forks_url": "https://api.github.com/repos/toge/conan-center-index/forks",
			"keys_url": "https://api.github.com/repos/toge/conan-center-index/keys{/key_id}",
			"collaborators_url": "https://api.github.com/repos/toge/conan-center-index/collaborators{/collaborator}",
			"teams_url": "https://api.github.com/repos/toge/conan-center-index/teams",
			"hooks_url": "https://api.github.com/repos/toge/conan-center-index/hooks",
			"issue_events_url": "https://api.github.com/repos/toge/conan-center-index/issues/events{/number}",
			"events_url": "https://api.github.com/repos/toge/conan-center-index/events",
			"assignees_url": "https://api.github.com/repos/toge/conan-center-index/assignees{/user}",
			"branches_url": "https://api.github.com/repos/toge/conan-center-index/branches{/branch}",
			"tags_url": "https://api.github.com/repos/toge/conan-center-index/tags",
			"blobs_url": "https://api.github.com/repos/toge/conan-center-index/git/blobs{/sha}",
			"git_tags_url": "https://api.github.com/repos/toge/conan-center-index/git/tags{/sha}",
			"git_refs_url": "https://api.github.com/repos/toge/conan-center-index/git/refs{/sha}",
			"trees_url": "https://api.github.com/repos/toge/conan-center-index/git/trees{/sha}",
			"statuses_url": "https://api.github.com/repos/toge/conan-center-index/statuses/{sha}",
			"languages_url": "https://api.github.com/repos/toge/conan-center-index/languages",
			"stargazers_url": "https://api.github.com/repos/toge/conan-center-index/stargazers",
			"contributors_url": "https://api.github.com/repos/toge/conan-center-index/contributors",
			"subscribers_url": "https://api.github.com/repos/toge/conan-center-index/subscribers",
			"subscription_url": "https://api.github.com/repos/toge/conan-center-index/subscription",
			"commits_url": "https://api.github.com/repos/toge/conan-center-index/commits{/sha}",
			"git_commits_url": "https://api.github.com/repos/toge/conan-center-index/git/commits{/sha}",
			"comments_url": "https://api.github.com/repos/toge/conan-center-index/comments{/number}",
			"issue_comment_url": "https://api.github.com/repos/toge/conan-center-index/issues/comments{/number}",
			"contents_url": "https://api.github.com/repos/toge/conan-center-index/contents/{+path}",
			"compare_url": "https://api.github.com/repos/toge/conan-center-index/compare/{base}...{head}",
			"merges_url": "https://api.github.com/repos/toge/conan-center-index/merges",
			"archive_url": "https://api.github.com/repos/toge/conan-center-index/{archive_format}{/ref}",
			"downloads_url": "https://api.github.com/repos/toge/conan-center-index/downloads",
			"issues_url": "https://api.github.com/repos/toge/conan-center-index/issues{/number}",
			"pulls_url": "https://api.github.com/repos/toge/conan-center-index/pulls{/number}",
			"milestones_url": "https://api.github.com/repos/toge/conan-center-index/milestones{/number}",
			"notifications_url": "https://api.github.com/repos/toge/conan-center-index/notifications{?since,all,participating}",
			"labels_url": "https://api.github.com/repos/toge/conan-center-index/labels{/name}",
			"releases_url": "https://api.github.com/repos/toge/conan-center-index/releases{/id}",
			"deployments_url": "https://api.github.com/repos/toge/conan-center-index/deployments",
			"created_at": "2021-02-24T18:12:59Z",
			"updated_at": "2023-01-31T18:45:30Z",
			"pushed_at": "2024-01-29T09:40:22Z",
			"git_url": "git://github.com/toge/conan-center-index.git",
			"ssh_url": "git@github.com:toge/conan-center-index.git",
			"clone_url": "https://github.com/toge/conan-center-index.git",
			"svn_url": "https://github.com/toge/conan-center-index",
			"homepage": "https://conan.io/center",
			"size": 70352,
			"stargazers_count": 0,
			"watchers_count": 0,
			"language": "Python",
			"has_issues": false,
			"has_projects": true,
			"has_downloads": true,
			"has_wiki": true,
			"has_pages": false,
			"has_discussions": false,
			"forks_count": 0,
			"mirror_url": null,
			"archived": false,
			"disabled": false,
			"open_issues_count": 1,
			"license": {
			  "key": "mit",
			  "name": "MIT License",
			  "spdx_id": "MIT",
			  "url": "https://api.github.com/licenses/mit",
			  "node_id": "MDc6TGljZW5zZTEz"
			},
			"allow_forking": true,
			"is_template": false,
			"web_commit_signoff_required": false,
			"topics": [
	  
			],
			"visibility": "public",
			"forks": 0,
			"open_issues": 1,
			"watchers": 0,
			"default_branch": "master"
		  }
		},
		"base": {
		  "label": "conan-io:master",
		  "ref": "master",
		  "sha": "44518a631b2a025bb485b9231720b019b5c6cd29",
		  "user": {
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
		  "repo": {
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
			"updated_at": "2024-01-26T03:54:11Z",
			"pushed_at": "2024-01-29T09:54:31Z",
			"git_url": "git://github.com/conan-io/conan-center-index.git",
			"ssh_url": "git@github.com:conan-io/conan-center-index.git",
			"clone_url": "https://github.com/conan-io/conan-center-index.git",
			"svn_url": "https://github.com/conan-io/conan-center-index",
			"homepage": "https://conan.io/center",
			"size": 51292,
			"stargazers_count": 862,
			"watchers_count": 862,
			"language": "Python",
			"has_issues": true,
			"has_projects": true,
			"has_downloads": true,
			"has_wiki": true,
			"has_pages": false,
			"has_discussions": true,
			"forks_count": 1516,
			"mirror_url": null,
			"archived": false,
			"disabled": false,
			"open_issues_count": 2281,
			"license": {
			  "key": "mit",
			  "name": "MIT License",
			  "spdx_id": "MIT",
			  "url": "https://api.github.com/licenses/mit",
			  "node_id": "MDc6TGljZW5zZTEz"
			},
			"allow_forking": true,
			"is_template": false,
			"web_commit_signoff_required": false,
			"topics": [
			  "conan",
			  "conan-center",
			  "conan-index",
			  "conan-packages",
			  "conan-recipe",
			  "cpp",
			  "cpp-library",
			  "dependencies",
			  "hacktoberfest",
			  "package-management",
			  "package-manager"
			],
			"visibility": "public",
			"forks": 1516,
			"open_issues": 2281,
			"watchers": 862,
			"default_branch": "master"
		  }
		},
		"_links": {
		  "self": {
			"href": "https://api.github.com/repos/conan-io/conan-center-index/pulls/22576"
		  },
		  "html": {
			"href": "https://github.com/conan-io/conan-center-index/pull/22576"
		  },
		  "issue": {
			"href": "https://api.github.com/repos/conan-io/conan-center-index/issues/22576"
		  },
		  "comments": {
			"href": "https://api.github.com/repos/conan-io/conan-center-index/issues/22576/comments"
		  },
		  "review_comments": {
			"href": "https://api.github.com/repos/conan-io/conan-center-index/pulls/22576/comments"
		  },
		  "review_comment": {
			"href": "https://api.github.com/repos/conan-io/conan-center-index/pulls/comments{/number}"
		  },
		  "commits": {
			"href": "https://api.github.com/repos/conan-io/conan-center-index/pulls/22576/commits"
		  },
		  "statuses": {
			"href": "https://api.github.com/repos/conan-io/conan-center-index/statuses/bb02be24067706b1696de3c62d5b1b40b1d87567"
		  }
		},
		"author_association": "CONTRIBUTOR",
		"auto_merge": null,
		"active_lock_reason": null,
		"merged": false,
		"mergeable": true,
		"rebaseable": true,
		"mergeable_state": "clean",
		"merged_by": null,
		"comments": 1,
		"review_comments": 0,
		"maintainer_can_modify": true,
		"commits": 1,
		"additions": 5,
		"deletions": 0,
		"changed_files": 2
	  }`)

	review, _, err := client.PullRequest.GetReviewSummary(context.Background(), "conan-io", "conan-center-index", &reviewers, pr)
	assert.Equal(t, nil, err)
	
	const layout = "2006-01-02 15:04:05 -0700 MST"
	createdAt, err := time.Parse(layout, "2024-01-29 00:31:19 +0000 UTC") // This is the debug time from `%+v` formatter
	assert.Equal(t, nil, err)
	lastCommitAt, err := time.Parse(layout, "2024-01-29 00:23:55 +0000 UTC")
	assert.Equal(t, nil, err)

	assert.Equal(t, &PullRequestSummary{
		Number:   22576,
		OpenedBy: "toge", CreatedAt: createdAt, Recipe: "fast_float", Change: EDIT, Weight: TINY,
		ReviewURL:     "https://github.com/conan-io/conan-center-index/pull/22576",
		LastCommitSHA: "bb02be24067706b1696de3c62d5b1b40b1d87567", LastCommitAt: lastCommitAt, CciBotPassed: true,
		Summary: Reviews{Count: 0, ValidApprovals: 0, TeamApproval: false, Approvals: nil,
			Blockers: nil, LastReview: nil,
			IsBump: true,
		},
	}, review)
}
