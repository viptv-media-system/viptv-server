package models

//	User object

type user struct {
	ID			string	`json:"id"`
	NAME		string	`json:"name"`
	EMAIL		string	`json:"email"`
	PASSWORD	string	`json:"password"`
	RECOVERY	string	`json:"recovery"`	//	Recovery password message
	AVATAR		string	`json:"avatar"`
}
