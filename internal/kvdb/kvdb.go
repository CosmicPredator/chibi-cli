package kvdb

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"github.com/CosmicPredator/chibi/internal"
	_ "modernc.org/sqlite"
)

type KV struct {
	db *sql.DB
}

func Open() (*KV, error) {
	path, err := os.UserConfigDir()
	if err != nil {
		return nil, fmt.Errorf("unable to get user config path: %w", err)
	}
	dbDirPath := filepath.Join(path, "chibi")
	if _, err := os.Stat(dbDirPath); err != nil {
		os.MkdirAll(dbDirPath, 0o755)
	}
	dbPath := filepath.Join(dbDirPath, internal.DB_PATH)
	
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
	}

	schema := `
	CREATE TABLE IF NOT EXISTS kv (
		key   TEXT PRIMARY KEY,
		value BLOB
	);
	`

	if _, err := db.Exec(schema); err != nil {
		return nil, err
	}

	return &KV{db: db}, nil
}


func (k *KV) Set(ctx context.Context, key string, value []byte) error {
	_, err := k.db.ExecContext(ctx,
		`INSERT INTO kv(key,value) VALUES(?,?)
		 ON CONFLICT(key) DO UPDATE SET value=excluded.value`,
		key, value,
	)
	return err
}

func (k *KV) Get(ctx context.Context, key string) ([]byte, error) {
	var v []byte
	err := k.db.QueryRowContext(ctx,
		`SELECT value FROM kv WHERE key=?`,
		key,
	).Scan(&v)

	return v, err
}

func (k *KV) Delete(ctx context.Context, key string) error {
	_, err := k.db.ExecContext(ctx,
		`DELETE FROM kv WHERE key=?`,
		key,
	)
	return err
}

func (k *KV) Close() error {
	return k.db.Close()
}