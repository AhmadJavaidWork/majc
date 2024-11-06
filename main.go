package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/ahmadjavaidwork/majc/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is maj\n", user.Username)
	fmt.Println("Feel free to type in commands")
	repl.Start(os.Stdin, os.Stdout)
}
