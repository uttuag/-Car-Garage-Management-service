package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// Car model
type Car struct {
	gorm.Model
	Brand  string
	model  string
	Status string
}

func main() {
	initDB()

	// Create router
	r := mux.NewRouter()

	// Define API routes
	r.HandleFunc("/cars", GetCars).Methods("GET")
	r.HandleFunc("/cars/{id}", GetCar).Methods("GET")
	r.HandleFunc("/cars", AddCar).Methods("POST")
	r.HandleFunc("/cars/{id}", UpdateCar).Methods("PUT")
	r.HandleFunc("/cars/{id}", DeleteCar).Methods("DELETE")

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", r))
}

func initDB() {
	var err error
	dsn := "root:my-secret-pw@tcp(localhost:3306)/cars?parseTime=true"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	// Auto Migrate the Car model
	db.AutoMigrate(&Car{})
}

// GetCars returns the list of cars
func GetCars(w http.ResponseWriter, r *http.Request) {
	// Fetch cars from the database
	var cars []Car
	db.Find(&cars)

	// Return JSON response
	w.Header().Set("Content-Type", "application/json")
	ServeJSON(w, cars)
}

// GetCar returns a single car by ID
func GetCar(w http.ResponseWriter, r *http.Request) {
	// Extract car ID from the request parameters
	params := mux.Vars(r)
	var car Car
	db.First(&car, params["id"])

	// Return JSON response
	w.Header().Set("Content-Type", "application/json")
	ServeJSON(w, car)
}

// AddCar adds a new car to the database
func AddCar(w http.ResponseWriter, r *http.Request) {
	// Parse the JSON request body
	var car Car
	ParseJSONBody(w, r, &car)

	// Create the car in the database
	db.Create(&car)

	// Return JSON response
	w.Header().Set("Content-Type", "application/json")
	ServeJSON(w, car)
}

// UpdateCar updates an existing car in the database
func UpdateCar(w http.ResponseWriter, r *http.Request) {
	// Extract car ID from the request parameters
	params := mux.Vars(r)

	// Find the car in the database
	var car Car
	db.First(&car, params["id"])

	// Parse the JSON request body
	ParseJSONBody(w, r, &car)

	// Update the car in the database
	db.Save(&car)

	// Return JSON response
	w.Header().Set("Content-Type", "application/json")
	ServeJSON(w, car)
}

// DeleteCar deletes a car from the database
func DeleteCar(w http.ResponseWriter, r *http.Request) {
	// Extract car ID from the request parameters
	params := mux.Vars(r)

	// Delete the car from the database
	db.Delete(&Car{}, params["id"])

	// Return success message
	w.Header().Set("Content-Type", "application/json")
	ServeJSON(w, map[string]string{"message": "Car deleted"})
}

// ServeJSON sends a JSON response with the provided data
func ServeJSON(w http.ResponseWriter, data interface{}) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

// ParseJSONBody parses the JSON request body into the provided interface
func ParseJSONBody(w http.ResponseWriter, r *http.Request, v interface{}) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(v)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
}