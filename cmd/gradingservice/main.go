package main

import (
	"context"
	"fmt"
	stlog "log"
	"student-grade/grades"
	"student-grade/log"
	"student-grade/registry"
	"student-grade/service"
)

func main() {
	host, port := "localhost", "6000"
	serviceAddress := "http://" + host + ":" + port

	var r registry.Registration
	r.ServiceName = registry.GradingService
	r.ServiceURL = serviceAddress
	r.RequiredServices = []registry.ServiceName{registry.LogService}
	r.ServiceUpdateURL = r.ServiceURL + "/services"

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
	if logProvider, err := registry.GetProvider(registry.LogService); err == nil {
		fmt.Printf("Logging service found at: %v\n", logProvider)
		log.SetClientLogger(logProvider, r.ServiceName)
	}

	<-ctx.Done()
	fmt.Println("shutdown grading service")
}
