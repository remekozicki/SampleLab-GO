package models

import "samplelab-go/src/dto"

type ProductGroup struct {
	ID   int64  `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"type:varchar(255)"`
}

func (ProductGroup) TableName() string {
	return "product_group"
}

func ProductGroupToDto(pg ProductGroup) dto.ProductGroupDto {
	return dto.ProductGroupDto{
		ID:   pg.ID,
		Name: pg.Name,
	}
}

func ProductGroupFromSaveDto(d dto.ProductGroupSaveDto) ProductGroup {
	return ProductGroup{
		Name: d.Name,
	}
}
