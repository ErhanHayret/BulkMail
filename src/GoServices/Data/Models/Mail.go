package Models

type Mail struct {
	Id				string		`json:"id"`
	MailSubject    	string 		`json:"mailSubject"`
	MailText 	   	string 		`json:"mailText"`
	SenderEmail    	string 		`json:"senderEmail"`
	SenderEmailPsw 	string 		`json:"senderEmailPsw`
	ArriveEmails   	[]string	`json:"arriveEmails"`
}