package zenhub

type Pipeline struct {
	WorkspaceID string `json:"workspace_id"`
	Name        string `json:"name"`
	PipelineID  string `json:"pipeline_id"`
}
