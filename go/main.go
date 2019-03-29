package main

import (
	"log"
	"regexp"
	"strings"
)

var words = []string{"kayak", "river", "civic", "level", "paper", "Repaper"}

var phrases = []string{"Don't nod.", "I did, did I?", "Who am I?", "Was it a cat I saw?"}

func main() {
	log.Printf("Words: %v\n", palindromes(words))
	log.Printf("Phrases: %v\n", palindromes(phrases))
}

func palindromes(elements []string) []string {
	regex, err := regexp.Compile("[^a-zA-Z]+")
	if err != nil {
		log.Fatal(err)
	}

	var palindromes []string
	for _, element := range elements {
		stripped := strings.ToLower(regex.ReplaceAllString(element, ""))
		endIndex := len(stripped) - 1
		palindrome := true
		for i := 0; i < len(stripped); i++ {
			if i == (endIndex - i) {
				break
			}
			if stripped[i] != stripped[endIndex-i] {
				palindrome = false
				break
			}
		}
		if palindrome {
			palindromes = append(palindromes, element)
		}
	}

	return palindromes
}
