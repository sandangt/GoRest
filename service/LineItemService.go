package service

import (
	"GoRest/repository"
	"GoRest/entity"
)

func ReadLineItemsByUserIdentityID(identityID string) interface{} {
	// Check if user with given identityID exists
	// Entities.User{} : empty struct
	if (repository.ReadUserByIdentityID(identityID) == entity.User{}) {
		return entity.Message { Content: "User not found" }
	}
	return repository.ReadLineItemsByUserIdentityID(identityID)
}
