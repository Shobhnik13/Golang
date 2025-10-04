package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name: ")

	// error handling instead try catch
	name, _ := reader.ReadString('\n')
	fmt.Println("Your name is: ", name)
	fmt.Printf("type of name is: %T \n", name)
}