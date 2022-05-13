package main

import "net/http"

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/products", productsHandler)
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the Go WebServer!"))
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Write([]byte("All the products will be served"))
	case "POST":
		w.Write([]byte("The given new product will be added"))
	case "PUT":
		/*
			get the id of the product from the URI (route parameter)
			read the data from the request stream
			deserialize the data into Go objects
			Manipulate
			serrialize the response
			send the response
		*/
		w.Write([]byte("The given product will be updated"))
	case "DELETE":
		w.Write([]byte("The given new product will be removed"))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
