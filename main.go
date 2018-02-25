package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	ID         string   `json:"id,omitempty"`
	First_name string   `json:"first_name,omitempty"`
	Last_name  string   `json:"last_name,omitempty"`
	Address    *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

func main() {
	people = append(people, Person{ID: "1", First_name: "John", Last_name: "Doe", Address: &Address{City: "City X", State: "State X"}})
	people = append(people, Person{ID: "2", First_name: "Jane", Last_name: "Doe", Address: &Address{City: "City Y", State: "State Y"}})
	people = append(people, Person{ID: "3", First_name: "Francis", Last_name: "Sunday"})

	router := mux.NewRouter()
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetPeople(w http.ResponseWriter, r *http.Request)    {}
func GetPerson(w http.ResponseWriter, r *http.Request)    {}
func CreatePerson(w http.ResponseWriter, r *http.Request) {}
func DeletePerson(w http.ResponseWriter, r *http.Request) {}
