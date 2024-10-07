package maps

import "errors"

/*
map[type of key] type of value
map[string]string

maps can return 2 things: the value, and a bool (true if key is found in dict)
*/

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word you were looking for")

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}
