package controller

import (
	"encoding/json"
	"fmt"
	"golang-app/model"
	"golang-app/view"
	"net/http"
)

// RegisterRoutes  Register all the routes
func RegisterRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", ping())
	mux.HandleFunc("/crud/", crudOperations())
	return mux
}

func ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Pong")
		}
	}
}

func crudOperations() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			data, err := model.GetAllEmployees()
			if err != nil {
				w.Write([]byte(err.Error()))
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(view.Res{Data: data})
		} else if r.Method == http.MethodPost {
			data := view.Employee{}
			json.NewDecoder(r.Body).Decode(&data)
			fmt.Println("Data: ", data)
			if err := model.AddEmployee(data.ID, data.Name); err != nil {
				w.Write([]byte(err.Error()))
			}
			w.WriteHeader(http.StatusAccepted)
			json.NewEncoder(w).Encode(view.Res{Data: data})
		} else if r.Method == http.MethodDelete {
			ID := r.URL.Path[6:]
			if err := model.DeleteEmployee(ID); err != nil {
				w.Write([]byte(err.Error()))
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(view.Res{Data: "Employee with ID: " + ID + " deleted successfully."})
		} else {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(view.Res{Data: "Not implemented"})
		}
	}
}
