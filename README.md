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

### SetUp Database
```
Currenty connected to the 'postgres_db' docker container.

- Setup the database configuration in the file config/config.go (only if needed):

	DB: &DBConfig{
		Dialect:  "postgres",		
		Host:     "postgres_db",						//Change the DB Host as needed
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

### Run
```
docker-compose up
```

## API Endpoints
```
1. http://localhost:3000/api/v1/rates/latest

2. http://localhost:3000/api/v1/rates/{YYYY-MM-DD}

3. http://localhost:3000/api/v1/rates/analyze

4. http://localhost:3000/api/v1/rates/USDSGD
- Returns exchange rate of SGD based on USD as base rate.
```