package main

import (
	"fmt"
	"go-for-beginner/scrapper"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

const fileName string = "comic_list.csv"

// Handler
func handleHello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

// Router & Handler
func handleGetComics(c echo.Context) error {
	target := scrapper.CleanString(c.Param("target"));

	defer os.Remove(fileName)

	fmt.Println(target)
	scrapper.Scrapper(target)
	c.Attachment(fileName, fileName)
	return c.String(http.StatusOK, "Get All list successful")
}

func main(){
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
  
	// Routes
	e.GET("/", handleHello)
	e.GET("/comics/:target", handleGetComics)

	// Start server
	e.Logger.Fatal(e.Start(":80"))
	
}