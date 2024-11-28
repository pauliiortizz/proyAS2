package client_admin

import (
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"time"
)

func AutoScale(service string) {

	log.Infof("Autoscaling %s", service)

	for {
		var avgCpuUsage float64

		stats, err := GetStatsByService(service)
		if err != nil {
			log.Errorf("Error getting %s stats: %v", service, err)
			continue
		}

		containersAmount := len(stats)

		for _, container := range stats {

			stringCPU := strings.Trim(container.CPU, "%")
			intCPU, err := strconv.ParseFloat(stringCPU, 64)
			if err != nil {
				log.Errorf("Error parsing string: %v", err)
				continue
			}

			avgCpuUsage += intCPU
		}

		avgCpuUsage = avgCpuUsage / float64(containersAmount)

		if avgCpuUsage >= 60 || containersAmount < 2 {
			instances, err := ScaleService(service)
			if err != nil {
				log.Errorf("Error creating %s container: %s", service, err)
				continue
			}

			log.Infof("Scaling up %s to %d instances", service, instances)

		} else if avgCpuUsage < 20 && containersAmount > 2 {

			err = DeleteContainer(stats[containersAmount-1].Id)
			if err != nil {
				log.Errorf("Error deleting %s container: %s", service, err)
				continue
			}

			log.Infof("Scaling down %s to %d instances", service, containersAmount-1)
		}

		time.Sleep(20 * time.Second)
	}
}
