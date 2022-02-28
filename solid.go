package main

import (
	"os"
	"bufio"
	"fmt"
)

func main() {
	fmt.Println("Welcome to Solid CLI")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("=> ")
		text, _ := reader.ReadString('\n')
		fmt.Println("do something with: " + text)
	}
}
