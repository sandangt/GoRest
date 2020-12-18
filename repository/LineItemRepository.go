package repository

import (
	"fmt"
	
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	
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
									initiativeName interface{},) []entity.LineItem {
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
		MATCH (lineItem)-[:BELONGS_TO]->(initiative:Initiative)-[:MARKETED_BY]->(brand:Brand)-[:OWNED_BY]->(brandCompany:Company)
		WHERE (lineItem.name=~'.*%v.*' AND creatorCompany.name=~'.*%v.*' AND brandCompany.name=~'.*%v.*' AND brand.name=~'.*%v.*' AND initiative.name=~'.*%v.*')
		OPTIONAL MATCH (lineItem)
		WHERE (lineItem.id='%v' AND lineItem.isContinuous='%v' AND lineItem.archived='%v' AND lineItem.publisher='%v')
		RETURN DISTINCT user.identityID as IdentityID, lineItem.id as LineItemID, creatorCompany.id as CreatorCompanyID, creatorCompany.name as CreatorCompanyName, brandCompany.id as BrandCompanyID, brandCompany.name as BrandCompanyName, brand.id as BrandID, brand.name as BrandName, initiative.id as InitiativeID, initiative.name as InitiativeName, lineItem.publisher as Publisher, lineItem.archived as Archived order by LineItemID;`, identityID, lineItemName, creatorCompanyName, brandCompanyName, brandName, initiativeName, lineItemID, isContinuous, archived, publisher)
		// , brandCompany.name=~'.*%v.*', brand.name=~'.*%v.*', initiative.name=~'.*%v.*'
		// , brandCompanyName, brandName, initiativeName)
		// , lineItemID, isContinuous, archived, publisher)
		// OPTIONAL MATCH (lineItem)
		// WHERE lineItem.id='%v', lineItem.isContinuous='%v', lineItem.archived='%v', lineItem.publisher='%v'
		// WHERE lineItem.id='%v', lineItem.name=~'.*%v.*', lineItem.isContinuous='%v', lineItem.archived='%v', lineItem.publisher='%v', creatorCompany.name=~'.*%v.*', brandCompany.name=~'.*%v.*', brand.name=~'.*%v.*', initiative.name=~'.*%v.*'
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
		identityID, _ := result.Record().Get("IdentityID")
		creatorCompanyID, _ := result.Record().Get("CreatorCompanyID")
		creatorCompanyName, _ := result.Record().Get("CreatorCompanyName")
		brandCompanyID, _ := result.Record().Get("BrandCompanyID")
		brandCompanyName, _ := result.Record().Get("BrandCompanyName")
		brandID, _ := result.Record().Get("BrandID")
		brandName, _ := result.Record().Get("BrandName")
		initiativeID, _ := result.Record().Get("InitiativeID")
		initiativeName, _ := result.Record().Get("InitiativeName")
		publisher, _ := result.Record().Get("Publisher")
		archived, _ := result.Record().Get("Archived")
		arr = append(arr, entity.LineItem {IdentityID: identityID, CreatorCompanyID: creatorCompanyID, CreatorCompanyName: creatorCompanyName, BrandCompanyID: brandCompanyID, BrandCompanyName: brandCompanyName, BrandID: brandID, BrandName: brandName, InitiativeID: initiativeID, InitiativeName: initiativeName, Publisher: publisher, Archived: archived})
	}
	return arr
}

