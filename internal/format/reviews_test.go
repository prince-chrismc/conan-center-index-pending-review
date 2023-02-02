package format

import (
	"encoding/json"
	"testing"

	"github.com/prince-chrismc/conan-center-index-pending-review/v4/pending_review"
	"github.com/stretchr/testify/assert"
)

func TestFormatTitles(t *testing.T) {
	assert.Equal(t, title(pending_review.NEW, "new-recipe"), ":new: new-recipe")
	assert.Equal(t, title(pending_review.EDIT, "edit-recipe"), ":memo: edit-recipe")
	assert.Equal(t, title(pending_review.DOCS, "docs"), ":green_book: docs")
	assert.Equal(t, title(pending_review.CONFIG, ".github"), ":gear: .github")
}

func TestFormatMarkdownRows(t *testing.T) {
	var rs []*pending_review.PullRequestSummary
	if err := json.Unmarshal([]byte(`[
		{
			"Number": 4556,
			"OpenedBy": "anton-danielsson",
			"Recipe": "protobuf",
			"Change": 1,
			"ReviewURL": "https://github.com/conan-io/conan-center-index/pull/4556",
			"LastCommitSHA": "6a14a091f3b63f0f7039520d03627c607e58f770",
			"LastCommitAt": "0001-01-01T00:00:00Z",
			"CciBotPassed": true,
			"Summary": {
				"Count": 36,
				"ValidApprovals": 1,
				"TeamApproval": false,
				"Approvals": [
					"prince-chrismc"
				],
				"Blockers": [
					"uilianries"
				],
				"LastReview": {
					"ReviewerName": "madebr",
					"SubmittedAt": "2021-04-09T23:49:10Z",
					"HTMLURL": "https://github.com/conan-io/conan-center-index/pull/4356#pullrequestreview-642778787"
				}
			}
		},
		{
			"Number": 4356,
			"OpenedBy": "prince-chrismc",
			"CreatedAt": "2021-01-25T15:14:40Z",
			"Recipe": "paho-mqtt-c",
			"Change": 1,
			"ReviewURL": "https://github.com/conan-io/conan-center-index/pull/4356",
			"LastCommitSHA": "f61a8a0b0c4171d8935fc5047c714b6761343346",
			"LastCommitAt": "0001-01-01T00:00:00Z",
			"CciBotPassed": true,
			"Summary": {
				"Count": 15,
				"ValidApprovals": 3,
				"TeamApproval": true,
				"Approvals": [
					"madebr", "SSE4", "SpaceIm"
				],
				"Blockers": []
			}
		}
	]`), &rs); err != nil {
		t.Fatal("Broken test - invalid JSON content:", err)
	}

	mergeRow, mergeCount := ReviewsToMarkdownRows(rs, true)
	assert.Equal(t, mergeCount, 1)
	assert.Equal(t, "[#4356](https://github.com/conan-io/conan-center-index/pull/4356)|[prince-chrismc](https://github.com/prince-chrismc)|Jan 25|:memo: paho-mqtt-c|15|madebr, SSE4, SpaceIm\n", mergeRow)

	reviewRow, reviewCount := ReviewsToMarkdownRows(rs, false)
	assert.Equal(t, reviewCount, 1)
	assert.Equal(t, "[#4556](https://github.com/conan-io/conan-center-index/pull/4556)|[anton-danielsson](https://github.com/anton-danielsson)|Jan 1|:memo: protobuf|:green_circle: XS|36|Apr 9 :bell:|uilianries|prince-chrismc\n", reviewRow)
}

func TestFormatMarkdownRowsDocs(t *testing.T) {
	var rs []*pending_review.PullRequestSummary
	if err := json.Unmarshal([]byte(`[
		{
			"Number": 7648,
			"OpenedBy": "jgsogo",
			"CreatedAt": "2021-10-11T10:45:20Z",
			"Recipe": "docs",
			"Change": 2,
			"ReviewURL": "https://github.com/conan-io/conan-center-index/pull/7648",
			"LastCommitSHA": "e9457d1319b4cdb57b732c54cc9e61db8adb398a",
			"LastCommitAt": "2021-10-11T10:39:04Z",
			"CciBotPassed": false,
			"Summary": {
			   "Count": 3,
			   "ValidApprovals": 3,
			   "TeamApproval": true,
			   "Approvals": [
				  "SSE4",
				  "uilianries",
				  "prince-chrismc"
			   ],
			   "Blockers": null,
			   "LastReview": {
				  "ReviewerName": "prince-chrismc",
				  "SubmittedAt": "2021-10-11T23:11:18Z",
				  "HTMLURL": "https://github.com/conan-io/conan-center-index/pull/7648#pullrequestreview-776761950"
			   }
			}
		 }
	]`), &rs); err != nil {
		t.Fatal("Broken test - invalid JSON content:", err)
	}

	mergeRow, mergeCount := ReviewsToMarkdownRows(rs, true)
	assert.Equal(t, mergeCount, 1)
	assert.Equal(t, "[#7648](https://github.com/conan-io/conan-center-index/pull/7648)|[jgsogo](https://github.com/jgsogo)|Oct 11|:green_book: docs|3|SSE4, uilianries, prince-chrismc\n", mergeRow)
}

func TestFormatMarkdownRowsCiPending(t *testing.T) {
	var rs []*pending_review.PullRequestSummary
	if err := json.Unmarshal([]byte(`[
		{
			"Number": 8557,
			"OpenedBy": "daravi",
			"CreatedAt": "2021-12-28T02:27:59Z",
			"Recipe": "libkmod",
			"Change": 0,
			"Weight": 6,
			"ReviewURL": "https://github.com/conan-io/conan-center-index/pull/8557",
			"LastCommitSHA": "88e25c1a89cbc4b130e37fb3d42fe7e16cf3b4ca",
			"LastCommitAt": "2021-12-28T02:26:17Z",
			"CciBotPassed": false,
			"Summary": {
				"Count": 0,
				"ValidApprovals": 0,
				"TeamApproval": false,
				"Approvals": [],
				"Blockers": [],
				"LastReview": null
			}
		}
	]`), &rs); err != nil {
		t.Fatal("Broken test - invalid JSON content:", err)
	}

	mergeRow, mergeCount := ReviewsToMarkdownRows(rs, false)
	assert.Equal(t, mergeCount, 1)
	assert.Equal(t, "[#8557](https://github.com/conan-io/conan-center-index/pull/8557)|[daravi](https://github.com/daravi)|Dec 28|:stopwatch: libkmod|???|0|:eyes:||\n", mergeRow)
}

func TestFormatMarkdownRowsCiSuccess(t *testing.T) {
	var rs []*pending_review.PullRequestSummary
	if err := json.Unmarshal([]byte(`[
		{
			"Number": 8557,
			"OpenedBy": "daravi",
			"CreatedAt": "2021-12-28T02:27:59Z",
			"Recipe": "libkmod",
			"Change": 0,
			"Weight": 2,
			"ReviewURL": "https://github.com/conan-io/conan-center-index/pull/8557",
			"LastCommitSHA": "88e25c1a89cbc4b130e37fb3d42fe7e16cf3b4ca",
			"LastCommitAt": "2021-12-28T02:26:17Z",
			"CciBotPassed": true,
			"Summary": {
				"Count": 0,
				"ValidApprovals": 0,
				"TeamApproval": false,
				"Approvals": [],
				"Blockers": [],
				"LastReview": null
			}
		}
	]`), &rs); err != nil {
		t.Fatal("Broken test - invalid JSON content:", err)
	}

	mergeRow, mergeCount := ReviewsToMarkdownRows(rs, false)
	assert.Equal(t, mergeCount, 1)
	assert.Equal(t, "[#8557](https://github.com/conan-io/conan-center-index/pull/8557)|[daravi](https://github.com/daravi)|Dec 28|:new: libkmod|M|0|:eyes:||\n", mergeRow)
}
