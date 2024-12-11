package main

import (
	"database/sql"
	"log"
	"net/http"
	"sever/controller"
	"sever/repository"

	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

func main() {
	con, err := sql.Open("postgres", "postgresql://auth_p959_user:YeiKsBKDQb2biAzqYDfwY8ygCw84vSzy@dpg-ct856vd2ng1s73f1ice0-a.oregon-postgres.render.com/auth_p959")
	if err != nil {
		panic(err)
	}
	defer con.Close()
	mux:= mux.NewRouter()
	authrepo := repository.NewAuthRepository(con)
	authcontr := controller.NewAuthController(authrepo)
	mux.HandleFunc("/register", authcontr.RegisterController)
	mux.HandleFunc("/login", authcontr.LoginController)
	PORT := ":8000"
	log.Printf("Server is running on port %v", PORT)
	log.Fatal(http.ListenAndServe(PORT, mux))
}
