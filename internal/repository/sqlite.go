package repository

import (
	"database/sql"

	"github.com/27vhd/raven-chat/internal/models"
)

type ChatRepository interface {
	SaveMessage(msg models.Message) error
	GetAllMessages() ([]models.Message, error)
}

type SQLiteRepository struct {
	db *sql.DB
}

func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{db}
}
