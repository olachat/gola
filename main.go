package main

import (
	"log"
	"os"

	"github.com/mitchellh/cli"
	"github.com/olachat/gola/v2/golalib"
)

func main() {
	c := cli.NewCLI("gola", golalib.VERSION)
	c.Args = os.Args[1:]

	c.Commands = map[string]cli.CommandFactory{
		"gen": golalib.GenCommandFactory,
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
