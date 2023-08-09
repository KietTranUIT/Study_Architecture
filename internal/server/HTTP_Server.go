package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type HTTPServer struct {
	Port   uint
	Ip     string
	server http.Server
}

func (server *HTTPServer) NewHTTPServer(mux *http.ServeMux, port uint, ip string) {
	server.Port = port
	server.Ip = ip

	server.server = http.Server{
		Addr:    fmt.Sprintf("%s:%d", server.Ip, server.Port),
		Handler: mux,
	}
}

func (server *HTTPServer) Start() {
	go func() {
		log.Printf("Server is running on port %d", server.Port)
		if err := server.server.ListenAndServe(); err != nil {
			log.Println("Error: ", err.Error())
			return
		}
	}()
	defer server.Stop()
}

func (server *HTTPServer) Stop() {
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	sig := <-signalCh
	log.Println("Received signal cancle: ", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.server.Shutdown(ctx); err != nil {
		log.Println("Error: ", err.Error())
	}
}
