package middleware

import (
	"backend/config"
	"backend/lib"
	"backend/models"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Add(r *gin.Engine, h ...gin.HandlerFunc) {
	for _, v := range h {
		r.Use(v)
	}
}

func AuthorizationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			lib.RespondWithError(c, http.StatusUnauthorized, "Unauthorized: Token is missing")
			return
		}
		c.Next()
	}
}

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			lib.RespondWithError(c, http.StatusUnauthorized, "Unauthorized: Token is missing")
			return
		}

		splitToken := strings.Split(token, " ")
		if len(splitToken) != 2 {
			lib.RespondWithError(c, http.StatusUnauthorized, "Unauthorized: Token has invalid syntax")
			return
		}

		resp, err := http.Post(config.App.AUTH_SERVICE+"/auth/checkToken", "application/json", strings.NewReader(`{"token":"`+splitToken[1]+`"}`))
		if err != nil {
			lib.RespondWithError(c, http.StatusUnauthorized, "Unauthorized: Token is invalid || No response from authentication service")
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			lib.RespondWithError(c, http.StatusUnauthorized, "Unauthorized: Token is invalid || Authentication service returned an error")
			return
		}

		text, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			lib.RespondWithError(c, http.StatusUnauthorized, "Unauthorized: Token is invalid || Couldn't read response body")
			return
		}

		var response models.CheckTokenResponse
		json.Unmarshal(text, &response)

		if response.Status != "OK" {
			lib.RespondWithError(c, http.StatusUnauthorized, "Unauthorized: Token is invalid || Token is not correct")
			return
		}

		c.Next()
	}
}
