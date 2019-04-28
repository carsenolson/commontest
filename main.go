package main

import (
	"os"
	"fmt"
	"log"
	"encoding/json"
	//"net/http"
	//"hightest/controllers"
)

func main() {
	fmt.Println("Hello world!")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

