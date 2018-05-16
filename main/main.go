package main

import (
	"os"
	"os/signal"
	"syscall"
	"strings"

	"github.com/chimanos/ticker/database"
	"github.com/chimanos/ticker/consumer"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)

	if os.Getenv("TICKER_ENV") != "" {
		log.SetLevel(log.InfoLevel)
	}
}

func main() {
	signalChannel := make(chan os.Signal, 2)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)

	clusterIPs := strings.Split(os.Getenv("TICKER_KAFKA_IPS"), ",")
	kafkaTopics := strings.Split(os.Getenv("TICKER_KAFKA_TOPICS"), ",")
	kafkaGroup := os.Getenv("TICKER_KAFKA_GROUPS")

	databaseHosts := strings.Split(os.Getenv("TICKER_DATABASE_HOSTS"), ",")
	databaseName := os.Getenv("TICKER_DATABASE_NAME")
	databasePort := os.Getenv("TICKER_DATABASE_PORT")
	databaseUser := os.Getenv("TICKER_DATABASE_USER")
	databasePassword := os.Getenv("TICKER_DATABASE_PASSWORD")

	database, err := database.NewDatabase(databaseHosts, databaseName, databasePort, databaseUser, databasePassword)

	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatal("Unable to start database connection")

		os.Exit(1)
	}

	consumer, err := consumer.NewConsumer(clusterIPs, kafkaTopics, kafkaGroup)

	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatal("Unable to instantiate ticker consumer")

		os.Exit(1)
	}

	go func() {
		<-signalChannel
		database.Close()
		consumer.Stop()
		close(signalChannel)
	}()

	messages := consumer.Start()
	database.Start(messages)

	log.Info("Candelabot ticker shutted down successfully")
	os.Exit(0)
}


