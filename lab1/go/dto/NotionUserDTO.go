package dto

type NotionUserDTO struct {
	Object    string `json:"object"`
	Id        string `json:"id"`
	Name      string `json:"name"`
	AvatarUrl string `json:"avatarUrl"`
	Type      string `json:"type"`
}
