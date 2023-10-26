package main

import (
	"fmt"
	"net/http"
	"strings"
)

type MyHttpHandler struct{}

func (h MyHttpHandler) print_http_request_detail(w http.ResponseWriter, r *http.Request) {
	fmt.Println("::Client address   : ", strings.Split(r.RemoteAddr, ":")[0])
	fmt.Println("::Client port      : ", strings.Split(r.RemoteAddr, ":")[1])
	fmt.Println("::Request command  : ", r.Method)
	fmt.Println("::Request line     : ", r.RequestURI)
	fmt.Println("::Request path     : ", r.URL.Path)
	fmt.Println("::Request version  : ", r.Proto)
}

func (h MyHttpHandler) send_http_response_header() {

}

func (h MyHttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func (h MyHttpHandler) simpleCalc(para1, para2 int) int {
	return para1 * para2
}

func (h MyHttpHandler) parameter_retrieval(msg string) {
	//var result []int

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
