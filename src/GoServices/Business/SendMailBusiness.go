package Business

import(
	//Local Packages
	"net/http"
	"encoding/json"
	
	//ThisProject Packages
	models "bulkmail/packages/Data/Models"
	myLog "bulkmail/packages/Utils/Logger"
	rabbit "bulkmail/packages/Utils/RabbitMQ"
)

func AddToQueue(w http.ResponseWriter, r *http.Request){
	var mail models.Mail 

	json.NewDecoder(r.Body).Decode(&mail)
	body, err := json.Marshal(mail)
	if err != nil{
		myLog.FailOnError(err, "Failed data can't converting")
	}

	rabbit.AddToQueue(body)
}