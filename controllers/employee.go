package controllers

import (
	"go_exercise/database"
	"go_exercise/helpers"
	"go_exercise/models"
	"log"
	"net/http"
)

//GetAllEmployees handler to retrieve all employees
func GetAllEmployees(w http.ResponseWriter, r *http.Request) {

	db := database.DbConn
	repository := models.Repository{Conn: db}

	employees, err := repository.GetEmployees()
	if err != nil {
		log.Printf("could not get employees: %v", err)
		helpers.WriteErrorJSON(w, http.StatusInternalServerError, "could not get employees")
		return
	}

	if len(employees) <= 0 {
		helpers.WriteErrorJSON(w, http.StatusNotFound, "could not find any employee")
		return
	}

	helpers.WriteJSON(w, http.StatusOK, employees)
}
