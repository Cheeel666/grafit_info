package main

import (
	"os"
	"os/signal"

	"go.uber.org/zap"
	"projects/grafit_info/config"
	"projects/grafit_info/internal/database/mongodb"
	"projects/grafit_info/internal/logger"
	"projects/grafit_info/internal/server"
)

func main() {
	// Init config.
	cfg := config.InitConfig()

	// Get logger.
	log := logger.NewLogger(cfg)

	// Connect to MongoDB.
	mongoClient := mongodb.NewMongoDB(cfg, log)

	err := mongoClient.Connect()
	if err != nil {
		log.Fatal("failed to connect to db:", zap.String("err", err.Error()))
	}
	defer mongoClient.Release()

	// Init web server.
	ws := server.NewServer(cfg, log)
	go ws.Run()

	// Wait for interrupt signal to gracefully shutdown the server with.
	interruptSignal := make(chan os.Signal, 1)
	signal.Notify(interruptSignal, os.Interrupt, os.Kill)
	takeSig := <-interruptSignal
	log.Info("Got signal to shutdown", zap.String("signal", takeSig.String()))
}
