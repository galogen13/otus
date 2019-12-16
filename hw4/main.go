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

	const top = 10

	wordMap := map[string]int{}
	result := []string{}

	// Разбивает фразу на слайс фраз по разделителю " ", добавляет полученные фразы в словарь, увеличивая счетчик
	words := strings.Split(str, " ")
	for _, word := range words {
		processedWord := processWord(word)
		if processedWord != "" {
			wordMap[processedWord]++
		}
	}

	type wordCount struct {
		word  string
		count int
	}
	wordSlice := []wordCount{}

	for key, value := range wordMap {
		wordSlice = append(wordSlice, wordCount{key, value})
	}

	sort.Slice(wordSlice, func(i, j int) bool {
		return wordSlice[i].count > wordSlice[j].count
	})

	i := 0
	for i < top && i < len(wordSlice) {
		result = append(result, wordSlice[i].word)
		i++
	}

	return result
}

// Приводит фразу к нижнему регистру и убирает все знаки препинания в конце
func processWord(word string) string {
	word = strings.ToLower(word)
	runes := []rune(word)
	for len(runes) != 0 && unicode.IsPunct(runes[len(runes)-1]) {
		runes = runes[:len(runes)-1]
	}
	return string(runes)
}
