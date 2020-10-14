package controllers

import (
	"github.com/gorilla/mux"
	"go_exercise/database"
	"go_exercise/helpers"
	"go_exercise/models"
	"log"
	"net/http"
)

//GetProductByCode handler to get product by product code
func GetProductByCode(w http.ResponseWriter, r *http.Request) {

	db := database.DbConn
	repository := models.Repository{Conn: db}
	muxVar := mux.Vars(r)
	productCode := muxVar["product_code"]

	product, err := repository.GetProductInfo(productCode)
	if err != nil {
		log.Printf("could not get product: %v", err)
		helpers.WriteErrorJSON(w, http.StatusInternalServerError, "could not get product")
		return
	}

	if product == nil {
		helpers.WriteErrorJSON(w, http.StatusNotFound, "could not find product")
		return
	}

	helpers.WriteJSON(w, http.StatusOK, product)
}
