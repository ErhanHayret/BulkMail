package Models

type TokenModel struct {
	AccessToken		string
	RefreshToken	string
	AtExpire		int64
	RtExpire		int64
}