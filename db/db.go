package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type FortuneDatabase interface {
	Initialize() error
	ReadRandom() (string, error)
	Create(fortune string) (int64, error)
	Update(index int, fortune string) error
	Delete(index int) error
}

type fortuneDatabase struct {
	conn *sql.DB
}

func Create(dbFilePath string) FortuneDatabase {
	db, err := sql.Open("sqlite3", dbFilePath)
	if err != nil {
		log.Fatal(err)
	}

	fortuneDb := fortuneDatabase{db}
	if err := fortuneDb.Initialize(); err != nil {
		log.Fatal(err)
	}

	return fortuneDb

}

func (db fortuneDatabase) Initialize() error {
	_, err := db.conn.Exec("CREATE TABLE IF NOT EXISTS fortunes (id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, fortune TEXT)")
	return err
}

func (db fortuneDatabase) ReadRandom() (string, error) {
	var fortune string
	err := db.conn.QueryRow("SELECT fortune FROM fortunes ORDER BY RANDOM() LIMIT 1").Scan(&fortune)
	return fortune, err
}

func (db fortuneDatabase) Create(fortune string) (int64, error) {
	result, err := db.conn.Exec("INSERT INTO fortunes(fortune) VALUES(?)", fortune)

	if result != nil {
		return result.LastInsertId()
	}
	return -1, err
}

func (db fortuneDatabase) Update(index int, fortune string) error {
	_, err := db.conn.Exec("UPDATE fortunes SET fortune=? WHERE id=?", fortune, index)
	return err
}

func (db fortuneDatabase) Delete(index int) error {
	_, err := db.conn.Exec("DELETE FROM fortunes WHERE id=?", index)
	return err
}
