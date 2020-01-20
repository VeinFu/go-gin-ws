package models

type UserResponse struct {
	Name string `json:"name"`
	Sex string `json:"sex"`
	Mobile string `json:"mobile"`
	Address string `json:"address"`
}

type NewUserResponse struct {
	Data UserResponse `json:"data"`
}

func GetUserResponse(resp UserResponse) *NewUserResponse {
	return &NewUserResponse{
		Data: resp,
	}
}
