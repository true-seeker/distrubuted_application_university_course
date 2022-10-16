package dto

type NotionSearchResultDTO struct {
	Id         string      `json:"id"`
	Title      interface{} `json:"title"`
	Properties interface{} `json:"properties"`
}
