package convert

import (
	"dataset/internal/application/dto"
	"dataset/internal/domain/dataset/model"
)

func DriverDimensionToDriverDimensionDTO(driverDimension *model.DriverDimension) *dto.DriverDimensionDTO {
	return &dto.DriverDimensionDTO{
		Id:         driverDimension.Id,
		Name:       driverDimension.Name,
		Label:      driverDimension.Label,
		DataType:   driverDimension.DataType,
		CreateTime: driverDimension.CreatedAt,
		UpdateTime: driverDimension.UpdatedAt,
	}
}
