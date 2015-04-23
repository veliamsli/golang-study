package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", helloHandle)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("listen and server: ", err.Error())
	}

}

func helloHandle(rw http.ResponseWriter, req *http.Request) {
	io.WriteString(rw, "Hello World!")
}
