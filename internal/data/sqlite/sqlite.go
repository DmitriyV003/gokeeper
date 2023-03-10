package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"os"
)

type SQLite struct {
	*sql.DB
}

func NewSQLite(ctx context.Context, path string) (*SQLite, error) {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		return nil, fmt.Errorf("create a SQLite database file: %w", err)
	}
	if err = file.Close(); err != nil {
		return nil, fmt.Errorf("close the database file: %w", err)
	}

	conn, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, fmt.Errorf("connect to the SQLite database: %w", err)
	}

	if err = conn.Ping(); err != nil {
		return nil, fmt.Errorf("ping the SQLite database: %w", err)
	}

	db := &SQLite{conn}
	if err = db.migrate(); err != nil {
		return nil, fmt.Errorf("migrate database: %w", err)
	}

	go func() {
		<-ctx.Done()
		db.Close()
	}()

	return db, nil
}

func (s *SQLite) migrate() error {
	if err := s.createSettingsTable(); err != nil {
		return err
	}

	if err := s.createUsersTable(); err != nil {
		return err
	}

	if err := s.createLoginSecretsTable(); err != nil {
		return err
	}

	if err := s.createCardSecretsTable(); err != nil {
		return err
	}

	return nil
}

func (s *SQLite) createSettingsTable() error {
	query := `CREATE TABLE IF NOT EXISTS settings (
		key   TEXT NOT NULL UNIQUE,
		value TEXT NOT NULL
	)`
	if _, err := s.Exec(query); err != nil {
		return err
	}

	return nil
}

func (s *SQLite) createUsersTable() error {
	query := `CREATE TABLE IF NOT EXISTS users (
		id         INTEGER PRIMARY KEY AUTOINCREMENT,
		login      TEXT NOT NULL UNIQUE,
		password   TEXT NOT NULL,
		aesSecret  TEXT NOT NULL,
		privateKey TEXT NOT NULL
	)`
	if _, err := s.Exec(query); err != nil {
		return err
	}

	return nil
}

func (s *SQLite) createLoginSecretsTable() error {
	query := `CREATE TABLE IF NOT EXISTS login_secrets (
		id              INTEGER PRIMARY KEY AUTOINCREMENT,
		website         TEXT NOT NULL,
		username        TEXT NOT NULL,
		password    TEXT NOT NULL,
		additional_data TEXT NOT NULL,
		user_id			TEXT NOT NULL
	)`
	if _, err := s.Exec(query); err != nil {
		return err
	}

	return nil
}

func (s *SQLite) createCardSecretsTable() error {
	query := `CREATE TABLE IF NOT EXISTS cards (
		id  INTEGER PRIMARY KEY AUTOINCREMENT,
   		cardholder_name TEXT NOT NULL,
    	type TEXT NOT NULL,
    	expire_date TEXT NOT NULL,
    	valid_from TEXT NOT NULL,
    	number TEXT NOT NULL,
    	secret_code TEXT NOT NULL,
    	additional_data TEXT NULL
	)`
	if _, err := s.Exec(query); err != nil {
		return err
	}

	return nil
}
