package reciver

import (
	"log"

	"github.com/gofiber/fiber/v2"
	c "github.com/rabbitmq/amqp091-go"
)

func Recive(cl *fiber.Ctx) error {
	/* 	var queue string
	   	fmt.Print("Enter the sender queuename: ")
	   	fmt.Scanf("%s", &queue) */

	queue := cl.FormValue("queue")

	conn, _ := c.Dial("amqp://guest:guest@localhost:5672/")
	defer conn.Close()

	ch, _ := conn.Channel()
	defer ch.Close()

	q, _ := ch.QueueDeclare(queue, false, false, false, false, nil)

	message, _ := ch.Consume(q.Name, "", true, false, false, false, nil)

	//loop := make(chan bool)
	log.Println(message)

	//var YA []byte
	go func() {
		for d := range message {

			//	YA = d.Body
			log.Printf("Received a message: %s", d.Body)

		}
	}()

	log.Printf(" message is arraving ....... ")

	//<-loop

	return cl.Render("./templates/reciver.html", fiber.Map{
		"tag":  "data is recving ..... ",
		"data": message,
	})
}

func ReciverTemplates(c *fiber.Ctx) error {
	// Render index template
	return c.Render("./templates/reciver.html", fiber.Map{
		"Title": "Hello, World!",
	})
}

/*
 */
