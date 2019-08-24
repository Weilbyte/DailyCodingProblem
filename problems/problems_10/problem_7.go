package problems_10

/* Problem Info
Asked by: Facebook

Given the mapping a = 1, b = 2, ... z = 26, and an encoded message, count the number of ways it can be decoded.
For example, the message '111' would give 3, since it could be decoded as 'aaa', 'ka', and 'ak'.
You can assume that the messages are decodable, for example, '001' is not allowed.

*/

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

//Add to unit test
func problem7solution(input string) {
	output := encode(input)
	fmt.Println(output)
	words := strings.Fields(output)
	for _, word := range words {
		fmt.Printf("Ways to solve %v - %v\n", word, decode(word))
	}

}

func decode(input string) int {
	if len(input) <= 1 {
		return 1
	} else {
		numb, _ := strconv.Atoi(input[:2])
		if (1 <= numb) && (numb <= 26) {
			return decode(input[1:]) + decode(input[2:])
		}
		return decode(input[1:])
	}
}

func encode(input string) string {
	var result string
	for _, x := range input {
		if unicode.IsSpace(x) {
			result += " "
			continue
		}
		if unicode.IsLetter(x) {
			result += string(chartonum(x))
		}
	}
	return result
}

func chartonum(char rune) string {
	charlower := strings.ToLower(string(char))
	result := int(charlower[0]) - 96
	return fmt.Sprintf("%d", result)
}
