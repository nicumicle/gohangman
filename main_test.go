package hangman

import (
	"strings"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestFindFirstPosition(t *testing.T) {

	name := "Gladys"

	position := findFirstPosition("i", strings.Split(name, ""))

	if position != -1 {
		t.Errorf("findFirstPosition i in  = %s; want -1", name)
	}
}

func TestReplaceInWord(t *testing.T) {
	word := "abc"
	initialArray := strings.Split("___", "")

	result := strings.Join(replaceInWord("a", strings.Split(word, ""), initialArray), "")

	expectedResult := "a__"
	if result != expectedResult {
		t.Errorf("replaceInWord want %s; get %s", expectedResult, result)
	}
}
