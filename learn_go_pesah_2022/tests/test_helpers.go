package tests

import (
	"fmt"
	"hello/helpers"
	"testing"
)

func ErrorResults(t *testing.T, err error) {
	if err != nil {
		calling_func := helpers.MyCaller()
		fmt.Printf("Correct results in %s - Error - %s \n", calling_func, err.Error())
	}
}

func FailedTestResults(t *testing.T, got interface{}, want interface{}) {
	if got != want {
		calling_func := helpers.MyCaller()
		t.Errorf("Error in %s - got %s want %s", calling_func, got, want)
	}
}

func CorrectResults(got interface{}, want interface{}) {
	if got == want {
		calling_func := helpers.MyCaller()
		fmt.Printf("Correct results in %s - got %s want %s \n", calling_func, got, want)
	}
}
