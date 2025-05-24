package fetchSalesData

import (
	salesFileInsert "Sales/Sales_File_Insert"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func ReloadSalesData(w http.ResponseWriter, r *http.Request) {

	log.Println("ReloadSalesData (+)")

	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	(w).Header().Set("Access-Control-Allow-Methods", http.MethodGet)
	(w).Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Content-Length,Accept-Encoding,X-CSTF-Token,Authorization")

	if r.Method == http.MethodGet {

		// var lInsertedIds string
		var lRespRec salesFileInsert.CommonRespStruct
		lRespRec.Status = "S"

		lErr := salesFileInsert.ReadFile_updatedata()
		if lErr != nil {
			// helperpkg.LogError(lErr)
			log.Println("Error : SRS01", lErr.Error())
			lRespRec.Status = "E"
			lRespRec.ErrMsg = "Error: SRS01" + lErr.Error()
		}

		lData, lErr := json.Marshal(lRespRec)
		if lErr != nil {
			log.Println("Error SCF02 :", lErr.Error())
		} else {
			fmt.Fprint(w, string(lData))
		}

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	log.Println("ReloadSalesData (-)")
}
