package mydict

import "errors"

var (
	errNotFound     = errors.New("Not Found")
	errCantUpdate   = errors.New("Cannot update non-exist word")
	errAlreadyExist = errors.New("Already Exist")
)

// The type is a kind of alias!
// And you can add method on type, like struct

// Dictionary type
type Dictionary map[string]string

// Search - find word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]

	if exists {
		return value, nil
	}

	return "", errNotFound
}

// Add - add word to dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)

	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errAlreadyExist
	}

	return nil
}

// Update - update definition of word
func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)

	switch err {
	case nil:
		d[word] = def
	case errNotFound:
		return errCantUpdate
	}

	return nil
}

// Delete - delete word
func (d Dictionary) Delete(word string) {
	// just using built-in function
	delete(d, word)
}
