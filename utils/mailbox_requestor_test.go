package utils

import "testing"

func TestMailRequestorListFail(*testing.T) {
	gmailClient := NewAuthenticatedGmailClient()
	requestor := NewMailboxRequestor(gmailClient)
	requestor.Init("../test_resources/", "../test_resources/credentials.json")
	requestor.GetMessages("inbox", "10")
}
func TestMailRequestorGetFail(*testing.T) {
	gmailClient := NewAuthenticatedGmailClient()
	requestor := NewMailboxRequestor(gmailClient)
	requestor.Init("../test_resources/", "../test_resources/credentials.json")
	requestor.GetMessage("")
}
func TestMailRequestorDeleteFail(*testing.T) {
	gmailClient := NewAuthenticatedGmailClient()
	requestor := NewMailboxRequestor(gmailClient)
	requestor.Init("../test_resources/", "../test_resources/credentials.json")
	requestor.DeleteMessage("")
}
