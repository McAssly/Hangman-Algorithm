package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// DataPack of data
type DataPack struct {
	DataPack []Data `json:"pack"`
}

// Data a type for stored result data
type Data struct {
	NumberOfWords int     `json:"words"`
	SuccessRate   float64 `json:"chance"`
}

// Set Data: set the new data to the old
func setData(data DataPack, chance float64, numberOfWords int) DataPack {
	for i := 0; i < len(data.DataPack); i++ {
		if data.DataPack[i].NumberOfWords == numberOfWords {
			data.DataPack[i].SuccessRate = (data.DataPack[i].SuccessRate + chance) / 2
			return data
		}
	}
	data.DataPack = append(data.DataPack, Data{numberOfWords, chance})
	return data
}

// Get Data: obtain data from JSON file, as usable STRUCT
func getData2() DataPack {
	jsonFile, err := os.Open("data.json")
	if err != nil {
		log.Fatal(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var data DataPack
	json.Unmarshal(byteValue, &data)
	return data
}

func toData2(data DataPack) {
	result, error := json.MarshalIndent(data, "", "    ")
	if error != nil {
		log.Fatal(error)
	}
	err := ioutil.WriteFile("data.json", result, 0644)
	if err != nil {
		log.Println(err)
	}
}
