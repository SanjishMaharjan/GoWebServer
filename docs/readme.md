# Go Web Server

This is a simple Go web server that serves static files and handles form submissions. The server responds to HTTP requests and provides routes for handling static content, form submissions, and basic routing.

## Features

- Serves static files from the `./static` directory.
- Handles POST requests to process form submissions.
- Simple routing to respond to `/hello` with a greeting.
- Error handling for unsupported paths and methods.

## Requirements

- Go 1.16 or higher.

## Setup

1. Clone the repository:

   ```bash
   https://github.com/SanjishMaharjan/GoWebServer.git
   ```

2. Navigate to the project directory:

   ```bash
   cd go-web-server
   ```

3. Install any dependencies:

   ```bash
   go mod tidy
   ```

4. Run the server:

   ```bash
   go run main.go
   ```

5. The server will be running on `http://localhost:8080`.

## File Structure

```bash
.
├── main.go         # Main Go server file
├── static/         # Directory for static files (HTML, CSS, JS, etc.)
├── go.mod          # Go module file
├── README.md       # Project README file
└── .gitignore      # Git ignore file
```
