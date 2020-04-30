package zenhub

import (
	"context"
	"fmt"
)

type EpicIssue struct {
	IssueNumber int64  `json:"issue_number"`
	RepoID      int64  `json:"repo_id"`
	IssueURL    string `json:"issue_url"`
}

func (c *Client) GetEpicsForRepository(ctx context.Context, id string) ([]EpicIssue, error) {
	path := fmt.Sprintf("/p1/repositories/%s/epics", id)

	var issues []EpicIssue
	if err := c.get(path, &issues); err != nil {
		return nil, err
	}

	return issues, nil
}
