package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// Words list of words in data form
type Words struct {
	Words []Word `json:"words"`
}

// Word a type for words used by algorithm
type Word struct {
	Content string `json:"word"`
	Usage   int    `json:"used"`
}

// Remove: remove element at i within Word array
func remove(i int, a []Word) []Word {
	// Remove the element at index i from a.
	copy(a[i:], a[i+1:]) // Shift a[i+1:] left one index.
	a[len(a)-1] = Word{} // Erase last element (write zero value).
	a = a[:len(a)-1]     // Truncate slice.
	return a
}

// Increase Usage: increase the 'used' data within the Words array
func increaseUsage(words []Word, word string) []Word {
	for i := 0; i < len(words); i++ {
		if word == words[i].Content {
			words[i].Usage++ // increase the used, for the given word
		}
	}
	return words // re-return the words array, that has been updated
}

// Has: detect if the words array HAS the given word within it
func has(word string, words []Word) bool {
	for i := 0; i < len(words); i++ {
		if word == words[i].Content {
			return true
		}
	}
	return false
}

// Get Data: obtain the data from the JSON file, as usable STRUCT
func getData() Words {
	jsonFile, err := os.Open("words.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var words Words

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &words)
	return words
}

// To Data: write words to data file
func toData(words Words) {
	result, error := json.MarshalIndent(words, "", "	") // transform into json data
	if error != nil {                                   // if it can't
		log.Println(error) // then return error
	}
	err := ioutil.WriteFile("words.json", result, 0644) // write to the JSON file
	if err != nil {
		log.Println(err)
	}
}
