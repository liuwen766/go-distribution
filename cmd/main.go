package main

import (
	"context"
	"distributed/log"
	"distributed/registry"
	"distributed/service"
	"fmt"
	stlog "log"
)

func main() {
	log.Run("./distr.log")
	host, port := "localhost", "4000"
	serviceAddress := fmt.Sprintf("http://%s:%s", host, port)
	r := registry.Registration{
		ServiceName: "Log Service",
		ServiceURL:  serviceAddress,
	}
	start, err := service.Start(
		context.Background(),
		host,
		port,
		r,
		log.RegisterHandler,
	)
	if err != nil {
		stlog.Fatalln(err)
	}
	<-start.Done()
	fmt.Println("shutting down log service.")
}
