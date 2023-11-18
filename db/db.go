package db

import (
	"database/sql"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/pressly/goose"

	_ "github.com/lib/pq"

	log "github.com/sirupsen/logrus"
)

type DBConfig struct {
	POSTGRES_USER     string
	POSTGRES_PASSWORD string
	POSTGRES_DB       string
	POSTGRES_PORT     string
	POSTGRES_HOST     string
	PORT              string
}

type DbHandler struct {
	SqlDB *sql.DB
}

func DbInit(dbConfig string) (DbHandler, error) {
	var logger = log.New()
	logger.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339,
	})
	logger.SetOutput(os.Stdout)
	logger.Infof("db url is %+v", dbConfig)
	db, err := sql.Open("postgres", dbConfig)

	cw, err := os.Getwd()
	if err != nil {
		logger.Errorf("getting error with new working directory %+v", err)
		return DbHandler{}, err
	}

	dir := fmt.Sprintf("%s/migrations", cw)
	logger.Infof("db dir is %+v", dir)
	if err := goose.Up(db, dir); err != nil {
		logger.Errorf("goose migration issue %+v", err)
		return DbHandler{}, err
	}

	logger.Infoln("Database migrations applied success!")

	return DbHandler{
		SqlDB: db,
	}, nil
}

func NewDatabaseConfig(config DBConfig) (string, error) {
	var dbParams []string
	dbParams = append(dbParams, fmt.Sprintf("user=%s", config.POSTGRES_USER))
	if host := config.POSTGRES_HOST; host == "" {
		dbParams = append(dbParams, fmt.Sprintf("host=%s", "localhost"))
	} else {
		dbParams = append(dbParams, fmt.Sprintf("host=%s", config.POSTGRES_HOST))
	}
	dbParams = append(dbParams, fmt.Sprintf("port=%s", config.POSTGRES_PORT))
	dbParams = append(dbParams, fmt.Sprintf("dbname=%s", config.POSTGRES_DB))
	if password := config.POSTGRES_PASSWORD; password != "" {
		dbParams = append(dbParams, fmt.Sprintf("password=%s", password))
	}

	dbParams = append(dbParams, fmt.Sprintf("sslmode=%s", "disable"))
	return strings.Join(dbParams, " "), nil
}
