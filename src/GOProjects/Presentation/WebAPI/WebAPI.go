package WebAPI

import(
    "net/http"
    "log"
    "sync"
    
    "github.com/gorilla/mux"

    mailBus "bulkmail/packages/Business/MailBusiness"
    userBus "bulkmail/packages/Business/UserBusiness"
    eLog "bulkmail/packages/Utils/Logger"

)

func WebAPI(wg *sync.WaitGroup) {
    defer wg.Done()
    
    router := mux.NewRouter().StrictSlash(true)

    router.HandleFunc("/SendMail", mailBus.SendMail).Methods("POST")
    router.HandleFunc("/GetAllMails", mailBus.GetAllMails).Methods("GET")
    router.HandleFunc("/GetMailById/{id}", mailBus.GetMailById).Methods("GET")
    router.HandleFunc("/GetMailBySender/{adress}", mailBus.GetMailsBySender).Methods("GET")

    router.HandleFunc("/GetToken", userBus.GetToken).Methods("POST")
    
    eLog.Print("Server Started. Port: 10000")
    log.Fatal(http.ListenAndServe(":10000", router))
}