package main

import (
	"encoding/json"
	"fmt"
	"github.com/bxcodec/faker/v3"
	uuid2 "github.com/google/uuid"
	"github.com/nats-io/stan.go"
	uuid "github.com/satori/go.uuid"
	"log"
	"testgo/internal/models"
)

const (
	clusterID = "test-cluster"
	clientID  = "test-publisher"
	channel   = "test"
)

const fakeDataSize = 100

func main() {
	nc, err := stan.Connect(clusterID, clientID)
	if err != nil {
		log.Println(err)
	}

	var fake models.Order
	err = faker.SetRandomMapAndSliceSize(4)
	if err != nil {
		log.Fatalln(err)
	}

	for i := 0; i < fakeDataSize; i++ {
		err = faker.FakeData(&fake)
		if err != nil {
			log.Println(err)
		}

		uid := uuid.NewV4()
		fake.OrderUuid = uuid2.UUID(uid)

		b, _ := json.Marshal(fake)
		var c []models.Order
		json.Unmarshal(b, &c)
		log.Println(c)
		err = nc.Publish(channel, b)
		if err != nil {
			log.Println(err)
		}
	}

	err = nc.Close()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("end")
}
