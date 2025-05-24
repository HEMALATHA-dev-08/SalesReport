package fetchSalesData

import (
	DBConn "Sales/DBConnection"
	salesFileInsert "Sales/Sales_File_Insert"
	"Sales/common"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func GetRevenue(w http.ResponseWriter, r *http.Request) {

	log.Println("GetRevenue (+) ")

	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, credentials")

	if r.Method == http.MethodPost {

		var linputRec salesFileInsert.ReqStruct
		var lGetRevenueDetailsRec salesFileInsert.RevenueStruct
		lGetRevenueDetailsRec.Status = common.SuccessCode

		lErr := json.NewDecoder(r.Body).Decode(&linputRec)
		if lErr != nil {
			log.Println("Error FGR01")
			return
		}

		if linputRec.FromDate != "" && linputRec.ToDate != "" {
			switch linputRec.Indicator {
			case "Total":
				lGetRevenueDetailsRec.Total_revenue, lErr = TotalRevenue(linputRec)

			case "Prod":
				lGetRevenueDetailsRec.TotProdRevenue, lErr = ProdWiseRevenue(linputRec)

			case "Cat":
				lGetRevenueDetailsRec.TotalcatRevenue, lErr = CatWiseData(linputRec)

			case "Reg":
				lGetRevenueDetailsRec.TotalRevenue_byreg, lErr = RegWiseData(linputRec)
			}

			if lErr != nil {
				log.Println("Error : FGR01 ", lErr.Error())
				return
			}

		} else {
			lGetRevenueDetailsRec.Status = "E"
			lGetRevenueDetailsRec.ErrMsg = "Missing Date: From And To Dates are Mandatory"
		}

		lData, lErr := json.Marshal(lGetRevenueDetailsRec)
		if lErr != nil {
			log.Println("Error : FJSON ", lErr.Error())
			return
		}
		fmt.Fprint(w, string(lData))

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	log.Println("GetRevenue (-) ")
}

func TotalRevenue(pSalesData salesFileInsert.ReqStruct) (string, error) {
	log.Println("TotalRevenue(+)")
	var totalRevenue string

	lErr := DBConn.GConnInst.Conn1_gorm.Table("OrderDetails").
		Joins("JOIN Orders ON Orders.OrderID = OrderDetails.OrderID").
		Where("Orders.DateOfSale BETWEEN ? AND ?", pSalesData.FromDate, pSalesData.ToDate).
		Select("SUM(OrderDetails.QuantitySold * OrderDetails.UnitPrice * (1 - OrderDetails.Discount)) AS total_revenue").
		Scan(&totalRevenue).Error

	if lErr != nil {
		log.Println("Error : FTR01", lErr.Error())
		return totalRevenue, lErr
	} else {
		log.Println("Total Revenue Fetched Successfully")
	}

	log.Println("TotalRevenue(-)")
	return totalRevenue, lErr
}

func ProdWiseRevenue(pSalesData salesFileInsert.ReqStruct) ([]salesFileInsert.ProductRevenue, error) {
	log.Println("ProdWiseRevenue(+)")

	var lProdWiseData []salesFileInsert.ProductRevenue

	lErr := DBConn.GConnInst.Conn1_gorm.Table("OrderDetails").
		Select("Products.ProductName, SUM(OrderDetails.QuantitySold * OrderDetails.UnitPrice * (1 - OrderDetails.Discount)) AS total_revenue").
		Joins("JOIN Orders ON Orders.OrderID = OrderDetails.OrderID").
		Joins("JOIN Products ON Products.ProductID = OrderDetails.ProductID").
		Where("Orders.DateOfSale BETWEEN ? AND ?", pSalesData.FromDate, pSalesData.ToDate).
		Group("OrderDetails.ProductID, Products.ProductName").
		Scan(&lProdWiseData).Error

	if lErr != nil {
		log.Println("Error : FPWR01", lErr.Error())
		return lProdWiseData, lErr
	} else {
		log.Println("Product wise Revenue Fetched Successfully")
	}

	log.Println("ProdWiseRevenue(-)")
	return lProdWiseData, lErr
}

func CatWiseData(pSalesData salesFileInsert.ReqStruct) ([]salesFileInsert.CategoryRevenue, error) {
	log.Println("CatWiseData(+)")

	var lCatWiseData []salesFileInsert.CategoryRevenue

	lErr := DBConn.GConnInst.Conn1_gorm.Table("OrderDetails").
		Select("Products.Category, SUM(OrderDetails.QuantitySold * OrderDetails.UnitPrice * (1 - OrderDetails.Discount)) AS total_revenue").
		Joins("JOIN Orders ON Orders.OrderID = OrderDetails.OrderID").
		Joins("JOIN Products ON Products.ProductID = OrderDetails.ProductID").
		Where("Orders.DateOfSale BETWEEN ? AND ?", pSalesData.FromDate, pSalesData.ToDate).
		Group("Products.Category").
		Scan(&lCatWiseData).Error

	if lErr != nil {
		log.Println("Error : FCWD01", lErr.Error())
		return lCatWiseData, lErr
	} else {
		log.Println("Catagory wise Revenue Fetched Successfully")
	}

	log.Println("CatWiseData(-)")
	return lCatWiseData, lErr
}

func RegWiseData(pSalesData salesFileInsert.ReqStruct) ([]salesFileInsert.RegionRevenue, error) {
	log.Println("RegWiseData(+)")

	var lRegWiseData []salesFileInsert.RegionRevenue

	lErr := DBConn.GConnInst.Conn1_gorm.Table("OrderDetails").
		Select("COALESCE(Orders.Region, 'Unknown') AS Region, SUM(OrderDetails.QuantitySold * OrderDetails.UnitPrice * (1 - OrderDetails.Discount)) AS Revenue").
		Joins("JOIN Orders ON Orders.OrderID = OrderDetails.OrderID").
		Where("Orders.DateOfSale BETWEEN ? AND ?", pSalesData.FromDate, pSalesData.ToDate).
		Group("COALESCE(Orders.Region, 'Unknown')").
		Scan(&lRegWiseData).Error

	if lErr != nil {
		log.Println("Error : FRWD01", lErr.Error())
		return lRegWiseData, lErr
	} else {
		log.Println("Region wise Revenue Fetched Successfully")
	}

	log.Println("RegWiseData(-)")
	return lRegWiseData, lErr
}
