package Dtos

type TokenRequestDto struct {
	UserName 	string 	`json:"username"`
	Password	string 	`json:"password"`
	IsAdmin 	bool 	`json:"isAdmin"`
}