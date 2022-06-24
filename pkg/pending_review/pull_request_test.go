package pending_review

import (
	"encoding/json"
	"testing"

	"github.com/google/go-github/v42/github"
	"github.com/stretchr/testify/assert"
)

func parsePrJSON(t *testing.T, str string) []*github.CommitFile {
	var files []*github.CommitFile

	if err := json.Unmarshal([]byte(str), &files); err != nil {
		t.Fatal()
	}

	return files
}

func TestFalseIsFalse(t *testing.T) {
	oneFile := parsePrJSON(t, `[
		{
		  "sha": "5cbce65d888e970205160de1ea33cb3dae4b948b",
		  "filename": "recipes/b2/portable/conandata.yml",
		  "status": "modified",
		  "additions": 3,
		  "deletions": 0,
		  "changes": 3,
		  "blob_url": "https://github.com/conan-io/conan-center-index/blob/7558ff23fa9eabd5ae08e90b89abc125f4a557e4/recipes/b2/portable/conandata.yml",
		  "raw_url": "https://github.com/conan-io/conan-center-index/raw/7558ff23fa9eabd5ae08e90b89abc125f4a557e4/recipes/b2/portable/conandata.yml",
		  "contents_url": "https://api.github.com/repos/conan-io/conan-center-index/contents/recipes/b2/portable/conandata.yml?ref=7558ff23fa9eabd5ae08e90b89abc125f4a557e4",
		  "patch": "@@ -17,3 +17,6 @@ sources:\n   \"4.6.0\":\n     url: \"https://github.com/bfgroup/b2/archive/4.6.0.tar.gz\"\n     sha256: \"3a308e0f79a039d8a9495b375f3292f5163000c19caa79c5687e4cb5b1938b49\"\n+  \"4.6.1\":\n+    url: \"https://github.com/bfgroup/b2/archive/4.6.1.tar.gz\"\n+    sha256: \"a3f3323eaeb2c27d7a3ca86842665c6c3bc3d93cc626ba362ae6d0c5a7bfbe2c\""
		}
	  ]`)

	assert.Equal(t, false, false)