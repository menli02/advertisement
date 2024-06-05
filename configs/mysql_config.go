package configs

import "database/sql"

type MySQLConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

func NewMysqlDB(cfg MySQLConfig) (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:pass123@tcp(127.0.0.1:3306)/advertisement?")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
