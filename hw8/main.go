package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"syscall"
)

func main() {

	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Fprintf(os.Stderr, "Bad args")
		os.Exit(1)
	}

	dir := args[0]
	cmd := args[1]

	env, err := ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
	exitCode := RunCmd(cmd, env)
	os.Exit(exitCode)

}

// ReadDir .
func ReadDir(dir string) (map[string]string, error) {

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	env := make(map[string]string)

	for _, fileInf := range files {

		file, err := os.Open(dir + "/" + fileInf.Name())
		if err != nil {
			return nil, err
		}

		scanner := bufio.NewScanner(file)
		value := ""
		if scanner.Scan() {
			value = scanner.Text()
		}
		file.Close()

		env[fileInf.Name()] = value
	}

	return env, nil

}

// RunCmd .
func RunCmd(cmd string, env map[string]string) int {

	const defaultFailedCode int = 1

	sysEnv := os.Environ()
	for key, value := range env {
		sysEnv = append(sysEnv, string(key+"="+value))
	}

	command := exec.Command(cmd)
	command.Env = sysEnv
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	err := command.Run()
	exitCode := 0
	if err != nil {

		if exitError, ok := err.(*exec.ExitError); ok {
			ws := exitError.Sys().(syscall.WaitStatus)
			exitCode = ws.ExitStatus()
		} else {
			exitCode = defaultFailedCode
		}
	} else {
		ws := command.ProcessState.Sys().(syscall.WaitStatus)
		exitCode = ws.ExitStatus()
	}

	return exitCode
}
