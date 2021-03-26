package email

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/m-ryan/tool-platform/constants"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func getEmail(c *gin.Context) {
	c.String(200, "pong")
}

type PostEmailData struct {
	ToEmail  string `json:"to_email" binding:"required,email"`
	Subject  string `json:"subject" binding:"required"`
	SubTitle string `json:"subtitle" binding:"required"`
	Html     string `json:"html" binding:"required"`
}

func sendEnail(c *gin.Context) {
	var data PostEmailData
	validErr := c.BindJSON(&data)
	if validErr != nil {
		c.JSON(400, gin.H{
			"message": validErr.Error(),
		})
		return
	}
	from := mail.NewEmail(constants.FROM_EMAIL, constants.FROM_EMAIL)
	to := mail.NewEmail(data.ToEmail, data.ToEmail)
	message := mail.NewSingleEmail(from, data.Subject, to, data.SubTitle, data.Html)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		c.JSON(response.StatusCode, gin.H{
			"message": err.Error(),
		})
	} else {

		c.JSON(200, gin.H{
			"message": "success",
		})
	}
}
