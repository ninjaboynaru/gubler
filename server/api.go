package server

import (
	"gubler/db"
	"html"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const minPostLength = 2
const maxPostLength = 200

func ping(context *gin.Context) {
	context.String(http.StatusOK, "PONG")
}

func getPosts(context *gin.Context) {
	var err error
	var posts []db.Post
	var reqBody getPostsRequest

	err = context.ShouldBindJSON(&reqBody)
	if err != nil {
		context.String(http.StatusBadRequest, "Error parsing request body: "+err.Error())
		return
	}

	if *(reqBody.Limit) > 20 {
		*(reqBody.Limit) = 20
	}

	posts, err = db.GetPosts(*reqBody.Limit, *reqBody.Offset)

	if err != nil {
		log.Printf("Error getting posts Limit: %d Offset: %d Error: %s", reqBody.Limit, reqBody.Offset, err.Error())
		context.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	context.JSON(http.StatusOK, gin.H{"posts": posts})
}

func createPost(context *gin.Context) {
	var err error
	var reqBody createPostRequest

	err = context.ShouldBindJSON(&reqBody)
	if err != nil {
		context.String(http.StatusBadRequest, "Error parsing request body: "+err.Error())
		return
	}

	reqBody.Body = strings.TrimSpace(reqBody.Body)

	if len(reqBody.Body) < minPostLength {
		context.String(http.StatusBadRequest, "Post body must be at least %d characters long", minPostLength)
		return
	} else if len(reqBody.Body) > maxPostLength {
		context.String(http.StatusBadRequest, "Post body can not be longer than %d characters", maxPostLength)
		return
	}

	reqBody.Body = html.EscapeString(reqBody.Body)

	err = db.CreatePost(reqBody.Body)

	if err != nil {
		log.Printf("Error creating post Post Body: %s\nError: %s", reqBody.Body, err.Error())
		context.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	context.String(http.StatusOK, "Post created")
}
