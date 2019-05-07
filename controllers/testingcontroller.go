package controllers

import (
	"io"
	"io/ioutil"
	"fmt"
	"strconv"
	"net/http"
	"encoding/json"
	"commontest/Test"
	"commontest/Config"
	"commontest/Result"
)

type TestingController struct {
	conf *Config.Config
	srv *http.Server
}

func NewTestingController(c *Config.Config) *TestingController {
	tc := new(TestingController)
	tc.conf = c
	return tc
}

func (tc *TestingController) Auth(rw http.ResponseWriter, req *http.Request) {
	tests, err := Test.GetAllTests(tc.conf.Test_path)
	if err != nil {
		fmt.Println(err)
	}
	tpl.ExecuteTemplate(rw, "auth.html", map[string]interface{}{ "Tests": tests })
}

func (tc *TestingController) StartTesting(rw http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	test, err := Test.NewTestFromFile(tc.conf.Test_path, req.Form["test"][0])
	if err != nil {
		fmt.Println(err)
	}
	err = tpl.ExecuteTemplate(rw, "starttesting.html", map[string]interface{}{
			"Test": test,
			"File_name": req.Form["test"][0],
			"Full_name": req.Form["fullname"][0],
			"Group": req.Form["group"][0]})
	if err != nil {
		fmt.Println(err)
	}
}

func compare(X, Y []int) []int {
	m := make(map[int]int)

	for _, y := range Y {
		m[y]++
	}

	var ret []int
	for _, x := range X {
		if m[x] > 0 {
			m[x]--
			continue
		}
		ret = append(ret, x)
	}

	return ret
}

func (tc *TestingController) Result(rw http.ResponseWriter, req *http.Request) {
	e := new(Result.Result)
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(data, &e)
	if err != nil {
		fmt.Println(err)
	}
	test, err := Test.NewTestFromFile(tc.conf.Test_path, e.File_name)
	if err != nil {
		fmt.Println(err)
	}
	commonLength := 0
	mistakes :=0
	true_answers := 0
	right_answers := 0
	picked_answers := 0
	fmt.Println(e.Picked_answers)
	for index, question := range test.Questions {
		commonLength += len(question.Answers)
		true_answers += len(question.True_answers)
		picked_answers += len(e.Picked_answers[index])
		right_answers += len(compare(e.Picked_answers[index], question.True_answers))
		mistakes += len(compare(question.True_answers, e.Picked_answers[index]))
	}
	fmt.Println("100 / questoinAmount: ", float64(100) / float64(commonLength))
	fmt.Println(float64(100) / float64(commonLength), " * ", float64(mistakes+right_answers))
	fmt.Println("100 - ", int((float64(100) / float64(commonLength)) * float64(mistakes+right_answers)))
	if picked_answers != 0 {
		e.Result = 100 - int((float64(100) / float64(commonLength)) * float64(mistakes+right_answers))
	} else {
		e.Result = 0
	}
	//e.Result = int((float64(mistakes+right_answers)/float64(commonLength) * float64(100) - float64(100)) * float64(-1))
	err = e.Save(tc.conf.Result_path)
	if err != nil {
		fmt.Println(err)
	}
	io.WriteString(rw, strconv.Itoa(e.Result))
}

