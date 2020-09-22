package lineageclass

import (
	"encoding/json"
	"io/ioutil"
)

type Class struct {
	Name             string `json:"name"`
	Description      string `json:"description"`
	Pros             string `json:"pros"`
	Cons             string `json:"cons"`
	Race             string `json:"race"`
	StarterClassName string `json:"starterClassName"`
	FirstClassName   string `json:"firstClassName"`
	SecondClassName  string `json:"secondClassName"`
	ThirdClassName   string `json:"thirdClassName"`
}

func LoadClass() ([]*Class, error) {

	data, err := ioutil.ReadFile("classdata.json")

	if err != nil {
		return nil, err
	}

	listClass := []*Class{}

	err = json.Unmarshal([]byte(data), &listClass)

	if err != nil {
		return nil, err
	}

	return listClass, nil
}
