package questions

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
func Routes(router *gin.Engine, questionSvc QuestionService) {
	subRouter := router.Group("/questions")
	{
		subRouter.POST("", func(c *gin.Context) {
			HandlerToCreateQuestions(questionSvc, c)
		})
		subRouter.PATCH("/:id", middleware.AuthMiddleware(), func(c *gin.Context) {
			HandlerToUpdateQuestions(questionSvc, c)
		})
		subRouter.GET("/:id", func(c *gin.Context) {
			HandlerToGetQuestionById(questionSvc, c)
		})
		subRouter.GET("", func(c *gin.Context) {
			HandlerToGetAllQuestions(questionSvc, c)
		})
		subRouter.DELETE("/:id", middleware.AuthMiddleware(), func(c *gin.Context) {
			HandlerToDeleteQuestionById(questionSvc, c)
		})
	}
}

type QuestionWithOptionsRequest struct {
	Questions    string   `json:"questions"`
	QuestionType string   `json:"question_type"`
	OptionNames  []string `json:"option_names"`
}

type UpdateQuestionRequest struct {
	Questions    string `json:"questions"`
	QuestionType string `json:"question_type"`
}

// HandlerToCreateQuestions godoc
// @Tags questions
// @Summary add questions
// @Description add a question
// @ID add-question
// @Accept  json
// @Produce  json
// @Param QuestionWithOptionsRequest body QuestionWithOptionsRequest true "question"
// @Success 200 {object} utils.ResponseMessage
// @Failure 400 {object} utils.ResponseMessage
// @Failure 404 {object} utils.ResponseMessage
// @Failure 500 {object} utils.ResponseMessage
// @Router /questions [post]
func HandlerToCreateQuestions(questionSvc QuestionService, c *gin.Context) {
	addQuestionRequest := QuestionWithOptionsRequest{}
	if err := c.ShouldBind(&addQuestionRequest); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseMessage{StatusCode: http.StatusBadRequest, Message: fmt.Sprintf(utils.InvalidIntegerValueLimitMessage, err), Data: nil})
		return
	}

	if err := validate.Struct(addQuestionRequest); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseMessage{StatusCode: http.StatusBadRequest, Message: fmt.Sprintf(utils.RequestSchemaInvalid, err), Data: nil})
		return
	}

	question := Question{
		Questions:    addQuestionRequest.Questions,
		QuestionType: addQuestionRequest.QuestionType,
	}

	createUser, err := questionSvc.AddQuestion(&question, addQuestionRequest.OptionNames)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseMessage{StatusCode: http.StatusInternalServerError, Message: "Something went wrong while creating question.", Data: nil})
		return
	}

	c.JSON(http.StatusOK, utils.ResponseMessage{StatusCode: http.StatusOK, Message: "question created successfully.", Data: createUser})
}

// HandlerToUpdateQuestions godoc
// @Tags questions
// @Summary update questions
// @Description update a question
// @ID update-question
// @Security ApiAuthKey
// @Accept  json
// @Produce  json
// @Param id path uint true "id"
// @Param UpdateQuestionRequest body UpdateQuestionRequest true "question"
// @Success 200 {object} utils.ResponseMessage
// @Failure 400 {object} utils.ResponseMessage
// @Failure 404 {object} utils.ResponseMessage
// @Failure 500 {object} utils.ResponseMessage
// @Router /questions/{id} [patch]
func HandlerToUpdateQuestions(questionSvc QuestionService, c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)

	updateQuestionRequest := UpdateQuestionRequest{}
	if err := c.ShouldBind(&updateQuestionRequest); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseMessage{StatusCode: http.StatusBadRequest, Message: fmt.Sprintf(utils.InvalidIntegerValueLimitMessage, err), Data: nil})
		return
	}
	if err := validate.Struct(&updateQuestionRequest); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseMessage{StatusCode: http.StatusBadRequest, Message: fmt.Sprintf(utils.RequestSchemaInvalid, err), Data: nil})
		return
	}
	existingQuestionData, err := questionSvc.GetQuestionById(uint(idInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseMessage{StatusCode: http.StatusInternalServerError, Message: fmt.Sprintf("Something went wrong while fetching data against given id: %v", err), Data: nil})
		return
	}

	_ = utils.UpdateEntity(existingQuestionData, updateQuestionRequest)

	err = questionSvc.UpdateQuestionByID(existingQuestionData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseMessage{StatusCode: http.StatusInternalServerError, Message: "Failed to update question.", Data: nil})
		return
	}

	c.JSON(http.StatusOK, utils.ResponseMessage{StatusCode: http.StatusOK, Message: "question updated successfully.", Data: nil})
}

// HandlerToGetQuestionById godoc
// @Tags questions
// @Summary Get user question by id
// @Description get question by id
// @ID get-question-by-id
// @Accept  json
// @Produce  json
// @Param id path uint true "id"
// @Success 200 {object} utils.ResponseMessage
// @Failure 400 {object} utils.ResponseMessage
// @Failure 404 {object} utils.ResponseMessage
// @Failure 500 {object} utils.ResponseMessage
// @Router /questions/{id} [get]
func HandlerToGetQuestionById(questionSvc QuestionService, c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)

	questionDetails, err := questionSvc.GetQuestionById(uint(idInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseMessage{StatusCode: http.StatusInternalServerError, Message: fmt.Sprintf("Cannot fetch user question against provided ID:", err), Data: nil})
		return
	}

	c.JSON(http.StatusOK, utils.ResponseMessage{StatusCode: http.StatusOK, Message: "Success", Data: questionDetails})
}

type QuestionResponse struct {
	Total           int64      `json:"total"`
	RecordsFiltered int        `json:"records_filtered"`
	Questions       []Question `json:"questions"`
}

// HandlerToGetAllQuestions godoc
// @Tags questions
// @Summary Get all questions
// @Description get all questions
// @ID get-all-questions
// @Accept  json
// @Produce  json
// @Param   limit    query     int     false  "example - 50"     limit(int)
// @Param   offset     query     int     false  "example - 0"     offset(int)
// @Param   orderBy     query     string     false  "example - created_at desc"  orderBy(string)
// @Success 200 {object} utils.ResponseMessage
// @Failure 400 {object} utils.ResponseMessage
// @Failure 404 {object} utils.ResponseMessage
// @Failure 500 {object} utils.ResponseMessage
// @Router /questions [get]
func HandlerToGetAllQuestions(questionSvc QuestionService, c *gin.Context) {
	limit := c.Request.URL.Query().Get("limit")
	offset := c.Request.URL.Query().Get("offset")
	orderBy := c.Request.URL.Query().Get("orderBy")

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

	allQuestions, total, err := questionSvc.GetAllQuestions(limitInt, offsetInt, orderBy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseMessage{StatusCode: http.StatusInternalServerError, Message: fmt.Sprintf("cannot fetch questions:", err), Data: nil})
		return
	}

	c.JSON(http.StatusOK, utils.ResponseMessage{StatusCode: http.StatusOK, Message: "Success", Data: QuestionResponse{Total: total, Questions: allQuestions, RecordsFiltered: len(allQuestions)}})
}

// HandlerToDeleteQuestionById godoc
// @Tags questions
// @Summary Delete question by id
// @Description delete question by id
// @ID delete-question-by-id
// @Security ApiAuthKey
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} utils.ResponseMessage
// @Failure 400 {object} utils.ResponseMessage
// @Failure 404 {object} utils.ResponseMessage
// @Failure 500 {object} utils.ResponseMessage
// @Router /questions/{id} [delete]
func HandlerToDeleteQuestionById(questionSvc QuestionService, c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)

	statusCode := http.StatusInternalServerError
	_, err := questionSvc.GetQuestionById(uint(idInt))
	if err == gorm.ErrRecordNotFound {
		statusCode = http.StatusNotFound
	}
	if err != nil {
		c.JSON(statusCode, utils.ResponseMessage{StatusCode: statusCode, Message: fmt.Sprintf("Something went wrong while fetching data against given id: %v", err), Data: nil})
		return
	}

	err = questionSvc.DeleteQuestionByID(uint(idInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseMessage{StatusCode: http.StatusInternalServerError, Message: fmt.Sprintf("unable to delete question against provided id:", err), Data: nil})
		return
	}

	c.JSON(http.StatusOK, utils.ResponseMessage{StatusCode: http.StatusOK, Message: "Success", Data: nil})
}
