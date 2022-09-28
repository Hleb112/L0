package nats

import (
	"encoding/json"
	"github.com/nats-io/stan.go"
	"log"
	"testgo/internal/cache"
	"testgo/internal/models"
	"testgo/internal/repository"
)

func HandleNewOrder(order *models.Order, repo *repository.Repository, cache *cache.Cache) error {
	err := repo.SaveOrder(order)
	if err != nil {
		log.Println("In HandleNewOrder: ", err)
		return err
	}

	cache.Set(order.OrderUuid, order, 5)

	return nil
}

func Sub(conn stan.Conn, repo *repository.Repository, cache *cache.Cache) (stan.Subscription, error) {
	handler := func(msg *stan.Msg) {
		var data models.Order

		err := json.Unmarshal(msg.Data, &data)
		if err != nil {
			log.Printf("error while decoding data from nats-pub: %v ", err)

			msg.Ack()
			if err != nil {
				log.Println(err)
				return
			}
			return
		}

		err = HandleNewOrder(&data, repo, cache)
		if err != nil {
			log.Printf("error while inserting(to db) data from nats-pub: %v ", err)

			msg.Ack()
			if err != nil {
				log.Println(err)
				return
			}

			return
		}

		if err := msg.Ack(); err != nil {
			log.Printf("failed ACK msg: %d", msg.Sequence)
			return
		}
	}

	sub, err := conn.Subscribe(
		"test",
		handler,
		stan.DurableName("durable-name"),
		stan.SetManualAckMode(),
	)
	if err != nil {
		return nil, err
	}

	return sub, nil
}
