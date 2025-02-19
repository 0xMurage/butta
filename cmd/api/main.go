package main

import (
	"butta/internal/app/api"
	"butta/pkg/logger"
	"context"
	"errors"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	api.Bootstrap()

	httpServer := &http.Server{
		Addr:    serverAddress(),
		Handler: api.Routes(), //register the app serve // mux
	}

	go runServer(httpServer) //run the app server in a goroutine

	awaitShutdownSignal() //wait for shutdown signal from the system

	api.Shutdown()                     //shutdown of the app
	serverGracefulShutdown(httpServer) //shutdown the app server
}

func serverAddress() string {
	port, found := os.LookupEnv("PORT")
	if !found {
		port = "8000"
	}

	host, found := os.LookupEnv("HOST")
	if !found {
		host = "0.0.0.0"
	}

	return net.JoinHostPort(host, port)
}

func runServer(httpServer *http.Server) {

	logger.Info("Starting server ...", "address", httpServer.Addr)

	if err := httpServer.ListenAndServe(); !errors.Is(http.ErrServerClosed, err) {
		logger.Error("error listening", "error", err)
	}
}

func awaitShutdownSignal() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-signalChan
	logger.Info("Received shutdown signal")
}

func serverGracefulShutdown(server *http.Server) {
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		logger.Fatal("error shutting down server", "err", err)
	}

	logger.Info("Server gracefully stopped")
}
