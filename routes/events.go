package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

// ErrorResponse represents a standard error response format
type ErrorResponse struct {
	Message string `json:"message"`
}

// EventResponse represents a successful event creation/update response
type EventResponse struct {
	Message string       `json:"message"`
	Event   models.Event `json:"event"`
}

// MessageResponse represents a simple success message response
type MessageResponse struct {
	Message string `json:"message"`
}

// getEvents godoc
// @Summary Get all events
// @Description Retrieve a list of all events
// @Tags events
// @Produce json
// @Success 200 {array} models.Event
// @Failure 500 {object} ErrorResponse "Something went wrong"
// @Router /events [get]
func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Something went wrong"})
		return
	}
	context.JSON(http.StatusOK, events)
}

// getEvent godoc
// @Summary Get event by ID
// @Description Retrieve an event by its ID
// @Tags events
// @Produce json
// @Param id path int true "Event ID"
// @Success 200 {object} models.Event
// @Failure 400 {object} ErrorResponse "Invalid ID"
// @Failure 500 {object} ErrorResponse "Something went wrong"
// @Router /events/{id} [get]
func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid ID"})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Something went wrong"})
		return
	}

	context.JSON(http.StatusOK, event)
}

// createEvent godoc
// @Summary Create a new event
// @Description Create a new event
// @Tags events
// @Produce json
// @Param event body models.Event true "Event data"
// @Success 201 {object} EventResponse "Event created successfully"
// @Failure 400 {object} ErrorResponse "Invalid request data"
// @Failure 500 {object} ErrorResponse "Something went wrong"
// @Router /events [post]
func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid request data"})
		return
	}

	userId := context.GetInt64("userId")
	event.UserID = userId

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Something went wrong"})
		return
	}

	context.JSON(http.StatusCreated, EventResponse{Message: "Event created!", Event: event})
}

// updateEvent godoc
// @Summary Update an existing event
// @Description Update an event by its ID
// @Tags events
// @Produce json
// @Param id path int true "Event ID"
// @Param event body models.Event true "Updated event data"
// @Success 200 {object} MessageResponse "Event updated successfully"
// @Failure 400 {object} ErrorResponse "Invalid ID or request data"
// @Failure 401 {object} ErrorResponse "Unauthorized"
// @Failure 500 {object} ErrorResponse "Something went wrong"
// @Router /events/{id} [put]
func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid ID"})
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventByID(eventId)
	fmt.Print(event, " ", err)
	if err != nil {
		context.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Something went wrong"})
		return
	}

	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, ErrorResponse{Message: "Unauthorized"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid request data"})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Something went wrong"})
		return
	}
	context.JSON(http.StatusOK, MessageResponse{Message: "Event updated successfully!"})
}

// deleteEvent godoc
// @Summary Delete an existing event
// @Description Delete an event by its ID
// @Tags events
// @Produce json
// @Param id path int true "Event ID"
// @Success 200 {object} MessageResponse "Event deleted successfully"
// @Failure 400 {object} ErrorResponse "Invalid ID"
// @Failure 401 {object} ErrorResponse "Unauthorized"
// @Failure 500 {object} ErrorResponse "Something went wrong"
// @Router /events/{id} [delete]
func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid ID"})
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Something went wrong"})
		return
	}

	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, ErrorResponse{Message: "Unauthorized"})
		return
	}

	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Something went wrong"})
		return
	}

	context.JSON(http.StatusOK, MessageResponse{Message: "Event deleted successfully!"})
}
