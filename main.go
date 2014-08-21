package main

import ("bytes"
		"net/smtp"
		)

func main(){
	type EmailUser struct {
	  Username    string
	  Password    string
	  EmailServer string
	  Port        int
	};

	emailUser := &EmailUser{'aalekh.nigam@gmail.com', '', 'smtp.gmail.com', 587};

	auth := smtp.PlainAuth("",
  		emailUser.Username,
  		emailUser.Password,
  		emailUser.EmailServer
	);
	type SmtpTemplateData struct {
  From    string
  To      string
  Subject string
  Body    string
}

const emailTemplate = `From: {{.From}}
To: {{.To}}
Subject: {{.Subject}}

{{.Body}}

Sincerely,

{{.From}}
`
var err error
var doc bytes.Buffer

context := &SmtpTemplateData{
  "SmtpEmailSender",
  "recipient@domain.com",
  "This is the e-mail subject line!",
  "Hello, this is a test e-mail body."
}
t := template.New("emailTemplate")
t, err = t.Parse(emailTemplate)
if err != nil {
  log.Print("error trying to parse mail template")
}
err = t.Execute(&doc, context)
if err != nil {
  log.Print("error trying to execute mail template")
}

}