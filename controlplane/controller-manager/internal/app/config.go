package app

import (
	"os"

	"github.com/reidx/dap/controlplane/pkg/apis"
)

type RuntimeConfig struct {
	LeaderElectionID string
	WatchNamespace   string
	DefaultRunPhase  apis.SparkRunPhase
}

func LoadRuntimeConfig() RuntimeConfig {
	ns := os.Getenv("WATCH_NAMESPACE")
	if ns == "" {
		ns = ""
	}

	leaderID := os.Getenv("LEADER_ELECTION_ID")
	if leaderID == "" {
		leaderID = "dap-controller-manager"
	}

	return RuntimeConfig{
		LeaderElectionID: leaderID,
		WatchNamespace:   ns,
		DefaultRunPhase:  apis.SparkRunPhaseQueued,
	}
}
