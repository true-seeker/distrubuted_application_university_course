openssl genrsa -out RMQ-CA-Key.pem
openssl req -new -key RMQ-CA-Key.pem -x509 -days 100 -out RMQ-CA-cert.pem
openssl genrsa -out RMQ-server-key.pem
openssl req -new -config /etc/ssl/openssl.cnf -key RMQ-server-key.pem -out RMQ-signingrequest.csr
openssl x509 -req -days 100 -in RMQ-signingrequest.csr -CA RMQ-CA-cert.pem -CAkey RMQ-CA-Key.pem -CAcreateserial -out RMQ-server-cert.pem
#cat RMQ-server-key.pem RMQ-server-cert.pem > RMQ-serverpemkeyfile.pem

