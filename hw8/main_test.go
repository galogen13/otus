package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestReadDir(t *testing.T) {

	/////// Несуществующий каталог

	if env, err := ReadDir("nosuchdir"); env != nil || err == nil {
		t.Fatalf("ReadDir: expected error, but error is nil")
	}

	dir := "test"
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		t.Fatalf("Can not create test directory")
	}
	defer os.RemoveAll("./" + dir)

	/////// Пустая директория
	env, err := ReadDir(dir)
	if env == nil || len(env) != 0 {
		t.Fatalf("ReadDir: not correct env: %v", env)
	}

	/////// 2  файла в директории
	value1 := "123"
	tmpFilename1 := createTempFile(value1, dir, t)

	value2 := "anothervar"
	tmpFilename2 := createTempFile(value2, dir, t)

	env, err = ReadDir(dir)
	if err != nil {
		t.Fatalf("ReadDir: unexpected error %v", err)
	}
	fileName1 := filepath.Base(tmpFilename1)
	resValue1, ok1 := env[fileName1]
	fileName2 := filepath.Base(tmpFilename2)
	resValue2, ok2 := env[fileName2]
	if env == nil || len(env) != 2 || !ok1 || resValue1 != value1 || !ok2 || resValue2 != value2 {
		t.Fatalf("ReadDir: not correct env: %v", env)
	}

}

func TestRunCmd(t *testing.T) {

	dir := "test"
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		t.Fatalf("Can not create test directory")
	}
	defer os.RemoveAll("./" + dir)

	value := "package main \n import \"os\" \n func main() { \n\n user, ok := os.LookupEnv(\"USER\") \n\nif !ok || user != \"TEST\" { \n os.Exit(33) \n return \n} \n os.Exit(0) \n}"
	fileName := createTempFile(value, dir, t)
	os.Chdir(dir)
	err = os.Rename(filepath.Base(fileName), "test.go")
	if err != nil {
		t.Fatalf("Error rename temp file: %v", err.Error())
	}

	cmd := exec.Command("go", "build", "test.go")
	if err := cmd.Run(); err != nil {
		t.Fatalf(err.Error())
	}

	env := make(map[string]string)
	env["USER"] = "TEST"
	exitCode := RunCmd("test", env)
	if exitCode != 0 {
		t.Fatalf("Exit code is not correct: %v. Expected: 0", exitCode)
	}

	env["USER"] = "EST"
	exitCode = RunCmd("test", env)
	if exitCode != 33 {
		t.Fatalf("Exit code is not correct: %v. Expected: 33", exitCode)
	}

	os.Chdir("../")

}

func createTempFile(value, dir string, t *testing.T) string {
	content := []byte(value)
	tmpfile, err := ioutil.TempFile(dir, "env")
	if err != nil {
		t.Fatalf("Can not create temp file")
	}
	if _, err := tmpfile.Write(content); err != nil {
		t.Fatalf("Can not write to tem file")
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatalf("Can not close test file")
	}
	return tmpfile.Name()
}
