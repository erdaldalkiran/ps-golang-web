package controllers

import (
	"net/http"
	"strconv"

	"github.com/erdalkiran/ps-golang-web/viewmodels"
	"github.com/gorilla/mux"
)

type productsController struct {
}

func (c *productsController) register() {
	router.HandleFunc("/products", c.get)
	router.HandleFunc("/products/category/{id}", c.getByCategoryId)
}

func (c *productsController) get(w http.ResponseWriter, r *http.Request) {
	vm := viewmodels.NewProducts()
	renderTemplate(w, "products", vm)
}

func (c *productsController) getByCategoryId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idRaw := vars["id"]
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		handleError(w, err)
		return
	}
	vm, err := viewmodels.NewProductsByCategoryId(id)
	if err != nil {
		handle404Error(w, err)
		return
	}
	renderTemplate(w, "products", vm)
}
