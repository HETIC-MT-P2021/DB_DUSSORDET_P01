package controllers

import (
	"github.com/gorilla/mux"
	"go_exercise/database"
	"go_exercise/helpers"
	"go_exercise/models"
	"log"
	"net/http"
	"strconv"
)

//GetCustomerByCustomerNumber handler for getting a customer by its customer number
func GetCustomerByCustomerNumber(w http.ResponseWriter, r *http.Request) {

	db := database.DbConn
	repository := models.Repository{Conn: db}
	muxVar := mux.Vars(r)
	strID := muxVar["customer_number"]
	employeeNumber, err := strconv.Atoi(strID)
	if err != nil {
		log.Printf("could not parse str to int: %v", err)
		helpers.WriteErrorJSON(w, http.StatusInternalServerError, "could not parse url")
		return
	}

	customer, err := repository.GetCustomerByNumber(employeeNumber)
	if err != nil {
		log.Printf("could not get customer: %v", err)
		helpers.WriteErrorJSON(w, http.StatusInternalServerError, "could not get customer")
		return
	}

	if customer == nil {
		helpers.WriteErrorJSON(w, http.StatusNotFound, "could not find customer")
		return
	}

	helpers.WriteJSON(w, http.StatusOK, customer)
}
