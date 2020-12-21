package controller

import (
	"net/http"
	"github.com/gorilla/mux"
	"GoRest/service"
	"GoRest/util"
)

func ReadLineItemsByUserIdentityID(res http.ResponseWriter, req *http.Request) {
	/**
	 * Request params: 
	 * - id
	 * - name
	 * - pacingType
	 * - isContinous
	 * - Archived
	 * - Publisher
	 * - CreatorCompanyName
	 * - BrandCompanyName
	 * - BrandName
	 * - InitiativeName
	 */
	identityID := mux.Vars(req)["identityID"]
	params := make(map[string]string)
	params["lineItemID"] = req.FormValue("lineItemID")
	params["lineItemName"] = req.FormValue("lineItemName")
	params["isContinuous"] = req.FormValue("isContinuous")
	params["archived"] = req.FormValue("archived")
	params["publisher"] = req.FormValue("publisher")
	params["creatorCompanyName"] = req.FormValue("creatorCompanyName")
	params["brandCompanyName"] = req.FormValue("brandCompanyName")
	params["brandName"] = req.FormValue("brandName")
	params["initiativeName"] = req.FormValue("initiativeName")
	
	data:= service.ReadLineItemsByUserIdentityID(identityID,params)
												
	util.PackingSendingData(res, req, http.StatusOK, &data)
}
