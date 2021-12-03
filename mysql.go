package crudgen

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type MysqlDatabase struct {
	Type     string
	Host     string
	Port     string
	Name     string
	Username string
	Password string
	dbc      *sql.DB
}

func New(host string, port string, name string, user string, password string) (*MysqlDatabase, error) {
	var db = MysqlDatabase{
		Type:     "mysql",
		Host:     host,
		Port:     port,
		Name:     name,
		Username: user,
		Password: password,
	}

	dbc, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, name))
	if err != nil {
		return nil, fmt.Errorf("Error opening mysql database: %v", err)
	}

	db.dbc = dbc
	if err = dbc.Ping(); err != nil {
		dbc.Close()
		return nil, fmt.Errorf("Error connecting database: %v", err)
	}

	return &db, nil
}

func (db *MysqlDatabase) GenerateCRUDFilesForTable(tablename string) (string, string, error) {
	var model string
	var impl string

	if db.dbc == nil {
		return "", "", errors.New("Cannot generate. MySQL database not open")
	}
	return model, impl, nil
}
