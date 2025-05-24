package salesFileInsert

import (
	DBConn "Sales/DBConnection"
	common "Sales/common/toml"
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"gorm.io/gorm"
)

func ReadFile_updatedata() error {
	log.Println("ReadFile_updatedata (+) ")

	lpathConfig := common.ReadTomlFile("./toml/serviceconfig.toml")
	lFilepath := fmt.Sprintf("%v", lpathConfig.(map[string]interface{})["FilePath"])

	lErr := CsvFile_Reader(lFilepath)
	if lErr != nil {
		log.Println("Error : SRF01 ", lErr)
		return lErr
	}

	log.Println("ReadFile_updatedata (-) ")
	return nil

}

func CsvFile_Reader(pFilepath string) error {
	log.Println("CsvFile_Reader(+)")
	//Opening the file
	lFile, lErr := os.Open(pFilepath)
	if lErr != nil {
		log.Println("Error :SCF01 ", lErr)
		return lErr
	}

	//Closing the file
	defer lFile.Close()

	// Create a new CSV reader
	lReader := csv.NewReader(lFile)

	//Reading all the rows in the file
	lRecords, lErr := lReader.ReadAll()
	if lErr != nil {
		log.Println("Error :SCF02 ", lErr)
		return lErr
	}

	for lIdx, lRows := range lRecords {

		if lIdx == 0 {
			continue
		}

		var lCustomerRec Customer
		var lProductsRec Product
		var lOrdersRec Order
		var lOrder_DetailRec OrderDetail

		lCustomerRec.CustomerID = lRows[2]
		lCustomerRec.Name = lRows[12]
		lCustomerRec.Email = lRows[13]
		lCustomerRec.Address = lRows[14]

		lErr := CreateCustomerIfExist(lCustomerRec)
		if lErr != nil {
			log.Println("Error :SCF03  ", lErr)
			return lErr
		}

		lProductsRec.ProductID = lRows[1]
		lProductsRec.ProductName = lRows[3]
		lProductsRec.Category = lRows[4]

		lErr = CreateProductIfExist(lProductsRec)
		if lErr != nil {
			log.Println("Error :SCF04  ", lErr)
			return lErr
		}

		Date, lErr := time.Parse("2006-01-02", lRows[6])
		if lErr != nil {
			log.Println("Error :SCF05  ", lErr)
			return lErr
		}

		ShippingCost, lErr := strconv.ParseFloat(lRows[10], 64)
		if lErr != nil {
			log.Println("Error :SCF06  ", lErr)
			return lErr
		}
		lOrdersRec.OrderID = lRows[0]
		lOrdersRec.CustomerID = lRows[2]
		lOrdersRec.Region = lRows[5]
		lOrdersRec.DateOfSale = Date
		lOrdersRec.PaymentMethod = lRows[11]
		lOrdersRec.ShippingCost = ShippingCost

		lExistID, lErr := CreateOrderIfExist(lOrdersRec)
		if lErr != nil {
			log.Println("Error :SCF07 ", lErr)
			return lErr
		}

		if lExistID != 0 {
			continue
		}

		lOrder_DetailRec.OrderID = lOrdersRec.OrderID
		lOrder_DetailRec.ProductID = lRows[1]
		lOrder_DetailRec.QuantitySold, lErr = strconv.Atoi(lRows[7])
		if lErr != nil {
			log.Println("Error :SCF08 ", lErr)
			return lErr
		}
		lOrder_DetailRec.UnitPrice, lErr = strconv.ParseFloat(lRows[8], 64)
		if lErr != nil {
			log.Println("Error :SCF09 ", lErr)
			return lErr
		}
		lOrder_DetailRec.Discount, lErr = strconv.ParseFloat(lRows[9], 64)
		if lErr != nil {
			log.Println("Error :SCF10 ", lErr)
			return lErr
		}

		lErr = DBConn.GConnInst.Conn1_gorm.Table("orderDetails").Create(&lOrder_DetailRec).Error
		if lErr != nil {
			log.Println("Error :SCF11 ", lErr)
			return lErr
		}
	}
	log.Println("CsvFile_Reader(-)")

	return nil
}

func CreateCustomerIfExist(pCustRec Customer) error {
	log.Println("CreateCustomerIfExist(+)")
	var existingID uint

	lErr := DBConn.GConnInst.Conn1_gorm.
		Table("customers").
		Select("id").
		Where("CustomerID = ?", pCustRec.CustomerID).
		Take(&existingID).Error

	if lErr == nil {
		// Found existing customer
		return nil
	}

	if errors.Is(lErr, gorm.ErrRecordNotFound) {
		// Not found, insert new customer
		if lErr := DBConn.GConnInst.Conn1_gorm.Create(&pCustRec).Error; lErr != nil {
			log.Println("Error : SCC01")
			return lErr
		}
		return nil
	}
	log.Println("CreateCustomerIfExist(+)")

	return lErr
}

func CreateProductIfExist(pProductRec Product) error {
	log.Println("CreateProductIfExist(+)")
	var existingID uint

	lErr := DBConn.GConnInst.Conn1_gorm.
		Table("products").
		Select("id").
		Where("ProductID = ?", pProductRec.ProductID).
		Take(&existingID).Error

	if lErr == nil {
		// Found existing customer
		return nil
	}

	if errors.Is(lErr, gorm.ErrRecordNotFound) {
		// Not found, insert new customer
		if lErr := DBConn.GConnInst.Conn1_gorm.Create(&pProductRec).Error; lErr != nil {
			log.Println("Error : SCP01")
			return lErr
		}
		return nil
	}

	log.Println("CreateProductIfExist(-)")
	return lErr
}

func CreateOrderIfExist(pOrderRec Order) (uint, error) {
	log.Println("CreateOrderIfExist(+)")
	var existingID uint

	lErr := DBConn.GConnInst.Conn1_gorm.
		Table("orders").
		Select("id").
		Where("OrderID = ?", pOrderRec.OrderID).
		Take(&existingID).Error

	if lErr == nil {
		// Found existing customer
		return existingID, nil
	}

	if errors.Is(lErr, gorm.ErrRecordNotFound) {
		// Not found, insert new customer
		if lErr := DBConn.GConnInst.Conn1_gorm.Create(&pOrderRec).Error; lErr != nil {
			return existingID, lErr
			log.Println("Error : SCO01")
		}
		return existingID, nil
	}

	log.Println("CreateOrderIfExist(-)")
	return existingID, lErr
}
