package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go-fiber-prisma-ex/api"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func NewServer() *fiber.App {
	app := fiber.New()
	api.New(app)
	return app
}

func main() {
	app := NewServer()
	// Listen from a different goroutine
	go func() {
		if err := app.Listen(":3000"); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)   // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	_ = <-c
	fmt.Println("Gracefully shutting down...")
	_ = app.Shutdown()

	fmt.Println("Running cleanup tasks...")

	// cleanup tasks go here
	// db.Close()
	// redisConn.Close()
	fmt.Println("Fiber was successful shutdown.")
}