package routes

import (
	"net/http"
	"rest_api/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"ok": false, "message": "Couldn't fetch events. Try again later"})
		return
	}
	context.JSON(http.StatusAccepted, gin.H{
		"ok":     true,
		"events": events,
	})
}

func getEventById(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"ok": false, "message": "Couldn't parse the event id"})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"ok": false, "message": "Couldn't get the event from the ID"})
		return
	}

	context.JSON(http.StatusAccepted, gin.H{
		"ok":    true,
		"event": event,
	})
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"ok":      false,
			"message": "Couldn't parse request data",
		})
		return
	}
	event.ID = 1
	event.UserID = 1
	event.DateTime = time.Now()

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"ok": false, "message": "Couldn't save event. Try again later"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "event created", "event": event})
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"ok":      false,
			"Message": "Couldn't parse event id!",
		})

		return
	}

	_, err = models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"ok":      false,
			"Message": "Couldn't fetch the event",
		})

		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"ok":      false,
			"Message": "Couldn't parse requested data!",
		})

		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"ok":      false,
			"Message": "Couldn't update the event",
		})

		return
	}

	context.JSON(http.StatusOK, gin.H{
		"ok":      true,
		"Message": "Event updated successfully",
	})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"ok":      false,
			"message": "Couldn't purse the id into Int",
		})

		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"ok":      false,
			"message": "Event not found",
		})

		return
	}

	err = event.Delete(eventId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"ok":      false,
			"message": "Failed to delete event",
		})

		return
	}

	context.JSON(http.StatusOK, gin.H{
		"ok":      true,
		"message": "Event deleted successfull",
	})
}
