package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/rxanders35/controlplane/server/internal/api" // Assuming module path
	"github.com/rxanders35/controlplane/server/internal/k8s" // Assuming module path
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// 1. Initialize Kubernetes Clients
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Could not determine home dir: %v", err)
	}
	kubeconfig := filepath.Join(homedir, ".kube", "config")

	cfg, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatalf("Failed to build config from ~/.kube/config: %v", err)
	}

	k8sClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize standard client set: %v", err)
	}
	dynamicClient, err := dynamic.NewForConfig(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize dynamic client: %v", err)
	}

	// 2. Wire Up Dependencies
	k8sService := k8s.NewService(k8sClient, dynamicClient)
	apiServer := api.NewServer(k8sService)

	// 3. Start the Server
	log.Println("Starting API server on :8080")
	if err := apiServer.Start(":8080"); err != nil {
		log.Fatalf("Failed to start API server: %v", err)
	}
}
