package main

import (
	"burmese_jewellery/server"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// Initialize the server
	svc := server.NewServer()

	// Run the server in a separate goroutine to allow for graceful shutdown
	go func() {
		if err := svc.Run(); err != nil {
			panic(err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)                    // Create a channel to listen for OS signals
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM) // Notify the channel on interrupt or termination signals
	<-quit                                             // Block until a signal is received
	svc.Shutdown()                                     // Call the shutdown method to gracefully stop the server

}
