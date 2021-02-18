package main

import (
	"fmt"
	"golang-app/controller"
	"golang-app/model"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	db := model.Connect()
	defer db.Close()
	mux := controller.RegisterRoutes()
	fmt.Println("Application is running...")
	http.ListenAndServe(":4000", mux)
}
