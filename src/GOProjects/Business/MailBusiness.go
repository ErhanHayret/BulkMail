package Business

import(
	//Local Packages
	"net/http"
	"encoding/json"
	"io/ioutil"
	
	//ThisProject Packages
	model	"bulkmail/packages/Data/Models"
	eLog 	"bulkmail/packages/Utils/Logger"
	rabbit 	"bulkmail/packages/Utils/RabbitMQ"
	mongo "bulkmail/packages/DataAccess/MongoDb"
	"go.mongodb.org/mongo-driver/bson/primitive"

	//Online Packages
	"github.com/gorilla/mux"
)

var collection = mongo.GetClient("MailDb", "Mail")

func AddToQueue(w http.ResponseWriter, r *http.Request){
	
	reqBody, _ := ioutil.ReadAll(r.Body)
	var mail model.Mail 
	json.Unmarshal(reqBody, &mail)
	mail.Id=primitive.NewObjectID()
	body, err := json.Marshal(mail)
	if err != nil{
		eLog.FailOnError(err, "Failed data can't converting")
	}

	rabbit.AddToQueue(body)
}

func GetAllMails(w http.ResponseWriter, r *http.Request){
	//collection := mongo.GetClient("MailDb", "Mail")
	result := mongo.FindAll(collection)
	json.NewEncoder(w).Encode(result)
}

func GetMailById(w http.ResponseWriter, r*http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	objId, _ := primitive.ObjectIDFromHex( id)
	result := mongo.FindById(collection, objId)
	json.NewEncoder(w).Encode(result)
}

func GetMailsBySender(w http.ResponseWriter, r*http.Request){
	vars := mux.Vars(r)
	adress := vars["adress"]
	result := mongo.FindBySender(collection, adress)
	json.NewEncoder(w).Encode(result)
}