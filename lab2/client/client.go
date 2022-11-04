package main

import (
	"flag"
	"fmt"
	"lab2/client/services"
)

func main() {
	communicationType := flag.String("c", "queue", "Communication type")
	encryptionType := flag.String("e", "aes", "Encryption type")

	fmt.Println(*communicationType, *encryptionType)

	if *communicationType == "queue" {
		services.SendDataViaQueue(*encryptionType)
	}
}
