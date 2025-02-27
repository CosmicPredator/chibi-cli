package db

const (
	QUERY_CREATE_TABLE = `CREATE TABLE IF NOT EXISTS config (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		key TEXT NOT NULL,
		value TEXT
	)`

	QUERY_INSERT_CONFIG = `INSERT INTO config (key, value) VALUES (?, ?)`

	QUERY_GET_CONFIG = `SELECT value FROM config WHERE key = ?`
)
