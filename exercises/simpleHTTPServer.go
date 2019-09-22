package main

import (
	"io"
	"net/http"
)

// Define Hello function that is called by the HandleFunc
func hello(response http.ResponseWriter, request *http.Request) {

	// set the response header
	response.Header().Set("Content-Type", "text/html")

	// set the html content [whih normally would be generated automatically]
	htmlContent := `<DOCTYPE html>
	<html>
		<head>
			<title> Welcome Page</title>
		</head>
		<body>
			Hello World, I'm just a simple HTTP Server !
		</body>
	</html>`

	//write the response conten "io.WriterString" as writer
	// "response" as the http Writer
	// a sample "http" body as an example of what is going to be written
	io.WriteString(response, htmlContent)
}

func root(response http.ResponseWriter, request *http.Request) {
	// set the response header
	response.Header().Set("Content-Type", "text/html")

	// set the html content [whih normally would be generated automatically]
	htmlContent := `<DOCTYPE html>
	<html>
		<head>
			<title>Root Page</title>
		</head>
		<body>
			This is a website server by a Go HTTP Server!
		</body>
	</html>`

	//write the response conten "io.WriterString" as writer
	// "response" as the http Writer
	// a sample "http" body as an example of what is going to be written
	io.WriteString(response, htmlContent)
}

func main() {
	http.HandleFunc("/", root)          // set the root route and its handler function
	http.HandleFunc("/welcome/", hello) // set the welcome route and its handler function
	//	fileserver := http.FileServer(http.Dir("httpFileServerData/"))  // define the FileServer file directory relatix to current code
	fileserver := http.FileServer(http.Dir("../testData/"))         // define the FileServer file directory relatix to current code
	http.Handle("/files/", http.StripPrefix("/files/", fileserver)) // set static asset route
	// http.Handle("/files/", http.FileServer(http.Dir("/httpFileServerData")))                              // set static asset route
	// http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir(".")))) // set static asset route
	http.ListenAndServe(":9001", nil) // set the port where the http server listens and serves
}
