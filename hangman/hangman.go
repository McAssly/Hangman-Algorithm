package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// TODO: POLISH IT DIM-WITT

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
