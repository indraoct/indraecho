
package main

import (
	"net/http"
	"math/rand"
	"github.com/labstack/echo"
	"time"
	"indraecho/models"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
        	return c.String(http.StatusOK, "Hello, World!")
	})

	// JSONP
	e.GET("/jsonp", func(c echo.Context) error {
		var content struct {
			Response  string    `json:"response"`
			Timestamp time.Time `json:"timestamp"`
			Random    int       `json:"random"`
		}
		content.Response = "Sent via JSONP"
		content.Timestamp = time.Now().UTC()
		content.Random = rand.Intn(1000)
		return c.JSON(1,content)
	})

	//JSONUser
	e.GET("jsonuser", func(c echo.Context) error{
		var content struct{
			Status string  `json:"status"`
			Data []map[string]interface{}  `json:"data"`
		}

		content.Status = "1"
                var data,_  = models.GetuserList()
		content.Data = data

		return c.JSON(1,&content)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
	
