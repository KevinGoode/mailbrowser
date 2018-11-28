# mailbrowser
This is a simple go app that lists email messages from a gmail
mail account. It can also display and delete individual email messages
It is based on code here: 
https://developers.google.com/gmail/api/quickstart/go

NOTE1: Code needs a valid credentials.json file in root directory .
Go to link above to find out how to create credentials.json.
The first time the code is executed, a token.json file is created.
After this , the code has permanent access to mailbox

NOTE2: The token.json only provides readonly access to the email account.
If you want to support delete then you will need to give your application
gmail.modify scope. To do this, go to your app page and set credentials.
IE Press 'API Console' link when you create credentials.json
EG link is something like:
 https://console.developers.google.com/apis/credentials/consent?project=PROJECTNAME


