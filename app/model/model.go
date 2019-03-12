package model

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type ExchangeRates struct {
	gorm.Model
	CurrencySymbol		string `gorm:"type:varchar(3)"`
	CurrencyRate 		int
	CurrencyDate		time.Time
}

// DBMigrate will create and migrate the table
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&ExchangeRates{})
	return db
}