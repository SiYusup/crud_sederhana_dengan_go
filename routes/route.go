package routes

import (
	"crud_sederhana/controller"
	"database/sql"
	"net/http"
)

func MapRoutes(server *http.ServeMux, db *sql.DB) {

	server.HandleFunc("/hello", controller.HelloWorldController())

	server.HandleFunc("/employee", controller.EmployeeController(db))

	server.HandleFunc("/employee/create", controller.CreateEmployeeController(db))

	server.HandleFunc("/employee/edit", controller.UpdateEmployeeController(db))

	server.HandleFunc("/employee/delete", controller.DeleteEmployeeController(db))
}
