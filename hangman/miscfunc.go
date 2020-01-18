package main

import (
	"crypto/rand"
	"math/big"
)

// Random: Get random value from 0->i
func random(i int) int64 {
	// Got this from somewhere, gonna be honest don't understand it
	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(i)))
	if err != nil { // Failure
		return -1
	}
	return nBig.Int64()
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
