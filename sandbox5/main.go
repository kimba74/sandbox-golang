package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	// Create a new application parser
	app := cli.NewApp()
	// Configure parameters shown in --help
	app.Name = "sandbox5"
	app.Usage = "A simple sandbox CLI"
	app.Description = "This simple sandbox application demonstrates how to " +
		"utilize the urfave/cli Go package from GitHub."
	app.Version = "0.1.0"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Steffen Krause",
			Email: "steffen@somewhere.io",
		},
	}
	app.Copyright = "(c) 2018 Somewhere.io"
	// Configure global options
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "verbose",
			Usage: "Specifies whether the sandbox application should run in verbose mode",
		},
	}
	// Configure commands for sandbox application
	app.Commands = []cli.Command{
		// Configure a 'greeter' command for the sandbox application
		cli.Command{
			Name:   "greet",
			Usage:  "Issues a greeting",
			Action: handlerGreeter,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name,n",
					Usage: "Sets the name to greet, if not specified the greeting goes to the world",
					Value: "world",
				},
			},
		},
		// Configure a 'goodbye' command for the sandbox application
		cli.Command{
			Name:   "goodbye",
			Usage:  "Issues a good-bye",
			Action: handlerGoodBye,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name,n",
					Usage: "Sets the name to say good-bye, if not specified the good-bye goes to the world",
					Value: "world",
				},
			},
		},
		// Configure a 'secret' (hidden) command for the sandbox application
		cli.Command{
			Name:   "secret",
			Usage:  "Issues a secret message",
			Hidden: true,
			Action: handlerSecret,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name,n",
					Usage: "Sets the name to give the secret message, if not specified the secret message goes to the world",
					Value: "world",
				},
				cli.IntFlag{
					Name:   "count,c",
					Usage:  "Sets the amount of iterations for the secret message",
					Value:  1,
					Hidden: true,
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
