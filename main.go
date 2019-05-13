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

	general := controllers.NewGeneral(config)
	test := controllers.NewTestController(config)
	result := controllers.NewResultController(config)

	r := mux.NewRouter()

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

	r.HandleFunc("/", general.Index).Methods("GET")
	r.HandleFunc("/", general.ActionHandle).Methods("POST")
	r.HandleFunc("/help", controllers.Help).Methods("GET")
	r.HandleFunc("/newtest", test.NewTest).Methods("GET")
	r.HandleFunc("/newtest", test.SaveTest).Methods("POST")
	r.HandleFunc("/test/{file_name}", test.ExistedTest).Methods("GET")
	r.HandleFunc("/results/", result.Origin).Methods("GET")
	r.HandleFunc("/results/", result.HandleResultPath).Methods("POST")
	r.HandleFunc("/results/{path}/", result.Files).Methods("GET")
	r.HandleFunc("/results/{path}/", result.DeleteRes).Methods("POST")
	r.HandleFunc("/results/{path}/{file_name}/", result.Result).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", r))
}
