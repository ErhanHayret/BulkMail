package UserBusiness

import(
	"net/http"
	"io/ioutil"
	"encoding/json"

	"github.com/gorilla/mux"

	"bulkmail/packages/Data/Models"
	"bulkmail/packages/Data/Dtos"
	mongo "bulkmail/packages/DataAccess/MongoDb"
)

var collection, clntResponse = mongo.GetClient("UserDb", "User")

func GetUser(w http.ResponseWriter, r *http.Request){
	if clntResponse.Status == false {
		w.WriteHeader(clntResponse.StatusCode)
		json.NewEncoder(w).Encode(clntResponse.Message)
		return
	}

	var userDto Dtos.UserDto

	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &userDto)

	user, response := mongo.FindUser(collection, userDto)
	if response.Status == false {
		w.WriteHeader(response.StatusCode)
		json.NewEncoder(w).Encode(response.Message)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func GetUserById(w http.ResponseWriter, r *http.Request){
	if clntResponse.Status == false {
		w.WriteHeader(clntResponse.StatusCode)
		json.NewEncoder(w).Encode(clntResponse.Message)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]
	user, response := mongo.FindUserById(collection, id)
	if response.Status == false {
		w.WriteHeader(response.StatusCode)
		json.NewEncoder(w).Encode(response.Message)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func AddUser(w http.ResponseWriter, r *http.Request){
	if clntResponse.Status == false {
		w.WriteHeader(clntResponse.StatusCode)
		json.NewEncoder(w).Encode(clntResponse.Message)
		return
	}
	
	var user Models.UserModel

	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &user)

	dbResponse := mongo.InsertUser(collection, user)
	if dbResponse.Status == false {
		w.WriteHeader(dbResponse.StatusCode)
		json.NewEncoder(w).Encode(dbResponse.Message)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Success")
}

func UpdateUser(w http.ResponseWriter, r *http.Request){
	if clntResponse.Status == false {
		w.WriteHeader(clntResponse.StatusCode)
		json.NewEncoder(w).Encode(clntResponse.Message)
		return
	}
	
	var user Models.UserModel

	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &user)

	dbResponse := mongo.UpdateUser(collection, user)
	if dbResponse.Status == false {
		w.WriteHeader(dbResponse.StatusCode)
		json.NewEncoder(w).Encode(dbResponse.Message)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Success")
}

func DeleteUser(w http.ResponseWriter, r *http.Request){
	if clntResponse.Status == false {
		w.WriteHeader(clntResponse.StatusCode)
		json.NewEncoder(w).Encode(clntResponse.Message)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

	dbResponse := mongo.DeleteUser(collection, id)
	if dbResponse.Status == false {
		w.WriteHeader(dbResponse.StatusCode)
		json.NewEncoder(w).Encode(dbResponse.Message)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Success")
}