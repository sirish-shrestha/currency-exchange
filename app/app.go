package app

import (
	"fmt"
	"log"
	"net/http"
	
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
    "github.com/jinzhu/gorm"

	"zumata-currency-exchange/app/model"
	"zumata-currency-exchange/app/dbseeds"
	"zumata-currency-exchange/app/handler"
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
	//defer db.Close()
	if err != nil {
		//fmt.Println(err)
		log.Fatal("Could not connect to database:")
	}
	//Run the Migrations
	print("Running Migrations..........")
	a.DB = model.DBMigrate(db)
	
	//Run the Database Seeds
	fmt.Println("Running Database Seeds")
	dbseeds.ImportRates(db)
	
	//Define new Chi Router
	a.Router = chi.NewRouter()
	a.setRouters()
}

// setRouters sets all required routers
func (a *App) setRouters() {

	a.Router.Use(
		render.SetContentType(render.ContentTypeJSON),	// Set content-type headers as application/json
		middleware.Logger,								// Log API request calls
		middleware.Recoverer,							// Recover from panics without crashing server
		middleware.DefaultCompress,						// Compress results, mostly gzipping assets and json
		middleware.RedirectSlashes,						// Redirect slashes to no slash URL versions
	)
	
	a.Router.Route("/api/v1", func (r chi.Router){		//Add API Versioning
		r.Mount("/rates", handler.RateRoutes(a.DB))		//Mount /rates paths to RateRoutes handler
	})	
}

// Run the app on it's router
func (a *App) Run(host string) {
	print("Server Listening at port :3000")
	log.Fatal(http.ListenAndServe(host, a.Router))
}