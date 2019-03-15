package model

import (
	

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//ExchangeRatesTbl define the database fields for the exchange_rates table
type ExchangeRates struct {
	CurrencyRateID		uint64 `gorm:"primary_key"`
	CurrencySymbol		string `gorm:"type:varchar(3)"`
	CurrencyRate 		int  `sql:"type:decimal(10,5);"`
	CurrencyDate		string `sql:"type:date"`
}

// DBMigrate will create and migrate the table
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&ExchangeRates{})
	return db
}