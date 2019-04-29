package main

import (
	//"os"
	"fmt"
	"log"
	//"encoding/json"
	"net/http"
	"commontest/Config"
	"commontest/controllers"
	"github.com/gorilla/mux"
	//"net/http"
	//"hightest/controllers"
)

func main() {
	config, err := Config.InitConfig()
	if err != nil {
		fmt.Println(err)
	}
	r := mux.NewRouter()
	// Define two dirs to serve statically for fetching the images and styles/js	
	r.PathPrefix("/images/").Handler(http.StripPrefix("/images/", http.FileServer(http.Dir(config.Test_path+"/images/"))))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	r.HandleFunc("/", controllers.Index)
	log.Fatal(http.ListenAndServe(":8081", r))
}

