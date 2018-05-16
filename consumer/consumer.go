package consumer

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
	"github.com/bsm/sarama-cluster"
	"github.com/Shopify/sarama"
)

type extra struct {
	Ask_size float64
	Bid_size float64
	Daily_change float64
	Daily_change_percent float64
	Higher_24h float64
	Lower_24h float64
	Volume_24h float64
}

type Message struct {
	Pair string
	Time string
	Price float64
	Market string
	Best_ask float64
	Best_bid float64
	Extra extra
}

type Consumer interface {
	Start() chan *Message
	Stop()
}


type ClusterConsumer struct {
	Consumer cluster.Consumer
	Messages chan *Message
	Topics []string
	quit chan struct{}
}

func NewConsumer(clusterIPs []string, topics []string, groups string) (Consumer, error) {
	log.WithFields(log.Fields{
		"IPs":   clusterIPs,
		"topics": topics,
	}).Info("Initializing ticker consumer")

	consumer, err := cluster.NewConsumer(clusterIPs, groups, topics, cluster.NewConfig())

	if err != nil {
		return nil, err
	}

	return &ClusterConsumer{
		Consumer: *consumer,
		Messages:	make(chan *Message),
		Topics: topics,
		quit: make(chan struct{}, 2),
	}, nil
}

func (k *ClusterConsumer) Start() chan *Message {
	go func() {
		log.Info("Ticker consumer listening")

		for listen := true; listen; {
			select {
			case msg := <-k.Consumer.Messages():
				k.publishMessage(msg)
				k.Consumer.MarkOffset(msg, "")	// mark message as processed
			case <-k.quit:
				listen = false
			}
		}

		log.Info("Ticker consumer stopped listening")
		k.quit <- struct{}{}
	}()

	return k.Messages
}

func (k *ClusterConsumer) Stop() {
	log.Info("Ticker consumer stopping")

	k.quit <- struct{}{}
	<-k.quit
	k.Consumer.Close()
	close(k.Messages)

	log.Info("Ticker consumer shutted down")
}

func (k *ClusterConsumer) publishMessage(msgFlux *sarama.ConsumerMessage) {
	var message *Message
	err := json.Unmarshal([]byte(msgFlux.Value), &message)

	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Warn("Unable to publish message")

		return
	}

	k.Messages <- message
}
