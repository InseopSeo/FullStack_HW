package main

import (
	"fmt"
	"net/http"
)

type MyHttpHandler struct{}

func (h MyHttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	server_name := "localhost"
	server_port := 8080
	address := fmt.Sprintf("%s%s%d", server_name, ":", server_port)

	http.Handle("/", MyHttpHandler{})

	fmt.Printf("HTTP server started at http://%s:%d\n", server_name, server_port)

	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println("Error:", err)
	}

}
