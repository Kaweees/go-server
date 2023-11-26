package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func handleGetExample(w http.ResponseWriter, r *http.Request) {
	// Set the response status code
	w.WriteHeader(http.StatusOK)

	// Set the response headers to indicate the content type
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// Build the response content as a JSON struct
	responseContent, err := requestHandler.GetExample()

	// Check if there was an error
	if err != nil {
		// Log the error
		Log.Error(fmt.Sprintf("Error handling request: %v", err))
		return
	} else {
		// Marshal the response content into JSON
		jsonData, err := json.Marshal(responseContent)
		if err != nil {
			Log.Error(fmt.Sprintf("Error marshaling JSON: %v", err))
			return
		}

		// Set the response status code
		w.WriteHeader(http.StatusOK)

		// Write the JSON response to the client
		w.Write(jsonData)
	}
}

func handlePostExample(w http.ResponseWriter, r *http.Request) {
	// Set the response status code
	w.WriteHeader(http.StatusOK)

	// Set the response headers to indicate the content type
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// Parse the request body into a WifiToggleRequest struct
	var request PostExampleRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse request body: %v", err), http.StatusBadRequest)
		return
	}

	// Build the response content as a JSON struct
	responseContent, err := requestHandler.PostExample(request.Value, request.AnotherValue)

	// Check if there was an error
	if err != nil {
		// Log the error
		Log.Error(fmt.Sprintf("Error handling request: %v", err))
		return
	} else {
		// Marshal the response content into JSON
		jsonData, err := json.Marshal(responseContent)
		if err != nil {
			Log.Error(fmt.Sprintf("Error marshaling JSON: %v", err))
			return
		}

		// Set the response status code
		w.WriteHeader(http.StatusOK)

		// Write the JSON response to the client
		w.Write(jsonData)
	}
}

var requestHandler *RequestHandler
var err error
var errs []error

func init() {
	// Initialize the logger
	initalizeLogger()

	// Initialize the Wi-Fi handler
	requestHandler, err = NewRequestHandler()
	if err != nil {
		Log.Error(fmt.Sprintf("Error initializing the request handler: %v", err))
		return
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/getExample", handleGetExample).Methods("GET")
	r.HandleFunc("/postExample", handlePostExample).Methods("POST")

	// Create a new HTTP server with the CORS (Cross-Origin Resource Sharing) middleware enabled
	corsOptions := handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Accept"}), // Allow only headers from the list
		handlers.AllowedOrigins([]string{"*"}),                                          // Allow requests from any origin
		handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}),                     // Allow only GET and POST methods
	)

	// Set up the server address and port
	const port = 8080
	const host = "localhost"

	// Start the server
	Log.Info(fmt.Sprintf("Server started at http://%s:%d.", host, port))
	http.Handle("/", corsOptions(r))
	http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), nil)
}