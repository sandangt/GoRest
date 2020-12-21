package entity

type LineItem struct {
	LineItemID interface{} `json:LineItemID`
	LineItemName interface{} `json:LineItemName`
	Publisher interface{} `json:"Publisher"`
	IsContinuous interface{} `json:"IsContinuous"`
	Archived interface{} `json:"Archived"`
	IdentityID interface{} `json:"IdentityID"`
	CreatorCompanyID interface{} `json:"CreatorCompanyID"`
	CreatorCompanyName interface{} `json:"CreatorCompanyName"`
	BrandCompanyID interface{} `json:"BrandCompanyID"`
	BrandCompanyName interface{} `json:"BrandCompanyName"`
	BrandID interface{} `json:"BrandID"`
	BrandName interface{} `json:"BrandName"`
	InitiativeID interface{} `json:"InitiativeID"`
	InitiativeName interface{} `json:"InitiativeName"`
}
