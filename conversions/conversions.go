package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("hello from conversions")

	// var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter your age:")
	age, _ := reader.ReadString('\n')

	fmt.Println("Your age is: ", age)
	fmt.Printf("type of age is: %T \n", age)

	// type conversions
	numAge, err := strconv.ParseFloat(strings.TrimSpace(age), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your age:", numAge+1)
	}

}
