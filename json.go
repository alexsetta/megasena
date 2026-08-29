package main

import (
	"fmt"
	"io/ioutil"
)

func ReadJson(fileName string) (string, error) {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", fmt.Errorf("ReadJson: %w", err)
	}
	return string(b), nil
}
