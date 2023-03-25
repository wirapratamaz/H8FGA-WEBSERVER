package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Employee struct {
	ID       int
	Name     string
	Age      int
	Division string
}

// variable dengan tipe data slice yg berisi struct
var employees = []Employee{
	{ID: 1, Name: "Airell", Age: 23, Division: "IT"},
	{ID: 2, Name: "Ai", Age: 22, Division: "Finance"},
	{ID: 3, Name: "Rell", Age: 24, Division: "IT"},
	{ID: 4, Name: "All", Age: 27, Division: "IT"},
}

var PORT = ":8080"

func main() {
	//handler untuk path employees
	http.HandleFunc("/employees", getEmployees)
	http.HandleFunc("/employee", createEmployees)

	fmt.Println("Application is listening on port", PORT)
	//menjalankan server aplikasi
	http.ListenAndServe(PORT, nil)
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	// Mengatur header response sebagai "application/json"
	// w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		//menggunakan package "encoding/json"
		//json.NewEndocer(w) -> menuliskan data JSON ke ResponseWriter "w"
		//Encode(employees) -> data employees dikonversi menjadi format JSON
		json.NewEncoder(w).Encode(employees)
		return
	}
	//*memparsing data ke html
	// tpl, err := template.ParseFiles("template.html")

	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// tpl.Execute(w, employees)
	// return

	// Jika request method bukan "GET", kirim response dengan status pesan error
	http.Error(w, "Invalid method", http.StatusBadRequest)
}

func createEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		name := r.FormValue("name")
		age := r.FormValue("age")
		division := r.FormValue("Division")

		convertAge, err := strconv.Atoi(age)

		if err != nil {
			http.Error(w, "Invalid age", http.StatusBadRequest)
			return
		}

		newEmployee := Employee{
			ID:       len(employees) + 1,
			Name:     name,
			Age:      convertAge,
			Division: division,
		}

		employees = append(employees, newEmployee)

		json.NewEncoder(w).Encode(newEmployee)
		return
	}

	http.Error(w, "Invalid method", http.StatusBadRequest)
}
