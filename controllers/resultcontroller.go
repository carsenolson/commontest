package controllers

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"commontest/Test"
	"commontest/Config"
	"commontest/Result"
	"github.com/gorilla/mux"
)

type ResultContr struct {
	conf *Config.Config
}

func NewResultController(c *Config.Config) *ResultContr {
	gen := new(ResultContr)
    gen.conf = c
    return gen
}

func (rs *ResultContr) Origin(rw http.ResponseWriter, req *http.Request) {
	files, err := Result.ListResult(rs.conf.Result_path)
	if err != nil {
		fmt.Println(err)
	}
	tpl.ExecuteTemplate(rw, "origin.html", map[string]interface{}{"Files": files})
}

func (rs *ResultContr) Files(rw http.ResponseWriter, req *http.Request) {
	path_name := mux.Vars(req)["path"]
	fmt.Println(mux.Vars(req))
	fmt.Println("opening path", rs.conf.Result_path+"/"+path_name)
	files, err := Result.ListResult(rs.conf.Result_path+"/"+path_name+"/")
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}
	tpl.ExecuteTemplate(rw, "origin.html", map[string]interface{}{"Files": files})
}

func (rs *ResultContr) Result(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	res, err := Result.GetResultFromFile(rs.conf.Result_path+"/"+vars["path"], vars["file_name"])
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
	test, err := Test.NewTestFromFile(rs.conf.Test_path, res.File_name)
	if err != nil {
		fmt.Println(err)
	}
	err = tpl.ExecuteTemplate(rw, "result.html", map[string]interface{}{"Test": test, "Result": res})
	if err != nil {
		fmt.Println(err)
	}
}

func (rs *ResultContr) DeletePath(rw http.ResponseWriter, req *http.Request) {
	e := &struct {
		File_name string
	}{}
	data, _ := ioutil.ReadAll(req.Body)
    fmt.Println(string(data))
	json.Unmarshal(data, &e)
    fmt.Println(e)
	if err := Result.DeleteResultPath(rs.conf.Result_path+"/"+e.File_name); err != nil {
		fmt.Println(err)
	}
}

func (rs *ResultContr) DeleteRes(rw http.ResponseWriter, req *http.Request) {
	e := &struct {
		File_name string
	}{}
	path := mux.Vars(req)["path"]
	data, _ := ioutil.ReadAll(req.Body)
    fmt.Println(string(data))
	json.Unmarshal(data, &e)
    fmt.Println(e)
	if err := Result.DeleteResult(rs.conf.Result_path+"/"+path, e.File_name); err != nil {
		fmt.Println(err)
	}
}
