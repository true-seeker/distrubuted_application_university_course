package dto

type ImportDTO struct {
	Data   ImportDataDTO   `json:"data,omitempty"`
	Fields []SheetFieldDTO `json:"fields,omitempty"`
}
