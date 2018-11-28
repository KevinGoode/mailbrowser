package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
)

// authenticatedGmailClient stores reference to oauth2 config
// Code in this source file based on this code: https://developers.google.com/gmail/api/quickstart/go
type authenticatedGmailClient struct {
	config    oauth2.Config
	tokenPath string
}

func (client *authenticatedGmailClient) Init(path string, credentialsFile string) {
	b, err := ioutil.ReadFile(credentialsFile)
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, gmail.GmailReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client.config = *config
	client.tokenPath = path
}

// Retrieve a token, saves the token, then returns the generated client.
func (client *authenticatedGmailClient) GetClient() *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := client.tokenPath + "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = client.getTokenFromWeb()
		saveToken(tokFile, tok)
	}
	httpClient := client.config.Client(context.Background(), tok)

	return httpClient
}

// Request a token from the web, then returns the retrieved token.
func (client *authenticatedGmailClient) getTokenFromWeb() *oauth2.Token {
	authURL := client.config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)
	fmt.Println("")
	fmt.Println("NOTE: THIS TOKEN WILL NOT PROVIDE WRITE ACCESS SO DELETE WILL NOT WORK.")
	fmt.Println("NOTE: TO ALLOW DELETE SIMPLY CHANGE 'readonly' IN ABOVE URL TO 'modify'")
	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := client.config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

// NewAuthenticatedGmailClient is constructor for authenticatedGmailClient that returns AuthenticatedGmailClientAPI
func NewAuthenticatedGmailClient() AuthenticatedGmailClientAPI {
	client := authenticatedGmailClient{}
	return &client
}
