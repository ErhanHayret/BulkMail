package main

import(
    //Local Packages
    "sync"
    //This Project Packages
	"bulkmail/packages/Presentation/WebAPI"
    "bulkmail/packages/Presentation/Consumer"
)

func main() {
    var wg sync.WaitGroup

    wg.Add(2)
    go WebAPI.WebAPI(&wg)
    go Consumer.Consumer(&wg)

    wg.Wait()
}