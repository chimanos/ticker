package database

import (
	"fmt"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/mxpetit/ticker/consumer"
	"github.com/mxpetit/ticker/utils"
	log "github.com/sirupsen/logrus"
)

type Message struct {
	ID uint `sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Pair string
	Market string
	Price float64
	BestAsk float64
	BestBid float64
	Time int64
}

type Database interface {
	Start(chan *consumer.Message)
	Close()
}


type DatabaseWrapper struct {
	DB *gorm.DB
	waitGroup sync.WaitGroup
	quit chan struct{}
}

func NewDatabase(hosts []string, dbname string, port string, user string, password string) (Database, error) {
	log.WithFields(log.Fields{
		"Hosts":   hosts,
		"Port": port,
		"DBName": dbname,
		"User": user,
	}).Info("Connecting ticker database")

	key := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s", hosts, port, dbname, user, password)
	db, err := gorm.Open("postgres", key)

	if err != nil {
		return nil, err
	}

	db.LogMode(true)
	db.SingularTable(true)

	return &DatabaseWrapper{
		DB: db,
		quit: make(chan struct{}, 2),
	}, nil
}

func (k *DatabaseWrapper) Start(channel chan *consumer.Message) {
	k.waitGroup.Add(1)
	go func() {
		for listen := true; listen; {
			select {
			case msg  := <- channel:
				k.AddMessage(Message{
					Pair: msg.Pair,
					Market: msg.Market,
					Price: msg.Price,
					BestAsk: msg.Best_ask,
					BestBid: msg.Best_bid,
					Time: utils.GetTimestampFromDate(msg.Time)})
			case <-k.quit:
				listen = false
			}
		}
		log.Info("Ticker database closing")
		k.quit <- struct{}{}
	}()
	k.waitGroup.Wait()
}

func (k *DatabaseWrapper) AddMessage(message Message) {
	k.DB.Create(&message)
	log.Info("Message added to database")
}

func (k *DatabaseWrapper) Close() {
	log.Info("Ticker database closing")

	k.quit <- struct{}{}
	<-k.quit
	k.DB.Close()

	log.Info("Ticker consumer shutted down")
}




