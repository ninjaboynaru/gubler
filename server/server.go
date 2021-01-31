package server

import (
	"github.com/gin-gonic/gin"
)

// RunServer ...
func RunServer() error {
	var err error
	var router *gin.Engine = gin.Default()

	if err != nil {
		return err
	}

	router.GET("api/ping", ping)
	router.GET("api/posts", getPosts)
	router.POST("api/posts", createPost)

	router.GET("/", staticPosts)
	router.GET("/posts", staticPosts)
	router.GET("/createpost", staticCreatePost)
	router.GET("/about", staticAbout)

	router.Static("/js", "static/js")
	router.Static("/css", "static/css")
	router.Static("/webfonts", "static/webfonts")

	err = router.Run(":8080")
	if err != nil {
		return err
	}

	return nil
}
