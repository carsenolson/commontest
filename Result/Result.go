package Result

import (
	"os"
	"io/ioutil"
	"time"
	"encoding/json"
	"strconv"
)

type Result struct {
	Full_name string
	File_name string
	Group string
	Picked_answers [][]int
	Result int
}

func ListResult(path string) (dirs []os.FileInfo, err error) {
	files, err := ioutil.ReadDir(path)
    if err != nil {
         return nil, err
    }
	return files, nil
}

func DeleteResultPath(path string) error {
	if err := os.RemoveAll(path); err != nil {
		return err
	}
	return nil
}

func DeleteResult(path, file_name string) error {
	if err := os.Remove(path+"/"+file_name); err != nil {
		return err
	}
	return nil
}

func GetResultFromFile(path, filename string) (*Result, error) {
	r := new(Result)
    content, err := ioutil.ReadFile(path+"/"+filename)
    if err != nil {
	   return nil, err
    }
    if err := json.Unmarshal(content, &r); err != nil {
        return nil, err
    }
    return r, nil
}

func (r *Result) Save(path string) error {
	dt := time.Now()
	dirName := dt.Format("01_02_2006")
	err := os.MkdirAll(path+"/"+dirName, 0700)
	if err != nil {
		return err
	}
	filename := r.Full_name+"_"+r.Group+"_"+r.File_name+"_"+strconv.Itoa(r.Result)+"_"+strconv.Itoa(dt.Hour())+"_"+strconv.Itoa(dt.Minute())+"_"+strconv.Itoa(dt.Second())
	file, err := os.Create(path+"/"+dirName+"/"+filename)
	if err != nil {
		return err
	}
	defer file.Close()
	data, err := json.Marshal(r)
	if err != nil {
		return err
	}
	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}
