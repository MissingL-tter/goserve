package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {

	args := os.Args
	path := "./"
	port := ":8080"

	if len(args) > 1 {
		arg := strings.Split(args[1], ":")

		if len(arg) > 0 {
			path = arg[0]
		}
		if len(arg) > 1 {
			port = ":" + arg[1]
		}
	}

	fmt.Printf("Starting server at %v on port %v\n", path, port)
	log.Fatal(http.ListenAndServe(port, http.FileServer(http.Dir(path))))
}
