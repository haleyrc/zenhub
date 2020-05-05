package zenhub

import (
	"context"
	"fmt"
)

func NewIssuesService(c *Client) *IssuesService {
	return &IssuesService{c: c}
}

type IssuesService struct {
	c *Client
}

type Issue struct {
	IssueNumber int64      `json:"issue_number"`
	IsEpic      bool       `json:"is_epic"`
	RepoID      int64      `json:"repo_id"`
	Estimate    Int64Value `json:"estimate"`
	Pipelines   []Pipeline `json:"pipelines"`
	Pipeline    *Pipeline  `json:"pipeline"`
}

func (s *IssuesService) GetIssue(ctx context.Context, repo, id int64) (*Issue, error) {
	path := fmt.Sprintf("/p1/repositories/%d/issues/%d", repo, id)

	var issue Issue
	if err := s.c.get(path, &issue); err != nil {
		return nil, err
	}

	return &issue, nil
}
