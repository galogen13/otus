package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestTop10(t *testing.T) {

	testSeq := map[string][]string{
		"cat and dog, cat and dog, cat and dog, CaT and. cat and, Cat and, CAT and, CAt!!! cAt, cat": []string{"cat", "and", "dog"},
		"Нога, нога, НОГА нога: НоГа!! - нОгА":                                                       []string{"нога"},
	}

	// Сгенерируем еще большую строку
	testWords := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "zero"}
	testSlice := []string{}
	for j := 0; j <= len(testWords); j++ {
		testSlice = append(testSlice, testWords[:j]...)
	}
	testSeq[strings.Join(testSlice, ", ")] = testWords

	for key, value := range testSeq {
		result := Top10(key)
		if !reflect.DeepEqual(value, result) {
			t.Fatalf("EXPECTED SLICE %s RESULT SLISE %s", value, result)
		}
	}

}
