package skills

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
func Routes(router *gin.Engine, skillSvc SkillService) {
	skillsRouter := router.Group("/skills")
	categoriesRouter := skillsRouter.Group("/categories")
	{
		categoriesRouter.POST("", func(c *gin.Context) {
			HandlerToCreateSkillCategories(c, skillSvc)
		})
		categoriesRouter.PATCH("/:id", func(c *gin.Context) {
			HandlerToUpdateSkillCategoryByID(c, skillSvc)
		})
		categoriesRouter.GET("", func(c *gin.Context) {
			HandlerToGetAllSkillCategories(c, skillSvc)
		})
		categoriesRouter.GET("/:id", func(c *gin.Context) {
			HandlerToGetSkillCategoryByID(c, skillSvc)
		})
		categoriesRouter.DELETE("/:id", func(c *gin.Context) {
			HandlerToDeleteSkillCategoryByID(c, skillSvc)
		})

	}
}

// HandlerToGetAllSkillCategories godoc
// @Tags Skills Categories
// @Summary Get all skill Categories
// @Description Get all skill Categories
// @ID Get-skill-categories
// @Security ApiAuthKey
// @Accept  json
// @Produce  json
// @Param   limit    query     int     false  "example - 50"     limit(int)
// @Param   offset     query     int     false  "example - 0"     offset(int)
// @Param   orderBy     query     string     false  "example - created_at desc,updated_at desc"     orderBy(string)
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Failure 500 {object} string
// @Router /skills/categories [get]
func HandlerToGetAllSkillCategories(c *gin.Context, skillSvc SkillService) {
	fmt.Println("HandlerToGetAllSkillCategories")
	baseQuery := c.Request.URL.Query()
	limit := baseQuery.Get("limit")
	offset := baseQuery.Get("offset")
	orderBy := baseQuery.Get("orderBy")

	if limit == "" {
		limit = utils.DefaultLimit
	}
	if offset == "" {
		offset = utils.DefaultOffset
	}
	if orderBy == "" {
		orderBy = utils.DefaultOrderBy
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
	skillCategoryList, totalRecords, err := skillSvc.FetchAllSkillCategories(limitInt, offsetInt, orderBy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseMessage{StatusCode: http.StatusInternalServerError, Message: fmt.Sprintf(utils.SomethingWentWrongWhileGettingSkillCategory, err), Data: nil})
		return
	}
	c.JSON(http.StatusOK, utils.ResponseMessage{StatusCode: http.StatusOK, Message: utils.Success, Data: utils.RecordsResponse{Total: totalRecords, RecordsFiltered: len(skillCategoryList), Data: skillCategoryList}})

}

// HandlerToUpdateSkillCategoryByID godoc
// @Tags Skills Categories
// @Summary Update skill category
// @Description Update skill category
// @ID update-skill-category
// @Security ApiAuthKey
// @Accept json
// @Param id path int true "Skill Category ID"
// @Param SkillCategoryUpdateRequest body SkillCategoryUpdateRequest true "Skills Categories"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Failure 500 {object} string
// @Router /skills/categories/{id} [patch]
func HandlerToUpdateSkillCategoryByID(c *gin.Context, skillSvc SkillService) {
	fmt.Println("HandlerToUpdateSkillCategoryByID")
	skillCategoryID := c.Param("id")
	skillCategoryIDInt, err := strconv.Atoi(skillCategoryID)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseMessage{StatusCode: http.StatusBadRequest, Message: fmt.Sprintf(utils.SomethingWentWrong, err), Data: nil})
		return
	}

	var skillCategoryUpdateRequest SkillCategoryUpdateRequest
	if err := c.ShouldBind(&skillCategoryUpdateRequest); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseMessage{StatusCode: http.StatusBadRequest, Message: fmt.Sprintf(utils.InvalidJsonBody, err), Data: nil})
		return
	}
	if err := validate.Struct(skillCategoryUpdateRequest); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseMessage{StatusCode: http.StatusBadRequest, Message: fmt.Sprintf(utils.RequestSchemaInvalid, err), Data: nil})
		return
	}
	fetchedSkillCategory, err := skillSvc.GetSkillCategoryById(uint(skillCategoryIDInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseMessage{StatusCode: http.StatusInternalServerError, Message: fmt.Sprintf(utils.SomethingWentWrongWhileGettingSkillCategory, err), Data: nil})
		return
	}
	utils.UpdateEntity(&fetchedSkillCategory, skillCategoryUpdateRequest)
	err = skillSvc.UpdateSkillCategory(fetchedSkillCategory)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseMessage{StatusCode: http.StatusInternalServerError, Message: fmt.Sprintf(utils.SomethingWentWrongWhileUpdatingSkillCategory, err), Data: nil})
		return
	}
	c.JSON(http.StatusOK, utils.ResponseMessage{StatusCode: http.StatusOK, Message: utils.SuccessfullyUpdatedSkillsCategory, Data: nil})

}

type SkillCategoryUpdateRequest struct {
	Name string `json:"name" validate:"required"`
}

// HandlerToDeleteSkillCategoryByID godoc
// @Tags Skills Categories
// @Summary Delete skill category
// @Description Delete skill category
// @ID delete-skill-category
// @Security ApiAuthKey
// @Accept  json
// @Param id path int true "Skill Category ID"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Failure 500 {object} string
// @Router /skills/categories/{id} [delete]
func HandlerToDeleteSkillCategoryByID(c *gin.Context, skillSvc SkillService) {
	fmt.Println("HandlerToDeleteSkillCategoryByID")
	skillCategoryID := c.Param("id")
	skillCategoryIDInt, err := strconv.Atoi(skillCategoryID)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseMessage{StatusCode: http.StatusBadRequest, Message: fmt.Sprintf(utils.SomethingWentWrong, err), Data: nil})
		return
	}
	err = skillSvc.DeleteSkillCategoryById(uint(skillCategoryIDInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseMessage{StatusCode: http.StatusInternalServerError, Message: fmt.Sprintf(utils.SomethingWentWrongWhileDeletingSkillCategories, err), Data: nil})
		return
	}

	c.JSON(http.StatusOK, utils.ResponseMessage{StatusCode: http.StatusOK, Message: utils.SuccessfullyDeletedSkillsCategory, Data: nil})

}

// HandlerToGetSkillCategoryByID godoc
// @Tags Skills Categories
// @Summary Get skill category
// @Description Get skill category
// @ID get-skill-category
// @Security ApiAuthKey
// @Accept  json
// @Param id path int true "Skill Category ID"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Failure 500 {object} string
// @Router /skills/categories/{id} [get]
func HandlerToGetSkillCategoryByID(c *gin.Context, skillSvc SkillService) {
	fmt.Println("HandlerToGetSkillCategoryByID")
	skillCategoryID := c.Param("id")
	skillCategoryIDInt, err := strconv.Atoi(skillCategoryID)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseMessage{StatusCode: http.StatusBadRequest, Message: fmt.Sprintf(utils.SomethingWentWrong, err), Data: nil})
		return
	}
	fetchedSkillCategory, err := skillSvc.GetSkillCategoryById(uint(skillCategoryIDInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseMessage{StatusCode: http.StatusInternalServerError, Message: fmt.Sprintf(utils.SomethingWentWrongWhileGettingSkillCategory, err), Data: nil})
		return
	}

	c.JSON(http.StatusOK, utils.ResponseMessage{StatusCode: http.StatusOK, Message: utils.Success, Data: fetchedSkillCategory})
}

// HandlerToCreateSkillCategories godoc
// @Tags Skills Categories
// @Summary Create skill Categories
// @Description Create skill Categories
// @ID Create-skill-categories
// @Security ApiAuthKey
// @Accept json
// @Produce json
// @Param CreateSkillCategoryRequest body CreateSkillCategoryRequest true "Skills Categories"
// @Success 200 {object} utils.ResponseMessage
// @Failure 400 {object} utils.ResponseMessage
// @Failure 404 {object} utils.ResponseMessage
// @Failure 500 {object} utils.ResponseMessage
// @Router /skills/categories [post]
func HandlerToCreateSkillCategories(c *gin.Context, skillSvc SkillService) {
	fmt.Println("HandlerToCreateSkillCategory")
	var createSkillCategoryRequest CreateSkillCategoryRequest
	if err := c.ShouldBind(&createSkillCategoryRequest); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseMessage{StatusCode: http.StatusBadRequest, Message: fmt.Sprintf(utils.InvalidJsonBody, err), Data: nil})
		return
	}

	if err := validate.Struct(createSkillCategoryRequest); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseMessage{StatusCode: http.StatusBadRequest, Message: fmt.Sprintf(utils.RequestSchemaInvalid, err), Data: nil})
		return
	}
	if len(createSkillCategoryRequest.Name) == 0 {
		c.JSON(http.StatusBadRequest, utils.ResponseMessage{StatusCode: http.StatusBadRequest, Message: fmt.Sprintf(utils.RequiredSkillCategoryNames), Data: nil})
		return
	}
	var skillsCategories []SkillCategory
	for _, skillCategoryName := range createSkillCategoryRequest.Name {
		skillsCategories = append(skillsCategories, SkillCategory{Name: skillCategoryName})
	}
	err := skillSvc.CreateSkillCategories(skillsCategories)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseMessage{StatusCode: http.StatusInternalServerError, Message: fmt.Sprintf(utils.SomethingWentWrongWhileCreatingSkillCategories, err), Data: nil})
		return
	}
	c.JSON(http.StatusOK, utils.ResponseMessage{StatusCode: http.StatusOK, Message: fmt.Sprintf(utils.SuccessfullyCreatedSkillsCategories), Data: nil})
}

type CreateSkillCategoryRequest struct {
	Name []string `json:"name" validate:"required"`
}
