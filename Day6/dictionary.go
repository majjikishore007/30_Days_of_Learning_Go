package day6

import "errors"

type Dictionary map[string]string

var ErrorNotFound = errors.New("could not find the word you were looking for")

func (d Dictionary) Search(key string) (string, error) {
	val, ok := d[key]
	if !ok {
		return "", ErrorNotFound
	}
	return val, nil
}
func (d Dictionary) Add(key, value string) {
	d[key] = value
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}
