package client_admin

import (
	dto "admin-api/dao_admin"
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os/exec"
	"strings"
)

var scalableServices = []string{"course", "search", "user", "inscripcion"}

func GetStats() (dto.ContainersStatsDto, error) {

	var containersStats dto.ContainersStatsDto

	command := exec.Command("docker", "stats", "--no-stream", "--format", "{{json .}}")

	stdout, err := command.StdoutPipe()
	if err != nil {
		return dto.ContainersStatsDto{}, err
	}

	err = command.Start()
	if err != nil {
		return dto.ContainersStatsDto{}, err
	}

	scanner := bufio.NewScanner(stdout)

	for scanner.Scan() {
		var containerStats dto.ContainerStatsDto

		err = json.Unmarshal(scanner.Bytes(), &containerStats)
		if err != nil {
			return dto.ContainersStatsDto{}, err
		}

		containersStats = append(containersStats, containerStats)
	}

	err = scanner.Err()
	if err != nil {
		return dto.ContainersStatsDto{}, err
	}

	err = command.Wait()
	if err != nil {
		return dto.ContainersStatsDto{}, err
	}

	return containersStats, nil
}

func GetStatsByService(service string) (dto.ContainersStatsDto, error) {

	if !serviceExists(service) {
		return dto.ContainersStatsDto{}, errors.New("service does not exist")
	}

	var containersStats dto.ContainersStatsDto

	containers, err := getContainersIdsByService(service)
	if err != nil {
		return dto.ContainersStatsDto{}, err
	}

	cmdArgs := append([]string{"stats", "--no-stream", "--format", "{{json .}}"}, containers...)

	command := exec.Command("docker", cmdArgs...)

	stdout, err := command.StdoutPipe()
	if err != nil {
		return dto.ContainersStatsDto{}, err
	}

	err = command.Start()
	if err != nil {
		return dto.ContainersStatsDto{}, err
	}

	scanner := bufio.NewScanner(stdout)

	for scanner.Scan() {
		var containerStats dto.ContainerStatsDto

		err = json.Unmarshal(scanner.Bytes(), &containerStats)
		if err != nil {
			return dto.ContainersStatsDto{}, err
		}

		containersStats = append(containersStats, containerStats)
	}

	err = scanner.Err()
	if err != nil {
		return dto.ContainersStatsDto{}, err
	}

	err = command.Wait()
	if err != nil {
		return dto.ContainersStatsDto{}, err
	}

	return containersStats, nil
}

func ScaleService(service string) (int, error) {

	if !serviceExists(service) {
		return 0, errors.New("service does not exist")
	}

	if !serviceScalable(service) {
		return 0, errors.New("service not scalable")
	}

	ids, err := getContainersIdsByService(service)
	if err != nil {
		return 0, err
	}

	currQty := len(ids)

	scaleCommand := exec.Command("docker-compose", "-f", "../docker-compose.yml", "up", "-d", "--scale", fmt.Sprintf("%s=%d", service, currQty+1))

	err = scaleCommand.Run()
	if err != nil {
		return 0, err
	}

	restartCommand := exec.Command("docker-compose", "-f", "../docker-compose.yml", "restart", fmt.Sprintf("%s%s", service, "nginx"))
	err = restartCommand.Run()
	if err != nil {
		return 0, err
	}

	return currQty + 1, err

}

func DeleteContainer(id string) error {

	service, err := getContainerServiceById(id)
	if err != nil {
		return err
	}

	if !serviceScalable(service) {
		return errors.New("you cant delete this service's containers")
	}

	containers, err := getContainersIdsByService(service)
	if len(containers) == 1 {
		return errors.New("there has to be at least one container running for each service")
	}

	deleteCommand := exec.Command("docker", "rm", "-f", id)
	err = deleteCommand.Run()
	if err != nil {
		return err
	}

	restartCommand := exec.Command("docker-compose", "-f", "../docker-compose.yml", "restart", fmt.Sprintf("%s%s", service, "nginx"))
	err = restartCommand.Run()
	if err != nil {
		return err
	}

	return nil
}

func GetScalableServices() []string {
	return scalableServices
}

func getContainersIdsByService(service string) ([]string, error) {

	command := exec.Command("docker-compose", "-f", "../docker-compose.yml", "ps", "-q", service)
	output, err := command.Output()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	ids := strings.TrimSpace(string(output))

	idsArray := strings.Split(ids, "\n")

	return idsArray, nil
}

func getContainerServiceById(id string) (string, error) {

	command := exec.Command("docker", "inspect", "--format", `{ "service": "{{ index .Config.Labels "com.docker.compose.service" }}" }`, id)
	output, err := command.Output()
	if err != nil {
		log.Error(err)
		return "", errors.New("container not found")
	}

	var containerService struct {
		Service string `json:"Service"`
	}

	err = json.Unmarshal(output, &containerService)
	if err != nil {
		log.Error(err)
		return "", err
	}

	return containerService.Service, nil
}

func serviceExists(service string) bool {

	command := exec.Command("docker-compose", "-f", "../docker-compose.yml", "config", "--services")
	output, err := command.Output()
	if err != nil {
		log.Error(err)
		return false
	}

	services := strings.TrimSpace(string(output))
	servicesArray := strings.Split(services, "\n")

	for _, serv := range servicesArray {

		if serv == service {
			return true
		}
	}

	return false
}

func serviceScalable(service string) bool {

	for _, serv := range scalableServices {
		if serv == service {
			return true
		}
	}

	return false
}
