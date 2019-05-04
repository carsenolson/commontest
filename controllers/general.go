package controllers

import (
	"fmt"
	"net/http"
	"commontest/Config"
	"commontest/Test"
	"io/ioutil"
	"encoding/json"
)

type General struct {
	conf *Config.Config
}

func NewGeneral(c *Config.Config) *General {
	gen := new(General)
	gen.conf = c
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
	}
	tpl.ExecuteTemplate(rw, "index.html", pageData)
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
		default:
			fmt.Println("don't know :(")
	}
}
