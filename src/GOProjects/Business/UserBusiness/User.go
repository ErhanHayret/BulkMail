package UserBusiness

import(
	"net/http"
	"encoding/json"
	"io/ioutil"

	"bulkmail/packages/Data/Dtos"
	tkn "bulkmail/packages/Business/JwtTokenBusiness"
)

func GetToken(w http.ResponseWriter, r *http.Request){
	var user Dtos.UserDto

	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &user)
	token, respone := tkn.CreateNewToken(user)
	if respone.Status == false {
		json.NewEncoder(w).Encode(respone)
		return
	}
	json.NewEncoder(w).Encode(token)
}