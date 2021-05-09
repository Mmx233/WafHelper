package util

import (
	"io/ioutil"
)

type file struct{}

var File file

func (a *file) Read(path string) (string, error) {
	t, e := ioutil.ReadFile(path)
	return string(t), e
}

func (a *file) Write(path string, data string) error {
	return ioutil.WriteFile(path, []byte(data), 700)
}
