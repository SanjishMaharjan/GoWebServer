package main

import (
	"fmt"
	"log"
	"net/http"
)

// formHandler processes the form data submitted via a POST request.
// It extracts the "name" and "email" values from the form and displays them on the response.
func formHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the form data from the POST request.
	if err := r.ParseForm(); err != nil {
		// If parsing the form fails, send an error message to the client.
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	// A success message indicating that the POST request was successful.
	fmt.Fprintf(w, "Post request successful\n")

	// Extract "name" and "email" fields from the form.
	name := r.FormValue("name")
	email := r.FormValue("email")

	// Write the extracted form values back to the response.
	fmt.Fprintf(w, "name: %s\n", name)
	fmt.Fprintf(w, "email: %s\n", email)
}

// helloHandler responds to the "/hello" route with a greeting message.
// It also checks the request method and URL path for proper handling.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the path is "/hello"; if not, return a 404 error.
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	// Ensure that the method is GET; if it's not, return a 405 error.
	if r.Method != "GET" {
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Respond with "Hello" when the "/hello" route is accessed via GET request.
	fmt.Fprintf(w, "Hello")
}

// The main function starts the HTTP server, handles routing, and serves static files.
func main() {
	// Create a file server to serve static files from the "./static" directory.
	fileserver := http.FileServer(http.Dir("./static"))

	// Handle the root URL ("/") by serving static files.
	http.Handle("/", fileserver)

	// Handle form submissions at the "/form" endpoint.
	http.HandleFunc("/form", formHandler)

	// Handle requests to "/hello" using the helloHandler function.
	http.HandleFunc("/hello", helloHandler)

	// Print a message to indicate that the server is running on port 8080.
	fmt.Printf("Listening on port 8080...\n")

	// Start the server on port 8080 and log any fatal errors.
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
