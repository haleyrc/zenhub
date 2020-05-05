package zenhub_test

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	"github.com/haleyrc/zenhub"
)

func TestGetEpics(t *testing.T) {
	ctx := context.Background()
	id := int64(172137510)
	c := zenhub.New(os.Getenv("ZENHUB_TOKEN"))
	epicsService := zenhub.NewEpicsService(c)

	issues, err := epicsService.GetEpics(ctx, id)
	if err != nil {
		t.Fatal(err)
	}

	// enc := json.NewEncoder(os.Stdout)
	// enc.SetIndent("", "    ")
	// enc.Encode(issues)

	epics := []*zenhub.Epic{}
	for _, issue := range issues {
		epic, err := epicsService.GetEpic(ctx, issue.RepoID, issue.IssueNumber)
		if err != nil {
			continue
		}
		epics = append(epics, epic)
	}
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "    ")
	enc.Encode(epics)
}
