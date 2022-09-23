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

	expectedResult := -1
	if position != expectedResult {
		t.Errorf("findFirstPosition want %d; get %d", expectedResult, position)
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

func TestPrepareWord(t *testing.T) {
	word := "testing"

	result := strings.Join(prepareWord(word), "")

	expectedResult := "t__t__g" // t _  _ t _ _ g
	if result != expectedResult {
		t.Errorf("prepareWord want %s; get %s", expectedResult, result)
	}

}
