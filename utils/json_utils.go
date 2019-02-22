package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func ReadAndUnmarshalFile(logger *log.Logger, file string, object interface{}) error {

	jsonFile, err := os.Open(file)
	// if we os.Open returns an error then handle it
	if err != nil {
		logger.Println(err)
		return err
	}
	logger.Println("successfully-opened-json-file")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal([]byte(byteValue), &object)
	return nil
}

func ReadFile(logger *log.Logger, file string) ([]byte, error) {

	jsonFile, err := os.Open(file)
	// if we os.Open returns an error then handle it
	if err != nil {
		logger.Println(err)
		return nil, err
	}
	logger.Println("successfully-opened-json-file")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	return byteValue, nil
}
