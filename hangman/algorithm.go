package main

import (
	"fmt"
	"sort"
	"strings"
)

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
