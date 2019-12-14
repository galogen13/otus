package main

import "testing"

func TestProcess1(t *testing.T) {

	type result struct {
		str string
		err int
	}

	testSeq := map[string]result{
		"a4bc2d5e": {"aaaabccddddde", 0},
		"abcd":     {"abcd", 0},
		"a22":      {"aaaaaaaaaaaaaaaaaaaaaa", 0},
		"45":       {"", 1}}
	for key, val := range testSeq {
		if str, err := Process1(key); str != val.str || err != val.err {
			t.Fatalf("EXPECTED STR %s RESULT STR %s; EXPECTED ERR %d RESULT ERR %d", val.str, str, val.err, err)
		}
	}
}

func TestProcess2(t *testing.T) {
	testSeq := map[string]string{
		"qwe\\45":   "qwe44444",
		"qwe\\4\\5": "qwe45",
		"qwe\\\\5":  "qwe\\\\\\\\\\",
	}
	for key, val := range testSeq {
		if str := Process2(key); str != val {
			t.Fatalf("EXPECTED STR %s RESULT STR %s", str, val)
		}
	}
}
