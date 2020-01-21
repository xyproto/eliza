package main

import (
	"bufio"
	"fmt"
	"github.com/xyproto/eliza"
	"os"
)

func getInput() string {
	fmt.Print("You: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return input
}

func main() {
	fmt.Println("Eliza: " + eliza.ElizaHi())
	for {
		statement := getInput()
		fmt.Println("Eliza: " + eliza.ReplyTo(statement))

		if eliza.IsQuitStatement(statement) {
			break
		}
	}
}


