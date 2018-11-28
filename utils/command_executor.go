package utils

import "fmt"

type commandExecutor struct {
	input     ArgumentReaderAPI
	requestor MailboxRequestorAPI
}

func (executor *commandExecutor) Run() {
	args := executor.input.GetArgs()
	executor.requestor.Init("./", args["credentials"])
	switch args["cmd"] {
	case "list":
		executor.requestor.GetMessages(args["folder"], args["num"])
	case "get":
		executor.requestor.GetMessage(args["id"])
	case "delete":
		executor.requestor.DeleteMessage(args["id"])
	case "help":
		executor.requestor.GetHelp()
		executor.input.GetHelp()
	default:
		fmt.Printf("Internal error. Unexpected command type")
	}
}

//NewCommandExecutor is a public constructor for command executor
func NewCommandExecutor(argsReader ArgumentReaderAPI, mailboxRequestor MailboxRequestorAPI) CommandExecutorAPI {
	executor := commandExecutor{}
	executor.input = argsReader
	executor.requestor = mailboxRequestor
	return &executor
}
