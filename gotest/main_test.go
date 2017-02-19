package main

import (
	"fmt"
	"testing"
)

const output = `
--- FAIL: TestRepoServiceReturnsEndpointsOnSuccess (0.00s)
	main_test.go:44: Expected https://api.github.com/rate_limit but got https://api.github.com/rate_limit
--- FAIL: TestRepoServiceReturnsErrorOnError (0.00s)
	main_test.go:53: Expected error
FAIL
FAIL	jordilin/learn/h	0.003s
--- FAIL: TestRepoServiceReturnsEndpointsOnSuccess (0.00s)
	main_test.go:44: Expected https://api.github.com/rate_limit but got https://api.github.com/rate_limit
FAIL
FAIL	jordilin/learn/t	0.003s
Makefile:4: recipe for target 'test' failed
make: *** [test] Error 1

`

func TestFullPath(t *testing.T) {
	got := parseTestOutput(output)
	expected := `
--- FAIL: TestRepoServiceReturnsEndpointsOnSuccess (0.00s)
	src/jordilin/learn/h/main_test.go:44: Expected https://api.github.com/rate_limit but got https://api.github.com/rate_limit
--- FAIL: TestRepoServiceReturnsErrorOnError (0.00s)
	src/jordilin/learn/h/main_test.go:53: Expected error
FAIL
FAIL	jordilin/learn/h	0.003s
--- FAIL: TestRepoServiceReturnsEndpointsOnSuccess (0.00s)
	src/jordilin/learn/t/main_test.go:44: Expected https://api.github.com/rate_limit but got https://api.github.com/rate_limit
FAIL
FAIL	jordilin/learn/t	0.003s
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
	got := parseTestOutput(testOutput)
	if got != testOutput {
		t.Errorf(fmt.Sprintf("Expected {%s} but got {%s}", testOutput, got))
	}
}
