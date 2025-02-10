package convert

import (
	"dataset/internal/application/dto"
	"dataset/internal/domain/dataset/model"
)

func DimensionToDimensionDTO(dimension *model.Dimension) *dto.DimensionDTO {
	return &dto.DimensionDTO{
		Id:         dimension.Id,
		Name:       dimension.Name,
		Label:      dimension.Label,
		DataType:   dimension.DataType,
		CreateTime: dimension.CreatedAt,
		UpdateTime: dimension.UpdatedAt,
	}
}
