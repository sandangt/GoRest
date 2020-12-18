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
	lineItemID := req.FormValue("lineItemID")
	lineItemName := req.FormValue("lineItemName")
	isContinuous := req.FormValue("isContinuous")
	archived := req.FormValue("archived")
	publisher := req.FormValue("publisher")
	creatorCompanyName := req.FormValue("creatorCompanyName")
	brandCompanyName := req.FormValue("brandCompanyName")
	brandName := req.FormValue("brandName")
	initiativeName := req.FormValue("initiativeName")
	data:= service.ReadLineItemsByUserIdentityID(identityID, 
												lineItemID,
												lineItemName,
												isContinuous,
												archived,
												publisher,
												creatorCompanyName,
												brandCompanyName,
												brandName,
												initiativeName,)
//	str := fmt.Sprintf("test:%v, test1:%v, test2:%v, test3:%v", req.FormValue("test"), req.FormValue("test1"), req.FormValue("test2"), req.FormValue("test3"))
//	fmt.Println(str)
	util.PackingSendingData(res, req, http.StatusOK, &data)
}
