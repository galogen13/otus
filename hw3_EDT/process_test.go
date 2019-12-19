package main

import "testing"

func TestUnpack(t *testing.T) {

	type result struct {
		str string
		err int
	}

	testSeq := map[string]result{
		"a4bc2d5e":                {"aaaabccddddde", 0},
		"abcd":                    {"abcd", 0},
		"a22":                     {"aaaaaaaaaaaaaaaaaaaaaa", 0},
		"45":                      {"", 1},
		"qwe\\45":                 {"qwe44444", 0},
		"qwe\\4\\5":               {"qwe45", 0},
		"a4bc2d5eqwe\\\\5j10\\l4": {"aaaabccdddddeqwe\\\\\\\\\\jjjjjjjjjjllll", 0}}
	for key, val := range testSeq {
		if str, err := Unpack(key); str != val.str || err != val.err {
			t.Fatalf("EXPECTED STR %s RESULT STR %s; EXPECTED ERR %d RESULT ERR %d", val.str, str, val.err, err)
		}
	}
}
