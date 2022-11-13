package main

import (
	"flag"
	"fmt"
	"lab2/server/services"
	"lab2/utils/orm"
)

func main() {
	communicationType := flag.String("c", "queue", "Communication type")
	encryptionType := flag.String("e", "tls", "Encryption type")
	flag.Parse()

	fmt.Println(*communicationType, *encryptionType)
	orm.Migrate()

	if *communicationType == "queue" {
		services.AcceptDataFromQueue(*encryptionType)
	} else if *communicationType == "socket" {
		services.AcceptDataFromSocket(*encryptionType)
	}

}
