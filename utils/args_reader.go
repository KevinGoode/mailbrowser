package utils

import (
	"fmt"
	"os"
	"strconv"
)

type cliOptionDefinition struct {
	name         string
	defaultValue string
	description  string
	required     bool
	number       bool
}
type cliOptionValue struct {
	definition cliOptionDefinition
	value      *string
}

//argReader struct simply contains cliArgs as a string/string key value map of cli arguments
type argReader struct {
	argsAPI ArgsAPI
	options []cliOptionValue
}

func (reader *argReader) GetArgs() cliArgs {
	//1.) Check for unrecognised args
	reader.argsAPI.Parse()
	//2.) Got this far so must be no unrecognised args. Now validate
	reader.validate()
	//3.) Return map of parameters
	args := reader.assembleArgs()
	return args
}
func (reader *argReader) assembleArgs() cliArgs {
	args := make(cliArgs)
	for i := 0; i < len(reader.options); i++ {
		args[reader.options[i].definition.name] = *reader.options[i].value
	}
	return args
}
func (reader *argReader) validate() {
	ok := true
	for i := 0; i < len(reader.options); i++ {
		//Check for required parameters
		if *reader.options[i].value == reader.options[i].definition.defaultValue && reader.options[i].definition.required {
			fmt.Println(reader.options[i].definition.description)
			ok = false
		}
		//Check numbers are numbers
		if reader.options[i].definition.number {
			//Ignore parsed value since we are just checking for error
			_, err := strconv.ParseInt(*reader.options[i].value, 10, 64)
			if err != nil {
				fmt.Println("Option: num is not a number. Num is the maximum number of message returned")
				ok = false
			}
		}
	}
	if !ok {
		os.Exit(1)
	}
	return
}

//NewArgumentReader is the private constructor and initialiser for for struct argReader
func NewArgumentReader(args ArgsAPI) ArgumentReaderAPI {
	//This code uses package flags. See info link.
	//https://gobyexample.com/command-line-flags
	reader := argReader{}
	reader.argsAPI = args
	reader.options = make([]cliOptionValue, 3, 3)
	reader.options[0].definition = cliOptionDefinition{"credentials", "credentials.json", "Must provide full path of credentials file by specifying -credentials option. See https://developers.google.com/gmail/api/quickstart/go for details", false, false}
	reader.options[1].definition = cliOptionDefinition{"folder", "inbox", "", false, false}
	reader.options[2].definition = cliOptionDefinition{"num", "10", "", false, true}
	reader.argsAPI.Init(reader.options)
	return &reader
}
