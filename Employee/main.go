package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Employee struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

var listEmployees []Employee

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hi! You came to Home! :)")
}

// Marshal --> Go struct --> JSON
// Unmarshal ---> JSON --> Go struct
func Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hi! You came to Home! :)")
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error in Reading Request", err)
	}
	employee := Employee{}
	err = json.Unmarshal(reqBody, &employee)
	if err != nil {
		fmt.Println("Error in Unmarshalling", err)
	}
	listEmployees = append(listEmployees, employee)
	fmt.Println(listEmployees)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(employee)
}
func Update(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hi! You came to Home! :)")
	idString := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idString)
	if err != nil {
		fmt.Println("Error in Conversion", err)
	}
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Println("Error in Reading Request", err)
	}
	employee := Employee{}
	err = json.Unmarshal(reqBody, &employee)
	if err != nil {
		fmt.Println("Error in Unmarshalling", err)
	}
	for i, emp := range listEmployees {
		if emp.ID == int32(id) {
			listEmployees[i].Name = employee.Name
			listEmployees[i].Role = employee.Role
			json.NewEncoder(w).Encode(listEmployees[i])
			break
		}
	}

	fmt.Println(listEmployees)
}
func Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hi! You came to Home! :)")
	idString := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idString)
	if err != nil {
		fmt.Println("Error in Conversion", err)
	}

	for i, emp := range listEmployees {
		if emp.ID == int32(id) {
			listEmployees = append(listEmployees[:i], listEmployees[i+1:]...)
			json.NewEncoder(w).Encode(listEmployees[i])
			break
		}
	}
	fmt.Println(listEmployees)
	json.NewEncoder(w).Encode(listEmployees)
}
func ReadAll(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(listEmployees)
}
func ReadOne(w http.ResponseWriter, r *http.Request) {
	idString := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idString)
	if err != nil {
		fmt.Println("Error in Conversion", err)
	}
	employee := Employee{}
	for _, emp := range listEmployees {
		if int32(id) == emp.ID {
			employee.ID = emp.ID
			employee.Name = emp.Name
			employee.Role = emp.Role
		}
	}
	json.NewEncoder(w).Encode(employee)
}
func main() {
	fmt.Println("Hi Hello! :)")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/home", Home)
	router.HandleFunc("/create", Create).Methods("POST")
	router.HandleFunc("/update/{id}", Update).Methods("PUT")
	router.HandleFunc("/read/{id}", ReadOne).Methods("GET")
	router.HandleFunc("/readall", ReadAll).Methods("GET")
	router.HandleFunc("/delete/{id}", Delete).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}

// A --> Request(Format in JSON) --> B
//  A <-- Response(Format in JSON) <-- B

// REST API GET/POST/PUT/DELETE
