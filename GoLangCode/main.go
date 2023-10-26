package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux" // You'll need to install this package with "go get -u github.com/gorilla/mux"
)

// Item struct represents an item in the API
type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Transaction struct {
	// ID int `json:"id"`
	// from   string `json:"name"`
	// to     string `json:"name"`
	// amount int
	//Name string
	AccountId int `json:"accountId"`
	//Timestamp int     `json:"timestamp"`
	Amount float64 `json:"amount"`
}

var items []Item
var currentID = 1

func main() {
	// Create a new router using the Gorilla Mux router.
	r := mux.NewRouter()

	// Define API endpoints
	// r.HandleFunc("/items", GetItems).Methods("GET")
	// r.HandleFunc("/items/{id}", GetItem).Methods("GET")
	// r.HandleFunc("/items/", CreateItem).Methods("POST")
	// r.HandleFunc("/items/{id}", UpdateItem).Methods("PUT")
	// r.HandleFunc("/items/{id}", DeleteItem).Methods("DELETE")

	r.HandleFunc("/transaction/", PostTransaction).Methods("POST")

	// Start the server on port 8080
	port := "localhost:8080"

	fmt.Printf("Server is running on port %s...\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}

// func PostTransaction(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Hello Tarun!")
// 	w.Header().Set("Content-Type", "application/json")
// 	//body, err := ioutil.ReadAll(r.Body)
// 	var transaction Transaction
// 	_ = json.NewDecoder(r.Body).Decode(&transaction)

// 	//fmt.Println("Body: ", string(body))
// 	// Log the transaction data to the console
// 	fmt.Printf("Received transaction data: %+v\n", transaction)
// 	// Process the transaction data as needed
// 	// Define the request URL with the query parameter
// 	requestURL := "http://localhost:8888/kafkaapp/post?msg=" + transaction.Name

// 	// Now you can log the request URL as well
// 	fmt.Printf("Sending POST request to: %s\n", requestURL)

// 	resp, err := http.Post(requestURL, "application/json", nil)
// 	if err != nil {
// 		http.Error(w, "Failed to send POST request", http.StatusInternalServerError)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode == http.StatusOK {
// 		// Handle the response data, if needed

// 	} else {
// 		http.Error(w, "POST request failed with status code: "+strconv.Itoa(resp.StatusCode), http.StatusBadGateway)
// 	}
// }

//26/5
// func PostTransaction(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Hello Tarun!")
// 	w.Header().Set("Content-Type", "application/json")
// 	//body, err := ioutil.ReadAll(r.Body)
// 	var transaction Transaction
// 	_ = json.NewDecoder(r.Body).Decode(&transaction)

// 	//fmt.Println("Body: ", string(body))
// 	// Log the transaction data to the console
// 	fmt.Printf("Received transaction data: %+v\n", transaction)
// 	// Process the transaction data as needed
// 	// Define the request URL with the query parameter
// 	requestURL := "http://localhost:8888/kafkaapp/post"

// 	// Now you can log the request URL as well
// 	fmt.Printf("Sending POST request to: %s\n", requestURL)

// 	resp, err := http.Post(requestURL, "application/json", nil)
// 	if err != nil {
// 		http.Error(w, "Failed to send POST request", http.StatusInternalServerError)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode == http.StatusOK {
// 		// Handle the response data, if needed

// 	} else {
// 		http.Error(w, "POST request failed with status code: "+strconv.Itoa(resp.StatusCode), http.StatusBadGateway)
// 	}
// }

func PostTransaction(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello Tarun!")
	w.Header().Set("Content-Type", "application/json")

	// Create a variable to hold the decoded Transaction data
	var transaction Transaction

	// Decode the JSON request body into the transaction variable
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	// Log the received transaction data to the console
	fmt.Printf("Received transaction data: %+v\n", transaction)

	// Process the transaction data as needed

	// Define the request URL for the POST request
	requestURL := "http://localhost:8888/kafkaapp/post"

	// Create a JSON payload from the transaction
	payload, err := json.Marshal(transaction)
	if err != nil {
		http.Error(w, "Failed to marshal JSON payload", http.StatusInternalServerError)
		return
	}

	// Now you can log the request URL as well
	fmt.Printf("Sending POST request to: %s\n", requestURL)

	// Send the POST request with the JSON payload
	resp, err := http.Post(requestURL, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		http.Error(w, "Failed to send POST request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		// Handle the response data, if needed
	} else {
		http.Error(w, "POST request failed with status code: "+strconv.Itoa(resp.StatusCode), http.StatusBadGateway)
	}
}

// GetItems returns a list of all items
func GetItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

// GetItem returns a specific item by ID
func GetItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for _, item := range items {
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	http.Error(w, "Item not found", http.StatusNotFound)
}

// CreateItem adds a new item
// func CreateItem(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var item Item
// 	_ = json.NewDecoder(r.Body).Decode(&item)
// 	item.ID = currentID
// 	currentID++
// 	items = append(items, item)
// 	json.NewEncoder(w).Encode(item)
// }

// PostMessage
// func PostTransaction(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var transaction Transaction
// 	_ = json.NewDecoder(r.Body).Decode(&transaction)
// 	// Process the transaction data as needed

// 	// Define the request URL with the query parameter
// 	requestURL := "http://localhost:8888/kafkaapp/post?msg"
// 	resp, err := http.Post(requestURL, "application/json", nil)
// 	if err != nil {
// 		http.Error(w, "Failed to send POST request", http.StatusInternalServerError)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode == http.StatusOK {
// 		// Handle the response data, if needed
// 	} else {
// 		http.Error(w, "POST request failed with status code: "+strconv.Itoa(resp.StatusCode), http.StatusBadGateway)
// 	}
// }

// UpdateItem updates an existing item
func UpdateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for i, item := range items {
		if item.ID == id {
			var updatedItem Item
			_ = json.NewDecoder(r.Body).Decode(&updatedItem)
			updatedItem.ID = item.ID
			items[i] = updatedItem
			json.NewEncoder(w).Encode(updatedItem)
			return
		}
	}

	http.Error(w, "Item not found", http.StatusNotFound)
}

// DeleteItem deletes an item by ID
func DeleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for i, item := range items {
		if item.ID == id {
			items = append(items[:i], items[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Item not found", http.StatusNotFound)
}
