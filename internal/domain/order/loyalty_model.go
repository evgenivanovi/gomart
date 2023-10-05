package order

/* __________________________________________________ */

type OrderLoyalty struct {
	id   OrderID
	data OrderData
}

func NewOrderLoyalty(
	id OrderID,
	data OrderData,
) *OrderLoyalty {
	return &OrderLoyalty{
		id:   id,
		data: data,
	}
}

func (e *OrderLoyalty) Identity() OrderID {
	return e.id
}

func (e *OrderLoyalty) Data() OrderData {
	return e.data
}

/* __________________________________________________ */
