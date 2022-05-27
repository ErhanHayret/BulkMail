package Consumer

import(
	"sync"
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"

	"bulkmail/packages/Data/Models"
	sender "bulkmail/packages/Utils/MailSender"
	myMongo "bulkmail/packages/DataAccess/MongoDb"
	myLog "bulkmail/packages/Utils/Logger"
)

func Consumer(wg *sync.WaitGroup) {
	defer wg.Done()
	//Rabbit Connection
	conn, err := amqp.Dial("amqp://root:root@localhost:5672/")
	myLog.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	//Rabbit Channel
	ch, err := conn.Channel()
	myLog.FailOnError(err, "Failed to open a channel")
	defer ch.Close()
	//Rabbit Queue
	q, err := ch.QueueDeclare(
		"BulkMail", // name
		false,   	// durable
		false,   	// delete when unused
		false,   	// exclusive
		false,   	// no-wait
		nil,     	// arguments
	)
	myLog.FailOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	myLog.FailOnError(err, "Failed to register a consumer")

	var forever chan struct{}
	
	go func() {
		for d := range msgs {
			myLog.PrintData("Received a message => ", d.Body)
			var mail Models.Mail
			err := json.Unmarshal(d.Body, &mail)
			if err != nil {
				myLog.PrintData("Consumer can't unmarshall data", err)
			}
			myLog.PrintData("recived data =>", mail)
			collection, dbResponse := myMongo.GetClient("MailDb", "Mail")
			if dbResponse.Status == true{
				myMongo.InsertOne(collection, mail)
				sender.Send(mail)
			} else {
				myLog.FailOnError(dbResponse.Error, dbResponse.Message)
			}
		}
	}()

	myLog.Print("To Close The Project Press To 'CTRL+C'")
	<-forever
}