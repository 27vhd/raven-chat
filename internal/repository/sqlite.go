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

func (s *SQLiteRepository) Init() error {
	query := `
	CREATE TABLE IF NOT EXISTS messages (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT,
    content TEXT,
    timestamp DATETIME
	);`

	_, err := s.db.Exec(query)
	return err
}

func (s *SQLiteRepository) SaveMessage(msg models.Message) error {
	query := `INSERT INTO messages (username, content, timestamp) VALUES (?, ?, ?)`
	_, err := s.db.Exec(query, msg.Username, msg.Content, msg.Timestamp)
	return err
}

func (s *SQLiteRepository) GetAllMessages() ([]models.Message, error) {
	query := `SELECT id, username, content, timestamp FROM messages`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []models.Message
	for rows.Next() {
		var msg models.Message
		err := rows.Scan(&msg.ID, &msg.Username, &msg.Content, &msg.Timestamp)
		if err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}

func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{db}
}
