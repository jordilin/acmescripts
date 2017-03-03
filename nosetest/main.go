package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var makeTarget string

func main() {
	flag.StringVar(&makeTarget, "t", "test", "make test target to execute")
	flag.Parse()
	if flag.NArg() != 0 {
		flag.Usage()
		os.Exit(2)
	}
	cmd := exec.Command("make", makeTarget)
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
