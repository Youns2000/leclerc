package pkg

type config struct {
	Url        string `json:"url"`
	ProductSku string `json:"productSku"`
	ExternalId string `json:"externalId"`
	Slug       string `json:"slug"`
}

type Detail struct {
	LastUpdateDate    string    `json:"lastUpdateDate"`
	LogisticClassCode string    `json:"logisticClassCode"`
	Label             string    `json:"label"`
	Variants          []Variant `json:"variants"`
}

type Variant struct {
	Id         string      `json:"id"`
	Sku        string      `json:"sku"`
	Label      string      `json:"label"`
	Slug       string      `json:"slug"`
	Attributes []Attribute `json:"attributes"`
	Offers     []Offer     `json:"offers"`
}

type Offer struct {
	Id         string `json:"id"`
	ExternalId string `json:"externalId"`
	SourceCode string `json:"sourceCode"`
	Locale     string `json:"locale"`
}

type Attribute struct {
	Code  string `json:"code"`
	Label string `json:"label"`
	Type  string `json:"type"`
	Value Value  `json:"value"`
}

type Value struct {
	Code     string `json:"code"`
	Label    string `json:"label"`
	Position int    `json:"position"`
}

type finalConfig struct {
	OperationName string `json:"operationName"`
	Variables     struct {
		Offers []struct {
			OfferId    string `json:"offerId"`
			ProductSku string `json:"productSku"`
			Quantity   int    `json:"quantity"`
			Slug       string `json:"slug"`
		} `json:"offers"`
	} `json:"variables"`
	Query string `json:"query"`
}
