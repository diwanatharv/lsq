package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// echo.context has all the url
func main() {

	//product is a slice of map
	products := []map[int]string{{1: "mobiles"}, {2: "TV"}, {3: "Laptop"}}

	//things which are variable are passed as/:id(field name)
	// contains the whole url

	e := echo.New() //creates an instance of echo
	e.GET("/", func(e echo.Context) error {
		return e.String(http.StatusOK, "well hello there")
	})
	//like in this case /: is params which is passed as parametr
	//products/apple then printed willl be apple
	//now another case comes of queryparams(olderthan)?olderthan=iphone10
	e.GET("/products/:vendors", func(e echo.Context) error {
		return e.JSON(http.StatusOK, e.Param("vendor"))
	})
	//agar post metod try karoge toh error dei dega ya path kuch aur bhi doge toh bhi error dega
	//creates a server

	//can use a simple method for listening
	e.Logger.Print("hello listening on this port")
	e.Logger.Fatal(e.Start(":8000"))
}
