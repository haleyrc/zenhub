package zenhub

import (
	"context"
	"fmt"
)

func NewEpicsService(c *Client) *EpicsService {
	return &EpicsService{c: c}
}

type EpicsService struct {
	c *Client
}

type Epic struct {
	Issues             []Issue    `json:"issues"`
	Estimate           Int64Value `json:"estimate"`
	Pipelines          []Pipeline `json:"pipelines"`
	TotalEpicEstimates Int64Value `json:"total_epic_estimates"`
	Pipeline           *Pipeline  `json:"pipeline"`
}

type EpicIssue struct {
	IssueNumber int64  `json:"issue_number"`
	RepoID      int64  `json:"repo_id"`
	IssueURL    string `json:"issue_url"`
}

func (s *EpicsService) GetEpic(ctx context.Context, repo, id int64) (*Epic, error) {
	path := fmt.Sprintf("/p1/repositories/%d/epics/%d", repo, id)

	var epic Epic
	if err := s.c.get(path, &epic); err != nil {
		return nil, err
	}

	return &epic, nil
}

func (s *EpicsService) GetEpics(ctx context.Context, repo int64) ([]*EpicIssue, error) {
	path := fmt.Sprintf("/p1/repositories/%d/epics", repo)

	var response struct {
		EpicIssues []*EpicIssue `json:"epic_issues"`
	}
	if err := s.c.get(path, &response); err != nil {
		return nil, err
	}

	return response.EpicIssues, nil
}
