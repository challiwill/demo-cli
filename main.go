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
	// if -h or --help is passed in, a ErrHelp is returned by the parser. This
	// lets us format our own help messages. We can even create a new parser that
	// calls specific help commands.
	if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
		fmt.Println("CUSTOM HELP MESSAGE")
		// This grabs the proper command field off of the cli struct so that we can
		// parse the struct tag. It does this by looking up the field that matches
		// the active command name. Potentially you would have to look up by
		// command alias as well.
		field, found := reflect.TypeOf(cli).FieldByNameFunc(
			func(fieldName string) bool {
				// We do not need to check four 'found' here because we are looping
				// over field names pulled off the struct.
				field, _ := reflect.TypeOf(cli).FieldByName(fieldName)
				return parser.Active.Name == field.Tag.Get("command")
			},
		)
		if !found {
			fmt.Println("FIELD NOT FOUND")
			os.Exit(1)
		}
		// we'd want to use load asset function to compile the translations
		// into the binary
		i18n.LoadTranslationFile("translations/fr-fr.all.json")
		// get translation function
		t, err := i18n.Tfunc("fr-FR")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		// use translation
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
