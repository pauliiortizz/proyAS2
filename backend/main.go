package main

import (
	"backend/clients"
	"backend/router"
	"fmt"
)

func main() {
	// Conectar a la base de datos
	clients.ConnectDatabase()

	// Migrar las entidades
	clients.MigrateEntities()

	// Otras operaciones de tu aplicación
	// ...

	fmt.Println("Aplicación lista para usar")
	// Aquí puedes continuar con la lógica de tu aplicación

	router.StartRoute()

}
