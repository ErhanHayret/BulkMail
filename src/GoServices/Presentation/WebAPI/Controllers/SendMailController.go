package Controllers

import(
	//Local Packages
	"net/http"
	
	//This Project Packages
	bsns "bulkmail/packages/Business"
)

func SendMail() {
    http.HandleFunc("/SendMail", bsns.AddToQueue)
}
