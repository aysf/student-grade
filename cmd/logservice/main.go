package main

import (
	"context"
	"fmt"
	stlog "log"
	"student-grade/log"
	"student-grade/service"
)

func main() {

	log.Run("./app.log")

	host, port := "localhost", "4000"

	ctx, err := service.Start(
		context.Background(),
		"log-service",
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
