package models

import (
	"encoding/json"
	"github.com/google/uuid"
)

type OrderDto struct {
	OrderUuid         uuid.UUID `json:"order_uuid" db:"order_uuid"`
	TrackNumber       string    `json:"track_number" db:"track_number" `
	Entry             string    `json:"entry" db:"entry"`
	Delivery          string    `json:"delivery" db:"delivery"`
	Payment           string    `json:"payment" db:"payment"`
	Items             string    `json:"items" db:"items"`
	Locale            string    `json:"locale" db:"locale"`
	InternalSignature string    `json:"internal_signature" db:"internal_signature"`
	CustomerId        string    `json:"customer_id" db:"customer_id"`
	DeliveryService   string    `json:"delivery_service" db:"delivery_service"`
	Shardkey          string    `json:"shardkey" db:"shardkey"`
	SmId              int       `json:"sm_id" db:"sm_id"`
	DateCreated       string    `json:"date_created" db:"date_created"`
	OofShard          string    `json:"oof_shard" db:"oof_shard"`
}

func (od OrderDto) ToModel() Order {
	var order = Order{
		OrderUuid:         od.OrderUuid,
		TrackNumber:       od.TrackNumber,
		Entry:             od.Entry,
		Locale:            od.Locale,
		InternalSignature: od.InternalSignature,
		CustomerId:        od.CustomerId,
		DeliveryService:   od.DeliveryService,
		Shardkey:          od.Shardkey,
		SmId:              od.SmId,
		DateCreated:       od.DateCreated,
		OofShard:          od.OofShard,
	}

	json.Unmarshal([]byte(od.Delivery), &order.Delivery)
	json.Unmarshal([]byte(od.Payment), &order.Payment)
	json.Unmarshal([]byte(od.Items), &order.Items)

	return order
}
