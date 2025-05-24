package main

import (
	DBconn "Sales/DBConnection"
	salesFileInsert "Sales/Sales_File_Insert"
	common "Sales/common/toml"
	fetchSalesData "Sales/fetch_Sales_Data"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {

	log.Println("Server started..")

	lFile, lErr := os.OpenFile("./log/logfile"+time.Now().Format("02012006.15.04.05.000000000")+".txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if lErr != nil {
		log.Fatalf("error opening file: %v", lErr)
	}
	defer lFile.Close()

	log.SetOutput(lFile)

	// Database connection process
	lErr = DBconn.GlobalDBConnection()
	if lErr != nil {
		log.Fatal(lErr)
	}
	defer DBconn.GConnInst.Conn1_sql.Close()

	lAutorunConfig := common.ReadTomlFile("./toml/serviceconfig.toml")
	lisAutorundaily := fmt.Sprintf("%v", lAutorunConfig.(map[string]interface{})["AutoRun"])
	if lisAutorundaily == "Y" {
		go func() {

			for {
				RefreshSalesData()
			}
		}()
	}

	http.HandleFunc("/ReloadSalesData", fetchSalesData.ReloadSalesData)
	http.HandleFunc("/GetRevenue", fetchSalesData.GetRevenue)

	http.ListenAndServe(":8080", nil)
	fmt.Println("server end....")

}

func RefreshSalesData() {
	log.Println("RefreshSalesData(+)")
	lNow := time.Now()

	lTimeConfig := common.ReadTomlFile("./toml/serviceconfig.toml")
	lHour := fmt.Sprintf("%v", lTimeConfig.(map[string]interface{})["hour"])
	lminute := fmt.Sprintf("%v", lTimeConfig.(map[string]interface{})["minute"])

	lHour_int, lErr := strconv.Atoi(lHour)
	if lErr != nil {
		log.Println("Error : MRSD01", lErr)
	}
	lminute_int, lErr := strconv.Atoi(lminute)
	if lErr != nil {
		log.Println("Error : MRSD02 ", lErr)

	}
	log.Println(lHour_int, ":", lminute_int)

	if lNow.Hour() == lHour_int && lNow.Minute() == lminute_int {

		lErr := salesFileInsert.ReadFile_updatedata()
		if lErr != nil {
			log.Println("Error : MRSD03 ", lErr)
		}

		time.Sleep(61 * time.Second)

	} else {
		time.Sleep(30 * time.Second)
	}
	log.Println("RefreshSalesData(-)")

}
