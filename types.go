package main

import "time"


type CreateAccountRequest struct {
	Name string `json:"name"`
}

type Account struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Balance int `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
}


func NewAccount(Name string) *Account {
	return &Account{
		Name: Name,
		CreatedAt: time.Now().UTC(),
	}
}