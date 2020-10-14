package controllers

import (
	"github.com/gorilla/mux"
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

	employees, err := repository.GetEmployees(false, "")
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

//GetOfficeWithEmployees handler to retrieve an office with its employees
func GetOfficeWithEmployees(w http.ResponseWriter, r *http.Request) {

	db := database.DbConn
	repository := models.Repository{Conn: db}

	muxVar := mux.Vars(r)
	officeCode := muxVar["office_code"]

	office, err := repository.GetOfficeInfo(officeCode)
	if err != nil {
		log.Printf("could not get office info: %v", err)
		helpers.WriteErrorJSON(w, http.StatusInternalServerError, "could not get office")
		return
	}

	if office == nil {
		helpers.WriteErrorJSON(w, http.StatusNotFound, "could not find an office by this code")
		return
	}

	employees, err := repository.GetEmployees(true, officeCode)
	if err != nil {
		log.Printf("could not get employees: %v", err)
		helpers.WriteErrorJSON(w, http.StatusInternalServerError, "could not get employees")
		return
	}

	if len(employees) >= 0 {
		office.Employees = employees
	}

	helpers.WriteJSON(w, http.StatusOK, office)
}
