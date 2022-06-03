package JwtTokenBusiness

import(
	"time"
	"fmt"
	"net/http"
	"encoding/json"
	"strings"
	"io/ioutil"

	"github.com/dgrijalva/jwt-go"

	"bulkmail/packages/Data/Models"
	"bulkmail/packages/Data/Dtos"
	mongo "bulkmail/packages/DataAccess/MongoDb"
)

var collection, dbResponse = mongo.GetClient("UserDb", "User")

func CreateToken(w http.ResponseWriter, r *http.Request){
	var userDto Dtos.UserDto
	var err error
	td := &Models.TokenModel{}

	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &userDto)

	user, response := mongo.FindUser(collection, userDto)
	if response.StatusCode == 200 {
		jwtKey := []byte("Bu_Asiri_Gizli_Bir_Sifredir")//Token secret need env
		//Expire times
		td.AtExpire = time.Now().Add(time.Minute * 30).Unix()
		td.RtExpire = time.Now().Add(time.Hour * 24 * 7).Unix()

		//Access Token
		atClaims := jwt.MapClaims{}
		atClaims["userName"] = user.UserName
		atClaims["isAdmin"] = user.IsAdmin

		aToken := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
		td.AccessToken, err = aToken.SignedString(jwtKey)
 		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode("Access token not signed")
			return
		}

		//Refresh Token
		rtClaims := jwt.MapClaims{}
		rtClaims["userName"] = user.UserName
		rtClaims["password"] = user.Password
		rtClaims["isAdmin"] = user.IsAdmin

		rToken := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
		td.RefreshToken, err = rToken.SignedString(jwtKey)
 		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode("Refresh token not signed")
			return
		}
	} else {
		w.WriteHeader(response.StatusCode)
		json.NewEncoder(w).Encode(response.Message)
		return
	}
	tokens := map[string]string{
		"accessToken": td.AccessToken,
		"refreshToken": td.RefreshToken,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tokens)
}

func TokenIsVaild(w http.ResponseWriter, r *http.Request){
	bearToken := r.Header.Get("Authorization")
	tokenArr := strings.Split(bearToken, " ")
	if len(tokenArr) != 2 {
		w.WriteHeader(http.StatusUnauthorized )
		json.NewEncoder(w).Encode("Token not found")
		return
	}

	secret := []byte("Bu_Asiri_Gizli_Bir_Sifredir")

	token, err := jwt.Parse(tokenArr[1], func(token *jwt.Token) (interface{}, error){
		if _, ok :=token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Token not parsed")
		}
		return secret, nil
	})
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized )
		json.NewEncoder(w).Encode("Token is expired")
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims["IsAdmin"] == true {
			w.WriteHeader(http.StatusOK )
			json.NewEncoder(w).Encode("admin")
			return
		} else if claims["IsAdmin"] == false {
			w.WriteHeader(http.StatusOK )
			json.NewEncoder(w).Encode("user")
			return
		}
	}
	w.WriteHeader(http.StatusUnauthorized )
	json.NewEncoder(w).Encode("Not Valid")
}