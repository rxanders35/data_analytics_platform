package k8s

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
)

type Service struct {
	k8sClient     kubernetes.Interface
	dynamicClient dynamic.Interface
}

func NewService(k8sClient kubernetes.Interface, dynClient dynamic.Interface) *Service {
	return &Service{
		k8sClient:     k8sClient,
		dynamicClient: dynClient,
	}
}

type PySparkJobRequest struct {
	Name                string   `json:"name"`
	Mode                string   `json:"mode"`
	Image               string   `json:"image"`
	SparkVersion        string   `json:"sparkVersion"`
	MainApplicationFile string   `json:"mainApplicationFile"`
	Arguments           []string `json:"arguments"`
	Driver              Driver   `json:"driver"`
	Executor            Executor `json:"executor"`
}

type Driver struct {
	Memory         string `json:"memory"`
	Cores          int32  `json:"cores"`
	ServiceAccount string `json:"serviceAccount"`
}

type Executor struct {
	Memory    string `json:"memory"`
	Cores     int32  `json:"cores"`
	Instances int32  `json:"instances"`
}

func (s *Service) ListPods(ctx context.Context) ([]string, error) {
	podlist, err := s.k8sClient.CoreV1().Pods("default").List(ctx, v1.ListOptions{})
	if err != nil {
		return nil, err
	}

	pods := make([]string, 0, len(podlist.Items))
	for _, pod := range podlist.Items {
		pods = append(pods, pod.Name)
	}

	return pods, nil
}

func (s *Service) SubmitSparkJob(ctx context.Context, req PySparkJobRequest) error {
	sparkApp := &unstructured.Unstructured{
		Object: map[string]any{
			"apiVersion": "sparkoperator.k8s.io/v1beta2",
			"kind":       "SparkApplication",
			"metadata": map[string]any{
				"name":      req.Name,
				"namespace": "default",
			},
			"spec": map[string]any{
				"type":                "Python",
				"pythonVersion":       "3",
				"mode":                req.Mode,
				"image":               req.Image,
				"sparkVersion":        req.SparkVersion,
				"mainApplicationFile": req.MainApplicationFile,
				"arguments":           req.Arguments,
				"restartPolicy": map[string]any{
					"type": "Never",
				},
				"driver": map[string]any{
					"memory":         req.Driver.Memory,
					"cores":          req.Driver.Cores,
					"serviceAccount": req.Driver.ServiceAccount,
				},
				"executor": map[string]any{
					"memory":    req.Executor.Memory,
					"cores":     req.Executor.Cores,
					"instances": req.Executor.Instances,
				},
			},
		},
	}

	gvr := schema.GroupVersionResource{
		Group:    "sparkoperator.k8s.io",
		Version:  "v1beta2",
		Resource: "sparkapplications",
	}

	_, err := s.dynamicClient.Resource(gvr).Namespace("default").Create(ctx, sparkApp, v1.CreateOptions{})
	return err
}
