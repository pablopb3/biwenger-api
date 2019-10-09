package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/pablopb3/biwenger-api/pkg/biwenger"

	"github.com/gin-gonic/gin"
)

const loginURL = "/auth/login"

func login(cli biwenger.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user biwenger.User
		if err := c.BindJSON(&user); err != nil {
			log.Println("error binding user", err)
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		userJSON, err := json.Marshal(user)
		if err != nil {
			log.Println("error marshalling user", err)
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		req := biwenger.Request{
			Method:   "POST",
			Body:     userJSON,
			Endpoint: loginURL,
		}

		body, err := cli.DoRequest(req)
		if err != nil {
			log.Println("error doing request for login", err)
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, string(body))
	}
}
