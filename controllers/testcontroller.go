package controllers

import (
	"io"
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"commontest/Config"
	"commontest/Test"
	"github.com/gorilla/mux"
)

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
	json.Unmarshal(data, &e)
	err := e.Test.Save(tc.conf.Test_path, e.File_name)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("saved test")
	}
	io.WriteString(rw, "okay")
}

func (tc *TestController) ExistedTest(rw http.ResponseWriter, req *http.Request) {
	file_name := mux.Vars(req)["file_name"]
	test, err := Test.NewTestFromFile(tc.conf.Test_path, file_name)
	if err != nil {
		fmt.Println(err)
	}
	err = tpl.ExecuteTemplate(rw, "existedtest.html", map[string]interface{}{"File_name": file_name, "Test": test})
	if err != nil {
		fmt.Println(err)
	}
}
