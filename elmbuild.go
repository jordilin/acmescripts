package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Provide a file to build\n")
		os.Exit(2)
	}
	fileName := os.Args[1]
	cmd := exec.Command("elm", "make", fileName)
	var errb bytes.Buffer
	cmd.Stdout = &errb
	cmd.Stderr = &errb
	err := cmd.Run()
	if err != nil {
		res := strings.Split(errb.String(), "\n")
		lineNumber := "0"
		for _, r := range res {

			numberFields := strings.Split(r, "|")
			if len(numberFields) == 2 {
				lineNumber = numberFields[0]
				fmt.Printf("\n%s:%s\n\n", fileName, lineNumber)
			}
			fmt.Println(r)
		}
	} else {
		fmt.Println("All good ;-)")
	}
}
