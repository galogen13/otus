package main

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestCopy(t *testing.T) {

	fromFileName := "test.txt"
	toFileName := "result.txt"
	var fileSize int64 = 10000000

	//////////// Обычная копия
	_ = generateFile(fromFileName, fileSize, t)

	err := Copy(fromFileName, toFileName, 0, 0)
	if err != nil {
		t.Fatalf("Simple copy: Unexpected error: %v", err)
	}

	result, err := os.Open(toFileName)
	if err != nil {
		t.Fatalf("Simple copy: failed to open result file info: %v", err)
	}

	fi, err := result.Stat()
	if err != nil {
		t.Fatalf("Simple copy: failed to get result file info: %v", err)
	}
	resultSize := fi.Size()
	if fileSize != resultSize {
		t.Fatalf("Simple copy: from file size = %v; to file size = %v", fileSize, resultSize)
	}

	result.Close()

	removeFile(fromFileName, t)
	removeFile(toFileName, t)

	//////////// Не существует источник
	err = Copy("qqqq.txt", toFileName, 0, 0)
	if err == nil {
		t.Fatalf("Source file not exist: expected error, but error is nil")
	}

	//////////// Смещение и лимит
	_ = generateFile(fromFileName, fileSize, t)
	var limit int64 = 5
	var offset int64 = 90
	err = Copy(fromFileName, toFileName, limit, offset)
	if err != nil {
		t.Fatalf("Offset and limit: Unexpected error: %v", err)
	}

	result, err = os.Open(toFileName)
	if err != nil {
		t.Fatalf("Offset and limit: failed to open result file info: %v", err)
	}

	fi, err = result.Stat()
	if err != nil {
		t.Fatalf("Offset and limit: failed to get result file info: %v", err)
	}
	resultSize = fi.Size()
	if resultSize != limit {
		t.Fatalf("Offset and limit: result file size = %v; limit = %v", resultSize, limit)
	}

	b, err := ioutil.ReadAll(result)
	if err != nil {
		t.Fatalf("Offset and limit:  cannot read file:  %v", err)
	}

	expStr := "12345"
	resStr := string(b)
	if resStr != expStr {
		t.Fatalf("Offset and limit:  EXPECTED STR: %v RESULT STR: %v", expStr, resStr)
	}

	result.Close()

	removeFile(fromFileName, t)
	removeFile(toFileName, t)

}

func generateFile(name string, size int64, t *testing.T) *os.File {

	file, err := os.Create(name)

	if err != nil {
		t.Fatalf("can not create test file")
	}
	defer file.Close()
	var strBuilder strings.Builder
	var i int64 = 1
	for i <= size {
		strRunes := []rune(strconv.FormatInt(int64(i), 10))
		strBuilder.WriteRune(strRunes[len(strRunes)-1])
		i++
	}

	file.WriteString(strBuilder.String())

	return file
}

func removeFile(fileName string, t *testing.T) {
	err := os.Remove(fileName)
	if err != nil {
		t.Fatalf("can not remove test file")
	}
}
