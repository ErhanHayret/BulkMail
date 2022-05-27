package WebAPI

import(
    "net/http"
    "log"
    "sync"

    "bulkmail/packages/Business"
    eLog "bulkmail/packages/Utils/Logger"

    "github.com/gorilla/mux"
)

func WebAPI(wg *sync.WaitGroup) {
    defer wg.Done()
    
    router := mux.NewRouter().StrictSlash(true)

    router.HandleFunc("/SendMail", Business.AddToQueue).Methods("POST")
    router.HandleFunc("/GetAllMails", Business.GetAllMails).Methods("GET")
    router.HandleFunc("/GetMailById/{id}", Business.GetMailById).Methods("GET")
    router.HandleFunc("/GetMailBySender/{adress}", Business.GetMailsBySender).Methods("GET")
    
    eLog.Print("Server Started. Port: 10000")
    log.Fatal(http.ListenAndServe(":10000", router))
}