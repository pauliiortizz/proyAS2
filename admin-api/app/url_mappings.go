package app

import (
	controller "admin-api/controller_admin"
	log "github.com/sirupsen/logrus"
)

func mapUrls() {

	// Add all methods and its mappings
	router.GET("/services", controller.GetServices)

	log.Info("Finishing mappings configurations")
}
