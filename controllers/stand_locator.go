package controllers

import (
	"net/http"

	"github.com/erdalkiran/ps-golang-web/viewmodels"
)

type standLocatorController struct {
}

func (c *standLocatorController) register() {
	http.HandleFunc("/stand_locator", c.get)
	http.HandleFunc("/stand_locator/search", c.search)
}

func (c *standLocatorController) get(w http.ResponseWriter, r *http.Request) {
	vm := viewmodels.NewStandLocatorVM()
	renderTemplate(w, r, "stand_locator", vm)
}

func (c *standLocatorController) search(w http.ResponseWriter, r *http.Request) {
	vm := viewmodels.GetStandLocations()
	renderJSON(w, r, vm)
}
