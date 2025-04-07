package rest

import (
	"Cinema_API/models"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func (r *Rest) Register(ctx *gin.Context) {
	param := models.UserRegister{}

	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = r.Usecase.UserUsecase.Register(param)
	if err != nil {
		if errors.Is(err, gorm.ErrInvalidData) {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "invalid data",
				"error":   err.Error(),
			})
		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "record not found",
				"error":   err.Error(),
			})
		} else if errors.Is(err, gorm.ErrDuplicatedKey) {
			ctx.JSON(http.StatusConflict, gin.H{
				"message": "User already exists",
				"error":   err.Error(),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
}

func (r *Rest) Login(ctx *gin.Context) {
	param := models.UserLogin{}

	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, err := r.Usecase.UserUsecase.Login(param)
	if err != nil {
		if errors.Is(err, gorm.ErrInvalidData) {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "invalid data",
				"error":   err.Error(),
			})
		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "record not found",
				"error":   err.Error(),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
		return
	}
	ctx.JSON(http.StatusOK, token)
}
