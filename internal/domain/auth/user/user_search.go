package user

import (
	"github.com/evgenivanovi/gomart/internal/domain/common"
	"github.com/evgenivanovi/gomart/pkg/search"
	"github.com/evgenivanovi/gomart/pkg/stdx"
)

/* __________________________________________________ */

const IDSearchKey search.Key = "id"
const UsernameSearchKey search.Key = "username"

/* __________________________________________________ */

func IdentityCondition(id common.UserID) search.SearchCondition {
	return *search.NewEquality(
		IDSearchKey,
		stdx.NewValue(id.ID()),
	)
}

func IdentitiesCondition(ids []common.UserID) search.SearchCondition {

	raws := make([]int64, 0)
	for _, id := range ids {
		raws = append(raws, id.ID())
	}

	return *search.NewContainsAny(
		IDSearchKey,
		stdx.NewValue(raws),
	)

}

func UsernameCondition(username string) search.SearchCondition {
	return *search.NewEquality(
		UsernameSearchKey,
		stdx.NewValue(username),
	)
}

func UsernamesCondition(usernames []string) search.SearchCondition {

	raws := make([]string, 0)
	raws = append(raws, usernames...)

	return *search.NewContainsAny(
		UsernameSearchKey,
		stdx.NewValue(raws),
	)

}

/* __________________________________________________ */
