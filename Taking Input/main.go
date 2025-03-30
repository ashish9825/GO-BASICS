package main

import (
	"fmt"
	"os"
	"bufio"

)

func main(){
	fmt.Println("hey, whats your  name : ")
	// var name string

	// fmt.Scan(&name)
	// fmt.Println("your name is ",name)

	reader := bufio.NewReader(os.Stdin)
	name, _ :=reader.ReadString('\n')
	fmt.Println("Hello, Mr.",name)
}