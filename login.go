package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct {
	Token string `json:"token"`
}

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

const loginURL = "/auth/login"

func (cli client) login(c *gin.Context) {
	var user User
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

	req := request{
		method:   "POST",
		body:     userJSON,
		endpoint: loginURL,
	}

	body, err := cli.doRequest(req)
	if err != nil {
		log.Println("error doing request for login", err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, Login{Token: string(body)})
}
