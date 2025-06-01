package models

import "samplelab-go/src/dto"

type ReportData struct {
	ID int64 `gorm:"primaryKey"`

	ManufacturerName      string
	ManufacturerAddressID int64
	ManufacturerAddress   Address `gorm:"foreignKey:ManufacturerAddressID"`

	ManufacturerCountry string

	SupplierName      string
	SupplierAddressID int64
	SupplierAddress   Address `gorm:"foreignKey:SupplierAddressID"`

	SellerName      string
	SellerAddressID int64
	SellerAddress   Address `gorm:"foreignKey:SellerAddressID"`

	RecipientName      string
	RecipientAddressID int64
	RecipientAddress   Address `gorm:"foreignKey:RecipientAddressID"`

	ProductionDate       string
	BatchNumber          int
	BatchSizeProd        string
	BatchSizeStorehouse  string
	SamplePacking        string
	SampleCollectionSite string
	SampleCollector      string
	JobNumber            int
	Mechanism            string
	DeliveryMethod       string
	CollectionDate       string
	ProtocolNumber       string
	SampleID             int64
}

func ToReportDataDto(m ReportData) dto.ReportDataDto {
	return dto.ReportDataDto{
		ID:                   m.ID,
		ManufacturerName:     m.ManufacturerName,
		ManufacturerAddress:  AddressToDto(m.ManufacturerAddress),
		ManufacturerCountry:  m.ManufacturerCountry,
		SupplierName:         m.SupplierName,
		SupplierAddress:      AddressToDto(m.SupplierAddress),
		SellerName:           m.SellerName,
		SellerAddress:        AddressToDto(m.SellerAddress),
		RecipientName:        m.RecipientName,
		RecipientAddress:     AddressToDto(m.RecipientAddress),
		ProductionDate:       m.ProductionDate,
		BatchNumber:          m.BatchNumber,
		BatchSizeProd:        m.BatchSizeProd,
		BatchSizeStorehouse:  m.BatchSizeStorehouse,
		SamplePacking:        m.SamplePacking,
		SampleCollectionSite: m.SampleCollectionSite,
		SampleCollector:      m.SampleCollector,
		JobNumber:            m.JobNumber,
		Mechanism:            m.Mechanism,
		DeliveryMethod:       m.DeliveryMethod,
		CollectionDate:       m.CollectionDate,
		ProtocolNumber:       m.ProtocolNumber,
		SampleID:             m.SampleID,
	}
}

func ToReportDataModel(d dto.ReportDataDto) ReportData {
	return ReportData{
		ID:                   d.ID,
		ManufacturerName:     d.ManufacturerName,
		ManufacturerAddress:  AddressToModel(d.ManufacturerAddress),
		ManufacturerCountry:  d.ManufacturerCountry,
		SupplierName:         d.SupplierName,
		SupplierAddress:      AddressToModel(d.SupplierAddress),
		SellerName:           d.SellerName,
		SellerAddress:        AddressToModel(d.SellerAddress),
		RecipientName:        d.RecipientName,
		RecipientAddress:     AddressToModel(d.RecipientAddress),
		ProductionDate:       d.ProductionDate,
		BatchNumber:          d.BatchNumber,
		BatchSizeProd:        d.BatchSizeProd,
		BatchSizeStorehouse:  d.BatchSizeStorehouse,
		SamplePacking:        d.SamplePacking,
		SampleCollectionSite: d.SampleCollectionSite,
		SampleCollector:      d.SampleCollector,
		JobNumber:            d.JobNumber,
		Mechanism:            d.Mechanism,
		DeliveryMethod:       d.DeliveryMethod,
		CollectionDate:       d.CollectionDate,
		ProtocolNumber:       d.ProtocolNumber,
		SampleID:             d.SampleID,
	}
}
