package server

import (
	"github.com/gin-gonic/gin"
)

func staticPosts(context *gin.Context) {
	serveTemplate("posts", context)
}

func staticCreatePost(context *gin.Context) {
	serveTemplate("createpost", context)
}

func staticAbout(context *gin.Context) {
	serveTemplate("about", context)
}
