package main

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	"io/ioutil"
	"log"
	"os"
	"testgo/internal/models"
)

const (
	clusterID = "test-cluster"
	clientID  = "test-publisher"
	channel   = "test"
)

func main() {
	js, err := os.Open("model.json")
	if err != nil {
		log.Fatalln(err)
	}
	defer js.Close()

	nc, err := stan.Connect(clusterID, clientID)
	if err != nil {
		fmt.Println("", err)
	}

	jsByte, _ := ioutil.ReadAll(js)

	var data []models.Order
	err = json.Unmarshal(jsByte, &data)
	if err != nil {
		log.Fatalln(err)
	}

	for _, value := range data {
		b, _ := json.Marshal(value)

		err = nc.Publish(channel, b)
		if err != nil {
			fmt.Println(err)
		}
	}

	err = nc.Close()
	if err != nil {
		log.Println(err)
	}
}
