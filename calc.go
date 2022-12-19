package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("Format: X arithmetic_operation Y. Example: 1 + 4 for arabic or X + VI for roman.")

	reader := bufio.NewReader(os.Stdin)

	var digitCheck = regexp.MustCompile(`^([1-9]|10)[+*/-]([1-9]|10)$`)
	var romanCheck = regexp.MustCompile(`^(IX|X|IV|V?I{0,3})?[+*/-](IX|X|IV|V?I{0,3})?$`)

	for {
		var isRoman bool
		var delimiter string

		fmt.Println("Input:")
		text, _ := reader.ReadString('\n')
		text = spaceMap(text) // strip spaces

		if digitCheck.MatchString(text) {
			isRoman = false
		} else if romanCheck.MatchString(text) {
			isRoman = true
		} else {
			fmt.Println("Wrong format!")
			os.Exit(1)
		}

		delimiter = getDelimiter(text)
		splitText := strings.Split(text, delimiter)

		fmt.Println("Output:\n", calculate(isRoman, splitText[0], splitText[1], delimiter))
	}
}

// IDEA: Check 'strings.Map' with rune type to strip spaces.
// RESULT: Looks like it works
func spaceMap(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}

func romanToInt(s string) int {
	values := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}
	a := values[s]
	return a
}

func intToRoman(i int) string {
	var roman = ""
	var numbers = []int{1, 4, 5, 9, 10, 40, 50, 90, 100}
	var romans = []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C"}

	var index = len(romans) - 1

	for i > 0 {
		for numbers[index] <= i {
			roman += romans[index]
			i -= numbers[index]
		}
		index -= 1
	}

	return roman
}

// Get delimiter to split string and use it for switching  ariphmetic operations
func getDelimiter(text string) string {
	re := regexp.MustCompile(`\+|-|\*|/`)
	delimiterIndex := re.FindStringIndex(text)
	delimiter := string(text[delimiterIndex[0]])
	return delimiter
}

func calculate(isRoman bool, a string, b string, delimiter string) string {
	x := 0
	y := 0
	if !isRoman {
		x, _ = strconv.Atoi(a)
		y, _ = strconv.Atoi(b)

		return strconv.Itoa(ariphmeticOperation(x, y, delimiter))
	} else {
		x = romanToInt(a)
		y = romanToInt(b)

		result := ariphmeticOperation(x, y, delimiter)
		if result <= 0 {
			return "There are no zero or negative numbers in the roman system!"
		}
		return intToRoman(result)
	}
}

func ariphmeticOperation(x int, y int, delimiter string) int {
	result := 0
	switch delimiter {
	case "+":
		result = x + y
	case "-":
		result = x - y
	case "*":
		result = x * y
	case "/":
		result = x / y
	}
	return result
}
