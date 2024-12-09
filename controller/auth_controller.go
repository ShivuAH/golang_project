package controller

import (
	"net/http"
	"sever/model"
)

type AuthController struct {
	authRepsotory model.AuthInterface
}

func NewAuthController(ai model.AuthInterface) *AuthController {
	return &AuthController{
		authRepsotory: ai,
	}
}
func (ac *AuthController) RegisterController(w http.ResponseWriter, r *http.Request) {
	mes, err := ac.authRepsotory.Register(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(mes))
}

func (ac *AuthController) LoginController(w http.ResponseWriter, r *http.Request) {
	mes, err := ac.authRepsotory.Register(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(mes))
}
