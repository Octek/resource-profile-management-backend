package user

import (
	"fmt"
	"github.com/Octek/resource-profile-management-backend.git/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
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
		subRouter.GET("/get-user-details/:id", func(c *gin.Context) {
			GetUserDetailsByUserIdHandler(userSvc, c)
		})
		subRouter.DELETE("/delete-user/:id", func(c *gin.Context) {
			DeleteUserByUserIdHandler(userSvc, c)
		})
		subRouter.PATCH("/update-user/:id", func(c *gin.Context) {
			UpdateUserByUserIdHandler(userSvc, c)
		})
		subRouter.POST("/add-user-education", func(c *gin.Context) {
			AddUserEducationHandler(userSvc, c)
		})
		subRouter.PATCH("/update-user-education/:id", func(c *gin.Context) {
			UpdateUserEducationByIdHandler(userSvc, c)
		})
		subRouter.DELETE("/delete-user-education/:id", func(c *gin.Context) {
			DeleteUserEducationByUserIdHandler(userSvc, c)
		})
		subRouter.GET("/get-user-education/:id", func(c *gin.Context) {
			GetUserEducationByUserIdHandler(userSvc, c)
		})
		subRouter.GET("/get-all-user-education/:id", func(c *gin.Context) {
			GetAllUserEducationHandler(userSvc, c)
		})
	}
}

type CreateUserRequest struct {
	FirstName      string `json:"first_name" validate:"required"`
	LastName       string `json:"last_name" validate:"required"`
	Email          string `json:"email" validate:"required"`
	MobileNumber   string `json:"mobile_number"`
	UserCategoryID uint   `json:"user_category_id"`
	JobTitle       string `json:"job_title"`
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
// @Router /user/create-user [post]
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
		UserCategoryID: createUserRequest.UserCategoryID,
		JobTitle:       createUserRequest.JobTitle,
	}
	createUser, err := userSvc.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseMessage{StatusCode: http.StatusInternalServerError, Message: "Something went wrong while creating user.", Data: nil})
		return
	}

	c.JSON(http.StatusOK, utils.ResponseMessage{StatusCode: http.StatusOK, Message: "user created successfully.", Data: createUser})
}

type GetAllUsers struct {
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
	//keyword := c.Request.URL.Query().Get("keyword")

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

	allUsers, total, err := userSvc.GetAllUser("", limitInt, offsetInt, orderBy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseMessage{StatusCode: http.StatusInternalServerError, Message: fmt.Sprintf("Cannot fetch Users:", err), Data: nil})
		return
	}

	c.JSON(http.StatusOK, utils.ResponseMessage{StatusCode: http.StatusOK, Message: "Success", Data: GetAllUsers{Total: total, User: allUsers, RecordsFiltered: len(allUsers)}})
}

// GetUserDetailsByUserIdHandler godoc
// @Tags user
// @Summary Get user details by id
// @Description get user details by id
// @ID get-user-details-by-id
// @Accept  json
// @Produce  json
// @Param id path uint true "id"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Failure 500 {object} string
// @Router /user/get-user-details/{id} [get]
func GetUserDetailsByUserIdHandler(userSvc UserService, c *gin.Context) {
	userId := c.Param("id")
	userIdInt, _ := strconv.Atoi(userId)
	userDetails, err := userSvc.GetUserDetailsByUserId(uint(userIdInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseMessage{StatusCode: http.StatusInternalServerError, Message: fmt.Sprintf("Cannot fetch user against provided ID:", err), Data: nil})
		return
	}

	c.JSON(http.StatusOK, utils.ResponseMessage{StatusCode: http.StatusOK, Message: "Success", Data: userDetails})
}

// DeleteUserByUserIdHandler godoc
// @Tags user
// @Summary Delete user by id
// @Description delete user by id
// @ID delete-user-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Failure 500 {object} string
// @Router /user/delete-user{id} [delete]
func DeleteUserByUserIdHandler(userSvc UserService, c *gin.Context) {
	userId := c.Param("id")
	userIdInt, _ := strconv.Atoi(userId)

	statusCode := http.StatusInternalServerError
	_, err := userSvc.GetUserDetailsByUserId(uint(userIdInt))
	if err == gorm.ErrRecordNotFound {
		statusCode = http.StatusNotFound
	}
	if err != nil {
		c.JSON(statusCode, utils.ResponseMessage{StatusCode: statusCode, Message: fmt.Sprintf("Something went wrong while fetching data against given id: %v", err), Data: nil})
		return
	}

	err = userSvc.DeleteUserByUserID(uint(userIdInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseMessage{StatusCode: http.StatusInternalServerError, Message: fmt.Sprintf("Unable to Delete user against provided id:", err), Data: nil})
		return
	}

	c.JSON(http.StatusOK, utils.ResponseMessage{StatusCode: http.StatusOK, Message: "Success", Data: nil})
}

type UpdateUser struct {
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email"`
	MobileNumber   string `json:"mobile_number"`
	Bio            string `json:"bio"`
	Location       string `json:"location"`
	VideoUrl       string `json:"video_url"`
	Certifications string `json:"certifications"`
	UserCategoryID uint   `json:"user_category_id"`
}

// UpdateUserByUserIdHandler godoc
// @Tags user
// @Summary Update user
// @Description Updates user
// @ID update-user
// @Accept  json
// @Produce  json
// @Param id path uint true "id"
// @Param UpdateUser body UpdateUser true "UpdateUser"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Failure 500 {object} string
// @Router /user/update-user/{id} [patch]
func UpdateUserByUserIdHandler(userSvc UserService, c *gin.Context) {
	userId := c.Param("id")
	userIdInt, _ := strconv.Atoi(userId)

	updateUserRequest := UpdateUser{}
	if err := c.ShouldBind(&updateUserRequest); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseMessage{StatusCode: http.StatusBadRequest, Message: fmt.Sprintf("Failed to bind user: %v", err), Data: nil})
		return
	}
	if err := validate.Struct(&updateUserRequest); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseMessage{StatusCode: http.StatusBadRequest, Message: fmt.Sprintf("Validation failed: %v", err), Data: nil})
		return
	}
	existingUserData, err := userSvc.GetUserDetailsByUserId(uint(userIdInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseMessage{StatusCode: http.StatusInternalServerError, Message: fmt.Sprintf("Something went wrong while fetching data against given id: %v", err), Data: nil})
		return
	}

	_ = utils.UpdateEntity(existingUserData, updateUserRequest)

	updatedUser, err := userSvc.UpdateUserByUserID(existingUserData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseMessage{StatusCode: http.StatusInternalServerError, Message: "Failed to update user.", Data: nil})
		return
	}

	c.JSON(http.StatusOK, utils.ResponseMessage{StatusCode: http.StatusOK, Message: "User updated successfully.", Data: updatedUser})
}

type AddUserEducation struct {
	UserID          uint      `json:"user_id" validate:"required"`
	InstitutionName string    `json:"institution_name" validate:"required"`
	Degree          string    `json:"degree"`
	FieldOfStudy    string    `json:"field_of_study"`
	Achievements    string    `json:"achievements"`
	StartDate       time.Time `json:"start_date" validate:"required"`
	EndDate         time.Time `json:"end_date" validate:"required"`
}

type UpdateUserEducation struct {
	InstitutionName string    `json:"institution_name" validate:"required"`
	Degree          string    `json:"degree"`
	FieldOfStudy    string    `json:"field_of_study"`
	Achievements    string    `json:"achievements"`
	StartDate       time.Time `json:"start_date" validate:"required"`
	EndDate         time.Time `json:"end_date"`
}

// AddUserEducationHandler godoc
// @Tags user
// @Summary add user education
// @Description add user education
// @ID add-user-education
// @Accept  json
// @Produce  json
// @Param AddUserEducation body AddUserEducation true "AddUserEducation"
// @Success 200 {object} utils.ResponseMessage
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Failure 500 {object} string
// @Router /user/add-user-education [post]
func AddUserEducationHandler(userSvc UserService, c *gin.Context) {
	addUserEducationReq := AddUserEducation{}

	if err := c.ShouldBindJSON(&addUserEducationReq); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseMessage{StatusCode: http.StatusBadRequest, Message: fmt.Sprintf("Failed to bind education request: %v", err), Data: nil})
		return
	}

	if err := validate.Struct(&addUserEducationReq); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseMessage{StatusCode: http.StatusBadRequest, Message: fmt.Sprintf("Validation failed: %v", err), Data: nil})
		return
	}

	statusCode := http.StatusInternalServerError
	_, err := userSvc.GetUserEducationByUserId(addUserEducationReq.UserID)
	if err == gorm.ErrRecordNotFound {
		statusCode = http.StatusNotFound
	}
	if err != nil {
		c.JSON(statusCode, utils.ResponseMessage{StatusCode: statusCode, Message: fmt.Sprintf("Something went wrong while fetching user education: %v", err), Data: nil})
		return
	}

	if addUserEducationReq.EndDate.Before(addUserEducationReq.StartDate) {
		c.JSON(http.StatusBadRequest, utils.ResponseMessage{StatusCode: http.StatusBadRequest, Message: "End date cannot be before the start date.", Data: nil})
		return
	}

	education := Education{
		UserID:          addUserEducationReq.UserID,
		InstitutionName: addUserEducationReq.InstitutionName,
		Degree:          addUserEducationReq.Degree,
		FieldOfStudy:    addUserEducationReq.FieldOfStudy,
		Achievements:    addUserEducationReq.Achievements,
		StartDate:       addUserEducationReq.StartDate,
		EndDate:         addUserEducationReq.EndDate,
	}

	createdExperiences, err := userSvc.AddUserEducation(education)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseMessage{StatusCode: http.StatusInternalServerError, Message: fmt.Sprintf("Failed to add education: %v", err), Data: nil})
		return
	}

	c.JSON(http.StatusOK, utils.ResponseMessage{StatusCode: http.StatusOK, Message: "education added successfully.", Data: createdExperiences})

}

// UpdateUserEducationByIdHandler godoc
// @Tags user
// @Summary Update user education
// @Description Update user education
// @ID update-user-education
// @Accept  json
// @Produce  json
// @Param id path uint true "id"
// @Param userId query uint true "userId"
// @Param UpdateUserEducation body UpdateUserEducation true "UpdateUserEducation"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Failure 500 {object} string
// @Router /user/update-user-education/{id} [patch]
func UpdateUserEducationByIdHandler(userSvc UserService, c *gin.Context) {
	eduId := c.Param("id")
	eduIdInt, _ := strconv.Atoi(eduId)
	userId := c.Request.URL.Query().Get("userId")
	userIdInt, _ := strconv.Atoi(userId)
	var updateEduRequest UpdateUserEducation

	if err := c.ShouldBindJSON(&updateEduRequest); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseMessage{StatusCode: http.StatusBadRequest, Message: fmt.Sprintf("Validation failed: %v", err), Data: nil})
		return
	}

	if err := validate.Struct(&updateEduRequest); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseMessage{StatusCode: http.StatusBadRequest, Message: fmt.Sprintf("Validation failed: %v", err), Data: nil})
		return
	}

	existingExperience, err := userSvc.GetEducationById(uint(eduIdInt))
	if err != nil {
		c.JSON(http.StatusNotFound, utils.ResponseMessage{StatusCode: http.StatusNotFound, Message: "Education not found", Data: nil})
		return
	}

	statusCode := http.StatusInternalServerError
	_, err = userSvc.GetUserEducationByUserId(uint(userIdInt))
	if err == gorm.ErrRecordNotFound {
		statusCode = http.StatusNotFound
	}
	if err != nil {
		c.JSON(statusCode, utils.ResponseMessage{StatusCode: statusCode, Message: fmt.Sprintf("Something went wrong while fetching user education: %v", err), Data: nil})
		return
	}

	_ = utils.UpdateEntity(existingExperience, updateEduRequest)
	if err = userSvc.UpdateEducation(existingExperience); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseMessage{StatusCode: http.StatusInternalServerError, Message: "Failed to update Education", Data: nil})
		return
	}

	c.JSON(http.StatusOK, utils.ResponseMessage{StatusCode: http.StatusOK, Message: "Education updated successfully", Data: nil})
}

// DeleteUserEducationByUserIdHandler godoc
// @Tags user
// @Summary Delete user education by user id
// @Description delete user education by user id
// @ID delete-user-education-by-user-id
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Failure 500 {object} string
// @Router /user/delete-user-education/{id} [delete]
func DeleteUserEducationByUserIdHandler(userSvc UserService, c *gin.Context) {
	userId := c.Param("id")
	userIdInt, _ := strconv.Atoi(userId)
	statusCode := http.StatusInternalServerError
	_, err := userSvc.GetUserEducationByUserId(uint(userIdInt))
	if err == gorm.ErrRecordNotFound {
		statusCode = http.StatusNotFound
	}
	if err != nil {
		c.JSON(statusCode, utils.ResponseMessage{StatusCode: statusCode, Message: fmt.Sprintf("Something went wrong while fetching user education: %v", err), Data: nil})
		return
	}

	err = userSvc.DeleteUserEducationByID(uint(userIdInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseMessage{StatusCode: http.StatusInternalServerError, Message: fmt.Sprintf("Unable to delete user education: %v", err), Data: nil})
		return
	}

	c.JSON(http.StatusOK, utils.ResponseMessage{StatusCode: http.StatusOK, Message: "Success", Data: nil})
}

// GetUserEducationByUserIdHandler godoc
// @Tags user
// @Summary Get user education details by user id
// @Description get user education details by user id
// @ID get-user-education-details-by-user-id
// @Accept  json
// @Produce  json
// @Param id path uint true "id"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Failure 500 {object} string
// @Router /user/get-user-education/{id} [get]
func GetUserEducationByUserIdHandler(userSvc UserService, c *gin.Context) {
	userId := c.Param("id")
	userIdInt, _ := strconv.Atoi(userId)

	expDetails, err := userSvc.GetUserEducationByUserId(uint(userIdInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseMessage{StatusCode: http.StatusInternalServerError, Message: fmt.Sprintf("Cannot fetch user education against provided ID:", err), Data: nil})
		return
	}

	c.JSON(http.StatusOK, utils.ResponseMessage{StatusCode: http.StatusOK, Message: "Success", Data: expDetails})
}

type GetAllUserEducation struct {
	RecordsFiltered int         `json:"records_filtered"`
	Total           uint        `json:"total"`
	Education       []Education `json:"education"`
}

// GetAllUserEducationHandler godoc
// @Tags user
// @Summary Get all user education
// @Description get all user education
// @ID get-all-user-education
// @Accept  json
// @Produce  json
// @Param id path uint true "id"
// @Param   limit    query     int     false  "example - 50"     limit(int)
// @Param   offset     query     int     false  "example - 0"     offset(int)
// @Param   orderBy     query     string     false  "example - created_at desc"  orderBy(string)
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Failure 500 {object} string
// @Router /user/get-all-user-education/{id} [get]
func GetAllUserEducationHandler(userSvc UserService, c *gin.Context) {
	limit := c.Request.URL.Query().Get("limit")
	offset := c.Request.URL.Query().Get("offset")
	orderBy := c.Request.URL.Query().Get("orderBy")
	userId := c.Param("id")
	userIdInt, _ := strconv.Atoi(userId)
	fmt.Println("userID", userIdInt)

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

	_, err = userSvc.GetUserEducationByUserId(uint(userIdInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseMessage{StatusCode: http.StatusInternalServerError, Message: fmt.Sprintf("Cannot fetch user education against provided ID:", err), Data: nil})
		return
	}

	allUserEducation, total, err := userSvc.GetAllUserEducation(uint(userIdInt), limitInt, offsetInt, orderBy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseMessage{StatusCode: http.StatusInternalServerError, Message: fmt.Sprintf("Cannot fetch Users:", err), Data: nil})
		return
	}

	c.JSON(http.StatusOK, utils.ResponseMessage{StatusCode: http.StatusOK, Message: "Success", Data: GetAllUserEducation{Total: total, Education: allUserEducation, RecordsFiltered: len(allUserEducation)}})
}
