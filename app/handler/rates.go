package handler

import (
	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
	"net/http"
	/*"github.com/go-chi/chi/middleware"*/
	"github.com/go-chi/render"
)

// Rate defines the currency and rate type
type Rate struct {
	Currency string
	Rate     float32
}

//RateRoutes define all the routes begining with rates/
/*func RateRoutes(db *gorm.DB) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/latest", GetLatestRates)
	//r.Get("/{date:[0-9-]+}", GetRatesByDate)
	//r.Get("/analyze", AnalyzeRates)
	return r
}*/

//RateRoutes define all the routes begining with rates/
func RateRoutes(db *gorm.DB) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/latest", func(w http.ResponseWriter, r *http.Request) {
		getrates := GetLatestRates(db)
		render.JSON(w, r, getrates)
	})
	//r.Get("/{date:[0-9-]+}", GetRatesByDate)
	//r.Get("/analyze", AnalyzeRates)
	return r
}

//GetLatestRates fetches latest exchange rate from database and returns the formatted output
func GetLatestRates(db *gorm.DB) map[string]interface{}{
	//Delete old data if exists
	queryText := `SELECT currency_symbol, currency_rate FROM exchange_rates
					  where currency_date IN (SELECT max(currency_date) from exchange_rates)
					  order by currency_symbol`
	rates := make(map[string]float32)
	outputFormat := map[string]interface{}{
		"base":  "EUR",
		"rates": map[string]float32{},
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
