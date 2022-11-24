package services

import (
	"crypto/tls"
	"crypto/x509"
	"lab3/utils/config"
	"os"
)

func GetClientCerts() *tls.Config {
	caCert, err := os.ReadFile(config.GetProperty("Ssl", "ca_cert_path"))
	failOnError(err, "Failed to open ca cert")
	cert, err := tls.LoadX509KeyPair(config.GetProperty("Ssl", "client_cert_path"),
		config.GetProperty("Ssl", "client_key_path"))
	failOnError(err, "Failed to load keypair")
	rootCas := x509.NewCertPool()
	rootCas.AppendCertsFromPEM(caCert)
	tlsConf := &tls.Config{RootCAs: rootCas,
		Certificates: []tls.Certificate{cert}}
	return tlsConf
}

func GetServerCerts() *tls.Config {
	caCert, err := os.ReadFile(config.GetProperty("Ssl", "ca_cert_path"))
	failOnError(err, "Failed to open ca cert")
	cert, err := tls.LoadX509KeyPair(config.GetProperty("Ssl", "server_cert_path"),
		config.GetProperty("Ssl", "server_key_path"))
	failOnError(err, "Failed to load keypair")
	rootCas := x509.NewCertPool()
	rootCas.AppendCertsFromPEM(caCert)
	tlsConf := &tls.Config{RootCAs: rootCas,
		Certificates: []tls.Certificate{cert}}
	return tlsConf
}
