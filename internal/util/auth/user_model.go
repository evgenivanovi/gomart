package auth

import "go.uber.org/zap/zapcore"

/* __________________________________________________ */

type User struct {
	UserID int64 `json:"user_id"`
}

func (u *User) MarshalLogObject(
	enc zapcore.ObjectEncoder,
) error {
	enc.AddInt64("user_id", u.UserID)
	return nil
}

func NewUser(userID int64) *User {
	return &User{
		UserID: userID,
	}
}

/* __________________________________________________ */
