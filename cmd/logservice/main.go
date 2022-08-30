package main

import (
	"context"
	"fmt"
	stlog "log"
	"student-grade/log"
	"student-grade/registry"
	"student-grade/service"
)

func main() {

	log.Run("./app.log")

	host, port := "localhost", "4000"
	var r registry.Registration
	r.ServiceName = registry.LogService
	r.ServiceURL = "http://" + host + ":" + port

	ctx, err := service.Start(
		context.Background(),
		r,
		host,
		port,
		log.RegisterHandlers,
	)
	if err != nil {
		stlog.Fatal(err)
	}

	<-ctx.Done()
	fmt.Println("shutdown service log")
}
