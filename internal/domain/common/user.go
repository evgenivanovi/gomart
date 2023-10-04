package common

/* __________________________________________________ */

type UserID struct {
	id int64
}

func NewUserID(id int64) *UserID {
	return &UserID{
		id: id,
	}
}

func (u UserID) ID() int64 {
	return u.id
}

/* __________________________________________________ */
