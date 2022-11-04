package main

import (
	"flag"
	"fmt"
)

func main() {
	communicationType := flag.String("c", "queue", "Communication type")
	encryptionType := flag.String("e", "aes", "Encryption type")

	fmt.Println(*communicationType, *encryptionType)
}
