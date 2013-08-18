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

	if err := AwaitSignals(); err != nil {
		log.Fatalln(err)
	}

	// Stop the root HTTP server, and any child servers that may
	// have been created.
	srv.Stop()

	log.Println("goodbye...")
}

func AwaitSignals() (err error) {
	// Set up channel on which to send signal notifications.
	// We must use a buffered channel or risk missing the signal
	// if we're not ready to receive when the signal is sent.
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	log.Println("CTRL-C to exit...")

	for {
		// Block until we receive a signal.
		sig := <-ch
		log.Println("Got: ", sig.String())

		switch sig {

		// TODO - Handle other signals that don't just stop the
		// process immediately.

		// SIGQUIT should exit gracefully.
		case syscall.SIGQUIT:
			return

		// SIGTERM should exit.
		case syscall.SIGTERM:
			return
		}
	}
}
