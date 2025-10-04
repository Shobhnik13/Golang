package main

import (
	"fmt"
)

func main() {
	var username string = "shobhnik"
	fmt.Println(username)
	fmt.Printf("type of username is: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("type of isLoggedIn is: %T \n", isLoggedIn)

	var age int = 25
	fmt.Println(age)
	fmt.Printf("type of age is: %T \n", age)

	var height float64 = 5.902939384938498394
	fmt.Println(height)
	fmt.Printf("type of height is: %T \n", height)

	// implicit type
	var webiste = "abcd.com"
	fmt.Println(webiste)
	fmt.Printf("type of webiste is: %T \n", webiste)

	// no var syntax
	// var numberOfUsers int = 2000
	numberOfUsers := 2000
	fmt.Println(numberOfUsers)
	fmt.Printf("type of numberOfUsers is: %T \n", numberOfUsers)
}
