package repository

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"sever/model"
	"golang.org/x/crypto/bcrypt"

	"github.com/google/uuid"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{
		db,
	}

}
func (ar *AuthRepository) Register(r *http.Request) (string, error) {
	var newUser model.AuthModel
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		return "", err
	}
	newUser.UserID = uuid.New().String()
	hashpass, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	_, err = ar.db.Exec("INSERT INTO users(user_id , user_name , password) VALUES($1,$2,$3)", newUser.UserID, newUser.UserName, string(hashpass))
	if err != nil {
		return "", err
	}
	return "Success", nil
}
func (ar *AuthRepository) Login() (string, error) {
	return "", nil
}
