package main

import (
	"net/http"

	"github.com/erdalkiran/ps-golang-web/controllers"
)

func main() {
	initilaize()

	controllers.Register()

	http.ListenAndServe(":8080", nil)

}

func initilaize() {
	controllers.Initialize()
}
