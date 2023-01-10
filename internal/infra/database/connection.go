package database

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Connection interface {
	QueryRow(stmt string, args ...any) (Row, error)
	Query(stmt string, args ...any) (Rows, error)
	Exec(stmt string, args ...any) error
	Close() error
}

type Row interface {
	Scan(dest ...any) error
}

type Rows interface {
	Close() error
	Next() bool
	Scan(dest ...any) error
}

type ConnectionAdapter struct {
	db *sql.DB
}

func NewConnection(driverName, dataSourceName string) (*ConnectionAdapter, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	return &ConnectionAdapter{db: db}, nil
}

func (c *ConnectionAdapter) QueryRow(stmt string, args ...any) (Row, error) {
	preparedStmt, err := c.db.Prepare(stmt)
	if err != nil {
		return nil, err
	}
	row := preparedStmt.QueryRow(args...)
	return row, nil
}

func (c *ConnectionAdapter) Query(stmt string, args ...any) (Rows, error) {
	preparedStmt, err := c.db.Prepare(stmt)
	if err != nil {
		return nil, err
	}
	return preparedStmt.Query(args...)
}

func (c *ConnectionAdapter) Exec(stmt string, args ...any) error {
	preparedStmt, err := c.db.Prepare(stmt)
	if err != nil {
		return err
	}
	_, err = preparedStmt.Exec(args...)
	return err
}

func (c *ConnectionAdapter) Close() error {
	return c.db.Close()
}
