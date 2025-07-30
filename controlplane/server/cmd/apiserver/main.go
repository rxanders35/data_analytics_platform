package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/rxanders35/controlplane/server/internal/api"
	"github.com/rxanders35/controlplane/server/internal/kubernetes"
	"k8s.io/client-go/dynamic"
	k8s "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Could not determine home dir: %v", err)
	}
	kubeconfig := filepath.Join(homedir, ".kube", "config")

	cfg, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatalf("Failed to build config from ~/.kube/config: %v", err)
	}

	k8sClient, err := k8s.NewForConfig(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize standard client set: %v", err)
	}
	dynamicClient, err := dynamic.NewForConfig(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize dynamic client: %v", err)
	}

	k8sOrchestrator := kubernetes.NewOrchestrator(k8sClient, dynamicClient)

	apiServer := api.NewServer(k8sOrchestrator)

	log.Println("Starting API server on :8080")
	if err := apiServer.Start(":8080"); err != nil {
		log.Fatalf("Failed to start API server: %v", err)
	}
}
