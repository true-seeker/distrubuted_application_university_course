package main

import (
	"flag"
	"fmt"
	"lab2/client/services"
)

func main() {
	communicationType := flag.String("c", "queue", "Communication type")
	encryptionType := flag.String("e", "aes", "Encryption type")
	flag.Parse()

	fmt.Println(*communicationType, *encryptionType)

	if *communicationType == "queue" {
		services.SendDataViaQueue(*encryptionType)
	} else if *communicationType == "socket" {
		services.SendDataViaSocket(*encryptionType)
	}
}
