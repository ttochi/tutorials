package main

import (
	"fmt"
	"strings"

	"github.com/ttochi/tutorials/go/test"
)

func multifly(a int, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeat(words ...string) {
	fmt.Println(words)
}

func lenAndUpper2(name string) (length int, uppercase string) {
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func lenAndUpper3(name string) (length int, uppercase string) {
	defer fmt.Println("Finished funciton!")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func superAdd(numbers ...int) int {
	for idx, number := range numbers {
		fmt.Print(idx, number)
	}

	for i := 0; i < len(numbers); i++ {
		fmt.Print(numbers[i])
	}

	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

func canIDrinkIfElse(age int) bool {
	if age < 18 {
		return false
	}

	// varaible expression: you can define local variable in if clause
	if koreanAge := age + 2; koreanAge < 19 {
		return false
	}

	return true
}

func canIDrinkSwitch(age int) bool {
	switch age {
	case 10:
		return false
	case 18:
		return true
	}

	switch {
	case age < 10:
		return false
	case age == 18:
		return true
	case age > 90:
		return false
	}

	// varaible expression: you can define local variable in switch clause
	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	}

	return false
}

func pointer() {
	// copying value
	a := 2
	b := a
	a = 10
	fmt.Println(a, b) // 10, 2

	// copying reference
	c := 2
	d := &c
	c = 10
	fmt.Println(&c, &d) // each variable's memory address
	fmt.Println(&c, d)  // same memory address
	fmt.Println(c, *d)  // 10, 10

	*d = 20
	fmt.Println(c, *d) // 20, 20
}

func arrayAndSlice() {
	// array: static array
	array := [5]string{"zero", "one"}
	array[2] = "two"
	fmt.Println(array)

	// slice: dynamic array
	slice := []string{"zero", "one"}
	slice = append(slice, "two")
	fmt.Println(slice)
}

func myMap() {
	// key-value pairs
	ttochi := map[string]string{"name": "yujin", "age": "30"}
	fmt.Println(ttochi)

	for key, value := range ttochi {
		fmt.Println(key, value)
	}
}

type person struct {
	name    string
	age     int
	favFood []string
}

func myStruct() {
	favFood := []string{"food1", "food2"}
	ttochi := person{"yujin", 30, favFood}
	ttochi = person{name: "yujin", age: 30, favFood: favFood}
	fmt.Println(ttochi)
	fmt.Println(ttochi.name)
	// there are no class and object concept in go
	// you need to use struct
	// there are no construct method. So, we need to run constructor ourselves
}

// 1. use main function and main package for compile
func main() {
	// 2. package export function is using UpperCase at function name
	test.SayHi()

	// 3. variable can be shorten (only in function!)
	var myVar1 string = "ttochi"
	myVar2 := "ttochi"
	fmt.Println(myVar1, myVar2)

	// 4. function declar in go: clarify data type of argument and return
	fmt.Println(multifly(9, 3))

	// 5. function can return multi-variables and you can ignore unuse variable with underscore
	_, name := lenAndUpper("ttochi")
	fmt.Println(name)

	// 6. function can receiver multi-parameter using ...datatype
	repeat("a", "b", "c", "d")

	// 7. naked return: don't have to return variable at the last line of function
	len7, name7 := lenAndUpper2("ttochi")
	fmt.Println(len7, name7)

	// 8. defer: kind of callback for function
	_, name8 := lenAndUpper3("ttochi")
	fmt.Println(name8)

	// 9. for, range
	total := superAdd(1, 2, 3, 4, 5)
	fmt.Println(total)

	// 10. if else
	drinkable1 := canIDrinkIfElse(18)
	fmt.Println(drinkable1)

	// 11. switch
	drinkable2 := canIDrinkSwitch(18)
	fmt.Println(drinkable2)

	// 12. pointer
	pointer()

	// 13. array and slice
	arrayAndSlice()

	// 14. map
	myMap()

	// 15. struct
	myStruct()
}
