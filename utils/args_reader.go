package utils

import (
	"fmt"
	"os"
	"strconv"
)

type optionValue struct {
	name        string
	description string
}
type cliOptionDefinition struct {
	name         string
	defaultValue string
	description  string
	required     bool
	number       bool
}
type cliOptionValue struct {
	definition  cliOptionDefinition
	validValues []optionValue
	value       *string
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
func (reader *argReader) GetHelp() {
	fmt.Println("Options:")
	for i := 0; i < len(reader.options); i++ {
		fmt.Println("-" + reader.options[i].definition.name + "(" + reader.options[i].definition.defaultValue + ")")
		fmt.Println("\t" + reader.options[i].definition.description)
		if len(reader.options[i].validValues) > 0 {
			fmt.Println("\tSub Options:")
			for j := 0; j < len(reader.options[i].validValues); j++ {
				fmt.Println("\t" + "-" + reader.options[i].definition.name + "=" + reader.options[i].validValues[j].name + " : " + reader.options[i].validValues[j].description)
			}
		}
		fmt.Println("")
	}
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
				fmt.Printf("Option: '%s' is not a number.\n", reader.options[i].definition.name)
				ok = false
			}
		}
		//Check valid values
		if len(reader.options[i].validValues) > 0 {
			valid := false
			for j := 0; j < len(reader.options[i].validValues); j++ {
				if reader.options[i].validValues[j].name == *reader.options[i].value {
					valid = true
					break
				}
			}
			if !valid {
				fmt.Printf("Option '%s' value '%s' is not valid. Valid values are '%s'\n. ", reader.options[i].definition.name, *reader.options[i].value, reader.options[i].validValues)
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
	reader.options = make([]cliOptionValue, 5, 5)
	reader.options[0].definition = cliOptionDefinition{"credentials", "credentials.json", "Full path of credentials file . See https://developers.google.com/gmail/api/quickstart/go for details", false, false}
	reader.options[1].definition = cliOptionDefinition{"folder", "inbox", "Folder to search. Only used when cmd=list. Otherwise ignored", false, false}
	reader.options[2].definition = cliOptionDefinition{"num", "10", "Number of emails to return. Only used when cmd=list. Otherwise ignored", false, true}
	reader.options[3].definition = cliOptionDefinition{"cmd", "list", "Command to execute.", false, false}
	reader.options[3].validValues = []optionValue{{"list", "Displays email list for given folder."}, {"get", "Displays an email message."}, {"delete", "Deletes an email."}, {"help", "Displays help."}}
	reader.options[4].definition = cliOptionDefinition{"id", "", "ID of message to display or delete. Only used when cmd=get/delete. Otherwise ignored", false, false}
	reader.argsAPI.Init(reader.options)
	return &reader
}
