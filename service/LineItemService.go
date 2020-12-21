package service

import (
	"GoRest/repository"
	"GoRest/entity"
)

func ReadLineItemsByUserIdentityID(identityID string, params map[string]string) interface{} {
	
	return repository.ReadLineItemsByUserIdentityID(identityID, params)
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
