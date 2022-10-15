package dto

type NotionCredentialsDTO struct {
	AccessToken   string         `json:"access_token"`
	TokenType     string         `json:"token_type"`
	BotId         string         `json:"bot_id"`
	WorkspaceName string         `json:"workspace_name"`
	WorkspaceIcon string         `json:"workspace_icon"`
	WorkspaceId   string         `json:"workspace_id"`
	Owner         NotionOwnerDTO `json:"owner"`
}
