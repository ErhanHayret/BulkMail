package Dtos

type TokenDto struct {
	AccessToken 		string	`json:"accessToken"`
	RefreshToken	string	`json:"refreshToken"`
}