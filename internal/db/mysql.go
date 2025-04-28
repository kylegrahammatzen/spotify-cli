package db

import (
	"database/sql"
	"errors"
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

func NewMySQLClient(config Config, sqlFilePath string) (*MySQLClient, error) {
	//fmt.Println("Opening database connection...")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/?multiStatements=true",
		config.User, config.Password, config.Host, config.Port)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening database connection: %v", err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Hour)

	//fmt.Println("Pinging database...")
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to MySQL: %v", err)
	}

	//fmt.Println("Ensuring database exists...")
	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", config.DBName))
	if err != nil {
		return nil, fmt.Errorf("error creating database: %v", err)
	}

	//fmt.Println("Selecting database...")
	_, err = db.Exec(fmt.Sprintf("USE %s", config.DBName))
	if err != nil {
		return nil, fmt.Errorf("error selecting database: %v", err)
	}

	client := &MySQLClient{db: db}

	//fmt.Println("Checking if 'users' table exists...")
	var tableName string
	err = db.QueryRow(`
		SELECT TABLE_NAME
		FROM INFORMATION_SCHEMA.TABLES
		WHERE TABLE_SCHEMA = ? AND TABLE_NAME = 'users'
	`, config.DBName).Scan(&tableName)

	if errors.Is(err, sql.ErrNoRows) {
		//fmt.Println("'users' table does not exist. Running SetupDB...")
		if setupErr := client.SetupDB(sqlFilePath); setupErr != nil {
			return nil, setupErr
		}
	} else if err != nil {
		return nil, fmt.Errorf("error checking for 'users' table: %v", err)
	}

	//fmt.Println("Database client is ready.")
	return client, nil
}

func (c *MySQLClient) Close() error {
	//fmt.Println("Closing database connection...")
	return c.db.Close()
}

func (c *MySQLClient) GetDB() *sql.DB {
	return c.db
}

func (c *MySQLClient) SetupDB(sqlFilePath string) error {
	fmt.Printf("Setting up database using %s...\n", sqlFilePath)

	content, err := ioutil.ReadFile(sqlFilePath)
	if err != nil {
		return fmt.Errorf("error reading SQL file: %v", err)
	}

	fmt.Println("Executing SQL setup script...")
	_, err = c.db.Exec(string(content))
	if err != nil {
		return fmt.Errorf("error executing SQL setup: %v", err)
	}

	fmt.Println("Database setup completed.")
	return nil
}
