package service

import (
	"strconv"
	"fmt"
	"strings"
	"GoRest/repository"
	"GoRest/entity"
)

func ReadLineItemsByUserIdentityID(identityID string, params map[string]string) interface{} {
	var optionalParams []string
	
	if params["publisher"] != "" {
		optionalParams = append(optionalParams, fmt.Sprintf("publisher:\"%v\"", params["publisher"]))
	}
	
	if params["lineItemID"] != "" {
		optionalParams = append(optionalParams, fmt.Sprintf("lineItemID:\"%v\"", params["lineItemID"]))
	}
	
	if params["isContinuous"] != "" {
		temp, _ := strconv.ParseBool(params["isContinuous"])
		optionalParams = append(optionalParams, fmt.Sprintf("isContinuous:%v", temp))
	}
	
	if params["archived"] != "" {
		temp, _ := strconv.ParseBool(params["archived"])
		optionalParams = append(optionalParams, fmt.Sprintf("archived:%v", temp))
	}
	
	return repository.ReadLineItemsByUserIdentityID(identityID, params, strings.Join(optionalParams,","))
}

func InvalidPublisher() interface{} {
	return entity.Message { Content: "Invalid publisher" }
}

func InvalidArchivedStatus() interface{} {
	return entity.Message { Content: "Invalid archived status" }
}

func InvalidContinuousStatus() interface{} {
	return entity.Message { Content: "Invalid continuous status" }
}

func ReadAllPublishers() []interface{} {
	return repository.ReadAllPublishers()
}
