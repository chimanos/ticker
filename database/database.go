package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"strconv"
	"fmt"
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

func Connect(host string, port int, user string, dbname string, password string) {
	var err error
	db, err = gorm.Open("postgres", "host=" + host + " port=" + strconv.Itoa(port) + " user=" + user + " dbname=" + dbname + " password=" + password)
	db.LogMode(true)
	db.SingularTable(true)
	if err != nil || db == nil {
		panic("Connexion to Ticker database failed.")
	}
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

func CloseConnexion() {
	if db != nil {
		defer db.Close()
	}
}




