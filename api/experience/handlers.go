package experience

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
func Routes(router *gin.Engine, experienceSvc ExperienceService) {
	subRouter := router.Group("/experience")
	{
		subRouter.POST("", func(c *gin.Context) {
			AddUserExperienceHandler(experienceSvc, c)
		})
		subRouter.GET("/:id", func(c *gin.Context) {
			GetUserExperienceByIdHandler(experienceSvc, c)
		})
		subRouter.DELETE("/:id", func(c *gin.Context) {
			DeleteUserExperienceByUserIdHandler(experienceSvc, c)
		})
		subRouter.PATCH("/:id", func(c *gin.Context) {
			UpdateUserExperienceByIdHandler(experienceSvc, c)
		})
	}

}

type AddUserExperienceRequest struct {
	SkillID     uint       `json:"skill_id"`
	UserID      uint       `json:"user_id" validate:"required"`
	Experiences ExpRequest `json:"experiences"`
}

type ExpRequest struct {
	Position           string    `json:"position" validate:"required"`
	Company            string    `json:"company" validate:"required"`
	Description        string    `json:"description"`
	StartDate          time.Time `json:"start_date" validate:"required"`
	EndDate            time.Time `json:"end_date"`
	IsCurrentlyWorking bool      `json:"is_currently_working"`
	Responsibilities   string    `json:"responsibilities"`
}

// AddUserExperienceHandler godoc
// @Tags experience
// @Summary Add experiences for user
// @Description Adds new experiences for a given user ID
// @ID add-experience
// @Accept json
// @Produce json
// @Param AddUserExperienceRequest body AddUserExperienceRequest true "AddUserExperienceRequest"
// @Success 200 {object} utils.ResponseMessage
// @Failure 400 {object} utils.ResponseMessage
// @Failure 404 {object} utils.ResponseMessage
// @Failure 500 {object} utils.ResponseMessage
// @Router /experience [post]
func AddUserExperienceHandler(experienceSvc ExperienceService, c *gin.Context) {
	addUserExpReq := AddUserExperienceRequest{}

	if err := c.ShouldBindJSON(&addUserExpReq); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseMessage{StatusCode: http.StatusBadRequest, Message: fmt.Sprintf("Failed to bind experience request: %v", err), Data: nil})
		return
	}

	if err := validate.Struct(&addUserExpReq); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseMessage{StatusCode: http.StatusBadRequest, Message: fmt.Sprintf("Validation failed: %v", err), Data: nil})
		return
	}

	if !addUserExpReq.Experiences.IsCurrentlyWorking && addUserExpReq.Experiences.EndDate.Before(addUserExpReq.Experiences.StartDate) {
		c.JSON(http.StatusBadRequest, utils.ResponseMessage{StatusCode: http.StatusBadRequest, Message: "End date cannot be before the start date.", Data: nil})
		return
	}

	experience := Experience{
		Position:           addUserExpReq.Experiences.Position,
		Company:            addUserExpReq.Experiences.Company,
		Description:        addUserExpReq.Experiences.Description,
		StartDate:          addUserExpReq.Experiences.StartDate,
		EndDate:            addUserExpReq.Experiences.EndDate,
		IsCurrentlyWorking: addUserExpReq.Experiences.IsCurrentlyWorking,
		Responsibilities:   addUserExpReq.Experiences.Responsibilities,
	}

	createdExperiences, err := experienceSvc.AddExperienceWithUserAndSkills(addUserExpReq.UserID, addUserExpReq.SkillID, experience)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseMessage{StatusCode: http.StatusInternalServerError, Message: fmt.Sprintf("Failed to add experiences: %v", err), Data: nil})
		return
	}

	c.JSON(http.StatusOK, utils.ResponseMessage{StatusCode: http.StatusOK, Message: "Experience added successfully.", Data: createdExperiences})
}

type UpdateExpRequest struct {
	UserId             uint      `json:"user_id"`
	Position           string    `json:"position" binding:"required"`
	Company            string    `json:"company" binding:"required"`
	Description        string    `json:"description"`
	StartDate          time.Time `json:"start_date" binding:"required"`
	EndDate            time.Time `json:"end_date"`
	IsCurrentlyWorking bool      `json:"is_currently_working"`
	Responsibilities   string    `json:"responsibilities"`
}

// UpdateUserExperienceByIdHandler godoc
// @Tags experience
// @Summary Update experience
// @Description Updates experience
// @ID update-experience
// @Accept  json
// @Produce  json
// @Param id path uint true "id"
// @Param UpdateExpRequest body UpdateExpRequest true "UpdateExpRequest"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Failure 500 {object} string
// @Router /experience/{id} [patch]
func UpdateUserExperienceByIdHandler(experienceSvc ExperienceService, c *gin.Context) {
	var updateExpRequest UpdateExpRequest

	if err := c.ShouldBindJSON(&updateExpRequest); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseMessage{StatusCode: http.StatusBadRequest, Message: fmt.Sprintf("Validation failed: %v", err), Data: nil})
		return
	}

	if err := validate.Struct(&updateExpRequest); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseMessage{StatusCode: http.StatusBadRequest, Message: fmt.Sprintf("Validation failed: %v", err), Data: nil})
		return
	}

	experienceId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseMessage{StatusCode: http.StatusBadRequest, Message: "Invalid experience ID", Data: nil})
		return
	}

	existingExperience, err := experienceSvc.GetExperienceById(uint(experienceId))
	if err != nil {
		c.JSON(http.StatusNotFound, utils.ResponseMessage{StatusCode: http.StatusNotFound, Message: "Experience not found", Data: nil})
		return
	}

	_, err = experienceSvc.GetUserExperienceByUserIdAndExperienceId(updateExpRequest.UserId, uint(experienceId))
	if err != nil {
		c.JSON(http.StatusForbidden, utils.ResponseMessage{StatusCode: http.StatusForbidden, Message: "You are not authorized to update this experience", Data: nil})
		return
	}

	_ = utils.UpdateEntity(existingExperience, updateExpRequest)
	if err = experienceSvc.UpdateExperience(existingExperience); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseMessage{StatusCode: http.StatusInternalServerError, Message: "Failed to update experience", Data: nil})
		return
	}

	c.JSON(http.StatusOK, utils.ResponseMessage{StatusCode: http.StatusOK, Message: "Experience updated successfully", Data: nil})
}

type GetExperienceRequest struct {
	UserId uint `json:"user_id"`
}

// GetUserExperienceByIdHandler godoc
// @Tags experience
// @Summary Get user experience details by id
// @Description get user experience details by id
// @ID get-user-experience-details-by-id
// @Accept  json
// @Produce  json
// @Param id path uint true "id"
// @Param GetExperienceRequest body GetExperienceRequest true "GetExperienceRequest"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Failure 500 {object} string
// @Router /experience/{id} [get]
func GetUserExperienceByIdHandler(experienceSvc ExperienceService, c *gin.Context) {
	expId := c.Param("id")
	expIdInt, _ := strconv.Atoi(expId)
	var getExpRequest GetExperienceRequest

	if err := c.ShouldBindJSON(&getExpRequest); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseMessage{StatusCode: http.StatusBadRequest, Message: fmt.Sprintf("Validation failed: %v", err), Data: nil})
		return
	}

	if err := validate.Struct(&getExpRequest); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseMessage{StatusCode: http.StatusBadRequest, Message: fmt.Sprintf("Validation failed: %v", err), Data: nil})
		return
	}
	expDetails, err := experienceSvc.GetAllUserExperienceList(uint(expIdInt), getExpRequest.UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseMessage{StatusCode: http.StatusInternalServerError, Message: fmt.Sprintf("Cannot fetch user experience against provided ID:", err), Data: nil})
		return
	}

	c.JSON(http.StatusOK, utils.ResponseMessage{StatusCode: http.StatusOK, Message: "Success", Data: expDetails})
}

// DeleteUserExperienceByUserIdHandler godoc
// @Tags experience
// @Summary Delete user experience by id
// @Description delete user experience by id
// @ID delete-user-experience-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Param GetExperienceRequest body GetExperienceRequest true "GetExperienceRequest"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Failure 500 {object} string
// @Router /experience/{id} [delete]
func DeleteUserExperienceByUserIdHandler(experienceSvc ExperienceService, c *gin.Context) {
	expId := c.Param("id")
	expIdInt, _ := strconv.Atoi(expId)
	var getExpRequest GetExperienceRequest

	if err := c.ShouldBindJSON(&getExpRequest); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseMessage{StatusCode: http.StatusBadRequest, Message: fmt.Sprintf("Validation failed: %v", err), Data: nil})
		return
	}

	if err := validate.Struct(&getExpRequest); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseMessage{StatusCode: http.StatusBadRequest, Message: fmt.Sprintf("Validation failed: %v", err), Data: nil})
		return
	}

	statusCode := http.StatusInternalServerError
	_, err := experienceSvc.GetUserExperienceByUserIdAndExperienceId(getExpRequest.UserId, uint(expIdInt))
	if err == gorm.ErrRecordNotFound {
		statusCode = http.StatusNotFound
	}
	if err != nil {
		c.JSON(statusCode, utils.ResponseMessage{StatusCode: statusCode, Message: fmt.Sprintf("Something went wrong while fetching user experience against given id: %v", err), Data: nil})
		return
	}

	err = experienceSvc.DeleteUserExperienceByID(uint(expIdInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseMessage{StatusCode: http.StatusInternalServerError, Message: fmt.Sprintf("Unable to Delete user experience against provided id:", err), Data: nil})
		return
	}

	c.JSON(http.StatusOK, utils.ResponseMessage{StatusCode: http.StatusOK, Message: "Success", Data: nil})

}
