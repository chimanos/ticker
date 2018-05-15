package database

import (
	"strconv"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
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

var db *gorm.DB

func Connect(host string, port int, dbname string, user string, password string) {
	var err error
	
	key := fmt.Sprintf("host=%s port=%v dbname=%s user=%s password=%s", host, port, dbname, user, password)
	db, err = gorm.Open("postgres", key)

	if err != nil || db == nil {
		panic("Connection to Ticker database failed.")
	}

	db.LogMode(true)
	db.SingularTable(true)
}

func AddMessage(message Message) {
	if db != nil {
		db.Create(&message)
		fmt.Println("Message added successfully.")
	} else {
		panic("Error with add message, database is null and not connected.")
	}
}

/*func GetPricesBetween(start string, end string) []float64 {
	prices := db.Where("time BETWEEN ? AND ?", start, end).Find(&Message{}).Value
	return make([]float64, prices)
}*/

func SelectAllMessage() interface{} {
	return db.Find(&Message{}).Value
}

func Close() {
	if db != nil {
		defer db.Close()
	}
}




