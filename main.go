package main

import ( 
	"log"
	"net/mail"
	"encoding/base64"
	"net/smtp"
	"fmt"
	"strings"
)
/* Structure to have all credentials for the SMTP Server. Parameters are:
	typeOfAuth : Type of Authentication to be used , for CRAMMD5Auth pass "CRAMMD5" whereas for PlainAuth use "PLAIN" as parameter value.
	smtpServer:  Pass the SMTP server with port number here, for example : "mail.example.com:345".	  
	username: Pass Username in here
	password: Pass Password in here.
*/	

type smtpServerCredential struct {
	typeOfAuth string
	smtpServer string
	username string
	password string
	crammd5String string
}

// smtpServerCredential provides authentication features currently supported autherisation types are CRAMMD5Auth and PlainAuth

func smtpServerCredential( smtpStruct smtpServerCredential ) string{
	if smtpStruct.typeOfAuth == "CRAMMD5"{
		auth := smtp.CRAMMD5Auth("", smtpStruct.username, smtpStruct.crammd5String)
	}else if smtpStruct.typeOfAuth == "PLAIN"{
		auth := smtp.PlainAuth("", smtpStruct.username, smtpStruct.password, smtpStruct.smtpServer)
	}
	return auth
}

type mailContent struct {
	receipentName string
	receipentEmail string
	sendersName string
	sendersEmail string
	subject string
	body string
}

func sendMail(structContent mailContent, authMail string){
		header := make(map[string]string)
		header["From"] = from.String()
		header["To"] = to.String()
		header["Subject"] = encodeRFC2047(title)
		header["MIME-Version"] = "1.0"
		header["Content-Type"] = "text/plain; charset=\"utf-8\""
		header["Content-Transfer-Encoding"] = "base64"
		 
		message := ""
		for meta, content := range header {
		message += fmt.Sprintf("%s: %s\r\n", meta, content)
		}
		message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(body))
		
}