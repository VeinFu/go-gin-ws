package models

import "fmt"

type UserResponse struct {
	// user name
	Name string `json:"name"`
	// user sex
	Sex string `json:"sex"`
	// user mobile phone number
	Mobile string `json:"mobile"`
	// user current living address
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

type NewTestAPIResponse struct {
	Description string `json:"desc"`
}

type TestAPIResponse struct {
	Name string `json:"name"`
	APIFeature string `json:"api_feature"`
}

func GetTestAPIResponse(resp TestAPIResponse) *NewTestAPIResponse {
	ret := fmt.Sprintf("Welcome %s for %s test.", resp.Name, resp.APIFeature)
	return &NewTestAPIResponse{
		Description: ret,
	}
}
