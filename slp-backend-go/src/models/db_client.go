package models

import "samplelab-go/src/dto"

type Client struct {
	ID          int64 `gorm:"primaryKey;autoIncrement"`
	WijharsCode string
	Name        string
	AddressID   int64
	Address     Address `gorm:"foreignKey:AddressID"`
}

func (Client) TableName() string {
	return "client"
}

func ClientToDto(client Client) dto.ClientDto {
	return dto.ClientDto{
		ID:          client.ID,
		WijharsCode: client.WijharsCode,
		Name:        client.Name,
		Address:     AddressToDto(client.Address),
	}
}

func ClientToModel(dto dto.ClientDto) Client {
	return Client{
		ID:          dto.ID,
		WijharsCode: dto.WijharsCode,
		Name:        dto.Name,
		Address:     AddressToModel(dto.Address),
	}
}
