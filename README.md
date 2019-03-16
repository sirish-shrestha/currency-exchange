# zumata-currency-exchange

## Introduction
Exchange Rate API based on European currency created in GO server
- Packages Used Chi, Gorm, pq
- Runs DB Migration(Table Schema) and Seeder(XML URL Source) during Server Initialization
- 

## Pre-requisites
go >= 1.12
postgresql >= 9.5

## How to run?
### Clone from github into your $GOPATH/src directory
```
git clone https://github.com/sirish-shrestha/zumata-currency-exchange.git
```
### Get packages for chi, gorm, pq  and testify/assert
chi - Efficient Routing with middlewares.
gorm - ORM library for GO having auto migrations support.
pq - Pure Go Postgres driver for database/sql, Used for prepared statements.
testify/assert - Testing tools for Go.

```
go get -u github.com/go-chi/chi
go get -u github.com/go-chi/render
go get -u github.com/jinzhu/gorm
go get -u github.com/lib/pq
go get -u github.com/stretchr/testify/assert
```

### SetUp Database
```
Currenty connected to the cloud PostgreSQL. If you want to connect to your own local database please follow the below steps:

1. Create am empty postgreSQL database
2. Setup the database configuration in the file config/config.go as needed:

	DB: &DBConfig{
		Dialect:  "postgres",		
		Host:     "otto.db.elephantsql.com",			//Change the DB Host as needed
		Port:     5432,									//Change the DB Port as needed.
		Username: "fzjbueum",							//Change the DB Username as needed.
		Password: "7b_4tN3JrmPNRVdN5971HOA8zALMiwRR",	//Change the DB Password as needed.
		Name:     "fzjbueum",							//Change the DB Name as needed.
	}

```

### Navigate to root project directory
```
cd /zumata-currency-exchange
```
### Build 
```
go build main.go
```
### Test
```
go test ./...
```
### Run
```
go run main.go
```