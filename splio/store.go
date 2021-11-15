package splio

type Store struct {
	ExternalId  *string        `json:"external_id,omitempty"`
	Name        string         `json:"name,omitempty"`
	Online      bool           `json:"online"`
	StoreType   string         `json:"store_type,omitempty"`
	Manager     string         `json:"manager,omitempty"`
	CustomField *[]CustomField `json:"custom_fields,omitempty"`
}
