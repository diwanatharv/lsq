package tronics

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var e = echo.New()
var v = validator.New()

func init() {
	err := cleanenv.ReadEnv(cfg)
	fmt.Printf("%+v", cfg)
	if err != nil {
		e.Logger.Fatal("unable to load configuration")
	}
}

//know about how to implement
//func servemessageHandlerFunc(c echo.Context) error {
//	fmt.Println("inside middleware")
//	return c.String(http.StatusOK, "middleware is here")
//}
// func serveMessagemiddleware(next echo.HandlerFunc) echo.HandlerFunc {
// return servemessageHandlerFunc()
//}

func serveMessage(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("i am inside middleware")
		return next(c) //??? how is this working
	}
}
func Start() {
	e.Use(middleware.BasicAuth())
	e.GET("/products", getProducts, serveMessage)
	e.GET("/products/:id", getProduct)
	e.DELETE("/product/:id", deleteProduct)
	e.PUT("/product/:id", updateProduct)
	e.POST("/products", createProduct)
	e.Logger.Print("hello listening on this port")
	e.Logger.Fatal(e.Start(":7200"))
}
