package queues

//go:generate easyjson

//easyjson:json
type CreatedUserData struct {
	UserID string `json:"userId"`
}

//easyjson:json
type DeletedUserData struct {
	UserID string `json:"userId"`
}
