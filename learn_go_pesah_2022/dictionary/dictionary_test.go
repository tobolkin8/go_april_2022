package dictionary

import (
	"hello/tests"
	"testing"
)

var dictionary = Dictionary{"test": "this is just a test"}

func TestSearchWordExist(t *testing.T) {

	got, _ := dictionary.Search("test")
	want := "this is just a test"
	tests.FailedTestResults(t, got, want)
	tests.CorrectResults(got, want)
}
func TestSearchWordNotExist(t *testing.T) {
	got, err := dictionary.Search("bu")
	want := "this is just a test"

	tests.ErrorResults(t, err)
	tests.CorrectResults(got, want)
}

func TestAddToSearch(t *testing.T) {
	dictionary.Add("nina", "nina is in the house")

	got, err := dictionary.Search("nina")
	want := "nina is in the house"
	tests.ErrorResults(t, err)
	tests.FailedTestResults(t, got, want)
	tests.CorrectResults(got, want)
}
func TestAddToSearchExisting(t *testing.T) {
	err := dictionary.Add("test", "this is just a test")
	tests.ErrorResults(t, err)
}

func TestUpdateSearch(t *testing.T) {
	dictionary.Update("test", "this is just a updated test")
	got, err := dictionary.Search("test")
	want := "this is just a updated test"
	tests.ErrorResults(t, err)
	tests.FailedTestResults(t, got, want)
	tests.CorrectResults(got, want)
}

func TestUpdateFailSearch(t *testing.T) {
	err := dictionary.Update("notExist", "this is just a test")
	tests.ErrorResults(t, err)
}

func TestDeleteSearch(t *testing.T) {
	dictionary.Delete("test")
	_, err := dictionary.Search("test")
	tests.ErrorResults(t, err)
}
