package repositories

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func InitDB(user, password, dbname, host, port string) (*sql.DB, error) {
    connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)

    db, err := sql.Open("postgres", connStr)
    if err != nil {
        return nil, err
    }

    // Test the connection
    if err = db.Ping(); err != nil {
        db.Close()
        return nil, err
    }

    return db, nil
}
