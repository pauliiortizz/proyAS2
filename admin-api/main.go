package main

import (
	app "admin-api/app"
	client "admin-api/client_admin"
)

func main() {

	services := client.GetScalableServices()

	for _, service := range services {

		go client.AutoScale(service)

	}

	app.StartRoute()
}
