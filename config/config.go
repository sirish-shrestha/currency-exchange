//Package config allows to set/get the database config
package config

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Dialect  string
	Host  string
	Port  int
	Username string
	Password string
	Name     string
}


func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "postgres",
			Host:     "otto.db.elephantsql.com",
			Port:     5432,	
			Username: "fzjbueum",
			Password: "7b_4tN3JrmPNRVdN5971HOA8zALMiwRR",
			Name:     "fzjbueum",
		},
	}
}