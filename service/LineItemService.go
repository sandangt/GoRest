package service

import (
	"strconv"
	"fmt"
	"strings"
	"GoRest/repository"
	"GoRest/entity"
)

func ReadLineItemsByUserIdentityID(identityID,
									lineItemID,
									lineItemName,
									isContinuous, 
									archived,
									publisher,
									creatorCompanyName,
									brandCompanyName,
									brandName,
									initiativeName string,) interface{} {
	var optionalParams []string
	if isContinuous != "" {
		temp, _ := strconv.ParseBool(isContinuous)
		optionalParams = append(optionalParams, fmt.Sprintf("isContinuous:%v", temp))
	}
	
	if archived != "" {
		temp, _ := strconv.ParseBool(archived)
		optionalParams = append(optionalParams, fmt.Sprintf("archived:%v", temp))
	}
	
	if publisher != "" {
		optionalParams = append(optionalParams, fmt.Sprintf("publisher:\"%v\"", publisher))
	}
	
	return repository.ReadLineItemsByUserIdentityID(identityID,
													lineItemID,
													lineItemName,
													creatorCompanyName,
													brandCompanyName,
													brandName,
													initiativeName,
													strings.Join(optionalParams,","),)
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
