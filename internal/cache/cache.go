package cache

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"github.com/CosmicPredator/chibi/internal"
	_ "modernc.org/sqlite"
)

type Cache struct {
	db *sql.DB
}

func Open() (*Cache, error) {
	dbDirPath, err := internal.ResolveDataPath()
	if err != nil {
		return nil, err
	}
	if err := os.MkdirAll(dbDirPath, 0o755); err != nil {
		return nil, fmt.Errorf("unable to create config path: %w", err)
	}
	dbPath := filepath.Join(dbDirPath, internal.DB_PATH)

	db, err := sql.Open("sqlite", dbPath)
	
	cache := &Cache{
		db: db,
	}
	if err != nil {
		return nil, err
	}
	
	return cache, nil
}
