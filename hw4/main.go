package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

func main() {
	str := "cat, and dog - one dog One two one; ONE cats!! and one: man, two. tree, buy cloud buy! record Simple one, after! Six, floor weather mother car! Car Cars dad hour watch, car, dad, weather mom mother believe"
	result := Top10(str)
	fmt.Println(result)

}

// Top10 Возвращает слайс, состоящий из 10 максимум самых частовстречаемых слов во фразе
//Слова упорядочены в порядке убывания частотности #
func Top10(str string) []string {

	top := 10

	wordMap := map[string]int{}

	str = strings.ToLower(str)

	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	words := strings.FieldsFunc(str, f)

	for _, word := range words {
		wordMap[word]++
	}

	type wordCount struct {
		word  string
		count int
	}
	wordSlice := make([]wordCount, 0, len(wordMap))

	for key, value := range wordMap {
		wordSlice = append(wordSlice, wordCount{key, value})
	}

	sort.Slice(wordSlice, func(i, j int) bool {
		return wordSlice[i].count > wordSlice[j].count
	})

	wordSliceLength := len(wordSlice)
	if top > wordSliceLength {
		top = wordSliceLength
	}

	result := make([]string, 0, top)
	for i := 0; i < top; i++ {
		result = append(result, wordSlice[i].word)
	}

	return result
}