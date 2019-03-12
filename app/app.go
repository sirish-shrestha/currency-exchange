package app

import (
	"fmt"
	"github.com/go-chi/chi"
    "github.com/go-chi/chi/middleware"
	"github.com/jinzhu/gorm"

	"github.com/sirish-shrestha/zumata-currency-exchange/app/model"
)

// App has router and db instances
type App struct {
	Router *chi.Mux
	DB     *gorm.DB
}


// Initialize initializes the app with predefined configuration
func (a *App) Initialize(config *config.Config) {
	dbURI := fmt.Sprintf("host=%s port=%i user=%s dbname=%s password=%s",
		config.DB.Host,
		config.DB.Port,
		config.DB.Username,
		config.DB.Name,
		config.DB.Password)

	db, err := gorm.Open(config.DB.Dialect, dbURI)
	if err != nil {
		log.Fatal("Could not connect to database")
	}

	a.DB = model.DBMigrate(db)
	a.Router = chi.NewRouter()
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}