package pending_review

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func parseSummaryJSON(t *testing.T, str string) *PullRequestSummary {
	var summary *PullRequestSummary

	if err := json.Unmarshal([]byte(str), &summary); err != nil {
		t.Fatal(err)
	}

	return summary
}

func TestEvaluateSummary(t *testing.T) {
	summary := parseSummaryJSON(t, `{"Number":16144,"OpenedBy":"SpaceIm","CreatedAt":"2023-02-19T15:10:36Z","Recipe":"re2","Change":1,"Weight":1,"ReviewURL":"https://github.com/conan-io/conan-center-index/pull/16144","LastCommitSHA":"e2aa65c961d48d688dd5450811229eb1d62649ba","LastCommitAt":"2023-02-19T15:10:08Z","CciBotPassed":true,"Summary":{"Count":2,"ValidApprovals":2,"TeamApproval":true,"Approvals":[{"Name":"toge","Tier":"community"},{"Name":"prince-chrismc","Tier":"community"}],"Blockers":null,"LastReview":{"ReviewerName":"prince-chrismc","SubmittedAt":"2023-03-11T06:46:57Z","HTMLURL":"https://github.com/conan-io/conan-center-index/pull/16144#pullrequestreview-1335829632"},"IsBump": false}}`)

	result := evaluateSummary(summary)
	assert.Equal(t, nil, result)
}

func TestEvaluateSummaryDocs(t *testing.T) {
	summary := parseSummaryJSON(t, `{"Number":16592,"OpenedBy":"prince-chrismc","CreatedAt":"2023-03-17T01:52:09Z","Recipe":"docs","Change":2,"Weight":0,"ReviewURL":"https://github.com/conan-io/conan-center-index/pull/16592","LastCommitSHA":"b27854be3d789d8e1303e899f84a996215bbd7ac","LastCommitAt":"2023-03-31T20:11:41Z","CciBotPassed":true,"Summary":{"Count":8,"ValidApprovals":0,"TeamApproval":false,"Approvals":null,"Blockers":null,"LastReview":{"ReviewerName":"SSE4","SubmittedAt":"2023-04-01T07:23:12Z","HTMLURL":"https://github.com/conan-io/conan-center-index/pull/16592#pullrequestreview-1367823599"},"IsBump": false}}`)

	result := evaluateSummary(summary)
	assert.Equal(t, nil, result)
}

func TestEvaluationSummary14707_1(t *testing.T) {
	// This is using the `"LastCommitAt": "<incorrect value>"`
	summary := parseSummaryJSON(t, `{"Number":14703,"OpenedBy":"bennyhuo","CreatedAt":"2022-12-12T23:10:29Z","Recipe":"tinycthreadpool","Change":0,"Weight":2,"ReviewURL":"https://github.com/conan-io/conan-center-index/pull/14703","LastCommitSHA":"6b173fd061c77e5eb51990f372d9c138f14bd7fa",
	"LastCommitAt":"2022-12-12T22:45:56Z","CciBotPassed":true,"Summary":{"Count":3,"ValidApprovals":0,"TeamApproval":false,"Approvals":null,"Blockers":null,
	"LastReview":{"ReviewerName":"uilianries","SubmittedAt":"2023-01-20T12:46:25Z","HTMLURL":"https://github.com/conan-io/conan-center-index/pull/14703#pullrequestreview-1263503119"},"IsBump": false}}`)

	err := evaluateSummary(summary)
	assert.EqualError(t, err, ErrWorkRequired.Error())
}

func TestEvaluationSummary14707_2(t *testing.T) {
	// This is using the correct value for "LastCommitAt"
	summary := parseSummaryJSON(t, `{"Number":14703,"OpenedBy":"bennyhuo","CreatedAt":"2022-12-12T23:10:29Z","Recipe":"tinycthreadpool","Change":0,"Weight":2,"ReviewURL":"https://github.com/conan-io/conan-center-index/pull/14703","LastCommitSHA":"6b173fd061c77e5eb51990f372d9c138f14bd7fa",
	"LastCommitAt":"2023-03-27T23:41:44Z","CciBotPassed":true,"Summary":{"Count":3,"ValidApprovals":0,"TeamApproval":false,"Approvals":null,"Blockers":null,
	"LastReview":{"ReviewerName":"uilianries","SubmittedAt":"2023-01-20T12:46:25Z","HTMLURL":"https://github.com/conan-io/conan-center-index/pull/14703#pullrequestreview-1263503119"},"IsBump": false}}`)

	result := evaluateSummary(summary)
	assert.Equal(t, nil, result)
}

func TestEvaluationSummary14707_3(t *testing.T) {
	// After I reviewed it
	summary := parseSummaryJSON(t, `{"Number":14703,"OpenedBy":"bennyhuo","CreatedAt":"2022-12-12T23:10:29Z","Recipe":"tinycthreadpool","Change":0,"Weight":2,"ReviewURL":"https://github.com/conan-io/conan-center-index/pull/14703","LastCommitSHA":"6b173fd061c77e5eb51990f372d9c138f14bd7fa","LastCommitAt":"2023-03-27T23:41:44Z","CciBotPassed":true,"Summary":{"Count":5,"ValidApprovals":0,"TeamApproval":false,"Approvals":null,"Blockers":null,"LastReview":{"ReviewerName":"prince-chrismc","SubmittedAt":"2023-04-07T21:13:24Z","HTMLURL":"https://github.com/conan-io/conan-center-index/pull/14703#pullrequestreview-1376588362"},"IsBump": false}}`)

	err := evaluateSummary(summary)
	assert.EqualError(t, err, ErrWorkRequired.Error())
}
