package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// now we will use echos validator
//echos validator is also a struct
//i dont think  i have to remember this
//just simply copy paste

type ProductValidator struct {
	validator *validator.Validate
}

func (p *ProductValidator) Validate(i interface{}) error {
	return p.validator.Struct(i)
}
func main() {
	products := []map[int]string{{1: "mobiles"}, {2: "TV"}, {3: "Laptop"}}

	//things which are variable are passed as/:id(field name)
	// contains the whole url

	e := echo.New() //creates an instance of echo
	v := validator.New()
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
	e.POST("/products", func(c echo.Context) error {
		fmt.Println("iske andar hai")
		type body struct {
			Name string `json:"product_name"`
		}
		var reqbody body
		e.Validator = &ProductValidator{validator: v}
		//bind methods
		if err := c.Bind(&reqbody); err != nil {
			return err
		}
		if err := c.Validate(reqbody); err != nil {
			return err
		}
		product := map[int]string{
			len(products) + 1: reqbody.Name,
		}
		products = append(products, product)
		return c.JSON(http.StatusOK, product)
	})
	//can use a simple method for listening
	e.Logger.Print("hello listening on this port")
	e.Logger.Fatal(e.Start(":8000"))

}
