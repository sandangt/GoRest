package repository

import (
	"fmt"
	
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	
	"GoRest/entity"
)
func ReadLineItemsByUserIdentityID(identityID string, params map[string]string, optionalParamsStr string) []entity.LineItem {
	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close()
	lineItems, err := session.ReadTransaction( func(tx neo4j.Transaction) (interface{}, error) {
		query := fmt.Sprintf(`MATCH (user:User)-[:EMPLOYED_BY]->(company:Company) 
		WHERE user.identityID='%v'
		WITH user, company
		MATCH (Company)<-[:WORKS_WITH]-(brand:Brand)<-[:MARKETED_BY]-(initiative:Initiative)<-[:BELONGS_TO]-(lineItem:LineItem)
		WHERE ( lineItem.publisher IN company.allowedAccess OR company.type='PROVIDER' )
		WITH user, lineItem
		MATCH (lineItem)-[:CREATED_BY]->(creator:User)-[:EMPLOYED_BY]->(creatorCompany:Company)
		WITH user, lineItem, creatorCompany
		MATCH (lineItem {%v})-[:BELONGS_TO]->(initiative:Initiative)-[:MARKETED_BY]->(brand:Brand)-[:OWNED_BY]->(brandCompany:Company)
		WHERE (
			lineItem.name=~'.*%v.*' 
			AND creatorCompany.name=~'.*%v.*' 
			AND brandCompany.name=~'.*%v.*' 
			AND brand.name=~'.*%v.*' 
			AND initiative.name=~'.*%v.*'
		)
		RETURN lineItem.id as LineItemID, lineItem.name as LineItemName, lineItem.publisher as Publisher, lineItem.isContinuous as IsContinuous, lineItem.archived as Archived, user.identityID as IdentityID, creatorCompany.id as CreatorCompanyID, creatorCompany.name as CreatorCompanyName, brandCompany.id as BrandCompanyID, brandCompany.name as BrandCompanyName, brand.id as BrandID, brand.name as BrandName, initiative.id as InitiativeID, initiative.name as InitiativeName order by LineItemID;`, identityID, optionalParamsStr, params["lineItemName"], params["creatorCompanyName"], params["brandCompanyName"], params["brandName"], params["initiativeName"])
		result, err := tx.Run(query, nil)
		if err != nil {
			return nil, err
		}
		return parseLineItems(result), err
	})
	if err != nil {
		return nil
	}
	return lineItems.([]entity.LineItem)
}

func parseLineItems(result neo4j.Result) []entity.LineItem {
	var arr []entity.LineItem
	for result.Next() {
		lineItemID, _ := result.Record().Get("LineItemID") 
		lineItemName, _ := result.Record().Get("LineItemName") 
		publisher, _ := result.Record().Get("Publisher")
		isContinuous, _ := result.Record().Get("IsContinuous")
		archived, _ := result.Record().Get("Archived")
		identityID, _ := result.Record().Get("IdentityID")
		creatorCompanyID, _ := result.Record().Get("CreatorCompanyID")
		creatorCompanyName, _ := result.Record().Get("CreatorCompanyName")
		brandCompanyID, _ := result.Record().Get("BrandCompanyID")
		brandCompanyName, _ := result.Record().Get("BrandCompanyName")
		brandID, _ := result.Record().Get("BrandID")
		brandName, _ := result.Record().Get("BrandName")
		initiativeID, _ := result.Record().Get("InitiativeID")
		initiativeName, _ := result.Record().Get("InitiativeName")
		arr = append(arr, entity.LineItem {LineItemID: lineItemID, LineItemName: lineItemName, Publisher: publisher, IsContinuous: isContinuous, Archived: archived, IdentityID: identityID, CreatorCompanyID: creatorCompanyID, CreatorCompanyName: creatorCompanyName, BrandCompanyID: brandCompanyID, BrandCompanyName: brandCompanyName, BrandID: brandID, BrandName: brandName, InitiativeID: initiativeID, InitiativeName: initiativeName})
	}
	return arr
}

func ReadAllPublishers() []interface{} {
	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close()
	publisherList, err := session.ReadTransaction( func(tx neo4j.Transaction) (interface{}, error) {
		var publisherList []interface{}
		query := `MATCH (li:LineItem) RETURN DISTINCT li.publisher as Publisher;`
		result, err := tx.Run(query, nil)
		if err != nil {
			return nil, err
		}
		for result.Next() {
			publisher, _ := result.Record().Get("Publisher")
			publisherList = append(publisherList, publisher)						
		}
		return publisherList, err
	})
	if err != nil {
		return nil
	}
	return publisherList.([]interface{})
}


