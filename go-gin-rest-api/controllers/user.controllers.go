package controllers

import (
	"fmt"
	"net/http"
	"rest-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

func NewPostController(DB *gorm.DB) UserController {
	return UserController{DB}
}

func (uc *UserController) GetUsers(ctx *gin.Context) {

	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "150")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit
	fmt.Println(offset)
	var user []models.User

	results := uc.DB.Raw(`select firstname, email, lastname from public."user"`).Scan(&user)

	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(user), "data": user})

}
