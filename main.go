package main

import (
	"crud_sederhana/database"
	"crud_sederhana/routes"
	"fmt"
	"net/http"
)

func main() {
	db := database.InitDatabase()
	defer db.Close()
	fmt.Println("Project CRUD Sederhana siap dimulai!")

	server := http.NewServeMux()

	routes.MapRoutes(server, db)

	http.ListenAndServe("localhost:8080", server)
}
