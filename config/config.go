//Package config allows to set/get the database config
package config

//Config holds the DBConfig
type Config struct {
	DB *DBConfig
}

//DBConfig holds the database configuration settings
type DBConfig struct {
	Dialect  string
	Host  string
	Port  int
	Username string
	Password string
	Name     string
}

//GetConfig returns the database configurations.
//Edit the Database configuration here as needed.
func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "postgres",		
			Host:     "otto.db.elephantsql.com",			//Change the DB Host as needed
			Port:     5432,									//Change the DB Port as needed.
			Username: "fzjbueum",							//Change the DB Username as needed.
			Password: "7b_4tN3JrmPNRVdN5971HOA8zALMiwRR",	//Change the DB Password as needed.
			Name:     "fzjbueum",							//Change the DB Name as needed.
		},
	}
}