package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)


type Customer struct {
	ID string
	Name string
	Role string
	Email string
	Phone int
	Contacted bool
}

var database = []Customer{{ID: uuid.New().String(), Name: "Bill", Role: "Data Analyst", Email: "jeremy@willborn.com", Phone: 5556667777, Contacted: true},
						{ID: uuid.New().String(), Name: "Will", Role: "Businessman", Email: "william@willborn.com", Phone: 5554447777, Contacted: false},
						{ID: uuid.New().String(), Name: "Chris", Role: "Software Engineer", Email: "chris@willborn.com", Phone: 777777777, Contacted: false}}

func getCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(database)
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	
	for _, member := range database {
		if member.ID == id {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(member)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]

	for i, member := range database {
		if member.ID == id {
			database = append(database[:i], database[i+1:]...)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(database)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newEntry Customer
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
        http.Error(w, "Error reading request body", http.StatusInternalServerError)
        return
    }

	json.Unmarshal(reqBody, &newEntry)
	if err != nil {
        http.Error(w, "Error unmarshalling JSON", http.StatusBadRequest)
        return
    }

	for _, member := range database {
		if member.ID == newEntry.ID {
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(map[string]string{"error": "Customer with ID " + member.ID + " already exists."})
			return
		}
	}
	newEntry.ID = uuid.New().String()
	database = append(database, newEntry)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(database)
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	
	var newEntry Customer
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
        http.Error(w, "Error reading request body", http.StatusInternalServerError)
        return
    }

	json.Unmarshal(reqBody, &newEntry)
	if err != nil {
        http.Error(w, "Error unmarshalling JSON", http.StatusBadRequest)
        return
    }

	for i, member := range database {
		if member.ID == id {
			w.WriteHeader(http.StatusOK)
			database = append(database[:i], newEntry)
			json.NewEncoder(w).Encode(database)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/customers", getCustomers).Methods("GET")
	router.HandleFunc("/customers", createCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", getCustomer).Methods("GET")
	router.HandleFunc("/customers/{id}", deleteCustomer).Methods("DELETE")
	router.HandleFunc("/customers/{id}", updateCustomer).Methods("PATCH")

	fileServer := http.FileServer(http.Dir("./static"))
	router.PathPrefix("/").Handler(http.StripPrefix("/", fileServer))

	fmt.Println("Server is starting on port 3000...")
	log.Fatal(http.ListenAndServe(":3000", router))
}