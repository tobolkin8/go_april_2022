package dictionary

import (
	"fmt"
	"hello/str_consts"
)

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {

	if v, ok := d[word]; !ok || v == "" {
		return "", fmt.Errorf(str_consts.DICTIONARY_ERROR_NOT_FOUND, word)
	}

	return d[word], nil
}
func (d Dictionary) Add(word string, definition string) error {

	_, err := d.Search(word)

	if err == nil {
		return fmt.Errorf(str_consts.DICTIONARY_ERROR_WORD_EXIST, word)
	}
	d[word] = definition
	return nil
}

func (d Dictionary) Update(word string, definition string) error {

	_, err := d.Search(word)

	if err != nil {
		return fmt.Errorf(str_consts.DICTIONARY_ERROR_WORD_NOT_EXIST, word)
	}
	d[word] = definition
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
