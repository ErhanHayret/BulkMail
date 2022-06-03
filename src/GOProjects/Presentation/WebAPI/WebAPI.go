package WebAPI

import(
    "net/http"
    "log"
    "sync"
    
    "github.com/gorilla/mux"

    mailBus "bulkmail/packages/Business/MailBusiness"
    userBus "bulkmail/packages/Business/UserBusiness"
    tokenBus "bulkmail/packages/Business/JwtTokenBusiness"
    eLog "bulkmail/packages/Utils/Logger"

)

func WebAPI(wg *sync.WaitGroup) {
    defer wg.Done()
    
    router := mux.NewRouter().StrictSlash(true)

    router.HandleFunc("/Mail/SendMail", mailBus.SendMail).Methods("POST")
    router.HandleFunc("/Mail/GetAllMails", mailBus.GetAllMails).Methods("GET")
    router.HandleFunc("/Mail/GetMailById/{id}", mailBus.GetMailById).Methods("GET")
    router.HandleFunc("/Mail/GetMailBySender/{adress}", mailBus.GetMailsBySender).Methods("GET")

    router.HandleFunc("/User/GetUser", userBus.GetUser).Methods("POST")
    router.HandleFunc("/User/GetUserById/{id}", userBus.GetUserById).Methods("GET")
    router.HandleFunc("/User/AddUser", userBus.AddUser).Methods("POST")
    router.HandleFunc("/User/UpdateUser", userBus.UpdateUser).Methods("PUT")
    router.HandleFunc("/User/DeleteUser/{id}", userBus.DeleteUser).Methods("DELETE")

    router.HandleFunc("/Token/GetToken", tokenBus.CreateToken).Methods("POST")//token not work
    router.HandleFunc("/Token/IsValid", tokenBus.TokenIsVaild).Methods("GET")//token not work
    
    eLog.Print("Server Started. Port: 10000")
    log.Fatal(http.ListenAndServe(":10000", router))
}