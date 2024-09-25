package k8s

import (
	"context"
	"log"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// NewK8sClient initializes and returns a Kubernetes clientset.
func NewK8sClient() *kubernetes.Clientset {
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		log.Fatalf("Failed to build kubeconfig: %v", err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Failed to create Kubernetes client: %v", err)
	}
	return clientset
}

func CleanupK8sResources(clientset *kubernetes.Clientset) {
	ctx := context.TODO()

	pods, err := clientset.CoreV1().Pods("").List(ctx, v1.ListOptions{})
	if err != nil {
		log.Fatalf("Failed to list pods: %v", err)
	}
	for _, pod := range pods.Items {
		if pod.DeletionTimestamp != nil {
			err := clientset.CoreV1().Pods(pod.Namespace).Delete(ctx, pod.Name, v1.DeleteOptions{})
			if err != nil {
				log.Println("Error deleting pod:", err)
			} else {
				log.Println("Deleted pod:", pod.Name)
			}
		}
	}

}
