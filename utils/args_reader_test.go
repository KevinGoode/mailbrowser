package utils

import "testing"

type argsStub struct {
}

//Init is the initialiser for the dummy struct Args
//For tests all arguments set to default value
func (args *argsStub) Init(options []cliOptionValue) {
	for i := 0; i < len(options); i++ {
		options[i].value = &options[i].definition.defaultValue
	}
}

//Parse is the main API that parses args.
func (args *argsStub) Parse() {
}

//NewArgsStub is  the constructor for the dummy struct argsStub.
func NewArgsStub() ArgsAPI {
	args := args{}
	return &args
}

func TestArgsReaderSuccess(t *testing.T) {
	apiStub := NewArgsStub()
	reader := NewArgumentReader(apiStub)
	args := reader.GetArgs()
	testArgValue := "credentials.json"
	if args["credentials"] != testArgValue {
		t.Errorf("Expected credentials argument to be %s. Was %s", testArgValue, args["credentials"])
	}
	testArgValue = "10"
	if args["num"] != testArgValue {
		t.Errorf("Expected credentials argument to be %s. Was %s", testArgValue, args["num"])
	}
	testArgValue = "inbox"
	if args["folder"] != testArgValue {
		t.Errorf("Expected credentials argument to be %s. Was %s", testArgValue, args["folder"])
	}
}
