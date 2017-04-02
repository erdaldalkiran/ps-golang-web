package controllers

import (
	"net/http"

	"github.com/erdalkiran/ps-golang-web/viewmodels"
)

type profileController struct{}

func (c *profileController) register() {
	router.HandleFunc("/profile", c.get).Methods("GET")
	router.HandleFunc("/profile", c.post).Methods("POST")

}

func (c *profileController) get(w http.ResponseWriter, r *http.Request) {
	vm := viewmodels.NewProfileVM()
	renderTemplate(w, r, "profile", vm)
}

func (c *profileController) post(w http.ResponseWriter, r *http.Request) {
	vm := viewmodels.NewProfileVM()

	vm.Content.Email = r.FormValue("email")
	vm.Content.FirstName = r.FormValue("first-name")
	vm.Content.LastName = r.FormValue("last-name")
	vm.Content.Stand.Address = r.FormValue("stand-address")
	vm.Content.Stand.City = r.FormValue("stand-city")
	vm.Content.Stand.State = r.FormValue("stand-state")
	vm.Content.Stand.Zip = r.FormValue("stand-zip")

	renderTemplate(w, r, "profile", vm)
}
