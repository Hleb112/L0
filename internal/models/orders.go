package models

import (
	"encoding/json"
	"github.com/google/uuid"
)

type Order struct {
	OrderUuid         uuid.UUID `json:"order_uuid" db:"order_uuid"`
	TrackNumber       string    `json:"track_number" db:"track_number" `
	Entry             string    `json:"entry" db:"entry"`
	Delivery          Delivery  `json:"delivery" db:"delivery"`
	Payment           Payment   `json:"payment" db:"payment"`
	Items             []Item    `json:"items" db:"items"`
	Locale            string    `json:"locale" db:"locale"`
	InternalSignature string    `json:"internal_signature" db:"internal_signature"`
	CustomerId        string    `json:"customer_id" db:"customer_id"`
	DeliveryService   string    `json:"delivery_service" db:"delivery_service"`
	Shardkey          string    `json:"shardkey" db:"shardkey"`
	SmId              int       `json:"sm_id" db:"sm_id"`
	DateCreated       string    `json:"date_created" db:"date_created"`
	OofShard          string    `json:"oof_shard" db:"oof_shard"`
}
type Delivery struct {
	Name    string `json:"name" db:"name"`
	Phone   string `json:"phone" db:"phone"`
	Zip     string `json:"zip" db:"zip"`
	City    string `json:"city" db:"city"`
	Address string `json:"address" db:"address"`
	Region  string `json:"region" db:"region"`
	Email   string `json:"email" db:"email"`
}

type Item struct {
	ChrtId      int     `json:"chrt_id" db:"chrt_id"`
	TrackNumber string  `json:"track_number" db:"track_number"`
	Price       float64 `json:"price" db:"price"`
	Rid         string  `json:"rid" db:"rid"`
	Name        string  `json:"name" db:"name"`
	Sale        int     `json:"sale" db:"sale"`
	Size        string  `json:"size" db:"size"`
	TotalPrice  float32 `json:"total_price" db:"total_price"`
	NmId        int     `json:"nm_id" db:"nm_id"`
	Brand       string  `json:"brand" db:"brand"`
	Status      int     `json:"status" db:"status"`
}

type Payment struct {
	Transaction  string  `json:"transaction" db:"transaction"`
	RequestId    string  `json:"request_id" db:"request_id"`
	Currency     string  `json:"currency" db:"currency"`
	Provider     string  `json:"provider" db:"provider"`
	Amount       float64 `json:"amount" db:"amount"`
	PaymentDt    int     `json:"payment_dt" db:"payment_dt"`
	Bank         string  `json:"bank" db:"bank"`
	DeliveryCost float64 `json:"delivery_cost" db:"delivery_cost"`
	GoodsTotal   int     `json:"goods_total" db:"goods_total"`
	CustomFee    float64 `json:"custom_fee" db:"custom_fee"`
}

func (o Order) GetDto() OrderDto {
	orderDto := OrderDto{
		OrderUuid:         o.OrderUuid,
		TrackNumber:       o.TrackNumber,
		Entry:             o.Entry,
		Locale:            o.Locale,
		InternalSignature: o.InternalSignature,
		CustomerId:        o.CustomerId,
		DeliveryService:   o.DeliveryService,
		Shardkey:          o.Shardkey,
		SmId:              o.SmId,
		DateCreated:       o.DateCreated,
		OofShard:          o.OofShard,
	}
	marshalDelivery, _ := json.Marshal(o.Delivery)
	marshalPayment, _ := json.Marshal(o.Payment)
	marshalItems, _ := json.Marshal(o.Items)
	orderDto.Delivery = string(marshalDelivery)
	orderDto.Payment = string(marshalPayment)
	orderDto.Items = string(marshalItems)
	return orderDto
}
