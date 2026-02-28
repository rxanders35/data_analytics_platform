package main

import (
	"log"
	"time"

	"github.com/reidx/dap/controlplane/controller-manager/internal/app"
)

func main() {
	cfg := app.LoadRuntimeConfig()
	log.Printf("controller-manager starting leaderElectionID=%s watchNamespace=%q", cfg.LeaderElectionID, cfg.WatchNamespace)

	// Step 3 scaffold: Step 7 adds controller-runtime manager startup.
	for {
		time.Sleep(30 * time.Second)
		log.Printf("controller-manager heartbeat defaultSparkRunPhase=%s", cfg.DefaultRunPhase)
	}
}
