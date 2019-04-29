package Test

import (
	"io"
	"os"
	"encoding/json"
	"io/ioutil"
)

type Question struct {
	Title string
	Image, Answers []string
	True_answers []int
}

type Test struct {
	Name string
	Time int
	Questions []Question
}

func (t *Test) AddQuestion(title string, images []string, answers []string, true_answers []int) {
	q := new(Question)
	q.Title = title
	q.Image = images
	q.Answers = answers
	q.True_answers = true_answers
	t.Questions = append(t.Questions, *q)
}

func (t *Test) DeleteQuestion(index int) {
	copy(t.Questions[index:], t.Questions[index+0:])
	t.Questions[len(t.Questions)-1] = *new(Question)
	t.Questions = t.Questions[:len(t.Questions)-1]
}

// there is no official way to copy file in golang, so I fetched that from stackoverflow 
func CopyFileContents(src, dst string) (err error) {
    in, err := os.Open(src)
    if err != nil {
        return
    }
    defer in.Close()
    out, err := os.Create(dst)
    if err != nil {
        return
    }
    defer func() {
        cerr := out.Close()
        if err == nil {
            err = cerr
        }
    }()
    if _, err = io.Copy(out, in); err != nil {
        return
    }
    err = out.Sync()
    return
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
	// analyze questions images and copy them into config.test_path/images 
	// directory for allowing import tests
	// TODO: optmize this process, it doesn't check if file already exists
	for _, value := range t.Questions {
		for _, image := range value.Image {
			err = CopyFileContents(image, path+"/images/"+image)
			if err != nil {
				return err
			}
		}
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

func NewTest(name string, time int) *Test {
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
	test, err := NewTestFromFile(path, file_name)
	if err != nil {
		return err
	}
	err = os.Remove(path+"/"+file_name)
	if err != nil { return err }
	for _, value := range test.Questions {
		for _, image := range value.Image {
			err = os.Remove(path+"/images/"+image)
			// Skip because of error "file doesn't exist" 	
			// It should not be like that, but it optimizes the process of 
			// deleting same files if they occurs in the same test	
			if err != nil {
				continue
			}
		}
	}
	return nil
}
