package models

// swagger:parameters auth signIn
type User struct {
	Password string `json:"password"`
	Username string `json:"username"`
}
