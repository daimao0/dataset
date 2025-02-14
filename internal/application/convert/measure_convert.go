package convert

import (
	"dataset/internal/application/dto"
	"dataset/internal/domain/dataset/model"
)

func MeasureToMeasureDTO(measure *model.Measure) *dto.MeasureDTO {
	return &dto.MeasureDTO{
		Id:         measure.Id,
		Name:       measure.Name,
		Label:      measure.Label,
		DataType:   measure.DataType,
		CreateTime: measure.CreatedAt,
		UpdateTime: measure.UpdatedAt,
	}
}
