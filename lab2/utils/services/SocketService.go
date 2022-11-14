package services

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io"
	"lab2/utils/dto"
	"net"
	"os"
)

func PutUnnormalizedDataToSocket(unnormalizedStudents []dto.UnnormalizedStudent) {
	caCert, err := os.ReadFile("../certs/ca_certificate.pem")
	failOnError(err, "Failed to open ca cert")
	cert, err := tls.LoadX509KeyPair("../certs/client_v1234281.hosted-by-vdsina.ru_certificate.pem",
		"../certs/client_v1234281.hosted-by-vdsina.ru_key.pem")
	failOnError(err, "Failed to load keypair")
	rootCas := x509.NewCertPool()
	rootCas.AppendCertsFromPEM(caCert)
	tlsConf := &tls.Config{RootCAs: rootCas,
		Certificates: []tls.Certificate{cert}}
	conn, err := tls.Dial("tcp", "176.124.200.41:9876", tlsConf)
	failOnError(err, "Failed to connect to socket")
	defer conn.Close()

	for _, elem := range unnormalizedStudents {
		byteData, err := json.Marshal(elem)
		reader := bytes.NewReader(byteData)
		failOnError(err, "Failed to marshal data")
		_, err = io.Copy(conn, reader)
		failOnError(err, "Failed to copy data")
	}
}

func GetUnnormalizedDataFromSocket() {
	caCert, err := os.ReadFile("../certs/ca_certificate.pem")
	failOnError(err, "Failed to open ca cert")
	cert, err := tls.LoadX509KeyPair("../certs/client_v1234281.hosted-by-vdsina.ru_certificate.pem",
		"../certs/client_v1234281.hosted-by-vdsina.ru_key.pem")
	failOnError(err, "Failed to load keypair")
	rootCas := x509.NewCertPool()
	rootCas.AppendCertsFromPEM(caCert)
	tlsConf := &tls.Config{RootCAs: rootCas,
		Certificates: []tls.Certificate{cert}}

	listener, err := tls.Listen("tcp", "localhost:9876", tlsConf)

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
