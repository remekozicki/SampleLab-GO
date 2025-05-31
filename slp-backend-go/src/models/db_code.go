package models

import "samplelab-go/src/dto"

type Code struct {
	ID   string `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}

func (Code) TableName() string {
	return "code"
}

func CodeToDto(code Code) dto.CodeDto {
	return dto.CodeDto{
		ID:   code.ID,
		Name: code.Name,
	}
}

func CodeToModel(dto dto.CodeDto) Code {
	return Code{
		ID:   dto.ID,
		Name: dto.Name,
	}
}
