package Config

import (
	"os"
	"fmt"
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Test_path, Result_path string
}

func (c *Config) Commit() error {
	err := os.Mkdir(c.Test_path, 0700)
	if err != nil {
		return err
	}
	err = os.Mkdir(c.Test_path+"/images", 0700)
	if err != nil {
		return err
	}
	err = os.Mkdir(c.Result_path, 0700)
	if err != nil {
		return err
	}
	return nil
}

func (c *Config) UpdateTestPath(path string) {
	c.Test_path = path
}

func (c *Config) UpdateResultPath(path string) {
	c.Result_path = path
}

func (c *Config) Save() error {
	file, err := os.Create("config.json")
	if err != nil {
		return err
	}
	defer file.Close()
	data, err := json.Marshal(c)
	if err != nil {
		return nil
	}
	_, err = file.Write(data)
	if err != nil {
		return nil
	}
	return nil
}

func InitConfig() (*Config, error) {
	conf := new(Config)
	if _, err := os.Stat("config.json"); err == nil {
		content, err := ioutil.ReadFile("config.json")
		if err != nil {
			return nil, err
		}
		if err = json.Unmarshal(content, &conf); err != nil {
			return nil, err
		}
		conf.Commit()
		return conf, nil
	} else if os.IsNotExist(err) {
		conf.Test_path = "tests"
		conf.Result_path = "results"
		file, err := os.Create("config.json")
		if err != nil {
			return nil, err
		}
		defer file.Close()
		data, err := json.Marshal(conf)
		if err != nil {
			return nil, err
		}
		_, err = file.Write(data)
		if err != nil {
			return nil, err
		}
		conf.Commit()
		return conf, nil
	} else {
		fmt.Println(err)
	}
	return conf, nil
}
