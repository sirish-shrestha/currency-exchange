package handler

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
	"github.com/go-chi/render"
)

// Rate defines the currency and rate type
type Rate struct {
	Currency string
	Rate     float32
}

//AggregateRate defines the  min, max and avg fields
//Used for the /analyze endpoint
type AggregateRate struct{
	Min float32
	Max float32
	Avg float32
}

//RateRoutes define all the routes begining with rates/
func RateRoutes(db *gorm.DB) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/latest", func(w http.ResponseWriter, r *http.Request) {
		getrates := GetLatestRates(db)
		render.JSON(w, r, getrates)
	})
	r.Get("/{date:\\d\\d\\d\\d-\\d\\d-\\d\\d}", func(w http.ResponseWriter, r *http.Request) {
		dates := chi.URLParam(r, "date")
		getrates := GetRatesByDate(db, dates)
		render.JSON(w, r, getrates)
	})
	r.Get("/analyze", func(w http.ResponseWriter, r *http.Request) {
		getrates := AnalyzeRates(db)
		render.JSON(w, r, getrates)
	})
	return r
}

//GetLatestRates fetches latest exchange rate from database and returns the formatted output
func GetLatestRates(db *gorm.DB) map[string]interface{}{
	queryText := `SELECT currency_symbol, currency_rate FROM exchange_rates
					  where currency_date IN (SELECT max(currency_date) from exchange_rates)
					  order by currency_symbol`
	rates := make(map[string]float32)
	outputFormat := map[string]interface{}{
		"base":"EUR",
		"rates":map[string]float32{},
	}

	rows, err := db.Raw(queryText).Rows()
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		rate := Rate{}
		err := rows.Scan(&rate.Currency, &rate.Rate)
		if err != nil {
			panic(err)
		}
		rates[rate.Currency] = rate.Rate
	}
	outputFormat["rates"] = rates
	return outputFormat
}


//GetRatesByDate fetches the date specific exchange rate from database and returns the formatted output
func GetRatesByDate(db *gorm.DB, date string) map[string]interface{}{
	queryText := fmt.Sprintf(`SELECT currency_symbol, currency_rate from exchange_rates 
							where currency_date = '%s' order by currency_symbol`,date);
	rates := make(map[string]float32)
	outputFormat := map[string]interface{}{
		"base":"EUR",
		"rates":map[string]float32{},
	}

	//datas := make(map[int]Rates)
	rows, err := db.Raw(queryText).Rows()
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		rate := Rate{}
		err := rows.Scan(&rate.Currency, &rate.Rate)
		if err != nil {
			panic(err)
		}
		rates[rate.Currency] = rate.Rate
	}
	outputFormat["rates"] = rates
	return outputFormat
}

//AnalyzeRates fetches the aggregate data from database and returns the formatted output
func AnalyzeRates(db *gorm.DB) map[string]interface{}{
	queryText := `SELECT currency_symbol, min(currency_rate) min, max(currency_rate) max, 
					TO_CHAR(AVG(currency_rate),'FM9999999999.9999999999999999') avg
					from exchange_rates group by currency_symbol order by currency_symbol`
	var currency string
	rates := make(map[string]AggregateRate)

	outputFormat := map[string]interface{}{
		"base":"EUR",
		"rates_analyze":map[string]AggregateRate{},
	}

	//datas := make(map[int]Rates)
	rows, err := db.Raw(queryText).Rows()
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		aggregateRate := AggregateRate{}
		err := rows.Scan(&currency, &aggregateRate.Min, &aggregateRate.Max, &aggregateRate.Avg)
		if err != nil {
			panic(err)
		}
		rates[currency] = aggregateRate
		outputFormat["rates_analyze"] = rates
	}
	return outputFormat
}