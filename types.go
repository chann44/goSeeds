package main

type Account struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Balance int `json:"balance"`
}


func NewAccount(Name string) *Account {
	return &Account{
		Name: Name,
	}
}