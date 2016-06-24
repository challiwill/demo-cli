package main

import (
	"errors"
	"fmt"
	"os"
	"reflect"

	"github.com/jessevdk/go-flags"
	"github.com/nicksnyder/go-i18n/i18n"
)

func main() {

	parser := flags.NewParser(&cli, flags.HelpFlag)
	_, err := parser.Parse()
	if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
		fmt.Println("CUSTOM HELP MESSAGE")
		field, found := reflect.TypeOf(cli).FieldByName("AppShow")
		if !found {
			fmt.Println("FIELD NOT FOUND")
			os.Exit(1)
		}
		i18n.LoadTranslationFile("translations/fr-fr.all.json")
		t, err := i18n.Tfunc("fr-FR")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Println(t(field.Tag.Get("description")))
	} else if err != nil {
		fmt.Println("ERROR: ", err.Error())
		os.Exit(1)
	}

	os.Exit(0)
}

var cli CFCLI

type CFCLI struct {
	AppShow AppCmd `command:"app-show" description:"show the app details"`
	Help    HelpCmd
}

type AppCmd struct{}

type HelpCmd struct{}

func (cmd AppCmd) Execute(args []string) error {
	if len(args) != 1 {
		return errors.New("Wrong number of arguments, expecting APP_NAME")
	}

	return nil
}
