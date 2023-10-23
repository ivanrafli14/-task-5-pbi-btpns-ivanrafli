package middleware

import (
	
	"net/http"

	"time"

	"github.com/gin-gonic/gin"

	"github.com/ivanrafli14/API-BTPN/database"
	"github.com/ivanrafli14/API-BTPN/helpers"
	"github.com/ivanrafli14/API-BTPN/models"
)

func CheckAuth() gin.HandlerFunc {
	return func(c *gin.Context){
		cookie, err := c.Cookie("Authorization")
		
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
				"status":  "fail",
			})
			return
		}

		claims, err := helpers.ParseToken(cookie);

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
				"status":  "fail",
			})
			return
		}
		
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		var user models.User
		database.DB.First(&user, claims["sub"]);

		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
				
		}

		c.Set("user", user)
		c.Next()

		
	}
	



	

}