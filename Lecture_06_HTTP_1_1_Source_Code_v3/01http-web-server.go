package main

import (
	_ "fmt"
	"net/http"
)

func main() {
	//server_name := "localhost"
	//sever_port:=8080

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Hello World"))
	})

	http.ListenAndServe(":8080", nil)

}
