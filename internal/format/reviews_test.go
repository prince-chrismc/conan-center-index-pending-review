package format

import (
	"encoding/json"
	"testing"

	"github.com/prince-chrismc/conan-center-index-pending-review/v2/pkg/pending_review"
	"github.com/stretchr/testify/assert"
)

func TestFormatTitles(t *testing.T) {
	assert.Equal(t, title(pending_review.ADDED, "new-recipe"), ":new: new-recipe")
	assert.Equal(t, title(pending_review.EDIT, "edit-recipe"), ":memo: edit-recipe")
	assert.Equal(t, title(pending_review.BUMP, "bump-recipe"), ":arrow_up: bump-recipe")
	assert.Equal(t, title(pending_review.DOCS, "docs"), ":green_book: docs")
}

func TestFormatMarkdownRows(t *testing.T) {
	var rs []*pending_review.ReviewSummary
	if err := json.Unmarshal([]byte(`[
		{
		   "Number": 4556,
		   "OpenedBy": "anton-danielsson",
		   "Recipe": "protobuf",
		   "Change": 1,
		   "ReviewURL": "https://github.com/conan-io/conan-center-index/pull/4556",
		   "LastCommitSHA": "6a14a091f3b63f0f7039520d03627c607e58f770",
		   "LastCommitAt": "0001-01-01T00:00:00Z",
		   "Reviews": 36,
		   "ValidApprovals": 1,
		   "IsMergeable": false,
		   "CciBotPassed": true,
		   "Approvals": [
			  "prince-chrismc"
		   ],
		   "Blockers": null
		},
		{
		   "Number": 4682,
		   "OpenedBy": "floriansimon1",
		   "Recipe": "protobuf",
		   "Change": 1,
		   "ReviewURL": "https://github.com/conan-io/conan-center-index/pull/4682",
		   "LastCommitSHA": "8b0f82031d2dd5099e33bea3cece524f084950f3",
		   "LastCommitAt": "0001-01-01T00:00:00Z",
		   "Reviews": 13,
		   "ValidApprovals": 1,
		   "IsMergeable": true,
		   "CciBotPassed": false,
		   "Approvals": [
			  "prince-chrismc"
		   ],
		   "Blockers": [
			  "uilianries"
		   ]
		}
	]`), &rs); err != nil {
		t.Fatal("Broken test - invalid JSON content:", err)
	}

	assert.Equal(t, ReviewsToMarkdownRows(rs, false), "[#4556](https://github.com/conan-io/conan-center-index/pull/4556)|[anton-danielsson](https://github.com/anton-danielsson)|:memo: protobuf|36||prince-chrismc\n")
	assert.Equal(t, ReviewsToMarkdownRows(rs, true), "[#4682](https://github.com/conan-io/conan-center-index/pull/4682)|[floriansimon1](https://github.com/floriansimon1)|:warning: protobuf|13|uilianries|prince-chrismc\n")
}
