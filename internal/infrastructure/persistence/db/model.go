package db

import (
	"gorm.io/gorm"
	"time"
)

// Model the table default columns
type Model struct {
	// ID the unique id
	ID uint64 `gorm:"column:id;primaryKey;type:BIGINT(20) UNSIGNED;"`

	// CreatedAt the create time
	CreatedAt time.Time `gorm:"column:created_at;type:DATETIME;not null;autoUpdateTime;index:idx_created_at"`

	// UpdatedAt the update time
	UpdatedAt time.Time `gorm:"column:updated_at;type:DATETIME;not null;autoUpdateTime;index:idx_created_at"`

	// DeletedAt the delete time
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:DATETIME;not null;index:idx_deleted_at"`
}
