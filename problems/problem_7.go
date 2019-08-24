package problems

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

//Problem7 prints a brief info about the problem and runs the solution
func Problem7() {
	problemInfo("Facebook", "Given the mapping a = 1, b = 2, ... z = 26, and an encoded message, count the number of ways it can be decoded. For example, the message '111' would give 3, since it could be decoded as 'aaa', 'ka', and 'ak'. You can assume that the messages are decodable. For example, '001' is not allowed.")
	fmt.Println("Message to encode: ")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	problem7solution(input.Text())
}

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
