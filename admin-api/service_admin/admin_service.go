package service_admin

import (
	"admin-api/client_admin"
	"admin-api/domain_admin"
	"fmt"
	"golang.org/x/net/context"
)

var dockerClient = client_admin.NewDockerClient()

func GetServices(ctx context.Context) domain_admin.ServicesResponse {
	fmt.Println(dockerClient.GetContainers(ctx))

	return domain_admin.ServicesResponse{}
}
