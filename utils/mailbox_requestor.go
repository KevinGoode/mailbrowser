package utils

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"

	"google.golang.org/api/gmail/v1"
)

type mailBoxRequestor struct {
	client AuthenticatedGmailClientAPI
}

//GetMessages gets all message from a given folder ordered from latest
//max is maximum number of message to receive
func (requestor *mailBoxRequestor) GetMessages(folderName string, max string) {
	//Basic code is based on this:
	//https://developers.google.com/gmail/api/quickstart/go
	//Need to get credentials.json (Credentials.json is this application (mailbrowser) credentials)
	//Need to give the application enough scope to delete emails
	// Scopes are explained here: https://developers.google.com/gmail/api/auth/scopes
	//I set mailbrowser scopes by going here:
	//NOTE NEED to logon to this account as this user: normangoode409@gmail.com
	//https://console.developers.google.com/apis/credentials/consent?project=mailbrowser-1542986963021&authuser=0
	//To support delete need this scope https://www.googleapis.com/auth/gmail.modify

	fmt.Printf("Executing request. Getting last %s messages from folder %s\n", max, folderName)
	maxValue, err := strconv.Atoi(max)
	srv, err := gmail.New(requestor.client.GetClient())
	if err == nil {
		messages, err := srv.Users.Messages.List("me").Do()
		if err == nil {
			count := 0
			for i := 0; i < len(messages.Messages); i++ {

				message, err := srv.Users.Messages.Get("me", messages.Messages[i].Id).Format("metadata").Do()
				if err == nil {
					if count < maxValue && requestor.isInFolder(message.LabelIds, folderName) {
						fmt.Println("*****************************************")
						fmt.Printf("MessageID: %s\n", messages.Messages[i].Id)
						for _, header := range message.Payload.Headers {
							switch name := header.Name; name {
							case "Date", "Subject", "From":
								fmt.Printf("%s: %s\n", header.Name, header.Value)
							default:
								//No nothing
							}
						}
						count++
					}
				}
			}
			if count > 0 {
				fmt.Println("*****************************************")
			}
		}
	}
}
func (requestor *mailBoxRequestor) GetMessage(messageID string) {
	//See following code specifically comment at the end that says to use 'base64.URLEncoding.DecodeString'
	//https://stackoverflow.com/questions/35510179/decoding-the-message-body-from-the-gmail-api-using-go
	fmt.Printf("Executing request. Getting message %s\n", messageID)
	fmt.Println("*****************************************")
	srv, err := gmail.New(requestor.client.GetClient())
	if err == nil {
		message, err := srv.Users.Messages.Get("me", messageID).Format("full").Do()
		if err == nil {
			if message.Payload != nil {
				for _, part := range message.Payload.Parts {
					if part.MimeType == "text/plain" {
						data, _ := base64.URLEncoding.DecodeString(part.Body.Data)
						text := string(data)
						fmt.Println(text)
					}

				}
			} else {
				fmt.Println("*ERROR* failed to get message.")
				fmt.Printf("Are you sure the message id '%s' is correct?\n", messageID)
			}
		}
	}
	fmt.Println("*****************************************")
}
func (requestor *mailBoxRequestor) DeleteMessage(messageID string) {
	fmt.Printf("Executing request. Delete message %s\n", messageID)
	fmt.Println("*****************************************")
	srv, err := gmail.New(requestor.client.GetClient())
	if err == nil {
		_, err := srv.Users.Messages.Trash("me", messageID).Do()
		if err == nil {
			fmt.Printf("*Success* . Message with message id '%s' has been deleted\n", messageID)
		} else {
			fmt.Println("*ERROR* failed to delete message.")
			fmt.Printf("Are you sure the message id '%s' is correct?\n", messageID)
			fmt.Println("Are you sure your token has write access to this email account? Try deleting token and re-generate token with 'modify' in URL")
		}
	}
	fmt.Println("*****************************************")
}
func (requestor *mailBoxRequestor) Init(path string, credentialsFile string) {
	requestor.client.Init(path, credentialsFile)
}
func (requestor *mailBoxRequestor) GetHelp() {
	fmt.Println("The mailbrowser allows browsing of a gmail mail account.")
	fmt.Println("Features:")
	fmt.Println(" - Lists last few emails in a particular folder")
	fmt.Println(" - Displays a particular email body")
	fmt.Println(" - Deletes a particular email")
	fmt.Println("")
}
func (requestor *mailBoxRequestor) isInFolder(labels []string, folder string) bool {
	for _, label := range labels {
		if strings.ToLower(label) == strings.ToLower(folder) {
			return true
		}
	}
	return false
}

//NewMailboxRequestor is a public constructor for mailBoxRequestor
func NewMailboxRequestor(authenticatedGmailClientAPI AuthenticatedGmailClientAPI) MailboxRequestorAPI {
	requestor := mailBoxRequestor{}
	requestor.client = authenticatedGmailClientAPI
	return &requestor
}
