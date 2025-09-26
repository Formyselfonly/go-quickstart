package service

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

// Init initializes all services
func Init(ctx context.Context) {
	// Initialize database connection
	initDatabase(ctx)
}

// initDatabase initializes the database connection and creates tables
func initDatabase(ctx context.Context) {
	db := g.DB()
	
	// Create users table if not exists
	createUsersTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username VARCHAR(50) NOT NULL UNIQUE,
		email VARCHAR(100) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL,
		status INTEGER DEFAULT 1,
		create_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		update_at DATETIME DEFAULT CURRENT_TIMESTAMP
	)
	`
	
	_, err := db.Exec(ctx, createUsersTableSQL)
	if err != nil {
		g.Log().Error(ctx, "Failed to create users table:", err)
		return
	}
	
	g.Log().Info(ctx, "Database initialized successfully")
}
