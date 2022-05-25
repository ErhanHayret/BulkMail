package MailSender

import(
	//Local Packages
	"net/smtp"

	//This Project Packages
	"bulkmail/packages/Data/Models"
	"bulkmail/packages/Utils/Logger"
)

var smtpHost = "smtp.gmail.com"
var smtpPort = "587"

func Send(mail Models.Mail){
	auth := smtp.PalinAuth("", mail.SenderEmail, mail.SenderEmailPsw, smtpHost)
	count := len(mail.ArriveEmails)
	for var i = 0 ; count > 0 ; i++ {
		err := smtp.SendMail(smtpHost + ":" + smtpPort, auth, mail.SenderEmail, mail.ArriveEmails[i], mail.MailText)
		if err != nil {
			Logger.FailOnError(err, "Mail Can't Send")
			return
		}
		Logger.Print("Mail Sended")
	}
}