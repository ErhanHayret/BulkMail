package Models

import(
	"github.com/dgrijalva/jwt-go"
)

type ClaimsModel struct {
	UserName 	string 	`json:"username"`
	Password	string 	`json:"password"`
	IsAdmin 	bool 	`json:"isAdmin"`
	jwt.StandardClaims
}