package zenhub

import (
	"context"
	"fmt"
)

type Epic struct {
	TotalEpicEstimates struct {
		Value int64 `json:"value"`
	} `json:"total_epic_estimates"`
	Estimate struct {
		Value int64 `json:"value"`
	} `json:"estimate"`
	Pipeline  *Pipeline  `json:"pipeline"`
	Pipelines []Pipeline `json:"pipelines"`
	Issues    []Issue    `json:"issues"`
}

func (c *Client) GetEpic(ctx context.Context, repoID, epicID string) (*Epic, error) {
	path := fmt.Sprintf("/p1/repositories/%s/epics/%s", repoID, epicID)

	var epic Epic
	if err := c.get(path, &epic); err != nil {
		return nil, err
	}

	return &epic, nil
}