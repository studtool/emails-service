package queues

//go:generate easyjson

//easyjson:json
type RegistrationEmailData struct {
	Email string `json:"email"`
	Token string `json:"token"`
}
