package db

import "time"

type Collection struct {
	Address string    `bson:"address"`
	Amount  string    `bson:"amount"`
	Time    time.Time `bson:"time"`
}
