package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"backend/config"
	"backend/models"

	"github.com/gin-gonic/gin"
)

func RegisterHandler(c *gin.Context) {

	// Get the body of the request
	var body models.RegisterRequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	postBody, _ := json.Marshal(body)
	postBodyReader := bytes.NewReader(postBody)

	resp, err := http.Post(config.App.AUTH_SERVICE+"/auth/register", "application/json", postBodyReader)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	defer resp.Body.Close()

	readBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var response models.RegisterResponse

	json.Unmarshal(readBody, &response)

	c.JSON(http.StatusOK, response)
}

func LoginHandler(c *gin.Context) {

	var body models.LoginRequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	postBody, _ := json.Marshal(body)
	postBodyReader := bytes.NewReader(postBody)

	resp, err := http.Post(config.App.AUTH_SERVICE+"/auth/login", "application/json", postBodyReader)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	defer resp.Body.Close()

	readBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var response models.LoginResponse

	json.Unmarshal(readBody, &response)

	if response.Message == "User not found!" {
		c.JSON(http.StatusNotFound, gin.H{
			"message": response.Message,
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
