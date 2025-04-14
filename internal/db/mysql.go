package db

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLClient struct {
	db *sql.DB
}

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

func NewMySQLClient(config Config) (*MySQLClient, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/?multiStatements=true",
		config.User, config.Password, config.Host, config.Port)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening database connection: %v", err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Hour)

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to MySQL: %v", err)
	}

	return &MySQLClient{db: db}, nil
}

func (c *MySQLClient) Close() error {
	return c.db.Close()
}

func (c *MySQLClient) GetDB() *sql.DB {
	return c.db
}

func (c *MySQLClient) SetupDB(dbName, sqlFilePath string) error {
	row := c.db.QueryRow(
		"SELECT SCHEMA_NAME FROM INFORMATION_SCHEMA.SCHEMATA WHERE SCHEMA_NAME = ?",
		dbName,
	)
	var existingDB string
	if err := row.Scan(&existingDB); err == nil && existingDB == dbName {
		return nil
	}

	content, err := ioutil.ReadFile(sqlFilePath)
	if err != nil {
		return fmt.Errorf("error reading SQL file: %v", err)
	}
	_, err = c.db.Exec(string(content))
	if err != nil {
		return fmt.Errorf("error executing SQL setup: %v", err)
	}
	return nil
}
