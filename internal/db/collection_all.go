package db

import "time"

type CollectionAll struct {
	Address     string    `bson:"address"`
	AmountETH   string    `bson:"amountETH"`
	AmountToken string    `bson:"amountToken"`
	Time        time.Time `bson:"time"`
}
