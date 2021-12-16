package pconfig

import (
	"encoding/json"
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

func LoadJson(path string, t interface{}) error {
	log.Println(t)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &t)
	if err != nil {
		return err
	}
	log.Println(t)
	return nil
}

func LoadYml(path string, t interface{}) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, t)
	if err != nil {
		return err
	}
	return nil
}
