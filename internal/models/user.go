package models

type User struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	SignupTime int64  `json:"signupTime"`
}
