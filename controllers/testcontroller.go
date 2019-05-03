package controllers

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
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

func (tc *TestController) SaveTest(rw http.ResponseWriter, req *http.Request) {
	e := &struct{
		File_name string
		Test Test.Test
	}{}
	data, _ := ioutil.ReadAll(req.Body)
	fmt.Println(string(data))
	json.Unmarshal(data, &e)
	e.Test.Save(tc.conf.Test_path, e.File_name+".json")
}
