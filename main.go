package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	var letter string
	var wordArray []string
	var word string

	colorReset := "\033[0m"
	colorGreen := "\033[32m"
	colorRed := "\033[31m"
	colorBlue := "\033[34m"
	colorYellow := "\033[33m"

	max_tries := 6
	wordList := []string{
		"linux",
		"github",
		"golang",
		"i-am-learning-golang-on-my-own",
	}

	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(wordList))

	word = wordList[index]

	tries := max_tries

	wordArray = PrepareWord(word)
	for {
		fmt.Print(colorYellow)
		fmt.Println("Tries left:", tries)
		drawHangman(max_tries - tries)
		fmt.Print(colorReset)

		fmt.Println("Your word is:", strings.Join(wordArray, " "))

		if strings.Join(wordArray, "") == word {
			fmt.Println(colorGreen, "\tYou won.", colorReset)
			break
		}
		if tries == 0 {
			fmt.Println(colorRed, "\tYou lost.\t", colorReset)
			fmt.Println("\tThe word is: ", colorBlue, word, colorReset)
			break
		}

		fmt.Print("Add a letter: ")
		fmt.Scanf("%s", &letter)

		pos := FindFirstPosition(letter, strings.Split(word, ""))
		if pos < 0 {
			tries--
		} else {
			wordArray = ReplaceInWord(letter, strings.Split(word, ""), wordArray)
		}
	}

}

/*
	Draw Hagman based on how many failed tries user has
*/
func drawHangman(step int) {

	if step >= 1 {
		fmt.Print("\t\t\t  0\t\t\t") //head
	}
	fmt.Println("\r") //newline and \r
	if step >= 2 {
		fmt.Print("\t\t\t /") //left arm
	}
	if step >= 3 {
		fmt.Print("|") //body
	}
	if step >= 4 {
		fmt.Print("\\ ") //right arm
	}
	fmt.Println("") //newline
	if step >= 5 {
		fmt.Print("\t\t\t /") //left leg
	}
	if step >= 6 {
		fmt.Print(" \\") //right leg
	}

	fmt.Println("") //newline
}

/*
	Replace character in word
*/
func ReplaceInWord(character string, word []string, array []string) []string {

	for i := 0; i < len(word); i++ {
		if word[i] == character {
			array[i] = character
		}
	}

	return array
}

/*
	Prepare the word for the first display
*/
func PrepareWord(word string) []string {
	var result []string

	chars := strings.Split(word, "")
	result = make([]string, len(chars))
	var i int
	for i = 0; i < len(chars); i++ {
		if chars[i] == "-" {
			result[i] = chars[i]
		} else {
			result[i] = "_"
		}

	}
	result = ReplaceInWord(chars[0], chars, result)
	result = ReplaceInWord(chars[len(chars)-1], chars, result)

	return result
}

/*
	Finds the position of a character
*/
func FindFirstPosition(character string, myArray []string) int {
	var i int
	for i = 0; i < len(myArray); i++ {
		if myArray[i] == character {
			return i
		}
	}

	return -1
}
