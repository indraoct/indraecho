
package main

import (
	"net/http"
	"github.com/labstack/echo"
	"indraecho/models"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
        	return c.String(http.StatusOK, "Hello, World!")
	})

	/**
	 * API JSON USER
	 */
	e.GET("jsonuser", func(c echo.Context) error{
		var content struct{
			Status string  `json:"status"`
			Msg string `json:"msg"`
			Data []map[string]interface{}  `json:"data"`
		}

		content.Status = "1"
		content.Msg = "Success"
		content.Data,_ = models.GetuserList()

		return c.JSON(1,&content)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
	
