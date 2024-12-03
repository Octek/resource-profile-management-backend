package contact_us

import (
	"fmt"
	"github.com/Octek/resource-profile-management-backend.git/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"net/smtp"
	"strings"
)

var validate = validator.New()

func Routes(router *gin.Engine) {
	subRouter := router.Group("/contact-us")
	{
		subRouter.POST("", func(c *gin.Context) {
			ContactUsHandler(c)
		})
	}
}

type ContactUsRequest struct {
	Name    string `json:"name" validate:"required"`
	Email   string `json:"email" validate:"required,email"`
	Subject string `json:"subject" validate:"required"`
	Message string `json:"message" validate:"required"`
}

// ContactUsHandler godoc
// @Tags contact-us
// @Summary Add Contact Us Info
// @Description Adds new Contact Us Info
// @ID add-contact-us
// @Accept json
// @Produce json
// @Param ContactUsRequest body ContactUsRequest true "ContactUsRequest"
// @Success 200 {object} utils.ResponseMessage
// @Failure 400 {object} utils.ResponseMessage
// @Failure 404 {object} utils.ResponseMessage
// @Failure 500 {object} utils.ResponseMessage
// @Router /contact-us [post]
func ContactUsHandler(c *gin.Context) {
	contactUsReq := ContactUsRequest{}
	if err := c.ShouldBind(&contactUsReq); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseMessage{StatusCode: http.StatusBadRequest, Message: fmt.Sprintf("Failed to bind: %v", err), Data: nil})
		return
	}
	if err := validate.Struct(&contactUsReq); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseMessage{StatusCode: http.StatusBadRequest, Message: fmt.Sprintf("Validation failed: %v", err), Data: nil})
		return
	}

	err := SendEmail(contactUsReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseMessage{StatusCode: http.StatusInternalServerError, Message: fmt.Sprintf("Failed to send email: %v", err), Data: nil})
		return
	}

	c.JSON(http.StatusOK, utils.ResponseMessage{StatusCode: http.StatusOK, Message: "Email sent successfully.", Data: nil})
}

func SendEmail(req ContactUsRequest) error {
	smtpHost := utils.GetSMTPHost()
	smtpPort := utils.GetSMTPPort()
	senderEmail := utils.GetSenderEmail()
	senderPass := utils.GetSenderPassword()
	ccEmails := utils.GetCCEmails()
	toEmail := utils.GetToEmail()
	ccToggle := utils.GetCCToggle()

	emailBody := fmt.Sprintf(
		"Name: %s\nEmail: %s\nMessage: %s",
		req.Name, req.Email, req.Message,
	)

	headers := make(map[string]string)
	headers["From"] = senderEmail
	headers["To"] = toEmail

	var recipients []string
	recipients = append(recipients, toEmail)

	if ccToggle == "true" && ccEmails != "" {
		ccAddresses := strings.Split(ccEmails, ",")
		for i, addr := range ccAddresses {
			ccAddresses[i] = strings.TrimSpace(addr)
		}
		headers["Cc"] = strings.Join(ccAddresses, ",")
		recipients = append(recipients, ccAddresses...)
	}

	headers["Subject"] = req.Subject

	msg := ""
	for key, value := range headers {
		msg += fmt.Sprintf("%s: %s\r\n", key, value)
	}
	msg += "\r\n" + emailBody

	auth := smtp.PlainAuth("", senderEmail, senderPass, smtpHost)
	addr := fmt.Sprintf("%s:%d", smtpHost, smtpPort)

	err := smtp.SendMail(addr, auth, senderEmail, recipients, []byte(msg))
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}
