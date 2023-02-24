package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

const k string = "4000"

func main() {
	//Make use of environment variables to replace port in stead of a hard-coded value.
	//sometimes our port mai busy
	//there should be some mechanism by which we can select
	//this environment is set using the terminal
	//export and os.getenv kei andar jo likha hai voh
	port := os.Getenv(k)
	if port == "" {
		port = "8080"
	}
	e := echo.New()
	e.GET("/", func(e echo.Context) error {
		return e.String(http.StatusOK, "all is well")
	})
	e.Logger.Print(fmt.Sprintf("listining on the port %s", port))
	e.Logger.Fatal(e.Start(port))
}
