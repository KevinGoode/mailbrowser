package utils

//MailboxRequestorAPI is main api to call mailbox apis
type MailboxRequestorAPI interface {
	Init(credentials string)
	GetMessages(folderName string, max string) string
}
