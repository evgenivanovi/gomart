package core

import "time"

/* __________________________________________________ */

type Metadata struct {
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

func (m Metadata) Equal(other any) bool {

	if another, ok := other.(Metadata); ok {

		isCreatedAtEquals := m.CreatedAt.Equal(another.CreatedAt)

		if !isCreatedAtEquals {
			return false
		}

		isUpdatedEquals := (m.UpdatedAt == nil && another.UpdatedAt == nil) ||
			(m.UpdatedAt != nil && another.UpdatedAt != nil &&
				m.UpdatedAt.Equal(*another.UpdatedAt))

		if !isUpdatedEquals {
			return false
		}

		isDeletedAtEquals := (m.DeletedAt == nil && another.DeletedAt == nil) ||
			(m.DeletedAt != nil && another.DeletedAt != nil &&
				m.DeletedAt.Equal(*another.DeletedAt))

		return isDeletedAtEquals

	}

	return false

}

func NewMetadata(
	createdAt time.Time,
	updatedAt *time.Time,
	deletedAt *time.Time,
) *Metadata {
	return &Metadata{
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		DeletedAt: deletedAt,
	}
}

func NewNowMetadata() *Metadata {
	return &Metadata{
		CreatedAt: time.Now(),
		UpdatedAt: nil,
		DeletedAt: nil,
	}
}

/* __________________________________________________ */
