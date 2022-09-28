package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/nats-io/stan.go"
	"log"
	cache "testgo/internal/cache"
	nats2 "testgo/internal/nats"
	"testgo/internal/repository"
	"testgo/internal/router"
	service2 "testgo/internal/service"
	"time"
)

func main() {
	db, err := sqlx.Connect("postgres", "user=postgres password=admin dbname=wb_new_order sslmode=disable")
	if err != nil {
		log.Fatalln("kek", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalln("kekPing", err)
	}

	defer db.Close()

	cache := cache.New(5*time.Minute, 10*time.Minute)
	repository := repository.New(db)
	service := service2.New(repository, cache)

	orders, err := repository.GetOrders()
	if err != nil {
		log.Fatal(err)
	}

	cache.InitFromDb(orders)

	nc, err := stan.Connect("test-cluster", "stan-sub")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	sub, err := nats2.Sub(nc, repository, cache)
	if err != nil {
		log.Println(err)
	}
	defer sub.Unsubscribe()

	srv := router.New(service)

	err = srv.Start()
	if err != nil {
		log.Fatalln(err)
	}

}
