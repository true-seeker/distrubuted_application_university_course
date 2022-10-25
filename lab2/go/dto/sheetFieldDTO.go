package dto

type SheetFieldDTO struct {
	FieldName    string           `json:"field_name,omitempty"`
	Title        string           `json:"title,omitempty"`
	EntityFields []EntityFieldDTO `json:"entity_fields,omitempty"`
}
