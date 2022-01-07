package main

import (
	"fmt"

	"github.com/ttochi/tutorials/go/mydict"
)

func typeMain() {
	dictionary := mydict.Dictionary{"first": "First word"}

	word := "word"
	definition := "The definition of new word"

	// we can use method on type
	res0, err := dictionary.Search(word)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res0)
	}

	// Add property on dictionary map using method, instead just using map
	// dictionary[word] = definition
	err = dictionary.Add(word, definition)
	if err != nil {
		fmt.Println(err)
	}
	res1, _ := dictionary.Search(word)
	fmt.Println(res1)

	err = dictionary.Update(word, "Tada! Updated definition!")
	if err != nil {
		fmt.Println(err)
	}
	res2, _ := dictionary.Search(word)
	fmt.Println(res2)

	dictionary.Delete(word)
	_, err = dictionary.Search(word)
	if err != nil {
		fmt.Println("Well deleted!")
	}
}
