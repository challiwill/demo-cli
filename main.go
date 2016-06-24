package main

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"
	"unicode"

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
		// parse the struct tag. We could handle this in at least two ways:
		//
		// - We could use FieldByName() and have a lookup from command name to
		// field name that we will pass in.
		// - We could convert from command name to field name (as I have done) and
		// keep consistent field names with command names.
		field, found := reflect.TypeOf(cli).FieldByNameFunc(
			func(fieldName string) bool {
				cmdName := parser.Active.Name
				splitCmd := strings.Split(cmdName, "-")
				camelSplitCmd := make([]string, len(splitCmd))
				for i, s := range splitCmd {
					a := []rune(s)
					a[0] = unicode.ToUpper(a[0])
					camelSplitCmd[i] = string(a)
				}
				camelCmd := strings.Join(camelSplitCmd, "")

				return camelCmd == fieldName
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
