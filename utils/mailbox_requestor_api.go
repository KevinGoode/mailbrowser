package utils

//MailboxRequestorAPI is main api to call mailbox apis
type MailboxRequestorAPI interface {
	Init(path string, credentials string)
	GetMessages(folderName string, max string)
	GetMessage(messageID string)
	DeleteMessage(messageID string)
	GetHelp()
}
