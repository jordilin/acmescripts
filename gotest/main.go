package main

// fullpath gotest
// adds src/packagename/ when there is a failure
// executing go test
// Allows to right click and open file in Acme editor.

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
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
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Run()
	fmt.Println(parseTestOutput(strings.TrimSpace(out.String())))
}

const FAIL = "FAIL"

type LineRpl struct {
	line     string
	found    bool
	filename string
	trace    string
	replaced bool
}

func parseTestOutput(output string) string {
	lines := strings.Split(output, "\n")
	lRpl := make(map[int]*LineRpl)
	counter := 0
	var line string
	for counter, line = range lines {
		lRpl[counter] = &LineRpl{line: line, found: false}
		fields := strings.Fields(line)
		if len(fields) > 0 {
			addr := strings.Split(fields[0], ":")
			if len(addr) > 2 {
				filename := strings.Join(addr[:2], ":")
				val, _ := lRpl[counter]
				val.filename = filename
				val.trace = strings.Join(fields[1:], " ")
				val.found = true
			} else if fields[0] == FAIL && len(fields) == 3 {
				pkgName := fields[1]
				for _, val := range lRpl {
					if val.found && !val.replaced {
						val.filename = filepath.Join("src", pkgName, val.filename)
						val.line = "\t" + val.filename + ": " + val.trace
						val.replaced = true
					}
				}
			}
		}
	}
	var newOutput []string
	count := 0
	for count <= counter {
		newOutput = append(newOutput, lRpl[count].line)
		count++
	}
	return strings.Join(newOutput, "\n")
}
