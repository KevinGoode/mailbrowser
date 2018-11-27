package utils

import "testing"

func TestMailRequestorFail(*testing.T) {
	gmailClient := NewAuthenticatedGmailClient()
	requestor := NewMailboxRequestor(gmailClient)
	requestor.Init("../test_resources/", "../test_resources/credentials.json")
	requestor.GetMessages("inbox", "10")
}
