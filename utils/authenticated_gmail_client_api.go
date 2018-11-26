package utils

import "net/http"

//AuthenticatedGmailClientAPI is main api to get authenticated client
type AuthenticatedGmailClientAPI interface {
	Init(credentialsFile string)
	GetClient() *http.Client
}
