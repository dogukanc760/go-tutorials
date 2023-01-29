package routes

import (
	"rest-api/controllers"

	"github.com/gin-gonic/gin"
)

type UserRouteController struct {
	userController controllers.UserController
}

func NewRouteUserController(userController controllers.UserController) UserRouteController {
	return UserRouteController{userController}
}

func (uc *UserRouteController) PostRoute(rg *gin.RouterGroup) {

	router := rg.Group("users")
	router.GET("/", uc.userController.GetUsers)
	// 
}