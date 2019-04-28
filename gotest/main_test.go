package main

import (
	"fmt"
	"testing"
)

const (
	output = `
--- FAIL: TestRepoServiceReturnsEndpointsOnSuccess (0.00s)
	main_test.go:44: Expected https://api.github.com/rate_limit but got https://api.github.com/rate_limit
--- FAIL: TestRepoServiceReturnsErrorOnError (0.00s)
	main_test.go:53: Expected error
FAIL
FAIL	github.com/jordilin/project/learn/h	0.003s
--- FAIL: TestRepoServiceReturnsEndpointsOnSuccess (0.00s)
	main_test.go:44: Expected https://api.github.com/rate_limit but got https://api.github.com/rate_limit
FAIL
FAIL	github.com/jordilin/project/learn/t	0.003s
Makefile:4: recipe for target 'test' failed
make: *** [test] Error 1

`
	baseDir = "project"
)

func TestFullPath(t *testing.T) {
	got := parseTestOutput(output, baseDir)
	expected := `
--- FAIL: TestRepoServiceReturnsEndpointsOnSuccess (0.00s)
	learn/h/main_test.go:44: Expected https://api.github.com/rate_limit but got https://api.github.com/rate_limit
--- FAIL: TestRepoServiceReturnsErrorOnError (0.00s)
	learn/h/main_test.go:53: Expected error
FAIL
FAIL	github.com/jordilin/project/learn/h	0.003s
--- FAIL: TestRepoServiceReturnsEndpointsOnSuccess (0.00s)
	learn/t/main_test.go:44: Expected https://api.github.com/rate_limit but got https://api.github.com/rate_limit
FAIL
FAIL	github.com/jordilin/project/learn/t	0.003s
Makefile:4: recipe for target 'test' failed
make: *** [test] Error 1

`
	if got != expected {
		t.Errorf(fmt.Sprintf("Expected {%s} but got {%s}", expected, got))
	}
}

func TestAllOk(t *testing.T) {
	testOutput := `
ok  	jordilin/learn/h	0.003s
ok  	jordilin/learn/t	0.003s

`
	expected := `All good ;-)`
	got := parseTestOutput(testOutput, baseDir)
	if got != expected {
		t.Errorf(fmt.Sprintf("Expected {%s} but got {%s}", expected, got))
	}
}
