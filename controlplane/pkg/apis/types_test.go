package apis

import "testing"

func TestWorkspaceReadyPhaseValue(t *testing.T) {
	if WorkspacePhaseReady != "Ready" {
		t.Fatalf("expected WorkspacePhaseReady=Ready, got %q", WorkspacePhaseReady)
	}
}
