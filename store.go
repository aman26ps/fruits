package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var fruits []Fruit

func getNextID() int {
	if len(fruits) == 0 {
		return 1
	}
	return fruits[len(fruits)-1].ID + 1
}

func loadFruits() {
	file, err := os.ReadFile("fruits.json")
	if err != nil {
		fruits = []Fruit{}
		return
	}
	json.Unmarshal(file, &fruits)
}

func saveFruits() {
	data, _ := json.MarshalIndent(fruits, "", "  ")
	ioutil.WriteFile("fruits.json", data, 0644)
}
