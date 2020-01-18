# Hangman Algorithm

WARNING: do not use capitol letters or symbols for the input

# Method 1 
Make a list of possible words based on the length of given unkowns
```go
listPossibleWords := []Word{}
for b := 0; b < len(words.Words); b++ {
  if len(words.Words[b].Content) == len(word) {
    listPossibleWords = append(listPossibleWords, words.Words[b])
  }
}
```

# Method 2 
(once guessed correctly) If a word in the list of possible words does not contain the same letter in that position, Then remove that word as possibility
```go
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
```
  
# Method 3
Store data on how many times a word has been used (given) and based on that data, increase in the possible words list the amount of said word based on its usage amount
```go
listPossibleWords := []Word{}
for b := 0; b < len(words.Words); b++ {
  if len(words.Words[b].Content) == len(word) {
    for n := 0; n < words.Words[b].Usage; n++ {
      listPossibleWords = append(listPossibleWords, words.Words[b])
    }
  }
}
```
  
# Method 4
Instead of a random letter from possible letter list, pick a random word from the possible word list and picks a random letter from the selected word

```go
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
```

For example:
* Say there are 100 words
  * 99 are happy
  * 1 is cool
  * 99:1 of selecting happy over selecting cool
