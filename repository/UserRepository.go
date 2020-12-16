package repository

import (
	"fmt"
	
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	
	"GoRest/entity"
)

func ReadUserByIdentityID(identityID string) entity.User {
	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close()
	user, err := session.ReadTransaction( func(tx neo4j.Transaction) (interface{}, error) {
		query := fmt.Sprintf(`MATCH (user:User) WHERE user.identityID='%v'
		RETURN DISTINCT user.id as ID, user.identityID as IdentityID, user.email as Email;`, identityID)
		result, err := tx.Run(query, nil)
		if err != nil {
			return nil, err
		}
		return parseUser(result), err
	})
	if err != nil {
		return entity.User{}
	}
	return user.(entity.User)
}

func parseUser(result neo4j.Result) entity.User {
	var user entity.User
	for result.Next() {
		id, _ := result.Record().Get("ID")
		identityID, _ := result.Record().Get("IdentityID")
		email, _ := result.Record().Get("Email")
		user = entity.User {ID: id, IdentityID: identityID, Email: email}
	}
	return user
}
