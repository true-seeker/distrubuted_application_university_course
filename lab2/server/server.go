package main

import (
	"flag"
	"fmt"
	"lab2/server/services"
	"lab2/utils/orm"
)

func main() {
	communicationType := flag.String("c", "queue", "Communication type")
	flag.Parse()

	fmt.Println(*communicationType)
	orm.Migrate()

	if *communicationType == "queue" {
		services.AcceptDataFromQueue()
	} else if *communicationType == "socket" {
		services.AcceptDataFromSocket()
	}

}
