package dictionary

import (
	"fmt"
)

type Dictionary map[string]string

var (
	ErrNotFound      = "the word %s not found"
	ErrWordExists    = "cannot add word %s because it already exists"
	ErrWordNotExists = "cannot update word %s because its not exists"
)

func (d Dictionary) Search(word string) (string, error) {

	if v, ok := d[word]; !ok || v == "" {
		return "", fmt.Errorf(ErrNotFound, word)
	}

	return d[word], nil
}
func (d Dictionary) Add(word string, definition string) error {

	_, err := d.Search(word)

	if err == nil {
		return fmt.Errorf(ErrWordExists, word)
	}
	d[word] = definition
	return nil
}

func (d Dictionary) Update(word string, definition string) error {

	_, err := d.Search(word)

	if err != nil {
		return fmt.Errorf(ErrWordNotExists, word)
	}
	d[word] = definition
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
