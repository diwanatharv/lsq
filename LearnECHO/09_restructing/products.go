package tronics

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ProductValidator struct {
	validator *validator.Validate
}

func (p *ProductValidator) Validate(i interface{}) error {
	return p.validator.Struct(i)
}

var products = []map[int]string{{1: "mobiles"}, {2: "TV"}, {3: "Laptop"}}

func getProducts(c echo.Context) error {
	return c.JSON(http.StatusOK, products)
}
func getProduct(c echo.Context) error {
	var product map[int]string
	for _, p := range products {
		for k := range p {
			//u can get the path paramter by e.param
			pid, err := strconv.Atoi(c.Param("id")) //we get the paramter and int is converted to
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
		return c.JSON(http.StatusNotFound, "product not found")

	}
	return c.JSON(http.StatusOK, product)
}

func deleteProduct(c echo.Context) error {
	var product map[int]string
	var index int
	for i, p := range products {
		for k := range p {
			//u can get the path paramter by e.param
			pid, err := strconv.Atoi(c.Param("id")) //we get the paramter and int is converted to
			// e.param is the url we extract id
			if err != nil {
				return err
			}
			if pid == k {
				product = p
				index = i
			}
		}
	}
	if product == nil {
		return c.JSON(http.StatusNotFound, "product not found")

	}

	//isme khud sei delete karna pardega delete function
	splice := func(s []map[int]string, index int) []map[int]string {
		//this step is easy
		//we are making a slice of elements before it
		//making a slice of elements after it
		//then merging it
		return append(s[:index], s[index+1:]...) //yei vala vo object ko remove karega through slicing
	}
	products = splice(products, index)
	return c.JSON(http.StatusOK, product)
}
func updateProduct(c echo.Context) error {
	var product map[int]string
	pid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	for _, p := range products {
		for k := range p {
			if pid == k {
				product = p
			}
		}
	}
	if product == nil {
		return c.JSON(http.StatusNotFound, "product not found")
		// now for validation same code  from  the post
	}
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
	product[pid] = reqbody.Name
	return c.JSON(http.StatusOK, product)
}
func createProduct(c echo.Context) error {
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
}
