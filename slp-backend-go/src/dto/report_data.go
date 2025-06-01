package dto

type ReportDataDto struct {
	ID                   int64      `json:"id"`
	ManufacturerName     string     `json:"manufacturerName"`
	ManufacturerAddress  AddressDto `json:"manufacturerAddress"`
	ManufacturerCountry  string     `json:"manufacturerCountry"`
	SupplierName         string     `json:"supplierName"`
	SupplierAddress      AddressDto `json:"supplierAddress"`
	SellerName           string     `json:"sellerName"`
	SellerAddress        AddressDto `json:"sellerAddress"`
	RecipientName        string     `json:"recipientName"`
	RecipientAddress     AddressDto `json:"recipientAddress"`
	ProductionDate       string     `json:"productionDate"`
	BatchNumber          int        `json:"batchNumber"`
	BatchSizeProd        string     `json:"batchSizeProd"`
	BatchSizeStorehouse  string     `json:"batchSizeStorehouse"`
	SamplePacking        string     `json:"samplePacking"`
	SampleCollectionSite string     `json:"sampleCollectionSite"`
	SampleCollector      string     `json:"sampleCollector"`
	JobNumber            int        `json:"jobNumber"`
	Mechanism            string     `json:"mechanism"`
	DeliveryMethod       string     `json:"deliveryMethod"`
	CollectionDate       string     `json:"collectionDate"`
	ProtocolNumber       string     `json:"protocolNumber"`
	SampleID             int64      `json:"sampleId"`
}
