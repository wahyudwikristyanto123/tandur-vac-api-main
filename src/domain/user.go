package domain

type User struct {
	Token    string `json:"token"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Name     string `json:"name"`
}
