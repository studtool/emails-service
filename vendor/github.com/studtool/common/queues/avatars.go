package queues

//go:generate easyjson

//easyjson:json
type AvatarToCreateData struct {
	UserID string `json:"userId"`
}

//easyjson:json
type AvatarToDeleteData struct {
	UserID string `json:"userId"`
}
