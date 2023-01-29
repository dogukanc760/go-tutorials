package services

// type UserService struct {
// 	DB *gorm.DB

// }

// func FindUsers(ctx *gin.Context) {
// 	var page = ctx.DefaultQuery("page", "1")
// 	var limit = ctx.DefaultQuery("limit", "150")

// 	intPage, _ := strconv.Atoi(page)
// 	intLimit, _ := strconv.Atoi(limit)
// 	offset := (intPage - 1) * intLimit

// 	var users []models.User
// 	userDB := *&gorm.DB{}
// 	results := userDB.Limit(intLimit).Offset(offset).Find(&users)
// 	if results.Error != nil {
// 		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(users), "data": users})
// }
