package MailBusiness

import(
	"net/http"
	"encoding/json"
	"io/ioutil"
	
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"

	model	"bulkmail/packages/Data/Models"
	rabbit 	"bulkmail/packages/Utils/RabbitMQ"
	mongo "bulkmail/packages/DataAccess/MongoDb"
)

var collection, dbResponse = mongo.GetClient("MailDb", "Mail")

func SendMail(w http.ResponseWriter, r *http.Request){
	var result model.ResultModel
	var mail model.MailModel 

	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &mail)
	mail.Id=primitive.NewObjectID()
	body, err := json.Marshal(mail)
	if err != nil{
		result.StatusCode = 500
		result.Message = "SendMail-Marshall failed"
		result.Status = false
		w.WriteHeader(result.StatusCode)
		json.NewEncoder(w).Encode(result)
	}

	rabbit.AddToQueue(body)
	result.StatusCode = 200
	result.Message = "Success"
	result.Status = true
	w.WriteHeader(result.StatusCode)
	json.NewEncoder(w).Encode(result)
}

func GetAllMails(w http.ResponseWriter, r *http.Request){
	result, dbResponse := mongo.FindAllMails(collection)
	
	w.WriteHeader(dbResponse.StatusCode)
	if dbResponse.Status == false {
		json.NewEncoder(w).Encode(dbResponse)
		return
	}
	json.NewEncoder(w).Encode(result)
}

func GetMailById(w http.ResponseWriter, r*http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	if len(id) < 20 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(nil)
		return
	}

	objId, _ := primitive.ObjectIDFromHex(id)
	result, dbResponse := mongo.FindMailById(collection, objId)	

	w.WriteHeader(dbResponse.StatusCode)
	if dbResponse.Status == false {
		json.NewEncoder(w).Encode(dbResponse.Message)
		return
	}
	json.NewEncoder(w).Encode(result)
}

func GetMailsBySender(w http.ResponseWriter, r*http.Request){
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	adress := vars["adress"]
	if len(adress) < 7 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(nil)
		return
	}

	result, dbResponse := mongo.FindBySender(collection, adress)

	w.WriteHeader(dbResponse.StatusCode)
	if dbResponse.Status == false {
		json.NewEncoder(w).Encode(dbResponse.Message)
		return
	}
	json.NewEncoder(w).Encode(result)
}