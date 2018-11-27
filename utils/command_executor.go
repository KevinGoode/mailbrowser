package utils

import (
	"fmt"
)

type commandExecutor struct {
	input     ArgumentReaderAPI
	requestor MailboxRequestorAPI
}

func (executor *commandExecutor) Run() {
	args := executor.input.GetArgs()
	//Ignore error when converting int because checked this already
	fmt.Printf("Executing request. Getting last %s messages from folder %s\n", args["num"], args["folder"])
	executor.requestor.Init("./", args["credentials"])
	executor.requestor.GetMessages(args["folder"], args["num"])

}

//NewCommandExecutor is a public constructor for command executor
func NewCommandExecutor(argsReader ArgumentReaderAPI, mailboxRequestor MailboxRequestorAPI) CommandExecutorAPI {
	executor := commandExecutor{}
	executor.input = argsReader
	executor.requestor = mailboxRequestor
	return &executor
}
