package controllers

import (
	"fmt"
	"net/http"
	"commontest/Config"
	"commontest/Test"
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
