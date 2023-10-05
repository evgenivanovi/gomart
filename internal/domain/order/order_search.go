package order

import (
	"github.com/evgenivanovi/gomart/internal/domain/common"
	"github.com/evgenivanovi/gomart/pkg/search"
	"github.com/evgenivanovi/gomart/pkg/stdx"
)

/* __________________________________________________ */

const IDSearchKey search.Key = "id"
const UserIDSearchKey search.Key = "user_id"
const StatusSearchKey search.Key = "status"

/* __________________________________________________ */

func IdentityCondition(id OrderID) search.SearchCondition {
	return *search.NewEquality(
		IDSearchKey,
		stdx.NewValue(id.ID()),
	)
}

func IdentitiesCondition(ids []OrderID) search.SearchCondition {

	raws := make([]string, 0)
	for _, id := range ids {
		raws = append(raws, id.ID())
	}

	return *search.NewContainsAny(
		IDSearchKey,
		stdx.NewValue(raws),
	)

}

func UserIDCondition(id common.UserID) search.SearchCondition {
	return *search.NewEquality(
		UserIDSearchKey,
		stdx.NewValue(id.ID()),
	)
}

func UserIDsCondition(ids []common.UserID) search.SearchCondition {

	raws := make([]int64, 0)
	for _, id := range ids {
		raws = append(raws, id.ID())
	}

	return *search.NewContainsAny(
		UserIDSearchKey,
		stdx.NewValue(raws),
	)

}

func StatusCondition(status OrderStatus) search.SearchCondition {
	return *search.NewEquality(
		StatusSearchKey,
		stdx.NewValue(status.String()),
	)
}

func NotStatusCondition(status OrderStatus) search.SearchCondition {
	return *search.NewInequality(
		StatusSearchKey,
		stdx.NewValue(status.String()),
	)
}

func StatusesCondition(statuses []OrderStatus) search.SearchCondition {

	raws := make([]string, 0)
	for _, status := range statuses {
		raws = append(raws, status.String())
	}

	return *search.NewContainsAny(
		StatusSearchKey,
		stdx.NewValue(raws),
	)

}

func NotStatusesCondition(statuses []OrderStatus) search.SearchCondition {

	raws := make([]string, 0)
	for _, status := range statuses {
		raws = append(raws, status.String())
	}

	return *search.NewNotContainsAll(
		StatusSearchKey,
		stdx.NewValue(raws),
	)

}

/* __________________________________________________ */

const CreatedAtOrderKey search.Key = "created_at"
const UpdatedAtIDOrderKey search.Key = "updated_at"
const DeletedAtIDOrderKey search.Key = "deleted_at"

/* __________________________________________________ */
