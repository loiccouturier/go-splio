package splio

import "github.com/shopspring/decimal"

type Product struct {
	ExternalId  *string           `json:"external_id,omitempty"`
	Name        string           `json:"name,omitempty"`
	Description string           `json:"description,omitempty"`
	Brand       string           `json:"brand,omitempty"`
	Price       *decimal.Decimal `json:"price,omitempty"`
	Sku         string           `json:"sku,omitempty"`
	ImageUrl    string           `json:"img_url,omitempty"`
	Category    string           `json:"category,omitempty"`
	CustomField *[]CustomField   `json:"custom_fields,omitempty"`
}
