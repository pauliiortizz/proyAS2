package service_admin

import (
	"admin-api/client_admin"
	"admin-api/domain_admin"
	"fmt"
	"golang.org/x/net/context"
	"strings"
)

var dockerClient = client_admin.NewDockerClient()

func GetServices(ctx context.Context) (domain_admin.ServicesResponse, error) {
	// Obtener los contenedores desde DockerClient
	containers, err := dockerClient.GetContainers(ctx)
	if err != nil {
		fmt.Println("Error al obtener contenedores:", err)
		return domain_admin.ServicesResponse{}, err
	}

	// Crear el slice de respuesta
	containerDomainList := make([]domain_admin.Service, 0, len(containers))

	expectedServices := map[string]bool{
		"search-api":        true,
		"cursos-api":        true,
		"inscripciones-api": true,
		"users-api":         true,
	}

	// Iterar sobre los contenedores
	for _, container := range containers {
		// Revisar cada nombre del contenedor
		for _, name := range container.Names {
			// Comprobar si el nombre contiene alguno de los servicios esperados
			for serviceName := range expectedServices {
				if strings.Contains(name, serviceName) && container.State == "running" {
					containerDomainList = append(containerDomainList, domain_admin.Service{
						Name:       []string{serviceName},
						Containers: []string{container.ID},
					})
					break // No es necesario seguir buscando en este contenedor
				}
			}
		}
	}

	return domain_admin.ServicesResponse{
		Services: containerDomainList,
	}, nil
}
