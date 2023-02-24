package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// echo.context has all the url
func main() {

	//product is a slice of map
	products := []map[int]string{{1: "mobiles"}, {2: "TV"}, {: "Laptop"}}

	//things which are variable are passed as/:id(field name)
	// contains the whole url

	e := echo.New() //creates an instance of echo
	e.GET("/", func(e echo.Context) error {
		return e.String(http.StatusOK, "well hello there")
	}) //this needs 3 path
	//agar post metod try karoge toh error dei dega ya path kuch aur bhi doge toh bhi error dega
	//creates a server

	e.GET("/products/:id", func(e echo.Context) error {
		var product map[int]string
		for _, p := range products {
			for k := range p {
				//u can get the path paramter by e.param
				pid, err := strconv.Atoi(e.Param("id")) //we get the paramter and int is converted to
				// e.param is the url we extract id
				if err != nil {
					return err
				}
				if pid == k {
					product = p
				}
			}
		}
		if product == nil {
			return e.JSON(http.StatusNotFound, "product not found")

		}
		return e.JSON(http.StatusOK, product)

	})
	//can use a simple method for listening
	e.Logger.Print("hello listening on this port")
	e.Logger.Fatal(e.Start(":8000"))
}
