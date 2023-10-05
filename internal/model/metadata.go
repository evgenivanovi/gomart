package model

import "time"

/* __________________________________________________ */

type Metadata struct {
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

/* __________________________________________________ */
