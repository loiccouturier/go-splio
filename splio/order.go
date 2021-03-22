package splio

import "github.com/shopspring/decimal"

type Order struct {
	ExternalId      *string          `json:"external_id,omitempty"`
	ContactId       string           `json:"contact_id"`
	CardCode        string           `json:"card_code,omitempty"`
	StoreExternalId string           `json:"store_external_id"`
	OrderAt         string           `json:"ordered_at,omitempty"`
	DiscountAmount  *decimal.Decimal `json:"discount_amount,omitempty"`
	ShippingAmount  *decimal.Decimal `json:"shipping_amount,omitempty"`
	TotalPrice      *decimal.Decimal `json:"total_price,omitempty"`
	TaxAmount       *decimal.Decimal `json:"tax_amount,omitempty"`
	Completed       bool             `json:"completed"`
	SalesPerson     string           `json:"sales_person,omitempty"`
	Currency        string           `json:"currency,omitempty"`
	Products        []OrderLine      `json:"products,omitempty"`
	CustomFields    *[]CustomField   `json:"custom_fields,omitempty"`
}

func (o *Order) setCustomField(fieldId int, value interface{}) {
	var customFields []CustomField
	if o.CustomFields != nil {
		customFields = *o.CustomFields
	}
	customFields = append(customFields, CustomField{Id: intPtr(fieldId), Value: value})
	o.CustomFields = &customFields
}

func (o *Order) SetHasGift(hasGift bool) {
	if hasGift {
		o.setCustomField(20, "1")
	} else {
		o.setCustomField(20, "0")
	}
}

type OrderLine struct {
	UnitPrice       *decimal.Decimal `json:"unit_price,omitempty"`
	Quantity        int              `json:"quantity,omitempty"`
	TotalLineAmount *decimal.Decimal `json:"total_line_amount,omitempty"`
	ProductId       string           `json:"product_id,omitempty"`
	CustomFields    *[]CustomField   `json:"custom_fields,omitempty"`
}

func (ol *OrderLine) setCustomField(fieldId int, value interface{}) {
	var customFields []CustomField
	if ol.CustomFields != nil {
		customFields = *ol.CustomFields
	}
	customFields = append(customFields, CustomField{Id: intPtr(fieldId), Value: value})
	ol.CustomFields = &customFields
}

func (ol *OrderLine) SetIsMainProduct(isRootProduct bool) {
	if isRootProduct {
		ol.setCustomField(4, "1")
	} else {
		ol.setCustomField(4, "0")
	}
}

func (ol *OrderLine) SetLabelInfo(title, text string) {
	ol.setCustomField(2, title)
	ol.setCustomField(3, text)
}

func (ol *OrderLine) SetGiftCardInfo(sender, message string) {
	ol.setCustomField(5, message)
	ol.setCustomField(6, sender)
}

func (ol *OrderLine) SetCustomTitle(title string) {
	ol.setCustomField(1, title)
}

func (ol *OrderLine) SetCustomVisual(visualUrl string) {
	ol.setCustomField(0, visualUrl)
}

func (ol *OrderLine) SetBundlePrice(price float64) {
	ol.setCustomField(7, price)
}

func intPtr(i int) *int {
	return &i
}