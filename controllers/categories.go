package controllers

import (
	"net/http"

	"github.com/erdalkiran/ps-golang-web/viewmodels"
)

type categoriesController struct {
}

func (c *categoriesController) register() {
	http.HandleFunc("/categories", c.get)
}

func (c *categoriesController) get(w http.ResponseWriter, r *http.Request) {
	vm := viewmodels.NewCategories()

	renderTemplate(w, r, "categories", vm)
}
