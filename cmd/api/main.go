package main

import (
	"butta/internal/app/http"
	"context"
	"errors"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	app.Bootstrap()

	httpServer := &http.Server{
		Addr:    serverAddress(),
		Handler: app.ApplicationRoutes(), //register the app routes
	}

	go runServer(httpServer) //run the app server in a goroutine

	awaitShutdownSignal() //wait for shutdown signal from the system

	app.Shutdown()                     //shutdown of the app
	serverGracefulShutdown(httpServer) //shutdown the app server
}

func serverAddress() string {
	port, found := os.LookupEnv("PORT")
	if !found {
		port = "8000"
	}

	host, found := os.LookupEnv("HOST")
	if !found {
		host = "localhost"
	}

	return net.JoinHostPort(host, port)
}

func runServer(httpServer *http.Server) {

	slog.Info("Starting server %s\n", httpServer.Addr)

	if err := httpServer.ListenAndServe(); !errors.Is(http.ErrServerClosed, err) {
		slog.Error("error listening: %s\n", err)
	}
}

func awaitShutdownSignal() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-signalChan
	slog.Info("Received shutdown signal\n")
}

func serverGracefulShutdown(server *http.Server) {
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		slog.Error("error shutting down server: %s\n", err)
		os.Exit(1)
	}

	slog.Info("Server gracefully stopped")
}
