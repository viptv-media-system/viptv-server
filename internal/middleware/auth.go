package middleware

//	middleware logon

type user struct {
	ID			string	`json:"id"`
	NAME		string	`json:"name"`
	PASSWORD	string	`json:"password"`
}
