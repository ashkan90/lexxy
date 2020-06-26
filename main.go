package main

import (
	"fmt"
	"lexxy/repl"
	"os"
	"os/user"
)

func main() {

	_user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is No-named Data Modelling Language!\n",
		_user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)

}
