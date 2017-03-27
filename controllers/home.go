package controllers

import (
	"net/http"

	"github.com/erdalkiran/ps-golang-web/viewmodels"
)

type homeController struct {
}

func (c *homeController) register() {
	http.HandleFunc("/home", c.get)
}

func (c *homeController) get(w http.ResponseWriter, r *http.Request) {
	vm := viewmodels.NewHome()
	renderTemplate(w, "home", vm)
}
