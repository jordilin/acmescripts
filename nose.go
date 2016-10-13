package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	cmd := exec.Command("make", "test")
	var err bytes.Buffer
	cmd.Stderr = &err
	cmd.Run()
	res := strings.Split(err.String(), "\n")
	for _, r := range res {
		fields := strings.Split(r, "File")
		if len(fields) == 2 {
			subfields := strings.Split(fields[1], ",")
			if len(subfields) == 3 {
				file := strings.Trim(subfields[0], " \"")
				lineNumF := strings.Split(subfields[1], "line")
				fmt.Printf("%s:%s: %s\n", file, strings.TrimLeft(lineNumF[1], " "), subfields[2])
				continue
			}
		}
		fmt.Println(r)
	}
}
