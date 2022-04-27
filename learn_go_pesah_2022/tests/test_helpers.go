package tests

import (
	"fmt"
	"hello/helpers"
	"hello/str_consts"
	"testing"
)

func ErrorResults(t *testing.T, err error) {
	if err != nil {
		calling_func := helpers.MyCaller()
		fmt.Printf(str_consts.TESTS_EXPECTED_ERROR_RESULT, calling_func, err.Error())
	}
}

func FailedTestResults(t *testing.T, got interface{}, want interface{}) {
	if got != want {
		calling_func := helpers.MyCaller()
		t.Errorf(str_consts.TESTS_ERROR, calling_func, got, want)
	}
}

func CorrectResults(got interface{}, want interface{}) {
	if got == want {
		calling_func := helpers.MyCaller()
		fmt.Printf(str_consts.TEST_CORRECT_RESULT, calling_func, got, want)
	}
}
