package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {

	db, err := sql.Open("postgres", "postgresql://shivu_user:jczvCepM28TDYReaCkBDnHcpzXfJoTz5@dpg-cta389qj1k6c738gs6m0-a.singapore-postgres.render.com/shivu")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(w, "Hello Shivu")
		log.Printf("Request recived on /login")

	})

	log.Println("Hello Shivu")

	log.Fatal(http.ListenAndServe(":8000", nil))
}
