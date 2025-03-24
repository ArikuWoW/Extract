package initializers

import (
	"github.com/ArikuWoW/extract/pkg/repository"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

func ConnectDB() *sqlx.DB {
	db, err := repository.NewPostgresDB()
	if err != nil {
		logrus.Fatalf("Failed to initialize db: %s", err.Error())
	}

	return db
}
