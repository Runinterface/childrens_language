package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

const data_json = "https://vl-testovac.s3.amazonaws.com/challenges/childspeak_v2/test.in.json"

// "b","c","d","f","g","h","j","k","l","m","n","p","q","r","s","t","v","w","x","z"
// "a","e","i","o","u","y"

func consonantsOrNo(value rune) bool {
	result := false
	vowels := make(map[rune]struct{})
	vowels['a'] = struct{}{}
	vowels['e'] = struct{}{}
	vowels['i'] = struct{}{}
	vowels['o'] = struct{}{}
	vowels['u'] = struct{}{}
	vowels['y'] = struct{}{}

	_, ok := vowels[value]
	if ok {
		result = true
	}
	return result
}

func delChar(s []rune, index int, len int) []rune {
	fmt.Println(string(s))
	return append(s[0:index], s[index+1:len]...)
}

func childrens_language(word string) (ch_word string) {
	// Make rune from word
	strings.TrimSuffix(word, "\n")
	wordRune := []rune(word)
	fmt.Println(string(wordRune))
	save_index := 0
	lastVowel := 0

	for i := 0; i < len(wordRune); i++ {
		if consonantsOrNo(wordRune[i]) {
			if i != len(wordRune)-1 {
				if consonantsOrNo(wordRune[i+1]) {
					wordRune[i] = wordRune[0]
					wordRune = delChar(wordRune, i+1, len(wordRune))
				}
			}
		}
	}

	for i := 0; i < len(wordRune); i++ {
		if !consonantsOrNo(wordRune[i]) {
			if i != len(wordRune)-1 {
				if !consonantsOrNo(wordRune[i+1]) {
					wordRune[i] = wordRune[0]
					wordRune = delChar(wordRune, i+1, len(wordRune))
				}
			}
		}
	}

	for i := 0; i < len(wordRune); i++ {
		if consonantsOrNo(wordRune[i]) {
			continue
		} else {
			wordRune[0], wordRune[i] = wordRune[i], wordRune[0]
			save_index = i
			break
		}
	}

	for i := save_index; i < len(wordRune); i++ {
		if !consonantsOrNo(wordRune[i]) {
			wordRune[i] = wordRune[0]
		}
	}

	for i := 0; i < len(wordRune); i++ {
		if !consonantsOrNo(wordRune[i]) {
			continue
		} else {
			lastVowel = i
		}
	}
	for i := lastVowel; i < len(wordRune); i++ {
		if consonantsOrNo(wordRune[i]) {
			wordRune[len(wordRune)-1], wordRune[i] = wordRune[i], wordRune[len(wordRune)-1]
		}
	}
	return string(wordRune)
}

func main() {
	wordlist := make(map[string]int)

	// Get data from remote JSON file.
	res, err := http.Get(data_json)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	// Read data from body.
	var arr []string
	data, err := ioutil.ReadAll(res.Body)
	err = json.Unmarshal([]byte(data), &arr)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	} else {
		log.Println("Reading data - [Successfully!]")
	}
	for i := 0; i < len(arr); i++ {
		current_word := childrens_language(arr[i])
		wordlist[current_word]++
		_, ok := wordlist[current_word]
		if ok {
			wordlist[current_word]++
		}
	}

	jsonStr, err := json.Marshal(wordlist)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		fmt.Println(string(jsonStr))
	}

	resultOutFile, err := os.OpenFile("./result.out", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	defer resultOutFile.Close()

}
