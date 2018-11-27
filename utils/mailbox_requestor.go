package utils

import (
	"fmt"
	"strconv"

	"google.golang.org/api/gmail/v1"
)

type mailBoxRequestor struct {
	client AuthenticatedGmailClientAPI
}

//GetMessages gets all message from a given folder ordered from latest
//max is maximum number of message to receive
func (requestor *mailBoxRequestor) GetMessages(folderName string, max string) string {
	//https://developers.google.com/gmail/api/quickstart/go
	maxValue, err := strconv.Atoi(max)
	srv, err := gmail.New(requestor.client.GetClient())
	if err == nil {
		messages, err := srv.Users.Messages.List("me").Do()
		if err == nil {
			for i := 0; i < len(messages.Messages); i++ {
				if i < maxValue {
					fmt.Println("Message ID: " + messages.Messages[i].Id)
					message, err := srv.Users.Messages.Get("me", messages.Messages[i].Id).Do()
					if err == nil {
						fmt.Println("Message   : " + message.Snippet)
					}
				}
			}
		}
	}
	return ""
}
func (requestor *mailBoxRequestor) Init(path string, credentialsFile string) {
	requestor.client.Init(path, credentialsFile)
}

//NewMailboxRequestor is a public constructor for mailBoxRequestor
func NewMailboxRequestor(authenticatedGmailClientAPI AuthenticatedGmailClientAPI) MailboxRequestorAPI {
	requestor := mailBoxRequestor{}
	requestor.client = authenticatedGmailClientAPI
	return &requestor
}
