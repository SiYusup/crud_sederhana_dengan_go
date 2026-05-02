package controller

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

type Employee struct {
	Id      int
	Name    string
	Npwp    string
	Address string
}

func EmployeeController(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, error := db.Query("SELECT * FROM employee;")

		if error != nil {
			http.Error(w, error.Error(), http.StatusInternalServerError)
			return
		}

		var employees []Employee
		for rows.Next() {
			var employee Employee
			err := rows.Scan(
				&employee.Id,
				&employee.Name,
				&employee.Npwp,
				&employee.Address,
			)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			employees = append(employees, employee)
		}
		defer rows.Close()

		data := make(map[string]any)
		data["employees"] = employees

		funcMap := template.FuncMap{
			"add": func(a, b int) int {
				return a + b
			},
		}

		indexPath := filepath.Join("views", "index.html")
		tmpl, error := template.New(filepath.Base(indexPath)).Funcs(funcMap).ParseFiles(indexPath)

		if error != nil {
			http.Error(w, error.Error(), http.StatusInternalServerError)
			return
		}

		error = tmpl.Execute(w, data)

		if error != nil {
			http.Error(w, error.Error(), http.StatusInternalServerError)
			return
		}
	}
}
