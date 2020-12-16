package entity

type LineItem struct {
	IdentityID interface{} `json:"IdentityID"`
	CreatorCompanyID interface{} `json:"CreatorCompanyID"`
	CreatorCompanyName interface{} `json:"CreatorCompanyName"`
	BrandCompanyID interface{} `json:"BrandCompanyID"`
	BrandCompanyName interface{} `json:"BrandCompanyName"`
	BrandID interface{} `json:"BrandID"`
	BrandName interface{} `json:"BrandName"`
	InitiativeID interface{} `json:"InitiativeID"`
	InitiativeName interface{} `json:"InitiativeName"`
	PublisherName interface{} `json:"PublisherName"`
	Archived interface{} `json:"Archived"`
}
