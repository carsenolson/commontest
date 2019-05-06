package controllers

import (
	"io"
	"fmt"
	"net"
	"time"
	"net/http"
	"context"
	"commontest/Config"
	"commontest/Test"
	"github.com/gorilla/mux"
	"io/ioutil"
	"encoding/json"
)

type General struct {
	conf *Config.Config
	c chan int
	isTestRunning bool
	tc *TestingController
}

func GetOutboundIP() net.IP {
    conn, err := net.Dial("udp", "8.8.8.8:80")
    if err != nil {
        fmt.Println(err)
    }
    defer conn.Close()
    localAddr := conn.LocalAddr().(*net.UDPAddr)
    return localAddr.IP
}

func NewGeneral(c *Config.Config) *General {
	gen := new(General)
	gen.conf = c
	gen.isTestRunning = false
	return gen
}

func (gen *General) Index(rw http.ResponseWriter, req *http.Request) {
	tests, err := Test.GetAllTests(gen.conf.Test_path)
	if err != nil {
		fmt.Println(err)
	}
	pageData := map[string]interface{} {
		"Config": gen.conf,
		"Tests": tests,
		"isTestRunning": gen.isTestRunning,
	}
	tpl.ExecuteTemplate(rw, "index.html", pageData)
}

func (gen *General) StartTesting() {
	gen.tc = NewTestingController(gen.conf)
	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	r.HandleFunc("/", gen.tc.Auth).Methods("GET")
	r.HandleFunc("/", gen.tc.StartTesting).Methods("POST")
	r.HandleFunc("/result", gen.tc.Result).Methods("POST")
	gen.tc.srv = &http.Server{
        Handler:      r,
		Addr:         "0.0.0.0:8080",
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }
	err := gen.tc.srv.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
	<-gen.c
}

func (gen *General) StopTesting() {
	err := gen.tc.srv.Shutdown(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	<-gen.c
}

func (gen *General) ActionHandle(rw http.ResponseWriter, req *http.Request) {
	e := &struct{
		Action string
        Option string
    }{}
    data, _ := ioutil.ReadAll(req.Body)
    fmt.Println(string(data))
    json.Unmarshal(data, &e)
	fmt.Println(e)
	switch act := e.Action; act {
		case "deleteTest":
			err := Test.DeleteTest(gen.conf.Test_path, e.Option)
			if err != nil {
				fmt.Println(err)
			}
		case "startTesting":
			gen.isTestRunning = true
			io.WriteString(rw, "http://"+GetOutboundIP().String()+":8080/")
			go gen.StartTesting()
		case "stopTesting":
			fmt.Println("got stop testing")
			gen.isTestRunning = false
			gen.StopTesting()
		default:
			fmt.Println("don't know :(")
	}
}

func Help(rw http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(rw, "help.html", nil)
}
