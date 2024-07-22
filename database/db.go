package database

import (
	"database/sql"
	"testServ/models"
)

var db *sql.DB

func SaveMessage(message *models.Message) error {
	query := `INSERT INTO messages (content, processed) VALUES ($1, $2) RETURNING id`
	return db.QueryRow(query, message.Text, message.Processed).Scan(&message.ID)
}

func GetProcessedMessagesCount() (int, error) {
	var count int
	query := `SELECT COUNT(*) FROM messages WHERE processed = true`
	err := db.QueryRow(query).Scan(&count)
	return count, err
}
