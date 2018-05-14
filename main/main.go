package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/bsm/sarama-cluster"
	"github.com/mxpetit/ticker/database"
	"encoding/json"
	"github.com/mxpetit/ticker/utils"
)

type ExtraFlux struct {
	Ask_size float64
	Bid_size float64
	Daily_change float64
	Daily_change_percent float64
	Higher_24h float64
	Lower_24h float64
	Volume_24h float64
}

type MessageFlux struct {
	Pair string
	Time string
	Price float64
	Market string
	Best_ask float64
	Best_bid float64
	Extra ExtraFlux
}

func main() {
	//Database connexion
	database.Connect("localhost", 5432, "postgres", "ticker", "root")

	// init (custom) config, enable errors and notifications
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true

	// init consumer
	brokers := []string{"127.0.0.1:9092"}
	topics := []string{"aggregator"}
	consumer, err := cluster.NewConsumer(brokers, "aggragator", topics, config)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	// trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// consume errors
	go func() {
		for err := range consumer.Errors() {
			log.Printf("Error: %s\n", err.Error())
		}
	}()

	// consume notifications
	go func() {
		for ntf := range consumer.Notifications() {
			log.Printf("Rebalanced: %+v\n", ntf)
		}
	}()

	// consume messages, watch signals
	for {
		select {
		case msg, ok := <-consumer.Messages():
			if ok {
				var msgFlux *MessageFlux
				json.Unmarshal([]byte(msg.Value), &msgFlux)
				if msgFlux != nil {
					//Save message in database
					database.AddMessage(database.Message{
						Pair: msgFlux.Pair,
						Market: msgFlux.Market,
						Price: msgFlux.Price,
						BestAsk: msgFlux.Best_ask,
						BestBid: msgFlux.Best_bid,
						Time: utils.GetTimestampFromDate(msgFlux.Time)})
				}
				consumer.MarkOffset(msg, "")	// mark message as processed
			}
		case <-signals:
			database.CloseConnexion()
			return
		}
	}

	database.CloseConnexion()
}
