package splio

type SearchFieldName string

const (
	Key          SearchFieldName = "key"
	FirstName                    = "firstname"
	LastName                     = "lastname"
	CreationDate                 = "creation_date"
	Language                     = "language"
	Email                        = "email"
	CellPhone                    = "cellphone"
	CardCode                     = "card code"
)

type SearchOperator string

const (
	Is       SearchOperator = "is"
	After                   = "after"
	Before                  = "before"
	Contains                = "contains"
	Ends                    = "ends"
	Equal                   = "equal"
	greater                 = "greater"
	IsNot                   = "isnot"
	Lower                   = "lower"
	NotEqual                = "notequal"
	Starts                  = "starts"
)

type SearchField struct {
	Key      SearchFieldName `json:"key"`
	Operator SearchOperator  `json:"operator"`
	Value    string          `json:"value"`
}

type SearchBody struct {
	PerPage    int           `json:"per_page"`
	PageNumber int           `json:"page_number"`
	Fields     []SearchField `json:"fields"`
}

type SearchResult struct {
	Count       int       `json:"count_element,omitempty"`
	CurrentPage int       `json:"current_page,omitempty"`
	PerPage     int       `json:"per_page,omitempty"`
	Sort        []string  `json:"sort,omitempty"`
	Elements    []Contact `json:"elements,omitempty"`
}
