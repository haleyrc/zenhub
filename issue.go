package zenhub

type Issue struct {
	IssueNumber int64      `json:"issue_number"`
	IsEpic      bool       `json:"is_epic"`
	RepoID      int64      `json:"repo_id"`
	Estimate    Int64Value `json:"estimate"`
	Pipelines   []Pipeline `json:"pipelines"`
	Pipeline    *Pipeline  `json:"pipeline"`
}
