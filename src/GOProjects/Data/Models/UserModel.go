package Models

import(
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

type UserModel struct {
	Id 			string	`json:"id,omitempty"	bson:"_id"`
	UserName	string	`json:"userName"		bson:"username"`
	Password	string 	`json:"password"		bson:"password"`
	IsAdmin 	bool 	`json:"isAdmin"			bson:"isadmin"`
}