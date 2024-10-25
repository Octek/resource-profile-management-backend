package bookings

import (
	"fmt"
	"github.com/Octek/resource-profile-management-backend.git/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
	"time"
)

var validate = validator.New()

// Routes Exports all routes handled by this service
func Routes(router *gin.Engine, bookingSvc BookingService) {
	subRouter := router.Group("/bookings")
	{
		subRouter.POST("", func(c *gin.Context) {
			AddUserBookingHandler(bookingSvc, c)
		})
		subRouter.GET("/:id", func(c *gin.Context) {
			GetUserBookingByIdHandler(bookingSvc, c)
		})
		subRouter.DELETE("/:id", func(c *gin.Context) {
			DeleteUserBookingByIdHandler(bookingSvc, c)
		})
		subRouter.PATCH("/:id", func(c *gin.Context) {
			UpdateUserBookingByIdHandler(bookingSvc, c)
		})
		subRouter.DELETE("/user/:id", func(c *gin.Context) {
			DeleteUserBookingByUserIdHandler(bookingSvc, c)
		})
		subRouter.GET("/user/:id", func(c *gin.Context) {
			HandlerToGetAllUserBookings(bookingSvc, c)
		})
	}

}

type AddBookingRequest struct {
	BookingRequest   BookingRequest `json:"booking_request" validate:"required"`
	SkillID          uint           `json:"skill_id"`
	QuestionOptionID uint           `json:"question_option_id"`
}

type BookingRequest struct {
	UserID          uint      `json:"user_id" validate:"required"`
	BookingDateTime time.Time `json:"booking_date_time" validate:"required"`
	MeetingLink     string    `json:"meeting_link" validate:"required"`
}

type UpdateUserBookingRequest struct {
	BookingDateTime time.Time `json:"booking_date_time" validate:"required"`
	MeetingLink     string    `json:"meeting_link" validate:"required"`
}

// AddUserBookingHandler godoc
// @Tags Booking
// @Summary Create bookings
// @Description Create bookings
// @ID Create-bookings
// @Accept json
// @Produce json
// @Param AddBookingRequest body AddBookingRequest true "Booking"
// @Success 200 {object} utils.ResponseMessage
// @Failure 400 {object} utils.ResponseMessage
// @Failure 404 {object} utils.ResponseMessage
// @Failure 500 {object} utils.ResponseMessage
// @Router /bookings [post]
func AddUserBookingHandler(bookingSvc BookingService, c *gin.Context) {
	var addBookingRequest AddBookingRequest
	if err := c.ShouldBind(&addBookingRequest); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseMessage{StatusCode: http.StatusBadRequest, Message: fmt.Sprintf(utils.InvalidJsonBody, err), Data: nil})
		return
	}

	if err := validate.Struct(addBookingRequest); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseMessage{StatusCode: http.StatusBadRequest, Message: fmt.Sprintf(utils.RequestSchemaInvalid, err), Data: nil})
		return
	}
	bookingObj := Booking{
		UserID:          addBookingRequest.BookingRequest.UserID,
		BookingDateTime: addBookingRequest.BookingRequest.BookingDateTime,
		MeetingLink:     addBookingRequest.BookingRequest.MeetingLink,
	}
	booking, err := bookingSvc.AddBooking(&bookingObj, addBookingRequest.QuestionOptionID, addBookingRequest.SkillID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseMessage{StatusCode: http.StatusInternalServerError, Message: fmt.Sprintf(utils.SomethingWentWrongWhileAddingBooking, err), Data: nil})
		return
	}
	c.JSON(http.StatusOK, utils.ResponseMessage{StatusCode: http.StatusOK, Message: fmt.Sprintf(utils.SuccessfullyAddedBooking), Data: booking})
}

// UpdateUserBookingByIdHandler godoc
// @Tags Booking
// @Summary Update booking
// @Description Update booking
// @ID update-booking
// @Accept json
// @Param id path int true "booking ID"
// @Param userId query uint true "userId"
// @Param UpdateUserBookingRequest body UpdateUserBookingRequest true "booking"
// @Success 200 {object} utils.ResponseMessage
// @Failure 400 {object} utils.ResponseMessage
// @Failure 404 {object} utils.ResponseMessage
// @Failure 500 {object} utils.ResponseMessage
// @Router /bookings/{id} [patch]
func UpdateUserBookingByIdHandler(bookingSvc BookingService, c *gin.Context) {
	userId := c.Request.URL.Query().Get("userId")
	userIdInt, _ := strconv.Atoi(userId)
	var updateUserBookingRequest UpdateUserBookingRequest

	if err := c.ShouldBindJSON(&updateUserBookingRequest); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseMessage{StatusCode: http.StatusBadRequest, Message: fmt.Sprintf(utils.InvalidJsonBody, err), Data: nil})
		return
	}

	if err := validate.Struct(&updateUserBookingRequest); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseMessage{StatusCode: http.StatusBadRequest, Message: fmt.Sprintf(utils.RequestSchemaInvalid, err), Data: nil})
		return
	}

	bookingId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseMessage{StatusCode: http.StatusBadRequest, Message: "Invalid Booking ID", Data: nil})
		return
	}

	existingBooking, err := bookingSvc.GetBookingById(uint(bookingId))
	if err != nil {
		c.JSON(http.StatusNotFound, utils.ResponseMessage{StatusCode: http.StatusNotFound, Message: "Booking not found", Data: nil})
		return
	}

	_, err = bookingSvc.GetUserBookingByUserIdAndBookingId(uint(userIdInt), uint(bookingId))
	if err != nil {
		c.JSON(http.StatusForbidden, utils.ResponseMessage{StatusCode: http.StatusForbidden, Message: "You are not authorized to update this Booking", Data: nil})
		return
	}

	_ = utils.UpdateEntity(existingBooking, updateUserBookingRequest)
	if err = bookingSvc.UpdateBooking(existingBooking); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseMessage{StatusCode: http.StatusInternalServerError, Message: "Failed to update Booking", Data: nil})
		return
	}

	c.JSON(http.StatusOK, utils.ResponseMessage{StatusCode: http.StatusOK, Message: "Booking updated successfully", Data: nil})
}

// GetUserBookingByIdHandler godoc
// @Tags Booking
// @Summary Get user booking details by id
// @Description get user booking details by id
// @ID get-user-booking-details-by-id
// @Accept  json
// @Produce  json
// @Param id path uint true "id"
// @Param userId query uint true "userId"
// @Success 200 {object} utils.ResponseMessage
// @Failure 400 {object} utils.ResponseMessage
// @Failure 404 {object} utils.ResponseMessage
// @Failure 500 {object} utils.ResponseMessage
// @Router /bookings/{id} [get]
func GetUserBookingByIdHandler(bookingSvc BookingService, c *gin.Context) {
	projectId := c.Param("id")
	projectIdInt, _ := strconv.Atoi(projectId)
	userId := c.Request.URL.Query().Get("userId")
	userIdInt, _ := strconv.Atoi(userId)

	expDetails, err := bookingSvc.GetAllUserBookingList(uint(projectIdInt), uint(userIdInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseMessage{StatusCode: http.StatusInternalServerError, Message: fmt.Sprintf("cannot fetch user booking against provided ID:", err), Data: nil})
		return
	}

	c.JSON(http.StatusOK, utils.ResponseMessage{StatusCode: http.StatusOK, Message: "Success", Data: expDetails})
}

// HandlerToGetAllUserBookings godoc
// @Tags Booking
// @Summary Get all user booking
// @Description Get all user booking
// @ID Get-all-user-booking
// @Accept  json
// @Produce  json
// @Param   limit    query     int     false  "example - 50"     limit(int)
// @Param   offset     query     int     false  "example - 0"     offset(int)
// @Param   orderBy     query     string     false  "example - created_at desc,updated_at desc"    orderBy(string)
// @Param id path int true "id"
// @Success 200 {object} utils.ResponseMessage
// @Failure 400 {object} utils.ResponseMessage
// @Failure 404 {object} utils.ResponseMessage
// @Failure 500 {object} utils.ResponseMessage
// @Router /bookings/user/{id} [get]
func HandlerToGetAllUserBookings(bookingSvc BookingService, c *gin.Context) {
	baseQuery := c.Request.URL.Query()
	limit := baseQuery.Get("limit")
	offset := baseQuery.Get("offset")
	orderBy := baseQuery.Get("orderBy")
	userId := c.Param("id")
	userIdInt, _ := strconv.Atoi(userId)
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
	bookingList, totalRecords, err := bookingSvc.GetAllUserBooking(uint(userIdInt), limitInt, offsetInt, orderBy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseMessage{StatusCode: http.StatusInternalServerError, Message: fmt.Sprintf(utils.SomethingWentWrongWhileGettingBooking, err), Data: nil})
		return
	}
	c.JSON(http.StatusOK, utils.ResponseMessage{StatusCode: http.StatusOK, Message: utils.Success, Data: utils.RecordsResponse{Total: int64(totalRecords), RecordsFiltered: len(bookingList), Data: bookingList}})
}

// DeleteUserBookingByIdHandler godoc
// @Tags Booking
// @Summary Delete user booking by id
// @Description delete user booking by id
// @ID delete-user-booking-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} utils.ResponseMessage
// @Failure 400 {object} utils.ResponseMessage
// @Failure 404 {object} utils.ResponseMessage
// @Failure 500 {object} utils.ResponseMessage
// @Router /bookings/{id} [delete]
func DeleteUserBookingByIdHandler(bookingSvc BookingService, c *gin.Context) {
	projectId := c.Param("id")
	projectIdInt, _ := strconv.Atoi(projectId)

	err := bookingSvc.DeleteUserBookingByID(uint(projectIdInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseMessage{StatusCode: http.StatusInternalServerError, Message: fmt.Sprintf("Unable to Delete user experience against provided id:", err), Data: nil})
		return
	}

	c.JSON(http.StatusOK, utils.ResponseMessage{StatusCode: http.StatusOK, Message: "Success", Data: nil})
}

// DeleteUserBookingByUserIdHandler godoc
// @Tags Booking
// @Summary Delete user booking by user id
// @Description delete user booking by user id
// @ID delete-booking-by-user-id
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} utils.ResponseMessage
// @Failure 400 {object} utils.ResponseMessage
// @Failure 404 {object} utils.ResponseMessage
// @Failure 500 {object} utils.ResponseMessage
// @Router /bookings/user/{id} [delete]
func DeleteUserBookingByUserIdHandler(bookingSvc BookingService, c *gin.Context) {
	userId := c.Param("id")
	userIdInt, _ := strconv.Atoi(userId)
	fmt.Println("userid", userIdInt)
	err := bookingSvc.DeleteUserBookingByUserID(uint(userIdInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseMessage{StatusCode: http.StatusInternalServerError, Message: fmt.Sprintf("Unable to Delete user booking against provided id:", err), Data: nil})
		return
	}

	c.JSON(http.StatusOK, utils.ResponseMessage{StatusCode: http.StatusOK, Message: "Success", Data: nil})
}
