package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"lab2/utils/dto"
	"net"
)

func PutUnnormalizedDataToSocket(unnormalizedStudents []dto.UnnormalizedStudent) {
	//conn, err := net.Dial("tcp", "176.124.200.41:9876")
	conn, err := net.Dial("tcp", "localhost:9876")
	failOnError(err, "Failed to connect to socket")

	for _, elem := range unnormalizedStudents {
		byteData, err := json.Marshal(elem)

		reader := bytes.NewReader(byteData)
		failOnError(err, "Failed to marshal data")

		_, err = io.Copy(conn, reader)
		failOnError(err, "Failed to copy data")
	}
}

func GetUnnormalizedDataFromSocket() {
	listener, _ := net.Listen("tcp", "localhost:9876") // открываем слушающий сокет
	for {
		conn, err := listener.Accept() // принимаем TCP-соединение от клиента и создаем новый сокет
		if err != nil {
			continue
		}
		go handleClient(conn) // обрабатываем запросы клиента в отдельной го-рутине
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close() // закрываем сокет при выходе из функции
	us := dto.UnnormalizedStudent{}
	d := json.NewDecoder(conn)
	for d.More() {
		_ = d.Decode(&us)
		fmt.Println(us)
		normalizeStudent(us)
	}

}
