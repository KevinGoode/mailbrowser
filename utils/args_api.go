package utils

//ArgsAPI is interface used to abstract away flag api
type ArgsAPI interface {
	Init(options []cliOptionValue)
	Parse()
}
