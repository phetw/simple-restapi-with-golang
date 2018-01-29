package main

import (
	"employee-api/models"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var employees []models.Employee

func main() {
	router := mux.NewRouter()

	// Mock data
	employees = append(employees, models.Employee{ID: "1", Firstname: "Wasuwat", Lastname: "Limsuparhat", Position: "Developer"})
	employees = append(employees, models.Employee{ID: "2", Firstname: "Suepsakun", Lastname: "Aiamlaoo", Position: "Developer"})
	employees = append(employees, models.Employee{ID: "3", Firstname: "Sittiphon", Lastname: "Songsaen", Position: "Developer"})
	employees = append(employees, models.Employee{ID: "4", Firstname: "Steve", Lastname: "Jobs", Position: "CEO"})

	// Prefix
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	// Routes
	subRouter.HandleFunc("/employee", getEmployees).Methods("GET")
	subRouter.HandleFunc("/employee/{id}", getEmployee).Methods("GET")
	subRouter.HandleFunc("/employee", createEmployee).Methods("POST")
	subRouter.HandleFunc("/employee/{id}", updateEmployee).Methods("PUT")
	subRouter.HandleFunc("/employee/{id}", deleteEmployee).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", router))
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(employees)
}

func getEmployee(w http.ResponseWriter, r *http.Request) {

}

func createEmployee(w http.ResponseWriter, r *http.Request) {

}

func updateEmployee(w http.ResponseWriter, r *http.Request) {

}

func deleteEmployee(w http.ResponseWriter, r *http.Request) {

}
