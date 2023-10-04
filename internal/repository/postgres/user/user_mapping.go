package user

import (
	"github.com/evgenivanovi/gomart/internal/domain/auth"
	"github.com/evgenivanovi/gomart/internal/domain/auth/user"
	"github.com/evgenivanovi/gomart/internal/domain/common"
	"github.com/evgenivanovi/gomart/internal/domain/core"
	"github.com/evgenivanovi/gomart/internal/repository/postgres/public/model"
	"github.com/evgenivanovi/gomart/pkg/std"
)

/* __________________________________________________ */

func ToUser(record *model.Users) *user.User {

	id := common.NewUserID(record.ID)

	var credentials *auth.Credentials

	if record.Hashed {
		credentials = auth.NewHashedCredentials(
			record.Username, record.Password,
		)
	} else {
		credentials = auth.NewCredentials(
			record.Username, record.Password,
		)
	}

	data := user.NewUserData(*credentials)

	metadata := core.NewMetadata(
		record.CreatedAt, record.UpdatedAt, record.DeletedAt,
	)

	return user.NewUser(*id, *data, *metadata)

}

func FromUser(entity user.User) model.Users {
	return model.Users{
		// ID
		ID: entity.Identity().ID(),
		// DATA
		Username: entity.Data().Credentials.Username(),
		Password: entity.Data().Credentials.Password(),
		Hashed:   entity.Data().Credentials.Hashed(),
		// METADATA
		CreatedAt: entity.Metadata().CreatedAt,
		UpdatedAt: entity.Metadata().UpdatedAt,
		DeletedAt: entity.Metadata().DeletedAt,
	}
}

func FromUserData(data user.UserData, metadata core.Metadata) model.Users {
	return model.Users{
		// DATA
		Username: data.Credentials.Username(),
		Password: data.Credentials.Password(),
		Hashed:   data.Credentials.Hashed(),
		// METADATA
		CreatedAt: metadata.CreatedAt,
		UpdatedAt: metadata.UpdatedAt,
		DeletedAt: metadata.DeletedAt,
	}
}

func ToUsers(records []*model.Users) []*user.User {
	entities := make([]*user.User, 0)
	for _, record := range records {
		entities = append(entities, ToUser(record))
	}
	return entities
}

func FromUsers(entities []user.User) []model.Users {
	records := make([]model.Users, 0)
	for _, entity := range entities {
		records = append(records, FromUser(entity))
	}
	return records
}

func FromUsersData(pairs []std.Pair[user.UserData, core.Metadata]) []model.Users {
	records := make([]model.Users, 0)
	for _, entity := range pairs {
		records = append(records, FromUserData(entity.First(), entity.Second()))
	}
	return records
}

/* __________________________________________________ */
