package user

import (
	"fmt"
	"github.com/Octek/resource-profile-management-backend.git/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
)

var validate = validator.New()

// Routes Exports all routes handled by this service
func Routes(router *gin.Engine, userSvc UserService) {
	subRouter := router.Group("/user")
	{
		subRouter.POST("/create-user", func(c *gin.Context) {
			CreateUserHandler(userSvc, c)
		})
		subRouter.GET("/get-all-user-list", func(c *gin.Context) {
			GetAllUsersListHandler(userSvc, c)
		})
	}
}

type CreateUserRequest struct {
	FirstName      string `json:"first_name" validate:"required"`
	LastName       string `json:"last_name" validate:"required"`
	Email          string `json:"email" validate:"required"`
	MobileNumber   string `json:"mobile_number" validate:"required"`
	Bio            string `json:"bio"`
	Location       string `json:"location"`
	VideoUrl       string `json:"video_url"`
	Certifications string `json:"certifications"`
	UserCategoryID uint   `json:"user_category_id"`
}

// CreateUserHandler godoc
// @Tags user
// @Summary Create user
// @Description creates a new complete user
// @ID create-user
// @Accept  json
// @Produce  json
// @Param CreateUserRequest body CreateUserRequest true "CreateUserRequest"
// @Success 200 {object} utils.ResponseMessage
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Failure 500 {object} string
// @Router /users/create-user [post]
func CreateUserHandler(userSvc UserService, c *gin.Context) {
	createUserRequest := CreateUserRequest{}
	if err := c.ShouldBind(&createUserRequest); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseMessage{StatusCode: http.StatusBadRequest, Message: fmt.Sprintf("Failed to create user: %v", err), Data: nil})
		return
	}

	if err := validate.Struct(createUserRequest); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseMessage{StatusCode: http.StatusBadRequest, Message: fmt.Sprintf("Failed to create user: %v", err), Data: nil})
		return
	}

	user := User{
		FirstName:      createUserRequest.FirstName,
		LastName:       createUserRequest.LastName,
		Email:          createUserRequest.Email,
		MobileNumber:   createUserRequest.MobileNumber,
		Bio:            createUserRequest.Bio,
		Location:       createUserRequest.Location,
		VideoUrl:       createUserRequest.VideoUrl,
		Certifications: createUserRequest.Certifications,
		UserCategoryID: createUserRequest.UserCategoryID,
	}
	createdOrgMetaData, err := userSvc.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseMessage{StatusCode: http.StatusInternalServerError, Message: "Something went wrong while creating user.", Data: nil})
		return
	}

	c.JSON(http.StatusOK, utils.ResponseMessage{StatusCode: http.StatusOK, Message: "user created successfully.", Data: createdOrgMetaData})
}

type OrgMetaDataContent struct {
	RecordsFiltered int    `json:"records_filtered"`
	Total           uint   `json:"total"`
	User            []User `json:"user"`
}

// GetAllUsersListHandler godoc
// @Tags user
// @Summary Get all user
// @Description get all user
// @ID get-all-user
// @Accept  json
// @Produce  json
// @Param   limit    query     int     false  "example - 50"     limit(int)
// @Param   keyword    query     string     false  "example - godjilla"     keyword(string)
// @Param   offset     query     int     false  "example - 0"     offset(int)
// @Param   orderBy     query     string     false  "example - created_at desc"  orderBy(string)
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Failure 500 {object} string
// @Router /user/get-all-user-list [get]
func GetAllUsersListHandler(userSvc UserService, c *gin.Context) {
	limit := c.Request.URL.Query().Get("limit")
	offset := c.Request.URL.Query().Get("offset")
	orderBy := c.Request.URL.Query().Get("orderBy")
	keyword := c.Request.URL.Query().Get("keyword")

	if limit == "" {
		limit = utils.DefaultLimit // default limit
	}
	if offset == "" {
		offset = utils.DefaultOffset // default offset
	}
	if orderBy == "" {
		orderBy = utils.DefaultOrderBy // default orderBy
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseMessage{StatusCode: http.StatusBadRequest, Message: fmt.Sprintf(utils.InvalidIntegerValueLimitMessage, err), Data: nil})
		return
	}
	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseMessage{StatusCode: http.StatusBadRequest, Message: fmt.Sprintf(utils.InvalidIntegerValueOffsetMessage, err), Data: nil})
		return
	}

	allUsers, total, err := userSvc.GetAllUser(keyword, limitInt, offsetInt, orderBy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseMessage{StatusCode: http.StatusInternalServerError, Message: fmt.Sprintf("Cannot fetch Users:", err), Data: nil})
		return
	}

	c.JSON(http.StatusOK, utils.ResponseMessage{StatusCode: http.StatusOK, Message: "Success", Data: OrgMetaDataContent{Total: total, User: allUsers, RecordsFiltered: len(allUsers)}})
}
