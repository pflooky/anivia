package functions

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func UnmarshalFile(file string, t interface{}) error {
	content, err := ReadFile(file)
	if err != nil {
		return errors.New(fmt.Sprintf("[file]: %s, [error]: %v", file, err))
	}
	err = UnmarshalYaml(content, t)
	if err != nil {
		return errors.New(fmt.Sprintf("[file]: %s, [error]: %v", file, err))
	}
	return nil
}

func ReadFile(file string) (*[]byte, error) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error reading file with msg %v", err))
	}

	return &content, nil
}

func UnmarshalYaml(content *[]byte, t interface{}) error {
	err := yaml.Unmarshal(*content, t)
	if err != nil {
		return errors.New(fmt.Sprintf("error unmarshalling to %T, with error %v", t, err))
	}
	return nil
}
