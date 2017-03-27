package controllers

import (
	"net/http"

	"github.com/erdalkiran/ps-golang-web/viewmodels"
)

type productController struct {
}

func (c *productController) register() {
	http.HandleFunc("/product", c.get)
}

func (c *productController) get(w http.ResponseWriter, r *http.Request) {
	vm := viewmodels.NewProduct()
	renderTemplate(w, "product", vm)
}
