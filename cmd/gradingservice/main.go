package main

import (
	"context"
	"fmt"
	stlog "log"
	"student-grade/grades"
	"student-grade/registry"
	"student-grade/service"
)

func main() {
	host, port := "localhost", "6000"
	serviceAddress := "http://" + host + ":" + port

	var r registry.Registration
	r.ServiceName = registry.GradingService
	r.ServiceURL = serviceAddress

	ctx, err := service.Start(
		context.Background(),
		r,
		host,
		port,
		grades.RegisterHandlers,
	)
	if err != nil {
		stlog.Fatal(err)
	}

	<-ctx.Done()
	fmt.Println("shutdown grading service")
}
