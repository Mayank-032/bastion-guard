package infrastructure

import "database/sql"

var DB *sql.DB

func InitDB(host, port, db string) error {
	return nil
}