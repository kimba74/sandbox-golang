package main

import (
	"fmt"

	"github.com/urfave/cli"
)

func handlerGoodBye(c *cli.Context) error {
	if c.NArg() > 0 {
		// return cli.NewExitError("command 'goodbye' cannot have arguments", 1)
		return NewError("command 'goodbye' cannot have arguments", 1)
	}
	if c.GlobalBool("verbose") {
		fmt.Println("[VERBOSE]")
	}
	fmt.Printf("Good-bye, %s!\n", c.String("name"))
	return nil
}

func handlerGreeter(c *cli.Context) error {
	if c.NArg() > 0 {
		// return cli.NewExitError("command 'greeter' cannot have arguments", 1)
		return NewError("command 'greeter' cannot have arguments", 1)
	}
	if c.GlobalBool("verbose") {
		fmt.Println("[VERBOSE]")
	}
	fmt.Printf("Hello, %s!\n", c.String("name"))
	return nil
}

func handlerSecret(c *cli.Context) error {
	if c.NArg() > 0 {
		// return cli.NewExitError("command 'greeter' cannot have arguments", 1)
		return NewError("command 'secret' cannot have arguments", 1)
	}
	if c.GlobalBool("verbose") {
		fmt.Println("[VERBOSE]")
	}
	for i := 0; i < c.Int("count"); i++ {
		fmt.Printf("Secret Hi, %s! *wink* *wink*\n", c.String("name"))
	}
	return nil
}
