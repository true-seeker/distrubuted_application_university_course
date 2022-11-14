package main

import (
	"flag"
	"fmt"
	"server/orm"
	"server/services"
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
