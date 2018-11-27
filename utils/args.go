package utils

import (
	"flag"
)

type args struct {
}

//Init is the initialiser for the dummy struct Args
func (args *args) Init(options []cliOptionValue) {
	for i := 0; i < len(options); i++ {
		options[i].value = flag.String(options[i].definition.name, options[i].definition.defaultValue, options[i].definition.description)
	}
}

//Parse is the main API that parses args. This is thin wrapper around flag component
func (args *args) Parse() {
	flag.Parse()
}

//NewArgs is  the constructor for the dummy struct Args. The whole of this source file simlpy provides
// a thin wrapper around flag component
func NewArgs() ArgsAPI {
	args := args{}
	return &args
}
