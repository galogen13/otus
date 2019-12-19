package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {

	start := ""
	res := ""
	err := 0

	start = "a4bc2d5eg\\f3"
	res, err = Unpack(start)
	fmt.Println(start, res, err)

}

// Unpack  #
func Unpack(str string) (string, int) {
	const escapeSymbol = '\\'
	runes := []rune(str)
	strLen := len(runes)
	var resBuilder strings.Builder
	for i := 0; i < strLen; {
		switch {
		case unicode.IsLetter(runes[i]):
			switch {
			case (i+1 < strLen && (runes[i+1] == escapeSymbol || unicode.IsLetter(runes[i+1]))) || i == strLen-1:
				{
					resBuilder.WriteRune(runes[i])
					i++
				}
			case i+1 < strLen && unicode.IsDigit(runes[i+1]):
				{
					c := runes[i]
					var count int
					count, i = getCount(runes, i+1)
					resBuilder.WriteString(strings.Repeat(string(c), count))
				}
			default:
				return "", 1
			}
		case runes[i] == escapeSymbol:
			{
				switch {
				case i+2 < strLen && unicode.IsDigit(runes[i+2]):
					{
						c := runes[i+1]
						var count int
						count, i = getCount(runes, i+2)
						resBuilder.WriteString(strings.Repeat(string(c), count))
					}
				case i+1 < strLen:
					{
						resBuilder.WriteRune(runes[i+1])
						i += 2
					}
				default:
					return "", 1
				}

			}
		default:
			return "", 1
		}
	}
	return resBuilder.String(), 0
}

func getCount(runes []rune, i int) (int, int) {
	var countBuilder strings.Builder
	strLen := len(runes)
	for j := i; j < strLen && unicode.IsDigit(runes[j]); j++ {
		countBuilder.WriteRune(runes[j])
		i++
	}
	count, _ := strconv.Atoi(countBuilder.String())
	return count, i
}
