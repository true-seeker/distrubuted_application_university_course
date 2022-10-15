package dto

type NotionOwnerDTO struct {
	Type string        `json:"type"`
	User NotionUserDTO `json:"user"`
}
