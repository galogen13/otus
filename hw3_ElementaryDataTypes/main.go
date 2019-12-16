package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {

	start := ""
	res := ""
	err := 0

	start = "a4bc2d5e"
	res, err = Process1(start)
	fmt.Println(start, res, err)

	start = "abcd"
	res, err = Process1(start)
	fmt.Println(start, res, err)

	start = "a22"
	res, err = Process1(start)
	fmt.Println(start, res, err)

	start = "45"
	res, err = Process1(start)
	fmt.Println(start, res, err)

	start = "qwe\\45"
	res = Process2(start)
	fmt.Println(start, res)

	start = "qwe\\4\\5"
	res = Process2(start)
	fmt.Println(start, res)

	start = "qwe\\\\5"
	res = Process2(start)
	fmt.Println(start, res)
}

func Process1(str string) (res string, err int) {
	var r rune
	var countBuilder strings.Builder
	var resBuilder strings.Builder

	for _, val := range str {
		switch {
		case r == 0 && isLetter(val):
			r = val
		case r != 0 && isNumeric(val):
			countBuilder.WriteRune(val)
		case r != 0 && isLetter(val):
			resBuilder, countBuilder, r = appendRune(resBuilder, countBuilder, r, val)
		default:
			return "", 1
		}

	}
	resBuilder, _, _ = appendRune(resBuilder, countBuilder, r, '1')
	return resBuilder.String(), 0
}

func Process2(str string) string {
	const escapeSymbol = '\\'
	runes := []rune(str)
	strLen := utf8.RuneCountInString(str)
	var resBuilder strings.Builder
	for i := 0; i < strLen; i++ {
		end := i + 3
		if i > strLen {
			end = strLen
		}
		part := runes[i:end]
		switch {
		case part[0] != escapeSymbol || len(part) == 1:
			resBuilder.WriteRune(part[0])
		case len(part) == 3 && part[0] == escapeSymbol && isNumeric(part[2]):
			count, _ := strconv.Atoi(string(part[2]))
			resBuilder.WriteString(strings.Repeat(string(part[1]), count))
			i = end
		case part[0] == escapeSymbol:
			resBuilder.WriteRune(part[1])
			i = end - 1
		}

	}
	return resBuilder.String()
}

func isLetter(r rune) bool {
	return 'a' <= r && r <= 'z'
}

func isNumeric(r rune) bool {
	return '0' <= r && r <= '9'
}

func appendRune(resBuilder, countBuilder strings.Builder, r, val rune) (strings.Builder, strings.Builder, rune) {
	if countBuilder.Len() == 0 {
		countBuilder.WriteRune('1')
	}
	count, _ := strconv.Atoi(countBuilder.String())
	resBuilder.WriteString(strings.Repeat(string(r), count))
	r = val
	countBuilder.Reset()
	return resBuilder, countBuilder, r
}
