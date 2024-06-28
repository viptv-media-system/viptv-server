package main

// import (
// 	"net/http"
// 	"github.com/gin-gonic/gin"
// )

// Routes structures
type show struct {
	ID	 	string	`json:"id"`
	NAME	string	`json:"name"`
	PATH	string	`json:"path"`	
	GROUP	string	`json:"group"`
}

type movie struct {
	ID	 	string	`json:"id"`
	NAME	string	`json:"name"`
	PATH	string	`json:"path"`
	GROUP	string	`json:"group"`
}

type tv struct {
	ID		string	`json:"id"`
	NAME	string	`json:"name"`
	PATH	string	`json:"path"`
	GROUP	string	`json:"group"`
}


// func getShow(c *gin.Context) {
//     c.IndentedJSON(http.StatusOK, shows)
// }
