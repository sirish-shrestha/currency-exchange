package app

import (
	"fmt"
	"log"
	"net/http"
	
	"github.com/go-chi/chi"
    
	"github.com/jinzhu/gorm"

	"zumata-currency-exchange/app/model"
	"zumata-currency-exchange/app/dbseeds"
	"zumata-currency-exchange/config"
)

// App has router and db instances
type App struct {
	Router *chi.Mux
	DB     *gorm.DB
}

// Initialize initializes the app with predefined configuration
func (a *App) Initialize(config *config.Config) {
	fmt.Println("Initializing Server..........")
	dbURI := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.DB.Host,
		config.DB.Port,
		config.DB.Username,
		config.DB.Password,
		config.DB.Name,
		)
	db, err := gorm.Open(config.DB.Dialect, dbURI)
	defer db.Close()
	if err != nil {
		//fmt.Println(err)
		log.Fatal("Could not connect to database:")
	}
	//Run the Migrations
	print("Running Migrations..........")
	a.DB = model.DBMigrate(db)
	println("Done.")

	//Run the Database Seeds
	fmt.Println("Running Database Seeds")
	dbseeds.ImportRates(db)
	println("Done.")

	//Define new Chi Router
	a.Router = chi.NewRouter()
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}