package Models

import(
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Mail struct {
	Id				primitive.ObjectID	`json:"id,omitempty"	bson:"_id"`
	MailSubject    	string 				`json:"mailSubject" 	bson:"mail_subject"`
	MailText 	   	string 				`json:"mailText" 		bson:"mail_test"`
	SenderEmail    	string 				`json:"senderEmail" 	bson:"sender_email"`
	SenderEmailPsw 	string				`json:"senderEmailPsw"	bson:"sender_email_psw"`
	ArriveEmails   	[]string			`json:"arriveEmails" 	bson:"arrive_emails"`
}