package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"
)

func main() {

	parser := flags.NewParser(&cli, flags.None)
	_, err := parser.Parse()
	if err != nil {
		fmt.Println("ERROR: ", err.Error())
		os.Exit(1)
	}

	os.Exit(0)
}

var cli CFCLI

type CFCLI struct {
	AppShow AppCmd `command:"app-show"`
}

type AppCmd struct{}

func (cmd AppCmd) Execute(args []string) error {
	if len(args) != 1 {
		return errors.New("Wrong number of arguments, expecting APP_NAME")
	}

	return nil
}
