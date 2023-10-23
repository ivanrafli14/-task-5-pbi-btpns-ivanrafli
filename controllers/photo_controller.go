package controllers

import (
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/ivanrafli14/API-BTPN/app"
	"github.com/ivanrafli14/API-BTPN/database"
	"github.com/ivanrafli14/API-BTPN/helpers"
	"github.com/ivanrafli14/API-BTPN/models"
)

func GetAllPhotos(c *gin.Context){
	tokenStr,err :=  c.Cookie("Authorization")
	

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status" : "fail",
			"message" : "Unauthorized",
		})
		return
	}

	claims, err := helpers.ParseToken(tokenStr);

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status" : "fail",
			"message" : "Unauthorized",
		})
		return
	}
	userID := claims["sub"]
	var photos []models.Photo

	database.DB.Where("user_id =?", userID).Find(&photos);

	c.JSON(http.StatusOK, gin.H {
		"status" : "success",
		"message" : "fetch all user photos",
		"data" : photos,
	})
}

func CreatePhoto (c *gin.Context){
	var photoReq app.PhotoCreate

	if err :=c.ShouldBindJSON(&photoReq); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H {
			"status" : "fail",
			"message" : err.Error(),
		})
		return
	}

	if _,err := govalidator.ValidateStruct(&photoReq); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H {
			"status" : "fail",
			"message" : err.Error(),
		})
		return
	}

	tokenStr,err :=  c.Cookie("Authorization")
	

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status" : "fail",
			"message" : "Unauthorized",
		})
		return
	}
	claims, err := helpers.ParseToken(tokenStr);
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status" : "fail",
			"message" : "Unauthorized",
		})
		return
	}
	
	UserID := claims["sub"].(float64)
	

	photo := models.Photo{
		Title:    photoReq.Title,
		Caption:  photoReq.Caption,
		PhotoUrl: photoReq.PhotoUrl,
		UserID:   uint(UserID),
	}

	
	if err := database.DB.Create(&photo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status" : "fail",
			"message" : err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status" : "successs",
		"message" : "successs create photo",
	})


}

func UpdatePhoto(c *gin.Context){
	var photoReq app.PhotoUpdate

	if err :=c.ShouldBindJSON(&photoReq); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H {
			"status" : "fail",
			"message" : err.Error(),
		})
		return
	}

	if _,err := govalidator.ValidateStruct(&photoReq); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H {
			"status" : "fail",
			"message" : err.Error(),
		})
		return
	}

	photoID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status" : "fail",
			"message" : "ID is not valid",
		})
		return
	}
	

	photo := models.Photo{
		Title:    photoReq.Title,
		Caption:  photoReq.Caption,
		PhotoUrl: photoReq.PhotoUrl,
	}

	if database.DB.Model(&photo).Where("id =?", photoID).Updates(&photo).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H {
			"message": "failed to update data", 
			"status": "fail",
		})
		return
	}

	

	c.JSON(http.StatusOK, gin.H{
		"status": "Success",
		"message": "Success update photo",
		
	})


}

func DeletePhoto(c *gin.Context) {
	photoID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status" : "fail",
			"message" : "ID is not valid",
		})
		return
	}

	var photo models.Photo 

	if database.DB.Unscoped().Delete(&photo, photoID).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H {
			"status": "fail",
			"message": "failed to delete data",
		})
		return 
	}

	c.JSON(200, gin.H {
		"status": "success",
		"message": "success delete photo",
	})
}