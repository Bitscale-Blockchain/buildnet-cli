package utils

import (
	"fmt"
	"os/exec"
)

func ShellExecute(command string, executeInDir string) {
	fmt.Println("Runnig command: ", command)
	// create a new *Cmd instance
	// here we pass the command as the first argument and the arguments to pass to the command as the
	// remaining arguments in the function
	cmd := exec.Command("/bin/sh", "-c", command)

	cmd.Dir = executeInDir

	// The `Output` method executes the command and
	// collects the output, returning its value
	out, err := cmd.Output()
	if err != nil {
		// if there was any error, print it here
		fmt.Println("could not run command: ", err)
	}
	// otherwise, print the output from running the command
	fmt.Println("Output: ", string(out))
}
