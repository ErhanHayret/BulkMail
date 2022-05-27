package Business

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

func AddToQueue(w http.ResponseWriter, r *http.Request){
	
	var result model.StatusResult
	var mail model.Mail 

	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &mail)
	mail.Id=primitive.NewObjectID()
	body, err := json.Marshal(mail)
	if err != nil{
		result.Error = err
		result.Message = "MailBusiness -> AddToQueue => Failed data can't converting"
		result.Status = false
	}
	rabbit.AddToQueue(body)
	json.NewEncoder(w).Encode(result)
}

func GetAllMails(w http.ResponseWriter, r *http.Request){
	result, dbResponse := mongo.FindAll(collection)
	if dbResponse.Status == false {
		json.NewEncoder(w).Encode(dbResponse)
		return
	}
	json.NewEncoder(w).Encode(result)
}

func GetMailById(w http.ResponseWriter, r*http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	objId, _ := primitive.ObjectIDFromHex(id)
	result, dbResponse := mongo.FindById(collection, objId)		
	if dbResponse.Status == false {
		json.NewEncoder(w).Encode(dbResponse)
		return
	}
	json.NewEncoder(w).Encode(result)
}

func GetMailsBySender(w http.ResponseWriter, r*http.Request){
	vars := mux.Vars(r)
	adress := vars["adress"]
	result, dbResponse := mongo.FindBySender(collection, adress)
	if dbResponse.Status == false {
		json.NewEncoder(w).Encode(dbResponse)
		return
	}
	json.NewEncoder(w).Encode(result)
}