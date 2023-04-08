package main

import (
	_ "armadabackend/logging"
	"armadabackend/routers"
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
)

func main() {
	ctx := context.Background()
	router := http.HandlerFunc(routers.Serve)
	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
		BaseContext: func(l net.Listener) context.Context {
			//modify context for the server here
			return ctx
		},
	}

	err := server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("server is closed")
	} else if err != nil {
		fmt.Printf("error listening for server: %s\n", err)
	}
}
