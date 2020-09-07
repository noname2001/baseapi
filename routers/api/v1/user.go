package v1

import (
	"baseapi/global"
	"baseapi/global/response"
	"baseapi/models"
	"baseapi/models/util"
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetFollowers(c *gin.Context) {
	var user models.User
	var followers []models.User
	userId := c.Param("id")
	global.BA_DB.First(&user, userId)
	global.BA_DB.Model(&user).Association("Followers").Find(&followers)
	fmt.Println(followers)
	response.OkWithData(followers, c)
}

func GetFollowees(c *gin.Context) {
	var user models.User
	var followees []models.User
	userId := c.Param("id")
	global.BA_DB.First(&user, userId)
	global.BA_DB.Model(&user).Association("Followees").Find(&followees)
	fmt.Println(user)
	fmt.Println(followees)
	response.OkWithData(followees, c)
}

func GetUserTags(c *gin.Context) {
	var user models.User
	var tags []models.Tag
	userId := c.Param("id")
	global.BA_DB.First(&user, userId)
	global.BA_DB.Model(&user).Association("Tags").Find(&tags)
	response.OkWithData(tags, c)
}

func FollowUser(c *gin.Context) {
	var user, currentUser models.User
	global.BA_DB.Find(&currentUser, 237)
	userId := c.Param("id")
	global.BA_DB.Find(&user, userId)
	global.BA_DB.Model(&currentUser).Association("Followers").Append(&user)
	response.Ok(c)
}

func UnfollowUser(c *gin.Context) {
	var user, currentUser models.User
	global.BA_DB.Find(&currentUser, 237)
	userId := c.Param("id")
	global.BA_DB.Find(&user, userId)
	global.BA_DB.Model(&currentUser).Association("Followers").Delete(&user)
	response.Ok(c)
}

// @Summary 新增用户
// @Produce  json
// @Param name query string true "Name"
// @Param password query string true "Password"
// @Param is_active query int false "IsActive"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/users [post]
func AddUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println(err)
		response.FailWithMessage(err.Error(), c)
		return
	}
	user = models.AddUser(user.Name, user.Password)
	response.OkWithData(user, c)
}

func Auth(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindQuery(&user); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := models.CheckAuth(user.Name, user.Password)
	if userId <= 0 {
		response.FailAuth(c)
		return
	}

	token, err := util.GenerateToken(user.Name, userId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(map[string]interface{}{"token": token}, c)
}
