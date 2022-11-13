package services

import (
	"bytes"
	"crypto/rand"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"lab2/utils/dto"
	"log"
	"net"
)

func PutUnnormalizedDataToSocket(unnormalizedStudents []dto.UnnormalizedStudent, encryptionType string) {
	//conn, err := net.Dial("tcp", "176.124.200.41:9876")

	if encryptionType == "tls" {
		cert, err := tls.LoadX509KeyPair("../certs/client.pem", "../certs/client.key")
		failOnError(err, "Failed to load keys")
		config := tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true}

		conn, err := tls.Dial("tcp", "localhost:9876", &config)
		failOnError(err, "Failed to connect to socket")
		defer conn.Close()
		state := conn.ConnectionState()
		log.Println("client: handshake: ", state.HandshakeComplete)

		for _, elem := range unnormalizedStudents {
			byteData, err := json.Marshal(elem)
			reader := bytes.NewReader(byteData)
			failOnError(err, "Failed to marshal data")
			_, err = io.Copy(conn, reader)
			failOnError(err, "Failed to copy data")
		}
	} else if encryptionType == "aes" {
		conn, err := net.Dial("tcp", "localhost:9876")
		failOnError(err, "Failed to connect to socket")
		defer conn.Close()
	}

}

func GetUnnormalizedDataFromSocket(encryptionType string) {
	if encryptionType == "tls" {
		cert, err := tls.LoadX509KeyPair("../certs/server.pem", "../certs/server.key")
		failOnError(err, "Failed to load keys")
		config := tls.Config{Certificates: []tls.Certificate{cert}}
		config.Rand = rand.Reader
		listener, err := tls.Listen("tcp", "localhost:9876", &config)

		for {
			conn, err := listener.Accept()
			if err != nil {
				continue
			}
			go handleClient(conn) // обрабатываем запросы клиента в отдельной го-рутине
		}
	} else if encryptionType == "aes" {

	}
	//listener, _ := net.Listen("tcp", "localhost:9876") // открываем слушающий сокет

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
