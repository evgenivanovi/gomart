package user

import (
	"github.com/evgenivanovi/gomart/internal/domain/common"
	"github.com/evgenivanovi/gomart/internal/domain/core"
	"github.com/evgenivanovi/gomart/internal/domain/order"
	"github.com/evgenivanovi/gomart/internal/repository/postgres/public/model"
	errx "github.com/evgenivanovi/gomart/pkg/err"
	slices "github.com/evgenivanovi/gomart/pkg/std/slice"
)

/* __________________________________________________ */

func ToUserOrders(records []*model.Orders) (*order.UserOrders, error) {

	if len(records) == 0 {
		return nil, nil
	}

	singleUser := slices.IsSingle(
		slices.Dedup(
			slices.Map(
				records,
				func(t *model.Orders) int64 {
					return t.UserID
				},
			),
		),
	)

	if !singleUser {
		return nil, errx.NewErrorWithEntityCodeMessage(
			order.ErrorOrderEntity, errx.ErrorInternalCode, errx.ErrorInternalMessage,
		)
	}

	orders, err := ToOrdersValues(records)
	if err != nil {
		return nil, errx.NewErrorWithEntityCodeMessage(
			order.ErrorOrderEntity, errx.ErrorInternalCode, errx.ErrorInternalMessage,
		)
	}

	return order.NewUserOrders(
		*common.NewUserID(records[0].UserID),
		orders,
	), nil

}

func ToUserOrder(record *model.Orders) *order.UserOrder {

	id := order.NewOrderID(
		record.ID,
	)

	status, _ := order.AsOrderStatus(record.Status)
	data := order.NewOrderData(status, record.Accrual)

	metadata := core.NewMetadata(
		record.CreatedAt, record.UpdatedAt, record.DeletedAt,
	)

	return order.NewUserOrder(
		*common.NewUserID(record.UserID),
		*order.NewOrder(*id, *data, *metadata),
	)

}

func FromUserOrder(entity order.UserOrder) model.Orders {
	return model.Orders{
		// ID
		ID: entity.Order.Identity().ID(),
		// DATA
		UserID:  entity.Owner.ID(),
		Status:  entity.Order.Status().String(),
		Accrual: entity.Order.Data().Accrual,
		// METADATA
		CreatedAt: entity.Order.Metadata().CreatedAt,
		UpdatedAt: entity.Order.Metadata().UpdatedAt,
		DeletedAt: entity.Order.Metadata().DeletedAt,
	}
}

func ToUserOrderSlice(records []*model.Orders) []*order.UserOrder {
	entities := make([]*order.UserOrder, 0)
	for _, record := range records {
		entities = append(entities, ToUserOrder(record))
	}
	return entities
}

func FromUserOrdersSlice(entities []order.UserOrder) []model.Orders {
	records := make([]model.Orders, 0)
	for _, entity := range entities {
		records = append(records, FromUserOrder(entity))
	}
	return records
}

/* __________________________________________________ */

func ToOrderPointer(record *model.Orders) (*order.Order, error) {

	id := order.NewOrderID(
		record.ID,
	)

	status, _ := order.AsOrderStatus(record.Status)
	data := order.NewOrderData(status, record.Accrual)

	metadata := core.NewMetadata(
		record.CreatedAt, record.UpdatedAt, record.DeletedAt,
	)

	return order.NewOrder(*id, *data, *metadata), nil

}

func ToOrderValue(record *model.Orders) (order.Order, error) {
	pointer, err := ToOrderPointer(record)
	if err != nil {
		return order.Order{}, err
	}
	return *pointer, err
}

func ToOrdersPointers(records []*model.Orders) ([]*order.Order, error) {
	entities := make([]*order.Order, 0)
	for _, record := range records {
		entity, err := ToOrderPointer(record)
		if err != nil {
			return nil, errx.NewErrorWithEntityCode(
				order.ErrorOrderEntity, errx.ErrorInternalMessage,
			)
		}
		entities = append(entities, entity)
	}
	return entities, nil
}

func ToOrdersValues(records []*model.Orders) ([]order.Order, error) {
	entities := make([]order.Order, 0)
	for _, record := range records {
		entity, err := ToOrderValue(record)
		if err != nil {
			return nil, errx.NewErrorWithEntityCode(
				order.ErrorOrderEntity, errx.ErrorInternalMessage,
			)
		}
		entities = append(entities, entity)
	}
	return entities, nil
}

/* __________________________________________________ */
