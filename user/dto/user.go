package dto

import "user-services/model"

type ReqSignUp struct {
	Email    string `validate:"required,email" json:"email,omitempty"`
	Password string `validate:"required,min=6" json:"password,omitempty"`
	Username string `validate:"required" json:"username,omitempty"`
}

type ReqSignIn struct {
	Email    string `validate:"required,email" json:"email,omitempty"`
	Password string `validate:"required" json:"password,omitempty"`
}

type ResSignUp struct {
	model.User
	Token string
}
