package projects

import (
	"fmt"
	"github.com/Octek/resource-profile-management-backend.git/api/middleware"
	"github.com/Octek/resource-profile-management-backend.git/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

var validate = validator.New()

// Routes Exports all routes handled by this service
func Routes(router *gin.Engine, projectSvc ProjectService) {
	subRouter := router.Group("/projects")
	subRouter.POST("", func(c *gin.Context) {
		AddProjectsHandler(projectSvc, c)
	})
	subRouter.POST("/:id", func(c *gin.Context) {
		AssociateProjectWithUser(projectSvc, c)
	})
	subRouter.PATCH("/:id", middleware.AuthMiddleware(), func(c *gin.Context) {
		UpdateProjectByIdHandler(projectSvc, c)
	})
	subRouter.DELETE("/:id", middleware.AuthMiddleware(), func(c *gin.Context) {
		DeleteUserProjectByUserIdHandler(projectSvc, c)
	})
	subRouter.GET("/:id", func(c *gin.Context) {
		GetUserProjectByIdHandler(projectSvc, c)
	})
	subRouter.GET("/all/:id", func(c *gin.Context) {
		GetAllUserProjectHandler(projectSvc, c)
	})
}

type AddProjectRequest struct {
	Name         string `json:"name" validate:"required"`
	Description  string `json:"description"`
	Link         string `json:"link"`
	Technologies string `json:"technologies" validate:"required"`
}

type UpdateProjectRequest struct {
	Name         string `json:"name" validate:"required"`
	Description  string `json:"description"`
	Link         string `json:"link"`
	Technologies string `json:"technologies" validate:"required"`
}

type AddProjectsInBulk struct {
	AddProject []AddProjectRequest `json:"add_project" validate:"required"`
}

// AddProjectsHandler godoc
// @Tags projects
// @Summary Create project
// @Description create project
// @ID create-project
// @Accept  json
// @Produce  json
// @Param AddProjectsInBulk body AddProjectsInBulk true "AddProjectRequest"
// @Success 200 {object} utils.ResponseMessage
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Failure 500 {object} string
// @Router /projects [post]
func AddProjectsHandler(projectSvc ProjectService, c *gin.Context) {
	addProjectRequest := AddProjectsInBulk{}
	if err := c.ShouldBind(&addProjectRequest); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseMessage{StatusCode: http.StatusBadRequest, Message: fmt.Sprintf("Failed to create user: %v", err), Data: nil})
		return
	}
	if err := validate.Struct(addProjectRequest); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseMessage{StatusCode: http.StatusBadRequest, Message: fmt.Sprintf("Failed to create user: %v", err), Data: nil})
		return
	}

	addProject, err := projectSvc.AddUserProject(addProjectRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseMessage{StatusCode: http.StatusBadRequest, Message: fmt.Sprintf("Failed to bind user: %v", err), Data: nil})
		return
	}

	c.JSON(http.StatusOK, utils.ResponseMessage{StatusCode: http.StatusOK, Message: "project added successfully.", Data: addProject})

}

// UpdateProjectByIdHandler godoc
// @Tags projects
// @Summary Update user project
// @Description Update user project
// @ID update-user-project
// @Security ApiAuthKey
// @Accept  json
// @Produce  json
// @Param id query uint true "project id"
// @Param id path uint true "user id"
// @Param UpdateProjectRequest body UpdateProjectRequest true "UpdateProjectRequest"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Failure 500 {object} string
// @Router /projects/{id} [patch]
func UpdateProjectByIdHandler(projectSvc ProjectService, c *gin.Context) {
	projId := c.Request.URL.Query().Get("id")
	projIdInt, _ := strconv.Atoi(projId)
	userId := c.Param("id")
	userIdInt, _ := strconv.Atoi(userId)
	var updateProjectRequest UpdateProjectRequest

	if err := c.ShouldBindJSON(&updateProjectRequest); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseMessage{StatusCode: http.StatusBadRequest, Message: fmt.Sprintf("Validation failed: %v", err), Data: nil})
		return
	}

	if err := validate.Struct(&updateProjectRequest); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseMessage{StatusCode: http.StatusBadRequest, Message: fmt.Sprintf("Validation failed: %v", err), Data: nil})
		return
	}

	existingExperience, err := projectSvc.GetProjectById(uint(projIdInt))
	if err != nil {
		c.JSON(http.StatusNotFound, utils.ResponseMessage{StatusCode: http.StatusNotFound, Message: "project not found", Data: nil})
		return
	}

	statusCode := http.StatusInternalServerError
	_, err = projectSvc.GetUserProjectByUserId(uint(userIdInt), uint(projIdInt))
	if err == gorm.ErrRecordNotFound {
		statusCode = http.StatusNotFound
	}
	if err != nil {
		c.JSON(statusCode, utils.ResponseMessage{StatusCode: statusCode, Message: fmt.Sprintf("Something went wrong while fetching user project: %v", err), Data: nil})
		return
	}

	_ = utils.UpdateEntity(existingExperience, updateProjectRequest)
	if err = projectSvc.UpdateProject(existingExperience); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseMessage{StatusCode: http.StatusInternalServerError, Message: "Failed to update project", Data: nil})
		return
	}

	c.JSON(http.StatusOK, utils.ResponseMessage{StatusCode: http.StatusOK, Message: "project updated successfully", Data: nil})
}

// DeleteUserProjectByUserIdHandler godoc
// @Tags projects
// @Summary Delete user project by user id
// @Description delete user project by user id
// @ID delete-user-project-by-user-id
// @Security ApiAuthKey
// @Accept  json
// @Produce  json
// @Param id path int true "user id"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Failure 500 {object} string
// @Router /projects/{id} [delete]
func DeleteUserProjectByUserIdHandler(projectSvc ProjectService, c *gin.Context) {
	userId := c.Param("id")
	userIdInt, _ := strconv.Atoi(userId)
	err := projectSvc.DeleteUserProjectByID(uint(userIdInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseMessage{StatusCode: http.StatusInternalServerError, Message: fmt.Sprintf("Unable to delete user project: %v", err), Data: nil})
		return
	}

	c.JSON(http.StatusOK, utils.ResponseMessage{StatusCode: http.StatusOK, Message: "Success", Data: nil})
}

// GetUserProjectByIdHandler godoc
// @Tags projects
// @Summary Get user project details by user id
// @Description get user project details by user id
// @ID get-user-project-details-by-user-id
// @Accept  json
// @Produce  json
// @Param id path uint true "project id"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Failure 500 {object} string
// @Router /projects/{id} [get]
func GetUserProjectByIdHandler(projectSvc ProjectService, c *gin.Context) {
	projId := c.Param("id")
	projIdInt, _ := strconv.Atoi(projId)

	expDetails, err := projectSvc.GetProjectById(uint(projIdInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseMessage{StatusCode: http.StatusInternalServerError, Message: fmt.Sprintf("Cannot fetch user project against provided ID:", err), Data: nil})
		return
	}

	c.JSON(http.StatusOK, utils.ResponseMessage{StatusCode: http.StatusOK, Message: "Success", Data: expDetails})
}

type AllUserProjects struct {
	RecordsFiltered int       `json:"records_filtered"`
	Total           uint      `json:"total"`
	Project         []Project `json:"projects"`
}

// GetAllUserProjectHandler godoc
// @Tags projects
// @Summary Get all user project
// @Description get all user project
// @ID get-all-user-project
// @Accept  json
// @Produce  json
// @Param id path uint true "user id"
// @Param   limit    query     int     false  "example - 50"     limit(int)
// @Param   offset     query     int     false  "example - 0"     offset(int)
// @Param   orderBy     query     string     false  "example - created_at desc"  orderBy(string)
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Failure 500 {object} string
// @Router /projects/get-all/{id} [get]
func GetAllUserProjectHandler(projectSvc ProjectService, c *gin.Context) {
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

	allUserProjects, total, err := projectSvc.GetAllUserProject(uint(userIdInt), limitInt, offsetInt, orderBy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseMessage{StatusCode: http.StatusInternalServerError, Message: fmt.Sprintf("cannot fetch user projects:", err), Data: nil})
		return
	}

	c.JSON(http.StatusOK, utils.ResponseMessage{StatusCode: http.StatusOK, Message: "Success", Data: AllUserProjects{Total: total, Project: allUserProjects, RecordsFiltered: len(allUserProjects)}})
}

type AddUserProjectRequest struct {
	ProjectID []uint `json:"project_id" gorm:"NOT NULL;index:project_id"`
}

// AssociateProjectWithUser godoc
// @Tags projects
// @Summary Add user projects
// @Description Add user projects
// @ID add-user-projects
// @Security ApiAuthKey
// @Accept json
// @Param id path int true "User Id"
// @Param AddUserProjectRequest body AddUserProjectRequest true "User Project"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Failure 500 {object} string
// @Router /projects/{id} [post]
func AssociateProjectWithUser(projectSvc ProjectService, c *gin.Context) {
	userId := c.Param("id")
	userIdInt, err := strconv.Atoi(userId)
	var addUserProjectRequest AddUserProjectRequest
	if err = c.ShouldBind(&addUserProjectRequest); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseMessage{StatusCode: http.StatusBadRequest, Message: fmt.Sprintf(utils.InvalidJsonBody, err), Data: nil})
		return
	}

	if err = validate.Struct(addUserProjectRequest); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseMessage{StatusCode: http.StatusBadRequest, Message: fmt.Sprintf(utils.RequestSchemaInvalid, err), Data: nil})
		return
	}

	_, err = projectSvc.AddBulkUserProject(uint(userIdInt), addUserProjectRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseMessage{StatusCode: http.StatusInternalServerError, Message: fmt.Sprintf(utils.SomethingWentWrongWhileCreatingUserSkill, err), Data: nil})
		return
	}
	c.JSON(http.StatusOK, utils.ResponseMessage{StatusCode: http.StatusOK, Message: fmt.Sprintf(utils.SuccessfullyCreatedUserSkill), Data: nil})
}
