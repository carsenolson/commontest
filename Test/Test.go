package Test

import (
	//"io"
	"os"
	"encoding/json"
	"io/ioutil"
)

type Question struct {
	Title string
	Image, Answers []string
	True_answers []int
}

func NewQuestion(title string, images []string, answers []string, true_answers []int) *Question {
	q := new(Question)
	q.Title = title
	q.Image = images
	q.Answers = answers
	q.True_answers = true_answers
	return q
}

type Test struct {
	Name string
	Time int
	Questions []Question
}

func (t *Test) DeleteQuestion(index int) {
	copy(t.Questions[index:], t.Questions[index+0:])
	t.Questions[len(t.Questions)-1] = *new(Question) // or the zero value of T
	t.Questions = t.Questions[:len(t.Questions)-1]
}

func (t *Test) Save(path string) error {
	file, err := os.Create(path)
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

func GetAllTests(path string) ([]os.FileInfo, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	return files, nil
}

func NewTest(name string, time int, questions []Question) *Test {
	t := new(Test)
	t.Name = name
	t.Time = time
	t.Questions = questions
	return t
}

func NewTestFromFile(filename string) (*Test, error) {
	t := new(Test)
	// Read file from path. The path should be from config	
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(content, &t); err != nil {
		return nil, err
	}
	return t, nil
}

func DeleteTest(path string) error {
	err := os.Remove(path)
	if err != nil { return err }
	return nil
}
