package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

type ResponseMessage struct {
	Message string `json:"message"`
}

// registerForEvent godoc
// @Summary Register user for an event
// @Description Register a user for a specific event using the event ID
// @Tags events
// @Produce json
// @Param id path int true "Event ID"
// @Success 201 {object} ResponseMessage
// @Failure 400 {object} ResponseMessage "Could not parse event id."
// @Failure 500 {object} ResponseMessage "Could not fetch event."
// @Failure 500 {object} ResponseMessage "Could not register user for event."
// @Router /events/{id}/register [post]
func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}

	err = event.Register(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register user for event."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Registered!"})
}

// cancelRegistration godoc
// @Summary Cancel registration for an event
// @Description Cancel the registration of a user for a specific event using the event ID
// @Tags events
// @Produce json
// @Param id path int true "Event ID"
// @Success 200 {object} ResponseMessage
// @Failure 400 {object} ResponseMessage "Could not parse event id."
// @Failure 500 {object} ResponseMessage "Could not cancel registration."
// @Router /events/{id}/cancel [delete]
func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	var event models.Event
	event.ID = eventId

	err = event.CancelRegistration(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel registration."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Cancelled!"})
}
