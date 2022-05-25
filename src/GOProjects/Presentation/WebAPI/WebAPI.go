package WebAPI

import(
    //Local Packages
    "net/http"
    "log"
    "sync"

    //This Project Packages
    "bulkmail/packages/Business"
    eLog "bulkmail/packages/Utils/Logger"

    //Online Packages
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