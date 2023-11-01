// Main code
//package main

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"strconv"
// 	"time"

// 	"github.com/gorilla/mux"
// )

// type Item struct {
// 	ID   int    `json:"id"`
// 	Name string `json:"name"`
// }

// type Transaction struct {
// 	AccountId int     `json:"accountId"`
// 	Amount    float64 `json:"amount"`
// 	DateTime  string  `json:"dateTime"`
// }

// var items []Item
// var currentID = 1

// func main() {
// 	r := mux.NewRouter()
// 	r.HandleFunc("/transaction/", PostTransaction).Methods("POST")

// 	port := "localhost:8080"

// 	fmt.Printf("Server is running on port %s...\n", port)
// 	log.Fatal(http.ListenAndServe(port, r))
// }

// func getHardCodedTransaction() []Transaction {
// 	// Create and return a hard-coded array of Transaction objects
// 	return []Transaction{
// 		{
// 			AccountId: 1,
// 			Amount:    100.0,
// 			DateTime:  "2023-10-30 12:00:00",
// 		},
// 		{
// 			AccountId: 2,
// 			Amount:    200.0,
// 			DateTime:  "2023-10-30 13:00:00",
// 		},
// 		{
// 			AccountId: 3,
// 			Amount:    200.0,
// 			DateTime:  "2023-10-30 13:00:00",
// 		},
// 		{
// 			AccountId: 1,
// 			Amount:    110.0,
// 			DateTime:  "2023-10-30 12:00:00",
// 		},
// 		{
// 			AccountId: 2,
// 			Amount:    201.0,
// 			DateTime:  "2023-10-30 13:00:00",
// 		},
// 		{
// 			AccountId: 3,
// 			Amount:    201.0,
// 			DateTime:  "2023-10-30 13:03:00",
// 		},
// 		{
// 			AccountId: 1,
// 			Amount:    111.0,
// 			DateTime:  "2023-10-30 12:00:00",
// 		},
// 		{
// 			AccountId: 2,
// 			Amount:    111.0,
// 			DateTime:  "2023-10-30 13:00:00",
// 		},
// 		{
// 			AccountId: 3,
// 			Amount:    290.0,
// 			DateTime:  "2023-10-30 13:00:00",
// 		},
// 	}

// }

// func PostTransaction(w http.ResponseWriter, r *http.Request) {

// 	fmt.Println("Hello Tarun!")
// 	w.Header().Set("Content-Type", "application/json")

// 	// Get the hard-coded array of transactions
// 	transactions := getHardCodedTransaction()

// 	// Log all received transactions
// 	for _, transaction := range transactions {
// 		currentTime := time.Now()
// 		transaction.DateTime = currentTime.Format("2006-01-02 15:04:05")
// 		fmt.Printf("Received transaction data: %+v\n", transaction)
// 	}

// 	// Process the transactions data as needed

// 	// Define the request URL for the POST request
// 	requestURL := "http://localhost:8888/kafkaapp/post"

// 	// Create a JSON payload from the array of transactions
// 	payload, err := json.Marshal(transactions)
// 	if err != nil {
// 		http.Error(w, "Failed to marshal JSON payload", http.StatusInternalServerError)
// 		return
// 	}

// 	fmt.Printf("Sending POST request to: %s\n", requestURL)

// 	resp, err := http.Post(requestURL, "application/json", bytes.NewBuffer(payload))
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

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Transaction struct {
	AccountId int     `json:"accountId"`
	Amount    float64 `json:"amount"`
	DateTime  string  `json:"dateTime"`
}

var items []Item
var currentID = 1

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/transaction/", PostTransaction).Methods("POST")

	port := "localhost:8080"

	fmt.Printf("Server is running on port %s...\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}

func getHardCodedTransaction() []Transaction {
	// Create and return a hard-coded array of Transaction objects
	return []Transaction{
		{
			AccountId: 1,
			Amount:    100.0,
			DateTime:  "2023-10-30 12:00:00",
		},
		{
			AccountId: 2,
			Amount:    200.0,
			DateTime:  "2023-10-30 13:00:00",
		},
		{
			AccountId: 3,
			Amount:    200.0,
			DateTime:  "2023-10-30 13:00:00",
		},
		{
			AccountId: 1,
			Amount:    110.0,
			DateTime:  "2023-10-30 12:00:00",
		},
		{
			AccountId: 2,
			Amount:    201.0,
			DateTime:  "2023-10-30 13:00:00",
		},
		{
			AccountId: 3,
			Amount:    201.0,
			DateTime:  "2023-10-30 13:03:00",
		},
		{
			AccountId: 1,
			Amount:    111.0,
			DateTime:  "2023-10-30 12:00:00",
		},
		{
			AccountId: 2,
			Amount:    111.0,
			DateTime:  "2023-10-30 13:00:00",
		},
		{
			AccountId: 3,
			Amount:    290.0,
			DateTime:  "2023-10-30 13:00:00",
		},
	}
}

func PostTransaction(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello Tarun!")
	w.Header().Set("Content-Type", "application/json")

	// Get the hard-coded array of transactions
	transactions := getHardCodedTransaction()

	// Process the transactions data as needed
	for _, transaction := range transactions {
		// Log the received transaction
		currentTime := time.Now()
		transaction.DateTime = currentTime.Format("2006-01-02 15:04:05")
		fmt.Printf("Received transaction data: %+v\n", transaction)

		// Define the request URL for the POST request
		requestURL := "http://localhost:8888/kafkaapp/post"

		// Create a JSON payload for the individual transaction
		payload, err := json.Marshal(transaction)
		if err != nil {
			http.Error(w, "Failed to marshal JSON payload", http.StatusInternalServerError)
			return
		}

		fmt.Printf("Sending POST request to: %s\n", requestURL)

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

		// Wait for 2 seconds before sending the next transaction
		time.Sleep(2 * time.Second)
	}
}
