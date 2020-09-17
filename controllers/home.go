package controllers

import (
	"fmt"
	"net/http"
)

//RenderHome renders a basic home message
func RenderHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World !")
	return
}
