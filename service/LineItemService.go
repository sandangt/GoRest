package service

import (
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
									initiativeName interface{},) interface{} {
	return repository.ReadLineItemsByUserIdentityID(identityID,
													lineItemID,
													lineItemName,
													isContinuous, 
													archived,
													publisher,
													creatorCompanyName,
													brandCompanyName,
													brandName,
													initiativeName,)
}

func InvalidPublisher() interface{} {
	return entity.Message { Content: "Invalid publisher" }
}

func InvalidArchivedStatus() interface{} {
	return entity.Message { Content: "Invalid archived status" }
}
