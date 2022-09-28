package repository

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"testgo/internal/models"
)

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r Repository) SaveOrder(order *models.Order) error {
	deliveryJson, _ := json.Marshal(order.Delivery)
	paymentJson, _ := json.Marshal(order.Payment)
	itemsJson, _ := json.Marshal(models.Item{})

	result, err := r.db.Exec("INSERT into new_order (order_uuid,"+
		"track_number,"+
		"entry,delivery,"+
		"paiment,"+
		"items,"+
		"locale,"+
		"internal_signature,"+
		"customer_id,"+
		"delivery_service,"+
		"shardkey,sm_id,date_created,oof_shard) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14)",
		order.OrderUuid,
		order.TrackNumber,
		order.Entry,
		deliveryJson,
		paymentJson,
		itemsJson,
		order.Locale,
		order.InternalSignature,
		order.CustomerId,
		order.DeliveryService,
		order.Shardkey,
		order.SmId,
		order.DateCreated,
		order.OofShard)

	if err != nil {
		return err
	}
	fmt.Print(result)
	return nil
}

func (r Repository) GetOrders() ([]models.Order, error) {
	var ordersDto []models.OrderDto
	err := r.db.Select(&ordersDto, "select * from new_order")
	if err != nil {
		return []models.Order{}, err
	}

	orders := make([]models.Order, 0, len(ordersDto))
	for _, dto := range ordersDto {
		orders = append(orders, dto.ToModel())
	}

	return orders, nil
}

func (r Repository) GetOrderByUuid(uuid uuid.UUID) (*models.Order, error) {
	var orderDto models.OrderDto

	err := r.db.Get(&orderDto, "select * from new_order where order_uuid = $1", uuid)
	if err != nil {
		return nil, err
	}

	order := orderDto.ToModel()
	return &order, nil
}
