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
	// Initially connect without specifying a database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/",
		config.User, config.Password, config.Host, config.Port)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening database connection: %v", err)
	}

	// Set connection pool settings
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Hour)

	// Check connection
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

func (c *MySQLClient) SetupDB(dbName string, sqlFilePath string) error {
	// Create database if it doesn't exist
	_, err := c.db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName))
	if err != nil {
		return fmt.Errorf("error creating database: %v", err)
	}

	// Switch to the database
	_, err = c.db.Exec(fmt.Sprintf("USE %s", dbName))
	if err != nil {
		return fmt.Errorf("error switching to database: %v", err)
	}

	// Run the SQL setup script
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
