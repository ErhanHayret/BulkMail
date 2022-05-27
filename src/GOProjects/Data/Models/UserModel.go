package Models

import(
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserModel struct {
	Id 			primitive.ObjectID	`json:"id,omitempty"	bson:"_id"`
	UserName	string				`json:"userName"		bson:"user_name"`
	Password	string 				`json:"password"		bson:"password"`
	IsAdmin 	bool 				`json:"isAdmin"			bson:"is_admin"`
}