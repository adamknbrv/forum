package db

import (
	"database/sql"
	"github.com/adamknbrv/forum/pkg/log"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/zap"
	_ "modernc.org/sqlite"
)

func SQLiteConnection() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "../forum.db")
	if err != nil {
		log.Send.Fatal("Ошибка подключения к базе данных",
			zap.Error(err))
		return nil, err
	}

	if err = db.Ping(); err != nil {
		log.Send.Fatal("Ошибка отключка от базы данных",
			zap.Error(err))
		db.Close()
		return nil, err
	}
	log.Send.Info("Успешное подключение к SQLite")

	return db, nil
}
