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

//easyjson:json
type ProfileToCreateData struct {
	UserID string `json:"userId"`
}

//easyjson:json
type ProfileToDeleteData struct {
	UserID string `json:"userId"`
}

//easyjson:json
type AvatarToCreateData struct {
	UserID string `json:"userId"`
}

//easyjson:json
type AvatarToDeleteData struct {
	UserID string `json:"userId"`
}

//easyjson:json
type DocumentUserToCreateData struct {
	UserID string `json:"userId"`
}

//easyjson:json
type DocumentUserToDeleteData struct {
	UserID string `json:"userId"`
}

//easyjson:json
type RegistrationEmailData struct {
	Email string `json:"email"`
	Token string `json:"token"`
}
