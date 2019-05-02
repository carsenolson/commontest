package controllers

import (
	"net/http"
	"commontest/Config"
	"commontest/Test"
)

var _ = Test.NewTest("hello", 20)

type TestController struct {
	conf *Config.Config
}

func NewTestController(c *Config.Config) *TestController {
	tc := new(TestController)
	tc.conf = c
	return tc
}

func (tc *TestController) NewTest(rw http.ResponseWriter, req *http.Request) {
	pageData := map[string]interface{}{}
	tpl.ExecuteTemplate(rw, "newtest.html", pageData)
}
