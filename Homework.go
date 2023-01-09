package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	arr := make([]string, 0)
	var text string
	for text != "q" {
		fmt.Print("Enter your text: ")
		scanner.Scan()
		text = scanner.Text()
		if text != "q" {
			fmt.Println("Your text was:  ", text)
			arr = append(arr, text)
		}
	}
	fmt.Println(arr)
}
