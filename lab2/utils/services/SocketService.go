package services

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"lab2/utils/dto"
	"log"
	"net"
)

func PutUnnormalizedDataToSocket(unnormalizedStudents []dto.UnnormalizedStudent) {
	tlsConf := GetServerCerts()

	conn, err := tls.Dial("tcp", "176.124.200.41:9876", tlsConf)
	failOnError(err, "Failed to connect to socket")
	defer conn.Close()

	for _, elem := range unnormalizedStudents {
		byteData, err := json.Marshal(elem)
		reader := bytes.NewReader(byteData)
		log.Printf(" [x] Sent %s\n", byteData)
		failOnError(err, "Failed to marshal data")
		_, err = io.Copy(conn, reader)
		failOnError(err, "Failed to copy data")
	}
}

func GetUnnormalizedDataFromSocket() {
	tlsConf := GetServerCerts()

	listener, _ := tls.Listen("tcp", "0.0.0.0:9876", tlsConf)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn) // обрабатываем запросы клиента в отдельной го-рутине
	}

}

func handleClient(conn net.Conn) {
	defer conn.Close()
	us := dto.UnnormalizedStudent{}
	d := json.NewDecoder(conn)
	for d.More() {
		_ = d.Decode(&us)
		fmt.Println(us)
		normalizeStudent(us)
	}

}
