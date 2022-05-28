package Logger

import(
	"log"
)

func FailOnError(err error, msg string) {
	if err != nil{
		log.Printf(" [MESSAGE] %s [ERROR] %s", msg, err)
	}
}

func ErrorPrint(msg string){
	log.Printf(" [ERROR] %s", msg)
}

func PrintData(str string, data any){
	log.Printf(" [LOG] %s %s\n", str, data)
}

func Print(str string){
	log.Printf("[LOG] %s\n",str)
}