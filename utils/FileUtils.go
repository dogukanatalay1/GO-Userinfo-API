package utils

import (
	"errors"
	"io/ioutil"
)

func ReadFile(fileName string) (string, error) {
  if IsEmpty(fileName) {
    return "", errors.New("Cannot be empty file name value!")
  }
  bytes, err := ioutil.ReadFile(fileName)
  CheckError(err)
  return string(bytes), nil
}
