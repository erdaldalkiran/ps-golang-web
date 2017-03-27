package controllers

import (
	"net/http"

	"strconv"

	"github.com/erdalkiran/ps-golang-web/viewmodels"
	"github.com/gorilla/mux"
)

type productController struct {
}

func (c *productController) register() {
	router.HandleFunc("/product/{id}", c.get)
}

func (c *productController) get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idRaw := vars["id"]
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		handleError(w, err)
		return
	}
	vm, err := viewmodels.NewProductVM(id)
	if err != nil {
		handle404Error(w, err)
		return
	}
	renderTemplate(w, "product", vm)
}
