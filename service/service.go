package service

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"student-grade/registry"
)

func Start(ctx context.Context, reg registry.Registration, host, port string, registerHandlersFunc func()) (context.Context, error) {

	registerHandlersFunc()

	ctx = startService(ctx, reg.ServiceName, host, port)
	err := registry.RegisterService(reg)
	if err != nil {
		return ctx, err
	}
	return ctx, nil
}

func startService(ctx context.Context, serviceName registry.ServiceName, host, port string) context.Context {

	ctx, cancel := context.WithCancel(ctx)

	var srv http.Server
	srv.Addr = ":" + port

	go func() {
		log.Println(srv.ListenAndServe())
		cancel()
	}()

	go func() {
		log.Printf("%v started on port %v. Press any key to stop\n", serviceName, port)
		var s string
		fmt.Scanln(&s)
		srv.Shutdown(ctx)
		cancel()
	}()

	return ctx
}
