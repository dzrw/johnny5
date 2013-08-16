package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	_ "time"
)

func main() {
	log.Println("starting....")

	// Setup the root HTTP server.
	srv := NewHttpServer(":8080", &WatchHandler{})
	//srv.ReadTimeout = 1 * time.Second

	// Start the root HTTP server.
	err := srv.Start()
	if err != nil {
		log.Fatal(err)
	}

	// Set up channel on which to send signal notifications.
	// We must use a buffered channel or risk missing the signal
	// if we're not ready to receive when the signal is sent.
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	log.Println("CTRL-C to exit...")

	// Block until we receive a signal.
	log.Println("Got signal: ", <-ch)

	// Stop the root HTTP server, and any child servers that may
	// have been created.
	srv.Stop()

	log.Println("goodbye...")
}
