package app

import (
	controller "admin-api/controller_admin"
	log "github.com/sirupsen/logrus"
)

func mapUrls() {

	// Add all methods and its mappings
	router.GET("/services", controller.GetScalableServices)
	router.GET("/stats", controller.GetStats)
	router.GET("/stats/:service", controller.GetStatsByService)
	router.POST("/scale/:service", controller.ScaleService)
	router.DELETE("/container/:id", controller.DeleteContainer)

	log.Info("Finishing mappings configurations")
}
