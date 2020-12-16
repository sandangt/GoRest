package service

import (
	"GoRest/repository"
	"GoRest/entity"
)

var publisherList = []string {"facebook","twitter","snapchat","pinterest","linkedin"}

func ReadLineItemsByUserIdentityID(identityID string) (interface{}, int) {
	// Check if user with given identityID exists
	// Entities.User{} : empty struct
	if (repository.ReadUserByIdentityID(identityID) == entity.User{}) {
		return entity.Message { Content: "User not found" }, 400
	}
	return repository.ReadLineItemsByUserIdentityID(identityID), 200
}
