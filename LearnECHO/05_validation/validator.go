package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// we are assuming that while in the post request the user always sends the correct data
// install go validator package
func main() {
	products := []map[int]string{{1: "mobiles"}, {2: "TV"}, {3: "Laptop"}}

	//things which are variable are passed as/:id(field name)
	// contains the whole url

	e := echo.New()      //creates an instance of echo
	v := validator.New() //validator will validaste things used  near the json one
	e.GET("/", func(e echo.Context) error {
		return e.String(http.StatusOK, "well hello there")
	}) //this needs 3 path
	//agar post metod try karoge toh error dei dega ya path kuch aur bhi doge toh bhi error dega
	//creates a server
	e.GET("/products", func(e echo.Context) error {
		return e.JSON(http.StatusOK, products)
	})
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
	e.POST("/products", func(e echo.Context) error {
		fmt.Println("iske andar hai")
		type body struct {
			//you have to learn all these things
			Name string `json:"product_name" validate:"required,min=4"`
			//for learning purpose
			Vendor  string `json:"vendor" validate: " min=5,max=10"`
			Email   string `json:"email", validate :"required_with=vendor,email"`
			Website string `json:"website" validate :"url"`
		}
		var reqbody body
		//bind methods
		if err := e.Bind(&reqbody); err != nil {
			return err
		}
		//it gets populated by the value of the bind and then we have to  check here
		//after this validation is checked
		err := v.Struct(reqbody)
		if err != nil {
			return err
		}
		product := map[int]string{
			len(products) + 1: reqbody.Name,
		}
		products = append(products, product)
		return e.JSON(http.StatusOK, product)
	})
	//can use a simple method for listening
	e.Logger.Print("hello listening on this port")
	e.Logger.Fatal(e.Start(":7000"))
}
