package services

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"lab2/utils/dto"
	"log"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
func PutUnnormalizedDataToQueue(unnormalizedStudents []dto.UnnormalizedStudent, encryptionType string) {
	//if encryptionType == "tls" {
	caCert, err := os.ReadFile("../certs/ca_certificate.pem")
	failOnError(err, "Failed to open ca cert")
	cert, err := tls.LoadX509KeyPair("../certs/client_v1234281.hosted-by-vdsina.ru_certificate.pem",
		"../certs/client_v1234281.hosted-by-vdsina.ru_key.pem")
	failOnError(err, "Failed to load keypair")
	rootCas := x509.NewCertPool()
	rootCas.AppendCertsFromPEM(caCert)
	tlsConf := &tls.Config{RootCAs: rootCas,
		Certificates: []tls.Certificate{cert}}

	conn, err := amqp.DialTLS("amqps://lab2:lab2@176.124.200.41:5671/", tlsConf)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	//} else if encryptionType == "aes" {
	//	conn, err := amqp.Dial("amqp://lab2:lab2@176.124.200.41:5672/")
	//	failOnError(err, "Failed to connect to RabbitMQ")
	//	defer conn.Close()
	//}

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"unnormalizedDataQueue", // name
		false,                   // durable
		false,                   // delete when unused
		false,                   // exclusive
		false,                   // no-wait
		nil,                     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	for _, elem := range unnormalizedStudents {
		byteData, err := json.Marshal(elem)

		failOnError(err, "Failed to marshal data")

		putDataToQueue(ch, ctx, q, byteData)
	}
}

func putDataToQueue(ch *amqp.Channel, ctx context.Context, q amqp.Queue, body []byte) {
	err := ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)

}

func GetUnnormalizedDataFromQueue(encryptionType string) {
	caCert, err := os.ReadFile("../certs/ca_certificate.pem")
	failOnError(err, "Failed to open ca cert")
	cert, err := tls.LoadX509KeyPair("../certs/client_v1234281.hosted-by-vdsina.ru_certificate.pem",
		"../certs/client_v1234281.hosted-by-vdsina.ru_key.pem")
	failOnError(err, "Failed to load keypair")
	rootCas := x509.NewCertPool()
	rootCas.AppendCertsFromPEM(caCert)
	tlsConf := &tls.Config{RootCAs: rootCas,
		Certificates: []tls.Certificate{cert}}

	conn, err := amqp.DialTLS("amqps://lab2:lab2@176.124.200.41:5671/", tlsConf)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"unnormalizedDataQueue", // name
		false,                   // durable
		false,                   // delete when unused
		false,                   // exclusive
		false,                   // no-wait
		nil,                     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			us := dto.UnnormalizedStudent{}
			json.Unmarshal(d.Body, &us)
			fmt.Println(us)
			normalizeStudent(us)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
