package models

import "samplelab-go/src/dto"

type Assortment struct {
	ID                 int64        `gorm:"primaryKey"`                                   // bigint, primary key
	Name               string       `gorm:"type:varchar(255)"`                            // varchar(255)
	OrganolepticMethod string       `gorm:"type:varchar(255);column:organoleptic_method"` // varchar(255), z nazwą kolumny
	GroupID            int64        `gorm:"column:group_id"`                              // bigint, FK do product_group
	Group              ProductGroup `gorm:"foreignKey:GroupID"`
}

func (Assortment) TableName() string {
	return "assortment"
}

func AssortmentToDto(a Assortment) dto.AssortmentDto {
	return dto.AssortmentDto{
		ID:                 a.ID,
		Name:               a.Name,
		OrganolepticMethod: a.OrganolepticMethod,
		GroupID:            a.GroupID,
	}
}

func AssortmentFromDto(dto dto.AssortmentDto) Assortment {
	return Assortment{
		ID:                 dto.ID,
		Name:               dto.Name,
		OrganolepticMethod: dto.OrganolepticMethod,
		GroupID:            dto.GroupID,
	}
}
