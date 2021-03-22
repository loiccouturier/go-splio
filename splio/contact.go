package splio

type Contact struct {
	Id           *int           `json:"id,omitempty"`
	LastName     *string        `json:"lastname,omitempty"`
	FirstName    *string        `json:"firstname,omitempty"`
	CreationDate *string        `json:"creation_date,omitempty"`
	Language     *string        `json:"language,omitempty"`
	Email        string         `json:"email"`
	CellPhone    *string        `json:"cellphone,omitempty"`
	Lists        *[]List        `json:"lists,omitempty"`
	CustomFields *[]CustomField `json:"custom_fields,omitempty"`
	Loyalty      *[]Loyalty     `json:"loyalty,omitempty"`
	//DoubleOptin  *DoubleOptin   `json:"double_optin"`
}

type DoubleOptin struct {
	Message       *string `json:"message,omitempty"`
	Reminder      *string `json:"reminder,omitempty"`
	ReminderDelay *int    `json:"reminder_delay,omitempty"`
}

type CustomField struct {
	Id       *int        `json:"id,omitempty"`
	Name     *string     `json:"name,omitempty"`
	Value    interface{} `json:"value,omitempty"`
	DataType string      `json:"data_type,omitempty"`
}

type Loyalty struct {
	CardCode  string `json:"card_code,omitempty"`
	IdProgram string `json:"id_program,omitempty"`
}
