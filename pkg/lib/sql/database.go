package sql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const schema = `
CREATE TABLE IF NOT EXISTS products (
	id INT AUTO_INCREMENT PRIMARY KEY,
	name VARCHAR(255) NOT NULL,
	value INT NOT NULL
);
`

func NewDatabase(connPath string) *sqlx.DB {
	db := sqlx.MustConnect("mysql", connPath)

	db.MustExec(schema)

	// EXAMPLE DATA
	db.MustExec("INSERT IGNORE INTO products (id, name, value) VALUES (1, 'Keyboard', 19900);")
	db.MustExec("INSERT IGNORE INTO products (id, name, value) VALUES (2, 'Mouse', 9900);")

	return db
}
