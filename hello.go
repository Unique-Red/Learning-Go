package main

// import other packages
import (
	"fmt"
)

func main() {
	//printing hello world
	//fmt.Println("Hello World!")

	//variables and constants
	//var firstName string = "Redx"
	// := shorthand version

	//lastName := "Noah"
	//age := 50

	//create but don't assign
	var fullName string
	var age int
	var price float32
	var tf bool
	fmt.Println(fullName, age, price, tf)

	fullName = "Redx Noah"
	age = 50
	price = 19.99
	tf = true
	fmt.Println(fullName, age, price, tf)

	//multiple variable declaration
	var name1, name2 string = "Reddd", "Redd"
	fmt.Println(name1, name2)

	name1 = "Reddington"
	fmt.Println(name1, name2)

	//Constants (variable that doesn't change)
	const PIZZA string = "Chicken BBQ"
	fmt.Println(PIZZA)

	//multiple constants
	const (
		PIZZA1 = "Cheese"
		PIZZA2 = "Pepperoni"
		MYNUM  = 77
	)

	fmt.Println(PIZZA2)

	//fmt.Println(firstName)
}
