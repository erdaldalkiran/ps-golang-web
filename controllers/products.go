package controllers

import (
	"net/http"

	"github.com/erdalkiran/ps-golang-web/viewmodels"
)

type productsController struct {
}

func (c *productsController) register() {
	http.HandleFunc("/products", c.get)
}

func (c *productsController) get(w http.ResponseWriter, r *http.Request) {
	vm := viewmodels.NewProducts()
	renderTemplate(w, "products", vm)
}
