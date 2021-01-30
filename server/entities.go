package server

type getPostsRequest struct {
	Limit  *int `json:"limit" binding:"required"`
	Offset *int `json:"offset" binding:"required"`
}

type createPostRequest struct {
	Body string `json:"body" binding:"required"`
}
