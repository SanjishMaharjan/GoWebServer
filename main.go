package main

import (
	"fmt"
	"log"
	"net/http"
)

// formHandler processes the form data submitted via a POST request.
// It extracts the "name" and "email" values from the form and displays them on the response.
func formHandler(res http.ResponseWriter, req *http.Request) {
	// Parse the form data from the POST request.
	if err := req.ParseForm(); err != nil {
		// If parsing the form fails, send an error message to the client.
		fmt.Fprintf(res, "ParseForm() err: %v", err)
		return
	}

	fmt.Fprintf(res, "Post request successful\n")

	name := req.FormValue("name")
	email := req.FormValue("email")

	fmt.Fprintf(res, "name: %s\n", name)
	fmt.Fprintf(res, "email: %s\n", email)
}

// helloHandler responds to the "/hello" route with a greeting message.
// It also checks the request method and URL path for proper handling.
func helloHandler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/hello" {
		http.Error(res, "404 Not Found", http.StatusNotFound)
		return
	}

	if req.Method != "GET" {
		http.Error(res, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(res, "Hello")
}

func main() {
	// Create a file server to serve static files from the "./static" directory.
	fileserver := http.FileServer(http.Dir("./static"))

	// Handle the root URL ("/") by serving static files.
	http.Handle("/", fileserver)

	// Handle form submissions at the "/form" endpoint.
	http.HandleFunc("/form", formHandler)

	// Handle requests to "/hello" using the helloHandler function.
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Listening on port 8080...\n")

	// Start the server on port 8080 and log any fatal errors.
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
