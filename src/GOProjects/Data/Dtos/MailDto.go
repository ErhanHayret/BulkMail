package Dtos

type MailDto struct{
	MailSubject string	`json:"mailSubject"`
	MailText 	string	`json:"mailText"`
	SenderEmail string	`json:"senderEmail"`
}