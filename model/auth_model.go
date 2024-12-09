package model

import "net/http"

type AuthModel struct {
	UserID   string `json:"userid"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthInterface interface {
	Login() (string, error)
	Register(*http.Request) (string, error)
}
