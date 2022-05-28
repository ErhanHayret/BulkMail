package SmtpWorkflow

import(
	"net/smtp"

	"bulkmail/packages/Data/Models"
	"bulkmail/packages/Utils/Logger"
)

var smtpHost = "smtp.gmail.com"//smtp.ethereal.email
var smtpPort = "587"

func Send(mail Models.MailModel){
	auth := smtp.PlainAuth("", mail.SenderEmail, mail.SenderEmailPsw, smtpHost)
	err := smtp.SendMail(smtpHost + ":" + smtpPort, auth, mail.SenderEmail, mail.ArriveEmails, []byte(mail.MailText))
	//Content-Type: text/plain; charset=utf-8\nFrom: name surname <info@test.email>\nTo: Name Surname <testmail@mail.com>\nSubject: Test Subject\n\nTest Body
	if err != nil {
		Logger.FailOnError(err, "Mail Can't Send")
	}
}