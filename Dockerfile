#Define base image from docker hub
FROM golang:1.11.2-stretch

#Install git on this light weight alpine linux
##RUN apk add --no-cache git

#Define the working directory inside the docker container
WORKDIR /go/src/zumata-currency-exchange

#COPY all codes from the host machine to the working directory. This is useful for production deployment. 
#For local, volume mapping is done in docker-compose.yml
COPY . .

# Download dependencies

#chi - Efficient Routing with middlewares.
#gorm - ORM library for GO having auto migrations support.
#pq - Pure Go Postgres driver for database/sql, Used for prepared statements.
#reflex - Hot Reloading

RUN go get "github.com/jinzhu/gorm"
RUN go get "github.com/go-chi/chi"
RUN go get "github.com/go-chi/render"
RUN go get "github.com/go-chi/chi/middleware"
RUN go get "github.com/lib/pq"
RUN go get "github.com/cespare/reflex"
#RUN go build .

COPY reflex.conf /
ENTRYPOINT ["reflex", "-c", "/reflex.conf"]

##CMD go get github.com/pilu/fresh && \	fresh; \