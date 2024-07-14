package main

import (
	"encoding/json"
	"os"
	"strings"
)

type ConfigData struct {
	UserName           string
	AdditionalProducts []Product
}

var Config ConfigData

func LoadConfig() (err error) {
	data, err := os.ReadFile("config.json")

	if err == nil {
		// Printfln(string(data))
		decoder := json.NewDecoder(strings.NewReader(string(data)))
		err = decoder.Decode(&Config)
	}

	return
}

func LoadConfigUsingFileStruct() (err error) {
	file, err := os.Open("config.json")

	if err == nil {
		defer file.Close()
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&Config)
	}

	return
}

func LoadConfigUsingSpecificLocation() (err error) {
	file, err := os.Open("config.json")

	if err == nil {
		defer file.Close()

		nameSlice := make([]byte, 5)
		file.ReadAt(nameSlice, 17)
		// Printfln("NameSlice: %v", string(nameSlice))
		Config.UserName = string(nameSlice)

		file.Seek(48, 0)

		decoder := json.NewDecoder(file)

		err = decoder.Decode(&Config.AdditionalProducts)
	}

	return
}

func init() {
	// err := LoadConfig()
	// err := LoadConfigUsingFileStruct()
	err := LoadConfigUsingSpecificLocation()

	if err != nil {
		Printfln("Error Loading Config: %v", err.Error())
	} else {
		Printfln("Username: %v", Config.UserName)
		Products = append(Products, Config.AdditionalProducts...)
	}
}
