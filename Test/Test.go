package Test

import (
	//"io"
	"os"
	"encoding/json"
	"io/ioutil"
)

type Question struct {
	Title string
	Image [][]string
	Answers []string
	True_answers []int
}

type Test struct {
	Name string
	Time string
	Questions []Question
}

func (t *Test) AddQuestion(title string, images [][]string, answers []string, true_answers []int) {
	q := new(Question)
	q.Title = title
	q.Image = images
	q.Answers = answers
	q.True_answers = true_answers
	t.Questions = append(t.Questions, *q)
}

func (t *Test) Save(path, filename string) error {
	// create and write json to file	
	file, err := os.Create(path+"/"+filename)
	if err != nil {
		return err
	}
	defer file.Close()
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}
	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func GetAllTests(path string) (tests []os.FileInfo, err error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		tests = append(tests, file)
	}
	return tests, nil
}

func NewTest(name, time string) *Test {
	t := new(Test)
	t.Name = name
	t.Time = time
	t.Questions = *new([]Question)
	return t
}

func NewTestFromFile(path, filename string) (*Test, error) {
	t := new(Test)
	// Read file from path. The path should be from config	
	content, err := ioutil.ReadFile(path+"/"+filename)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(content, &t); err != nil {
		return nil, err
	}
	return t, nil
}

func DeleteTest(path, file_name string) error {
	err := os.Remove(path+"/"+file_name)
	if err != nil { return err }
	return nil
}
