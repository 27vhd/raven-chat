package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "modernc.org/sqlite"

	"github.com/27vhd/raven-chat/internal/handlers"
	"github.com/27vhd/raven-chat/internal/repository"
)

func main() {
	db, err := sql.Open("sqlite", "./chat.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := repository.NewSQLiteRepository(db)
	if err := repo.Init(); err != nil {
		log.Fatal(err)
	}

	handler := handlers.NewChatHandler(repo)

	http.HandleFunc("/", handler.Index)
	http.HandleFunc("/send", handler.SendMessage)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
