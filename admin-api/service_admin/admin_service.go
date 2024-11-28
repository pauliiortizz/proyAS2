package service_admin

import (
	"admin-api/dao_admin"
	"context"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

type DockerService struct {
	Client *client.Client
}

func NewDockerService() (*DockerService, error) {
	apiClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}
	return &DockerService{Client: apiClient}, nil
}

// ListContainers devuelve todos los contenedores Docker
func (ds *DockerService) ListContainers(ctx context.Context) ([]dao_admin.Container, error) {
	containers, err := ds.Client.ContainerList(ctx, container.ListOptions{All: true})
	if err != nil {
		return nil, err
	}

	var result []dao_admin.Container
	for _, ctr := range containers {
		result = append(result, dao_admin.Container{
			ID:     ctr.ID,
			Image:  ctr.Image,
			Status: ctr.Status,
		})
	}

	return result, nil
}
