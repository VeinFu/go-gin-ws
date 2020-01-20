package swagger

import "go-gin-ws/models"

// UserResponse retrieve specified user
// swagger:response UserResponse
type swagUserResponse struct {
	// in:body
	Body struct {
		Data models.UserResponse `json:"data"`
	}
}

// ResponseError is an error that is used when the request fail
// swagger:response ResponseError
type swagResponseError struct {
	// in:body
	Body models.ResponseError
}
