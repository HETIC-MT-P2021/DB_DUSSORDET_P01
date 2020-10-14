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

//GetOrderProductList handler to get order's product list
func GetOrderProductList(w http.ResponseWriter, r *http.Request) {

	db := database.DbConn
	repository := models.Repository{Conn: db}
	muxVar := mux.Vars(r)
	strID := muxVar["order_number"]
	orderNumber, err := strconv.Atoi(strID)
	if err != nil {
		log.Printf("could not parse str to int: %v", err)
		helpers.WriteErrorJSON(w, http.StatusInternalServerError, "could not parse url")
		return
	}

	productList, err := repository.GetOrderProductList(orderNumber)
	if err != nil {
		log.Printf("could not get products: %v", err)
		helpers.WriteErrorJSON(w, http.StatusInternalServerError, "could not get products from order")
		return
	}

	if len(productList) <= 0 {
		helpers.WriteErrorJSON(w, http.StatusNotFound, "could not find any product")
		return
	}

	helpers.WriteJSON(w, http.StatusOK, productList)
}
