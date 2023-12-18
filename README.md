Car Management API
This project implements a simple HTTP API for managing cars using GoLang, Gorm, and MySQL.

Table of Contents
Requirements
Setup
API Endpoints
Usage
Contributing

Requirements
Make sure you have the following installed:

Go (https://golang.org/)
MySQL Database
Setup
Clone the repository:

bash
Copy code
git clone https://github.com/uttuag/car-management-api.git
Change into the project directory:

bash
Copy code
cd car-management-api
Install dependencies:

bash
Copy code
go get -u github.com/gorilla/mux
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
Configure the MySQL Database:

Create a MySQL database named cars.

Update the database connection details in main.go:

go
Copy code
dsn := "sql:database@tcp(localhost:3306)/cars?parseTime=true"
Run the application:

bash
Copy code
go run main.go
The application will start at http://localhost:8080.

API Endpoints
GET /cars: Retrieve the list of all cars.
GET /cars/{id}: Retrieve a single car by ID.
POST /cars: Add a new car to the database.
PUT /cars/{id}: Update an existing car in the database.
DELETE /cars/{id}: Delete a car from the database.
Usage
Use tools like Postman or curl to interact with the API.

Example API request (using curl):

bash
Copy code
curl -X GET http://localhost:8080/cars
Contributing
Contributions are welcome! If you find any issues or have suggestions for improvement, please open an issue or create a pull request.
