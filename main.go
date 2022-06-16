package main

import (
	"fmt"
	"log"
)

func main() {
	server := configure()

	fmt.Println(">>> Starting server.")
	log.Fatal(server.ListenAndServe())
}
