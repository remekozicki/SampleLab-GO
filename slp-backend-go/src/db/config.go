package db

//
//import (
//	"log"
//	"time"
//
//	"gorm.io/gorm"
//	"samplelab-go/src/models"
//)
//
//func InitializeData(db *gorm.DB) {
//	var sampleCount int64
//	if err := db.Model(&models.Sample{}).Count(&sampleCount).Error; err != nil {
//		log.Fatalf("Failed to count samples: %v", err)
//	}
//
//	if sampleCount > 0 {
//		log.Println("Sample data already initialized. Skipping.")
//		return
//	}
//
//	db.Transaction(func(tx *gorm.DB) error {
//		// Add Codes
//		codes := []models.Code{
//			{ID: "Kp", Name: "Kontrolne"},
//			{ID: "Kd", Name: "Kontrolne"},
//			{ID: "O", Name: "Ocena Jakości Handlowej"},
//			{ID: "KW", Name: "Kontrola wewnętrzna"},
//			{ID: "IH", Name: "Inspekcja Handlowa"},
//		}
//		tx.Create(&codes)
//
//		// Add Clients
//		clients := []models.Client{
//			{
//				Name:        "WIJHARS Kielce",
//				WijharsCode: "ki",
//				Address: models.Address{
//					Street:  "Aleja IX Wieków Kielc 3",
//					ZipCode: "25-516",
//					City:    "Kielce",
//				},
//			},
//			{
//				Name:        "WIJHARS Kraków",
//				WijharsCode: "kr",
//				Address: models.Address{
//					Street:  "Ujastek 7",
//					ZipCode: "31-752",
//					City:    "Kraków",
//				},
//			},
//		}
//		tx.Create(&clients)
//
//		// Add Inspections
//		inspections := []models.Inspection{
//			{Name: "Kontrola planowa"},
//			{Name: "Kontrola dorażna GI"},
//			{Name: "Kontrola doraźna WIJHARS"},
//			{Name: "W ramach skargi"},
//			{Name: "RASFF"},
//			{Name: "Kontrola graniczna"},
//		}
//		tx.Create(&inspections)
//
//		// Add Sampling Standards
//		standards := []models.SamplingStandard{
//			{Name: "PN-72/A-74001"},
//			{Name: "PN-EN ISO 542"},
//			{Name: "PN-EN ISO 24333:2012"},
//		}
//		tx.Create(&standards)
//
//		// Add Product Groups and Assortments (simplified)
//		group := models.ProductGroup{Name: "Przetwory zbożowe"}
//		tx.Create(&group)
//
//		assortment := models.Assortment{
//			Name:  "Kasza",
//			Group: group,
//		}
//		tx.Create(&assortment)
//
//		// Add Sample
//		address := models.Address{
//			Street:  "Kawiory 21",
//			ZipCode: "30-055",
//			City:    "Kraków",
//		}
//		tx.Create(&address)
//
//		report := models.ReportData{
//			ManufacturerName:     "manufacturer1",
//			ManufacturerAddress:  address,
//			SupplierName:         "supplier1",
//			SupplierAddress:      address,
//			SellerName:           "seller1",
//			SellerAddress:        address,
//			RecipientName:        "recipient1",
//			RecipientAddress:     address,
//			JobNumber:            11,
//			Mechanism:            "mechanism1",
//			DeliveryMethod:       "deliveryMethod1",
//			CollectionDate:       time.Now().Format("2006-01-02"),
//			ProductionDate:       time.Now().Format("2006-01-02"),
//			SamplePacking:        "opakowanie",
//			BatchNumber:          1,
//			BatchSizeProd:        "1kg",
//			BatchSizeStorehouse:  "1kg",
//			SampleCollectionSite: "site",
//			SampleCollector:      "collector",
//			ProtocolNumber:       "123",
//		}
//		tx.Create(&report)
//
//		sample := models.Sample{
//			CodeID:             "Kp",
//			ClientID:           clients[0].ID,
//			AssortmentID:       assortment.ID,
//			InspectionID:       inspections[0].ID,
//			SamplingStandardID: standards[0].ID,
//			AdmissionDate:      time.Now().Format("2006-01-02"),
//			ExpirationDate:     time.Now().AddDate(0, 6, 0).Format("2006-01-02"),
//			Size:               "500g",
//			State:              "OK",
//			Analysis:           true,
//			ReportDataID:       report.ID,
//			ProgressStatus:     "DONE",
//		}
//		tx.Create(&sample)
//
//		log.Println("Sample data initialized.")
//		return nil
//	})
//}
