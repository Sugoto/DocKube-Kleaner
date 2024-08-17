package main

import (
	"log"
	"time"

	"github.com/sugoto/dockube-kleaner/internal/scheduler"
	"github.com/sugoto/dockube-kleaner/pkg/docker"
	"github.com/sugoto/dockube-kleaner/pkg/k8s"
)

func main() {
	dockerClient, err := docker.NewDockerClient()
	if err != nil {
		log.Fatalf("Failed to create Docker client: %v", err)
	}

	k8sClient := k8s.NewK8sClient()

	scheduler.ScheduleCleanup(24*time.Hour, func() {
		docker.CleanupDockerResources(dockerClient)
		k8s.CleanupK8sResources(k8sClient)
	})

	select {}
}
