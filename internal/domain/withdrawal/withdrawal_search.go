package withdrawal

import (
	"github.com/evgenivanovi/gomart/internal/domain/common"
	"github.com/evgenivanovi/gomart/pkg/search"
	"github.com/evgenivanovi/gomart/pkg/stdx"
)

/* __________________________________________________ */

const IDSearchKey search.Key = "id"
const UserIDSearchKey search.Key = "user_id"

/* __________________________________________________ */

func IdentityCondition(id WithdrawalID) search.SearchCondition {
	return *search.NewEquality(
		IDSearchKey,
		stdx.NewValue(id.ID()),
	)
}

func IdentitiesCondition(ids []WithdrawalID) search.SearchCondition {

	raws := make([]int64, 0)
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

/* __________________________________________________ */
