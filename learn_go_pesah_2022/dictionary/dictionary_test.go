package dictionary

import (
	"hello/str_consts"
	"hello/tests"
	"testing"
)

var dictionary = Dictionary{str_consts.DICTIONARY_TEST_SEARCH_WORD: str_consts.DICTIONARY_TEST_SEARCH_TEXT}

func TestSearchWordExist(t *testing.T) {

	got, _ := dictionary.Search(str_consts.DICTIONARY_TEST_SEARCH_WORD)
	want := str_consts.DICTIONARY_TEST_SEARCH_TEXT
	tests.FailedTestResults(t, got, want)
	tests.CorrectResults(got, want)
}
func TestSearchWordNotExist(t *testing.T) {
	got, err := dictionary.Search(str_consts.DICTIONARY_TEST_NOT_EXIST_SEARCH_WORD)
	want := str_consts.DICTIONARY_TEST_SEARCH_TEXT

	tests.ErrorResults(t, err)
	tests.CorrectResults(got, want)
}

func TestAddToSearch(t *testing.T) {
	dictionary.Add(str_consts.DICTIONARY_TEST_ADD_SEARCH_WORD, str_consts.DICTIONARY_TEST_ADD_SEARCH_TEXT)

	got, err := dictionary.Search(str_consts.DICTIONARY_TEST_ADD_SEARCH_WORD)
	want := str_consts.DICTIONARY_TEST_ADD_SEARCH_TEXT
	tests.ErrorResults(t, err)
	tests.FailedTestResults(t, got, want)
	tests.CorrectResults(got, want)
}
func TestAddToSearchExisting(t *testing.T) {
	err := dictionary.Add(str_consts.DICTIONARY_TEST_SEARCH_WORD, str_consts.DICTIONARY_TEST_SEARCH_TEXT)
	tests.ErrorResults(t, err)
}

func TestUpdateSearch(t *testing.T) {
	dictionary.Update(str_consts.DICTIONARY_TEST_SEARCH_WORD, str_consts.DICTIONARY_TEST_UPDATED_SEARCH_TEXT)
	got, err := dictionary.Search(str_consts.DICTIONARY_TEST_SEARCH_WORD)
	want := str_consts.DICTIONARY_TEST_UPDATED_SEARCH_TEXT
	tests.ErrorResults(t, err)
	tests.FailedTestResults(t, got, want)
	tests.CorrectResults(got, want)
}

func TestUpdateFailSearch(t *testing.T) {
	err := dictionary.Update(str_consts.DICTIONARY_TEST_NOT_EXIST_SEARCH_WORD, str_consts.DICTIONARY_TEST_SEARCH_TEXT)
	tests.ErrorResults(t, err)
}

func TestDeleteSearch(t *testing.T) {
	dictionary.Delete(str_consts.DICTIONARY_TEST_SEARCH_WORD)
	_, err := dictionary.Search(str_consts.DICTIONARY_TEST_SEARCH_WORD)
	tests.ErrorResults(t, err)
}
