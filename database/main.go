package database

import (
	"database/sql"
	"log"
	"os"
)

func Make() (*sql.DB, error) {
	os.Remove("ws-task.db")
	log.Println("Creating ws-task.db...")
	file, err := os.Create("ws-task.db")
	if err != nil {
		return nil, err
	}
	file.Close()
	log.Println("ws-task.db created")

	sqliteDatabase, err := sql.Open("sqlite3", "./ws-task.db")
	if err != nil {
		return nil, err
	}
	createTable(sqliteDatabase)
	return sqliteDatabase, nil
}

func createTable(db *sql.DB) error {
	createStudentTableSQL := `CREATE TABLE messages (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
        "user_id" VARCHAR(255) NOT NULL,
		"content" TEXT,
        "seen" BOOLEAN,
		"created_at" TEXT);`

	log.Println("Create message table...")
	statement, err := db.Prepare(createStudentTableSQL)
	if err != nil {
		return err
	}
	statement.Exec()
	log.Println("messages table created")
	return nil
}
