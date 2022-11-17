package main

import (
	"flag"
	"fmt"
	"lab2/client/services"
)

func main() {
	communicationType := flag.String("c", "queue", "Communication type")
	flag.Parse()

	fmt.Println(*communicationType)

	if *communicationType == "queue" {
		services.SendDataViaQueue()
	} else if *communicationType == "socket" {
		services.SendDataViaSocket()
	} else {
		fmt.Println(fmt.Sprintf("Ключ  %s не найден", *communicationType))
	}
}
