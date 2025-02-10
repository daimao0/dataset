package service

import "dataset/internal/domain/dataset/model"

// FieldService is the interface of filed service
type FieldService interface {

	// GenerateFieldByMap generate field by map
	GenerateFieldByMap(map[string]any) ([]model.Field, error)
}
