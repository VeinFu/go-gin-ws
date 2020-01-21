package swagger

import "go-gin-ws/models"

// UserResponse retrieve specified user
// swagger:response UserResponse
type swagUserResponse struct {
	// in:body
	Body models.NewUserResponse
}

// ResponseError is an error that is used when the request fail
// swagger:response ResponseError
type swagResponseError struct {
	// in:body
	Body models.ResponseError
}

// TestAPIResponse retrieve test result
// swagger:response TestAPIResponse
type swagTestAPIResponse struct {
	// in:body
	Body models.NewTestAPIResponse
}
