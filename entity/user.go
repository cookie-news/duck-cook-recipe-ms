package entity

type User struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	User  string `json:"user"`
	Name  string `json:"name"`
	Image string `json:"image"`
}
