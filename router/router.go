package router

import (
	"github.com/gorilla/mux"
	"go_exercise/controllers"
	"net/http"
)

// Route struct defining all of this project routes
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes slice of Route
type Routes []Route

// NewRouter registers public routes
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{
		Name:        "Home",
		Method:      "GET",
		Pattern:     "/",
		HandlerFunc: controllers.RenderHome,
	},
	Route{
		Name:        "Get customer by customer number",
		Method:      "GET",
		Pattern:     "/customers/{customer_number}",
		HandlerFunc: controllers.GetCustomerByCustomerNumber,
	},
	Route{
		Name:        "Get orders by customer number",
		Method:      "GET",
		Pattern:     "/orders/{customer_number}",
		HandlerFunc: controllers.GetCustomerOrders,
	},
	Route{
		Name:        "Get an order's product list by order number",
		Method:      "GET",
		Pattern:     "/orders/{order_number}/products",
		HandlerFunc: controllers.GetOrderProductList,
	},
	Route{
		Name:        "Get a product by code",
		Method:      "GET",
		Pattern:     "/products/{product_code}",
		HandlerFunc: controllers.GetProductByCode,
	},
	Route{
		Name:        "Get all employees",
		Method:      "GET",
		Pattern:     "/employees",
		HandlerFunc: controllers.GetAllEmployees,
	},
}
