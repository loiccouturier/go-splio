package splio

type MultiStatusItem struct {
	ContactId   string `json:"contact_id"`
	Code        int    `json:"code"`
	Description string `json:"description"`
}

type MultiStatus struct {
	Time       int               `json:"time"`
	ErrorCount int               `json:"errors"`
	Items      []MultiStatusItem `json:"items"`
}
