package app

import "testing"

func TestLoadRuntimeConfigDefaults(t *testing.T) {
	t.Setenv("WATCH_NAMESPACE", "")
	t.Setenv("LEADER_ELECTION_ID", "")

	cfg := LoadRuntimeConfig()
	if cfg.LeaderElectionID != "dap-controller-manager" {
		t.Fatalf("unexpected leader election id: %s", cfg.LeaderElectionID)
	}
	if cfg.DefaultRunPhase != "Queued" {
		t.Fatalf("unexpected default run phase: %s", cfg.DefaultRunPhase)
	}
}
