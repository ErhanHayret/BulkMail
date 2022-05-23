package Logger

import(
	//Local Packages
	"log"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf(" [MESSAGE] %s [ERROR] %s", msg, err)
	}
}

func PrintData(str string, data any){
	log.Printf(" [LOG] %s %s\n", str, data)
}

func Print(str string){
	log.Printf("[LOG] %s\n",str)
}