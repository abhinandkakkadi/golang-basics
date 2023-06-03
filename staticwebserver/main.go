package main

import (
	"net/http"
)

func main() {
	// a router is created
	mux := http.NewServeMux()
	// then a handler is returned which should be the second parameter to mux.Handler since it need a type which implements http.Handler interface
	file := http.FileServer(http.Dir("public"))
	mux.Handle("/", file)
	// we are setting the port to listen to all the above mentioned routes
	http.ListenAndServe(":8080", mux)

	// here the static website was able to show because the name of html file was index.html
	// if the name was not index.html it shows the website in form a file and an option to select each of the file present inside the directory public
}
