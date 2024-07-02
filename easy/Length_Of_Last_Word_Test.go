package easy

import (
	"fmt"
	"testing"
)

func Test_Length_Of_Last_Word(t *testing.T) {
	input1 := "Hello World"
	input2 := "   fly me   to   the moon  "
	input3 := "luffy is still joyboy"
	fmt.Println(Length_Of_Last_Word(input1))
	fmt.Println(Length_Of_Last_Word(input2))
	fmt.Println(Length_Of_Last_Word(input3))
}

func TestReverseWord(t *testing.T) {
	input1 := "Hello World"
	input2 := "   fly me   to   the moon  "
	input3 := "luffy is still joyboy"
	fmt.Println(reverseWord(input1))
	fmt.Println(reverseWord(input2))
	fmt.Println(reverseWord(input3))
}
