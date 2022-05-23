package WebAPI

import(
    //Local Packages
    "net/http"
    "log"
    "sync"
    //This Project Packages
	"bulkmail/packages/Presentation/WebAPI/Controllers"
    "bulkmail/packages/Utils/Logger"
)

func WebAPI(wg *sync.WaitGroup) {
    defer wg.Done()
    Controllers.SendMail()
    Logger.Print("Server Started")
    log.Fatal(http.ListenAndServe(":10000", nil))
}