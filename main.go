package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}

	cwd, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", http.FileServer(http.Dir(cwd)))
	fmt.Println("Listening to port " + port)
	fmt.Println("Serving path " + cwd)
	http.ListenAndServe(":"+port, nil)
}
