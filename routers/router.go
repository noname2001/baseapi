package routers

import (
	"baseapi/global"

	"baseapi/middleware"
	v1 "baseapi/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	app := global.BA_CONFIG.App
	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(app.RunMode)
	apiv1 := r.Group("/api/v1")
	apiv1.GET("/auth", v1.Auth)
	apiv1.Use(middleware.JWT())
	{
		apiv1.GET("/tags", v1.GetTags)
		apiv1.POST("/tags", v1.AddTag)
		apiv1.PUT("/tags/:id", v1.UpdateTag)
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
	}
	{
		apiv1.POST("/users", v1.AddUser)
		apiv1.POST("/follow_user/:id", v1.FollowUser)
		apiv1.DELETE("/follow_user/:id", v1.UnfollowUser)
		apiv1.GET("/users/:id/followers", v1.GetFollowers)
		apiv1.GET("/users/:id/followees", v1.GetFollowees)
		apiv1.GET("/users/:id/tags", v1.GetUserTags)
	}
	return r
}
