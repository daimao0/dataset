package po

import "dataset/internal/infrastructure/persistence/db"

type DatasetPO struct {
	// Model default columns
	db.Model

	// Name of dataset
	Name string `gorm:"column:name;type:VARCHAR(255);not null;unique"`

	// Description of dataset
	Description string `gorm:"column:description;type:VARCHAR(255);not null"`
}

func (*DatasetPO) TableName() string {
	return "t_dataset"
}
