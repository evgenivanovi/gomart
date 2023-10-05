package order

import (
	"github.com/evgenivanovi/gomart/internal/domain/core"
	errx "github.com/evgenivanovi/gomart/pkg/err"
)

/* __________________________________________________ */

type OrderID struct {
	id string
}

func NewOrderID(id string) *OrderID {
	return &OrderID{
		id: id,
	}
}

func (u OrderID) ID() string {
	return u.id
}

/* __________________________________________________ */

type OrderStatus struct {
	value string
}

func (os OrderStatus) String() string {
	return os.value
}

func AsOrderStatus(value string) (OrderStatus, error) {

	switch value {
	case New.value:
		return New, nil
	case Invalid.value:
		return Invalid, nil
	case Processed.value:
		return Processed, nil
	case Processing.value:
		return Processing, nil
	case Registered.value:
		return Registered, nil
	}

	return Invalid, errx.NewErrorWithEntityCodeMessage(
		ErrorOrderEntity, errx.ErrorInternalCode, "unknown order status",
	)

}

const (
	NewOrderStatus        = "NEW"
	RegisteredOrderStatus = "REGISTERED"
	ProcessingOrderStatus = "PROCESSING"
	ProcessedOrderStatus  = "PROCESSED"
	InvalidOrderStatus    = "INVALID"
)

var (
	New        = OrderStatus{NewOrderStatus}
	Registered = OrderStatus{RegisteredOrderStatus}
	Processing = OrderStatus{ProcessingOrderStatus}
	Processed  = OrderStatus{ProcessedOrderStatus}
	Invalid    = OrderStatus{InvalidOrderStatus}
)

/* __________________________________________________ */

type OrderData struct {
	Status  OrderStatus
	Accrual float64
}

func NewOrderData(status OrderStatus, accrual float64) *OrderData {
	return &OrderData{
		Status:  status,
		Accrual: accrual,
	}
}

func AsNewOrderData() *OrderData {
	return NewOrderData(New, 0)
}

/* __________________________________________________ */

type Order struct {
	id       OrderID
	data     OrderData
	metadata core.Metadata
}

func NewOrder(id OrderID, data OrderData, metadata core.Metadata) *Order {
	return &Order{
		id:       id,
		data:     data,
		metadata: metadata,
	}
}

func AsNewOrder(id OrderID) *Order {
	return NewOrder(id, *AsNewOrderData(), *core.NewNowMetadata())
}

func NewEmptyOrder() Order {
	return Order{}
}

func NewEmptyOrders() []Order {
	return nil
}

func NewEmptyPointerOrder() *Order {
	return nil
}

func NewEmptyPointerOrders() []*Order {
	return nil
}

func ToOrderPointers(entities []Order) []*Order {
	result := make([]*Order, 0)
	for _, entity := range entities {
		result = append(result, &entity)
	}
	return result
}

func ToOrderValues(entities []*Order) []Order {
	result := make([]Order, 0)
	for _, entity := range entities {
		result = append(result, *entity)
	}
	return result
}

func (e *Order) Identity() OrderID {
	return e.id
}

func (e *Order) Data() OrderData {
	return e.data
}

func (e *Order) Metadata() core.Metadata {
	return e.metadata
}

func (e *Order) WithIdentity(id OrderID) *Order {
	return NewOrder(id, e.Data(), e.Metadata())
}

func (e *Order) WithData(data OrderData) *Order {
	return NewOrder(e.Identity(), data, e.Metadata())
}

func (e *Order) WithMetadata(metadata core.Metadata) *Order {
	return NewOrder(e.Identity(), e.Data(), metadata)
}

func (e *Order) Status() OrderStatus {
	return e.data.Status
}

func (e *Order) ToNew() *Order {
	data := e.Data()
	data.Status = New
	data.Accrual = 0
	return e.WithData(data)
}

func (e *Order) IsNew() bool {
	return e.data.Status.String() == New.String()
}

func (e *Order) ToRegistered() *Order {
	data := e.Data()
	data.Status = Registered
	data.Accrual = 0
	return e.WithData(data)
}

func (e *Order) IsRegistered() bool {
	return e.data.Status.String() == Registered.String()
}

func (e *Order) ToProcessing() *Order {
	data := e.Data()
	data.Status = Processing
	return e.WithData(data)
}

func (e *Order) IsProcessing() bool {
	return e.data.Status.String() == Processing.String()
}

func (e *Order) ToProcessed(accrual float64) *Order {
	data := e.Data()
	data.Status = Registered
	data.Accrual = accrual
	return e.WithData(data)
}

func (e *Order) IsProcessed() bool {
	return e.data.Status.String() == Processed.String()
}

func (e *Order) ToInvalid() *Order {
	data := e.Data()
	data.Status = Invalid
	return e.WithData(data)
}

func (e *Order) IsInvalid() bool {
	return e.data.Status.String() == Invalid.String()
}

/* __________________________________________________ */
