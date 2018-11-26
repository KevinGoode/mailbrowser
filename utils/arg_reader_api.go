package utils

//cliArgs is a key value map that holds cli arguments
//Note that key should not include the '-' character . eg key for '-h' option  is 'h'
type cliArgs map[string]string

//ArgumentReaderAPI is main api to read cli arguments
type ArgumentReaderAPI interface {
	GetArgs() cliArgs
}
