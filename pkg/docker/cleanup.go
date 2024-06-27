package docker

import (
	"context"
	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/volume"
	"github.com/docker/docker/client"
)

func NewDockerClient() (*client.Client, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}
	return cli, nil
}

// CleanupDockerResources cleans up unused Docker containers, images, and volumes.
func CleanupDockerResources(cli *client.Client) {
	ctx := context.Background()

	// Remove stopped containers
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{All: true})
	if err != nil {
		log.Fatal(err)
	}
	for _, container := range containers {
		if container.State == "exited" {
			err := cli.ContainerRemove(ctx, container.ID, types.ContainerRemoveOptions{RemoveVolumes: true})
			if err != nil {
				log.Println("Error removing container:", err)
			} else {
				log.Println("Removed container:", container.ID)
			}
		}
	}

	// Remove dangling images
	imgFilters := filters.NewArgs()
	imgFilters.Add("dangling", "true")
	images, err := cli.ImageList(ctx, types.ImageListOptions{Filters: imgFilters})
	if err != nil {
		log.Fatal(err)
	}
	for _, image := range images {
		_, err := cli.ImageRemove(ctx, image.ID, types.ImageRemoveOptions{Force: true})
		if err != nil {
			log.Println("Error removing image:", err)
		} else {
			log.Println("Removed image:", image.ID)
		}
	}

	// Remove unused volumes
	volumes, err := cli.VolumeList(ctx, volume.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	for _, vol := range volumes.Volumes {
		err := cli.VolumeRemove(ctx, vol.Name, true)
		if err != nil {
			log.Println("Error removing volume:", err)
		} else {
			log.Println("Removed volume:", vol.Name)
		}
	}
}