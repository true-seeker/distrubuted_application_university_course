package services

import (
	"bytes"
	"crypto/rand"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"lab2/utils/dto"
	"net"
)

func PutUnnormalizedDataToSocket(unnormalizedStudents []dto.UnnormalizedStudent, encryptionType string) {
	//conn, err := net.Dial("tcp", "176.124.200.41:9876")
	var conn net.Conn
	if encryptionType == "tls" {
		cert, err := tls.LoadX509KeyPair("../certs/client.pem", "../certs/client.key")
		failOnError(err, "Failed to load keys")
		config := tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true}
		conn, err = tls.Dial("tcp", "localhost:9876", &config)
		failOnError(err, "Failed to connect to socket")
		defer conn.Close()

		//state := conn.ConnectionState()
		//log.Println("client: handshake: ", state.HandshakeComplete)

	} else if encryptionType == "aes" {
		conn, err := net.Dial("tcp", "localhost:9876")
		failOnError(err, "Failed to connect to socket")
		defer conn.Close()
	}
	for _, elem := range unnormalizedStudents {
		byteData, err := json.Marshal(elem)
		reader := bytes.NewReader(byteData)
		failOnError(err, "Failed to marshal data")
		_, err = io.Copy(conn, reader)
		failOnError(err, "Failed to copy data")
	}
}

func GetUnnormalizedDataFromSocket(encryptionType string) {
	var listener net.Listener
	if encryptionType == "tls" {
		cert, err := tls.LoadX509KeyPair("../certs/server.pem", "../certs/server.key")
		failOnError(err, "Failed to load keys")
		config := tls.Config{Certificates: []tls.Certificate{cert}}
		config.Rand = rand.Reader
		listener, err = tls.Listen("tcp", "localhost:9876", &config)

	} else if encryptionType == "aes" {
		listener, _ = net.Listen("tcp", "localhost:9876") // открываем слушающий сокет
	}

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
