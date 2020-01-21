// Package main Demo Web Server based on Gin
//
// This documentation describes a demo web server based on Gin APIs and code will be found under https://github.com/VeinFu/go-gin-ws
//
//     Schemes: http
//     BasePath:
//     Version: 1.0.0
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: vien.fu <vien.fu@ucloud.cn>
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rakyll/statik/fs"

	_ "go-gin-ws/cmd/swaggerui"
	"go-gin-ws/models"
	_ "go-gin-ws/models/swagger"
)

type userGetParams struct {
	Name string `uri:"name" json:"name" binding:"required"`
}

type userCreateReqParams struct {
	// User Name
	Name string `json:"name"`
	// User Sex
	Sex string `json:"sex"`
	// User Mobile Phone Number
	Mobile string `json:"mobile"`
	// User Address
	Address string `json:"address"`
}

// userCreateReq is the createUser Operation parameters
// swagger:parameters createUser
type userCreateReq struct {
	// in:body
	Body userCreateReqParams
}

// testAPIReq is the getTestAPI parameters
// swagger:parameters getTestAPI
type testAPIReq struct {
	Name string `form:"name" json:"name"`
	APIFeature string `form:"api_feature" json:"api_feature"`
}

func main() {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	// swagger:operation GET /api/v1/users/{name} users getUser
	// ---
	// summary: get a specified user
	// description: retrieve a specified user message according the user name
	// parameters:
	// - name: name
	//   in: path
	//   type: string
	//   description: User Name
	//   required: true
	//
	// responses:
	//   "200":
	//     "$ref": "#/responses/UserResponse"
	//   "400":
	//     "$ref": "#/responses/ResponseError"
	v1.GET("/users/:name", getUser)
	// swagger:route POST /api/v1/users users createUser
	//
	// create new user
	//
	// create a new user based on some user parameters
	//
	//     Responses:
	//		 200: UserResponse
	//       default: ResponseError
	v1.POST("/users", createUser)

	test := router.Group("/api/test")
	// swagger:route GET /api/test test getTestAPI
	//
	// just a test api
	//
	// a test for gin query restful api
	//
	//     Responses:
	//		 200: TestAPIResponse
	//       default: ResponseError
	test.GET("/", getTestAPI)

	swaggerUI(router)

	// router.Run(":8990")
	httpServer := &http.Server{
		Addr:              ":8990",
		Handler:           router,
	}
	log.Fatal(httpServer.ListenAndServe())
}

func getTestAPI(c *gin.Context) {
	var req testAPIReq
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.NewResponseError(err.Error()))
		return
	}

	resp := models.TestAPIResponse{
		Name:       req.Name,
		APIFeature: req.APIFeature,
	}
	c.JSON(http.StatusOK, models.GetTestAPIResponse(resp))
}

func getUser(c *gin.Context) {
	var req userGetParams
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.NewResponseError(err.Error()))
		return
	}

	resp := models.UserResponse{
		Name:    req.Name,
		Sex:     "FEMALE",
		Mobile:  "13062772964",
		Address: "Moon",
	}
	c.JSON(http.StatusOK, models.GetUserResponse(resp))
}

func createUser(c *gin.Context) {
	var req userCreateReqParams
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.NewResponseError(err.Error()))
		return
	}

	resp := models.UserResponse{
		Name:    req.Name,
		Sex:     req.Sex,
		Mobile:  req.Mobile,
		Address: req.Address,
	}
	c.JSON(http.StatusOK, models.GetUserResponse(resp))
}

func swaggerUI(e *gin.Engine) {
	statikFS, err := fs.New()
	if err != nil {
		panic(err)
	}
	staticServer := http.FileServer(statikFS)
	sh := http.StripPrefix("/api/v1/swaggerui/", staticServer)
	e.GET("/api/v1/swaggerui/*swagger", gin.WrapH(sh))
}
