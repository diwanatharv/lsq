package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New() //creates an instance of echo
	e.GET("/", func(e echo.Context) error {
		return e.String(http.StatusOK, "well hello there")
	}) //this needs 3 path
	//agar post metod try karoge toh error dei dega ya path kuch aur bhi doge toh bhi error dega
	//creates a server

	//can use a simple method for listening
	e.Logger.Print("hello listening on this port")
	e.Logger.Fatal(e.Start(":8000")) //this returns an error and this error can be resolved

}
