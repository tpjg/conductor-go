package conductor_http_client

import (
	"context"
	"testing"

	"github.com/conductor-sdk/conductor-go/pkg/conductor_client/conductor_http_client"
	"github.com/conductor-sdk/conductor-go/tests/e2e/conductor_client"
)

func TestUpdateTaskRefByName(t *testing.T) {
	workflowName := "workflow_with_go_task_example_from_code"
	taskReferenceName := "go_task_example_from_code_ref_0"
	workflowId, err := StartWorkflow(workflowName)
	if err != nil {
		t.Error(err)
	}
	apiClient := conductor_http_client.NewAPIClient(
		conductor_client.GetAuthenticationSettings(),
		conductor_client.GetHttpSettingsWithAuth(),
	)
	taskClient := *&conductor_http_client.TaskResourceApiService{
		APIClient: apiClient,
	}
	_, _, err = taskClient.UpdateTaskByRefName(
		context.Background(),
		map[string]interface{}{"hello": "world"},
		workflowId,
		taskReferenceName,
		"COMPLETED",
	)
	if err != nil {
		t.Error(err)
	}
	// TODO check response and workflow task status
}
