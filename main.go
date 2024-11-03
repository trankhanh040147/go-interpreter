package main

import (
	"fmt"
	"os"
	"os/user"

	"go-interpreter/repl"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Welcome, %s.", usr.Username)
	fmt.Println("Type any commands...")
	repl.Start(os.Stdin, os.Stdout)
}
