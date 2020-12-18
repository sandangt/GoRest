package service

import (
	"GoRest/entity"
	"GoRest/repository"
)

func UserNotFound() interface{} {
	return entity.Message { Content: "User not found" }
}

func ReadUserByIdentityID(identityID string) interface{} {
	return repository.ReadUserByIdentityID(identityID)
}
