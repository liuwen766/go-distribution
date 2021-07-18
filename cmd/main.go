package main

import (
	"context"
	"distributed/log"
	"distributed/service"
	"fmt"
	stlog "log"
)

func main() {
	log.Run("./distr.log")
	host, port := "localhost", "4000"
	start, err := service.Start(
		context.Background(),
		"My Log",
		host,
		port,
		log.RegisterHandler,
	)
	if err != nil {
		stlog.Fatalln(err)
	}
	<-start.Done()
	fmt.Println("shutting down log service.")
}
