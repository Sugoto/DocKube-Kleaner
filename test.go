package main

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/sugoto/dockube-kleaner/pkg/docker"
	"github.com/sugoto/dockube-kleaner/pkg/k8s"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1core "k8s.io/client-go/kubernetes/typed/core/v1"
)

func TestDockerCleanup(t *testing.T) {
	log.Println("Starting Docker cleanup test...")

	// Initialize Docker client
	dockerClient, err := docker.NewDockerClient()
	if err != nil {
		t.Fatalf("Failed to create Docker client: %v", err)
	}

	ctx := context.Background()

	// Create a dummy container
	containerName := "test-container"
	log.Printf("Creating Docker container: %s", containerName)
	_, err = dockerClient.ContainerCreate(ctx, &container.Config{
		Image: "alpine",
		Cmd:   []string{"sleep", "300"},
	}, nil, nil, nil, containerName)
	if err != nil {
		t.Fatalf("Failed to create Docker container: %v", err)
	}

	// Start the container
	log.Printf("Starting Docker container: %s", containerName)
	err = dockerClient.ContainerStart(ctx, containerName, types.ContainerStartOptions{})
	if err != nil {
		t.Fatalf("Failed to start Docker container: %v", err)
	}

	// Stop the container to make it eligible for cleanup
	time.Sleep(3 * time.Second)
	log.Printf("Stopping Docker container: %s", containerName)
	err = dockerClient.ContainerStop(ctx, containerName, nil)
	if err != nil {
		t.Fatalf("Failed to stop Docker container: %v", err)
	}

	// Run the garbage collector
	log.Println("Running Docker garbage collector...")
	docker.CleanupDockerResources(dockerClient)

	// Verify that the container is removed
	log.Printf("Verifying that Docker container %s is removed...", containerName)
	containers, err := dockerClient.ContainerList(ctx, types.ContainerListOptions{All: true})
	if err != nil {
		t.Fatalf("Failed to list Docker containers: %v", err)
	}
	for _, c := range containers {
		if c.Names[0] == "/"+containerName {
			t.Fatalf("Docker container %s was not removed by the garbage collector", containerName)
		}
	}

	log.Println("Docker cleanup test completed successfully.")
}

func TestKubernetesCleanup(t *testing.T) {
	log.Println("Starting Kubernetes cleanup test...")

	// Initialize Kubernetes client
	k8sClient := k8s.NewK8sClient()

	ctx := context.TODO()

	// Create a dummy pod
	podName := "test-pod"
	namespace := "default"
	log.Printf("Creating Kubernetes pod: %s in namespace: %s", podName, namespace)
	podClient := k8sClient.CoreV1().Pods(namespace)
	pod := &v1core.Pod{
		ObjectMeta: v1.ObjectMeta{
			Name: podName,
		},
		Spec: v1core.PodSpec{
			Containers: []v1core.Container{
				{
					Name:  "sleep",
					Image: "alpine",
					Command: []string{
						"sleep", "300",
					},
				},
			},
		},
	}

	_, err := podClient.Create(ctx, pod, v1.CreateOptions{})
	if err != nil {
		t.Fatalf("Failed to create Kubernetes pod: %v", err)
	}

	// Simulate pod termination by deleting it with grace period 0
	time.Sleep(3 * time.Second)
	log.Printf("Deleting Kubernetes pod: %s in namespace: %s", podName, namespace)
	err = podClient.Delete(ctx, podName, v1.DeleteOptions{GracePeriodSeconds: new(int64)})
	if err != nil {
		t.Fatalf("Failed to delete Kubernetes pod: %v", err)
	}

	// Run the garbage collector
	log.Println("Running Kubernetes garbage collector...")
	k8s.CleanupK8sResources(k8sClient)

	// Verify that the pod is removed
	log.Printf("Verifying that Kubernetes pod %s is removed...", podName)
	pods, err := podClient.List(ctx, v1.ListOptions{})
	if err != nil {
		t.Fatalf("Failed to list Kubernetes pods: %v", err)
	}
	for _, p := range pods.Items {
		if p.Name == podName {
			t.Fatalf("Kubernetes pod %s was not removed by the garbage collector", podName)
		}
	}

	log.Println("Kubernetes cleanup test completed successfully.")
}

func TestMain(m *testing.M) {
	// Initialize logger
	log.Println("Initializing test suite...")

	// Run tests
	result := m.Run()

	// Cleanup after tests
	log.Println("Test suite completed.")
}