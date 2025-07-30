package kubernetes // Updated package name

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
)

type Orchestrator struct {
	k8sClient     kubernetes.Interface
	dynamicClient dynamic.Interface
}

func NewOrchestrator(k8sClient kubernetes.Interface, dynClient dynamic.Interface) *Orchestrator {
	return &Orchestrator{
		k8sClient:     k8sClient,
		dynamicClient: dynClient,
	}
}

func (o *Orchestrator) ListPods(ctx context.Context) ([]string, error) {
	podlist, err := o.k8sClient.CoreV1().Pods("default").List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	pods := make([]string, 0, len(podlist.Items))
	for _, pod := range podlist.Items {
		pods = append(pods, pod.Name)
	}

	return pods, nil
}

type WorkspaceRequest struct {
	Name string `json:"name"`
}

func (o *Orchestrator) CreateWorkspace(ctx context.Context, req WorkspaceRequest) error {
	n := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: req.Name,
		},
	}
	_, err := o.k8sClient.CoreV1().Namespaces().Create(ctx, n, metav1.CreateOptions{})
	return err
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

func (o *Orchestrator) SubmitSparkJob(ctx context.Context, req PySparkJobRequest) error {
	sparkApp := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "sparkoperator.k8s.io/v1beta2",
			"kind":       "SparkApplication",
			"metadata": map[string]interface{}{
				"name":      req.Name,
				"namespace": "default",
			},
			"spec": map[string]interface{}{
				"type":                "Python",
				"pythonVersion":       "3",
				"mode":                req.Mode,
				"image":               req.Image,
				"sparkVersion":        req.SparkVersion,
				"mainApplicationFile": req.MainApplicationFile,
				"arguments":           req.Arguments,
				"restartPolicy":       map[string]interface{}{"type": "Never"},
				"driver": map[string]interface{}{
					"memory":         req.Driver.Memory,
					"cores":          req.Driver.Cores,
					"serviceAccount": req.Driver.ServiceAccount,
				},
				"executor": map[string]interface{}{
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

	_, err := o.dynamicClient.Resource(gvr).Namespace("default").Create(ctx, sparkApp, metav1.CreateOptions{})
	return err
}
