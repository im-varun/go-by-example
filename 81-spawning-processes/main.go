package main

import (
	"fmt"
	"os/exec"
)

func main() {
	gitCmd := exec.Command("git", "--version")
	gitVersion, err := gitCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> git --version")
	fmt.Println(string(gitVersion))

	_, err = exec.Command("date").Output()
	if err != nil {
		switch e := err.(type) {
		case *exec.Error:
			fmt.Println("failed executing:", err)
		case *exec.ExitError:
			fmt.Println("command exit rc =", e.ExitCode())
		default:
			panic(err)
		}
	}
}
