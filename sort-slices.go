package main

import (
	"fmt"
	"strings"
)

func main() {
	numbers := []int{2, 3, 6, 7, 5, 4, 1}
	words := []string{"aaaa", "aaaab", "apple", "boy", "aaa", "aba", "aaaba", "bake", "cake", "cast", "bask", "stew", "stee"}

	sortSlice(numbers)
	sortSlice(words)
}

//check type of interface
func sortSlice(i interface{}) {
	var sortedAny interface{}
	switch i.(type) {
	case []int:
		sortedAny = sortIntSlice(i.([]int))
	case []string:
		sortedAny = sortString(i.([]string))
	default:
		fmt.Println("Type not supported")
	}
	fmt.Println(sortedAny)
}

// sort slices of ints
func sortIntSlice(numbers []int) []int {
	lengthOfN := len(numbers)

	for comparisons := 0; comparisons < lengthOfN; comparisons++ { //outer loop keeps count of passes
		for i := 0; i < lengthOfN-comparisons-1; i++ { //inner loop sorts through till the last sorted element in the previous pass
			if numbers[i] > numbers[i+1] {
				numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
			}
		}
	}
	return numbers
}

//sort slices of strings
func sortString(words []string) []string {
	lengthOfWords := len(words)
	var firstWord string
	var secondWord string

	for comparisons := 0; comparisons < lengthOfWords; comparisons++ {
		for i := 0; i < lengthOfWords-comparisons-1; i++ {
			firstWord = strings.ToLower(words[i])
			secondWord = strings.ToLower(words[i+1])

			//take the shorter of both strings and use its length to compare the characters of both strings,
			//this is necessary in order to avoid index out of bounds
			var maxCharComparison int
			if len(firstWord) < len(secondWord) {
				maxCharComparison = len(firstWord)
			} else {
				maxCharComparison = len(secondWord)
			}

			//compare individual characters in both strings to determine order
			for c := 0; c < maxCharComparison; c++ {
				if firstWord[c] > secondWord[c] { //swap words if the character being compared in both words are in wrong order
					words[i], words[i+1] = words[i+1], words[i]
					break
				} else if firstWord[c] < secondWord[c] { // if the characters in the first word is in the correct order break from the loop, already in correct order
					break
				} else if c == maxCharComparison-1 && maxCharComparison == len(secondWord) { //swap if both words have same characters but the succeeding word is shorter
					words[i], words[i+1] = words[i+1], words[i]
				}
			}
		}
	}
	return words
}
