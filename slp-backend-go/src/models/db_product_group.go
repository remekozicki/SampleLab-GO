package models

import "samplelab-go/src/dto"

type ProductGroup struct {
	ID                int64              `gorm:"primaryKey;autoIncrement"`
	Name              string             `gorm:"type:varchar(255)"`
	SamplingStandards []SamplingStandard `gorm:"many2many:product_group_sampling_standard;joinForeignKey:ProductGroupID;JoinReferences:SamplingStandardID"`
	Assortments       []Assortment       `gorm:"foreignKey:GroupID"`
}

func (ProductGroup) TableName() string {
	return "product_group"
}

func ProductGroupToDto(g ProductGroup) dto.ProductGroupDto {
	return dto.ProductGroupDto{
		ID:                g.ID,
		Name:              g.Name,
		SamplingStandards: MapSamplingStandardsToDto(g.SamplingStandards),
		Assortments:       MapAssortmentsToDto(g.Assortments),
	}
}

func ProductGroupToModel(pg dto.ProductGroupDto) ProductGroup {
	return ProductGroup{
		ID:   pg.ID,
		Name: pg.Name,
	}
}

func ProductGroupFromSaveDto(d dto.ProductGroupSaveDto) ProductGroup {
	return ProductGroup{
		Name: d.Name,
	}
}

func MapSamplingStandardsToDto(list []SamplingStandard) []dto.SamplingStandardDto {
	result := make([]dto.SamplingStandardDto, len(list))
	for i, s := range list {
		result[i] = SamplingStandardToDto(s)
	}
	return result
}

func MapAssortmentsToDto(list []Assortment) []dto.AssortmentDto {
	result := make([]dto.AssortmentDto, len(list))
	for i, a := range list {
		result[i] = AssortmentToDto(a)
	}
	return result
}
