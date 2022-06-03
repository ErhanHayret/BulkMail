package Consumer

import(
	"sync"
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"

	"bulkmail/packages/Data/Models"
	sender "bulkmail/packages/Utils/SmtpWorkflow"
	eLog "bulkmail/packages/Utils/Logger"
	bus "bulkmail/packages/Business/ConsumerBusiness"
)

func Consumer(wg *sync.WaitGroup) {
	defer wg.Done()

	//Rabbit Connection
	conn, err := amqp.Dial("amqp://root:root@localhost:5672/")
	eLog.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	//Rabbit Channel
	ch, err := conn.Channel()
	eLog.FailOnError(err, "Failed to open a channel")
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
	eLog.FailOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	eLog.FailOnError(err, "Failed to register a consumer")

	var forever chan struct{}
	
	go func() {
		for d := range msgs {
			var mail Models.MailModel

			err := json.Unmarshal(d.Body, &mail)
			if err != nil {
				eLog.FailOnError(err, "Consumer can't unmarshall data")
			}
			
			dbResponse := bus.Insert(mail)
			sender.Send(mail)
			if dbResponse.Status == false{
				eLog.FailOnError(nil, dbResponse.Message)
			}
		}
	}()

	eLog.Print("To Close The Project Press To 'CTRL+C'")
	<-forever
}