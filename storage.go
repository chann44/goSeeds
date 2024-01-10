package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)


type Storage interface {
	CreateAccount (*Account) error
	UpdateAccount (*Account) error
	DeleteAccount (int) error
	GetAccountByID (int) (*Account, error)
}


type postgresStore struct {
	db *sql.DB
}


func NewPostgresStore() (*postgresStore, error) {

	connStr := "user=postgres dbname=postgres password=goSeeds sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil , err
	}
	
	if err:= db.Ping(); err != nil {{
		return nil, err
	}}

	return &postgresStore{
		db: db,
	}, nil
}


// database initialization and table cation 

func (s *postgresStore) init() error {
	return s.CreateAccountTable()
}


func (s *postgresStore) CreateAccountTable() error {
	query := `CREATE TABLE IF NOT EXISTS account (
		id SERIAL PRIMARY KEY,
		name VARCHAR(50),
		balance DECIMAL,
		created_at TIMESTAMP
	);`

	_, err := s.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}




// stuff related to account

func (s *postgresStore) CreateAccount(*Account) error {
	return nil
}

func (s *postgresStore) UpdateAccount(*Account) error {
	return nil
}

func (s *postgresStore) DeleteAccount(id int) error {
	return nil
}


func (s *postgresStore) GetAccountByID(id int) (*Account, error) {
	return nil, nil
}