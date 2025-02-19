package po

import "dataset/internal/infrastructure/persistence/dataset_db"

type DatasetPO struct {
	// Model default columns
	dataset_db.Model

	// Name of dataset
	Name string `gorm:"column:name;type:VARCHAR(255);not null;unique"`

	// Description of dataset
	Description string `gorm:"column:description;type:VARCHAR(255);not null"`
}

func (*DatasetPO) TableName() string {
	return "t_dataset"
}
