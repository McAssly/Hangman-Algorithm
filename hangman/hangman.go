package main

import (
	"bufio"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"sort"
	"strings"
)

// TODO: POLISH IT DIM-WITT

// Words list of words in data form
type Words struct {
	Words []Word `json:"words"`
}

// Word a type for words used by algorithm
type Word struct {
	Content string `json:"word"`
	Usage   int    `json:"used"`
}

// DataPack of data
type DataPack struct {
	DataPack []Data `json:"pack"`
}

// Data a type for stored result data
type Data struct {
	NumberOfWords int     `json:"words"`
	SuccessRate   float64 `json:"chance"`
}

// Random: Get random value from 0->i
func random(i int) int64 {
	// Got this from somewhere, gonna be honest don't understand it
	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(i)))
	if err != nil { // Failure
		return -1
	}
	return nBig.Int64()
}

// Remove: remove element at i within Word array
func remove(i int, a []Word) []Word {
	// Remove the element at index i from a.
	copy(a[i:], a[i+1:]) // Shift a[i+1:] left one index.
	a[len(a)-1] = Word{} // Erase last element (write zero value).
	a = a[:len(a)-1]     // Truncate slice.
	return a
}

// Del: delete a string(f) within a string array(a)
func del(a []string, f string) []string {
	b := []string{} // init resulting array
	for i := 0; i < len(a); i++ {
		if a[i] != f { // if it isn't the f string
			b = append(b, a[i]) // add to resulting array
		}
	}
	return b
}

// List: convert string into string array
func list(s string) []string {
	a := []string{} // resulting array
	for i := 0; i < len(s); i++ {
		a = append(a, string(s[i])) // append character as string, to result
	}
	return a
}

// Compare: compare two different string arrays
func compare(a []string, b []string) bool {
	if len(a) != len(b) { // if their length is not equal
		return false // then they are obviously not the same
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] { // if there is one element unequal
			return false // then it is false
		}
	}
	return true // otherwise true
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

// Lengthen: lengthen an integer array based on given size
func lengthen(a []int, s int) []int {
	for i := 0; i < s-1; i++ {
		a = append(a, a[0])
	}
	return a
}

// Lengthen: lengthen a string array based on the given size
func lengthen2(a []string, s int) []string {
	for i := 0; i < s-1; i++ {
		a = append(a, a[0])
	}
	return a
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

func has2(a []string, x string) bool {
	for i := 0; i < len(a); i++ {
		if x == a[i] {
			return true
		}
	}
	return false
}

// Contains: detect if a string contains a byte
func contains(a string, x byte) bool {
	// Detects if a string contains a character (or sub string)
	for _, n := range a {
		if x == byte(n) { // if the byte is in there
			return true // then return true
		}
	}
	return false // otherwise false
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

// Algorithm: the JUICY PART
func algorithm(letters string, words Words, lose int, wordsInSeq []string) (int, Words) {
	// Output: [Some, Duck]
	masterLetters := list(letters)
	correctSoFar := []string{}
	for i := 0; i < len(wordsInSeq); i++ {
		correctSoFar = append(correctSoFar, strings.Join(lengthen2([]string{"-"}, len(wordsInSeq[i])), ""))
	}
	// Output: [----, ----]
	for i := 0; i < len(wordsInSeq); i++ {
		word := wordsInSeq[i]
		c := list(correctSoFar[i])
		if len(word) <= 0 {
			return 0, words
		}
		// MAKE a list of POSSIBLE words (based on length of unkown)
		listPossibleWords := []Word{}
		numberGuessed := lose
		for b := 0; b < len(words.Words); b++ {
			if len(words.Words[b].Content) == len(word) {
				listPossibleWords = append(listPossibleWords, words.Words[b])
			}
		}
		for {
			cc := []string{}
			for i := 0; i < len(c); i++ {
				if c[i] != "-" {
					cc = append(cc, string(c[i]))
				}
			}
			// IF a word in the possible words list doesnt have the letter in such position
			// THEN remove such word
			if len(cc) > 0 {
				listPossibleWords2 := []Word{}
				for i := 0; i < len(listPossibleWords); i++ {
					pw := listPossibleWords[i].Content
					numCC := 0
					for j := 0; j < len(pw); j++ {
						if c[j] != "-" {
							if c[j] == string(pw[j]) {
								numCC++
							}
						}
					}
					if numCC >= len(cc) {
						listPossibleWords2 = append(listPossibleWords2, listPossibleWords[i])
					}
				}
				listPossibleWords = listPossibleWords2
			}
			sort.Slice(listPossibleWords[:], func(i, j int) bool {
				return listPossibleWords[i].Usage > listPossibleWords[j].Usage
			})
			guess := ""
			/*
				Instead of random letter from possible letter list, pick random word
				from possible word list and pick a random letter from said word

				lets say 100 words
					> 99 words are :: Happy
					> 1 is :: Cool

						- with old sys :: 50/50 chance of H or C

						+ with new sys :: 99/1 chance of H over C
			*/
			if len(listPossibleWords) > 0 {
				guess = listPossibleWords[random(len(listPossibleWords))].Content
			} else {
				guess = words.Words[random(len(words.Words))].Content
			}
			listPossibleLetters := []string{}
			for i := 0; i < len(guess); i++ {
				for j := 0; j < len(masterLetters); j++ {
					if masterLetters[j][0] == guess[i] {
						listPossibleLetters = append(listPossibleLetters, string(guess[i]))
					}
				}
			}
			g := "-"
			if len(listPossibleLetters) <= 0 {
				g = string(masterLetters[random(len(masterLetters))])
			} else {
				g = listPossibleLetters[random(len(listPossibleLetters))]
			}
			numberGuessedSubtractor := 0
			maxLetters := 0
			for j := 0; j < len(wordsInSeq); j++ {
				for m := 0; m < len(wordsInSeq[j]); m++ {
					if wordsInSeq[j][m] == g[0] {
						okay := list(correctSoFar[j])
						okay[m] = g
						correctSoFar[j] = strings.Join(okay, "")
					} else {
						numberGuessedSubtractor++
					}
					maxLetters++
				}
			}
			if numberGuessedSubtractor > maxLetters-1 {
				numberGuessed--
			}
			listPossibleLetters = del(listPossibleLetters, g)
			masterLetters = del(masterLetters, g)
			fmt.Printf("\n%v :: %s\n", correctSoFar, g)
			if numberGuessed <= 0 {
				fmt.Printf("\nInput: %v", strings.Join(wordsInSeq, " "))
				fmt.Printf("\nYOU LOSE! %v\n", strings.Join(lengthen2([]string{"!"}, 25), ""))
				return 0, words
			}
			if has2(wordsInSeq, correctSoFar[i]) {
				break
			}
		}
	}
	fmt.Printf("\nInput: %v", strings.Join(wordsInSeq, " "))
	fmt.Printf("\nYOU WIN!  %v\n", strings.Join(lengthen2([]string{"$"}, 25), ""))
	return 1, words
}

func main() {
	words := getData()
	data := getData2()
	letters := "abcdefghijklmnopqrstuvwxyz"
	sampleSize := 100
	lose := 6
	// get input from user
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Sequence: ")
	text, _ := reader.ReadString('\n')
	// set the sequence to what the user gave, ex: Some Duck
	sequence := text
	// get words within the sequence
	wordsInSeq := []string{}
	index := 0
	for index < len(sequence) {
		// loop through sequence
		w := []string{} // init a word
		for index < len(sequence) {
			if sequence[index] == " "[0] || sequence[index] == "\n"[0] {
				// stop at a space or new line (word seperators)
				index++
				break
			}
			w = append(w, string(sequence[index])) // add letter to word
			index++
		}
		wordsInSeq = append(wordsInSeq, strings.Join(w, "")) // add word to words in sequence array
	}
	for i := 0; i < len(wordsInSeq); i++ {
		word := wordsInSeq[i]
		if !has(word, words.Words) {
			words.Words = append(words.Words, Word{word, 1})
		} else if has(word, words.Words) {
			words.Words = increaseUsage(words.Words, word)
		}
	}
	// Begin Algorithm
	times := 0
	for i := 0; i < sampleSize; i++ {
		result, rWords := algorithm(letters, words, lose, wordsInSeq)
		words = rWords
		if result == -1 {
			result, rWords = algorithm(letters, words, lose, wordsInSeq)
			words = rWords
		} else if result == 1 {
			times++
		}
	}
	var chance float64
	chance = float64(times) / float64(sampleSize) * 100
	fmt.Printf("\n\nTotal number of wins in %v iterations was: %v\n\n", sampleSize, times)
	fmt.Printf("\nChance of winning: %v%%\n\n", chance)
	data = setData(data, chance, len(wordsInSeq))
	// Update JSON data
	toData(words)
	toData2(data)
}
