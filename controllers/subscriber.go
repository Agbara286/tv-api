package controllers

import (
	"fmt"
	"net/http"
	"os"
	"state-tv-api/config"
	"state-tv-api/models"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/resend/resend-go/v2"
)

type SubscribeInput struct {
	Email string `json:"email" binding:"required,email"`
}

func SubscribeToNewsletter(c *gin.Context) {
	var input SubscribeInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please provide a valid email address."})
		return
	}

	cleanEmail := strings.ToLower(input.Email)
	subscriber := models.Subscriber{Email: cleanEmail}

	result := config.DB.Create(&subscriber)

	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "unique constraint") || strings.Contains(result.Error.Error(), "duplicate key") {
			c.JSON(http.StatusConflict, gin.H{"message": "You are already in the inner circle!"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save subscription."})
		return
	}

	// 👇 THE WELCOME HANDSHAKE (RESEND API) 👇

	// 1. Grab the API key from your .env file
	apiKey := os.Getenv("RESEND_API_KEY")
	client := resend.NewClient(apiKey)

	// 2. Draft the beautiful HTML Email
	htmlBody := fmt.Sprintf(`
		<div style="font-family: sans-serif; max-w-xl; margin: 0 auto; padding: 40px; background-color: #0f172a; color: white; border-radius: 12px;">
			<h1 style="color: #3b82f6; font-size: 28px; margin-bottom: 10px;">Welcome to BCOS.</h1>
			<p style="font-size: 16px; color: #cbd5e1; line-height: 1.6;">
				You are officially on the list. When critical news breaks in Oyo State, you will be the first to know.
			</p>
			<hr style="border-color: #1e293b; margin: 30px 0;" />
			<p style="font-size: 12px; color: #64748b;">
				Sent securely from the BCOS Newsroom.
			</p>
		</div>
	`)

	// 3. Configure the delivery details
	params := &resend.SendEmailRequest{
		From:    "BCOS News <onboarding@resend.dev>", // DO NOT CHANGE THIS YET
		To:      []string{cleanEmail},
		Subject: "Welcome to the Inner Circle",
		Html:    htmlBody,
	}

	// 4. Fire the email! (We do this in a goroutine so it doesn't slow down the user's webpage)
	go func() {
		_, err := client.Emails.Send(params)
		if err != nil {
			fmt.Println("Warning: Failed to send welcome email:", err)
		}
	}()
	// 👆 -------------------------------------- 👆

	// 5. Success! Tell the Next.js UI to show the checkmark.
	c.JSON(http.StatusOK, gin.H{"message": "Successfully subscribed to BCOS News!"})
}