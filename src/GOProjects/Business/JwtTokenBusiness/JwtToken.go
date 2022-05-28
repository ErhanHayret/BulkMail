package JwtTokenBusiness

import(
	"time"

	"github.com/dgrijalva/jwt-go"

	"bulkmail/packages/Data/Models"
	"bulkmail/packages/Data/Dtos"
	mongo "bulkmail/packages/DataAccess/MongoDb"
)

var claim Models.ClaimsModel
var result Models.ResultModel
var collection, dbResponse = mongo.GetClient("UserDb", "User")

func CreateNewToken(userDto Dtos.UserDto) (Models.TokenModel, Models.ResultModel) {
	var tokenModel Models.TokenModel
	var tokenString string
	var accessTokenString string
	var err error
	user, respone := mongo.FindUser(collection, userDto)

	if respone.StatusCode == 200 {
		jwtKey := []byte("Bu_Asiri_Gizli_Bir_Sifredir")//Token secret

		//Token
		tokenExpTime := time.Now().Add(time.Minute * 30)//It's define expire time to 30 minutes
		tokenAttribute := &Models.ClaimsModel{
			StandardClaims: jwt.StandardClaims{
			 	ExpiresAt: tokenExpTime.Unix(),
			},
			UserName: user.UserName,
			Password: user.Password,
			IsAdmin: user.IsAdmin,
		}
		tokenClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenAttribute)
		tokenString, err = tokenClaim.SignedString(jwtKey)
 		if err != nil {
			result.Message = "Token can't created"
			result.Status = false
			result.StatusCode = 500
	  		return tokenModel, result
		}

		//AccessToken
		tokenExpTime = time.Now().Add(time.Hour * 24 * 7)//It's define expire time to 7 days
		accessAttribute := &Models.ClaimsModel{
			StandardClaims: jwt.StandardClaims{
			 	ExpiresAt: tokenExpTime.Unix(),
			},
			UserName: user.UserName,
			Password: user.Password,
			IsAdmin: user.IsAdmin,
		}
		accessClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, accessAttribute)
		accessTokenString, err = accessClaim.SignedString(jwtKey)
 		if err != nil {
			result.Message = "Access token can't created"
			result.Status = false
			result.StatusCode = 500
	  		return tokenModel, result
		}
	} else {
		result.Message = "User not found"
		result.Status = false
		result.StatusCode = 404
		return tokenModel, result
	}
	
	tokenModel.Token = tokenString
	tokenModel.AccessToken = accessTokenString
	result.Message = "Success"
	result.Status = true
	result.StatusCode = 200
	return tokenModel, result
}