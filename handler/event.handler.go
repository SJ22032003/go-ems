package handler

import (
	"net/http"
	"strconv"
	"time"

	model "github.com/SJ22032003/go-ems/models"
	service "github.com/SJ22032003/go-ems/services"
	gin "github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type EventHandler struct {}

func (e *EventHandler) GetEventById(ctx *gin.Context) {
	
	id, err := strconv.ParseInt(ctx.Params.ByName("id"), 10, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request, Please check the request parameter and try again.",
			"error":   err.Error(),
		})
		return
	}

	event, err := service.GetEventByIdService(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Could not get event, does not exist",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"event":   *event,
	})
}

func (e *EventHandler) CreateEvents(ctx *gin.Context) {
	var event model.Event

	err := ctx.ShouldBindJSON(&event)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request, Please check the request body and try again.",
			"error":   err.Error(),
		})
		return
	}

	user, ok := ctx.Get("user")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	userId, _ := user.(jwt.MapClaims)["id"].(int64)
	
	event.UserId = userId
	event.Date = time.Now().Format("2006-01-02T15:04:05.000Z")

	err = service.CreateEventService(&event)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error, could not create event",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Event Created",
		"event":   event,
	})

}

func (e *EventHandler) GetEvents(ctx *gin.Context) {
	events, err := service.GetAllEventsService()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error, could not get events",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"events":  events,
	})
}

func (e *EventHandler) UpdateEventById(ctx *gin.Context) {

	var event *model.Event
	var err error
	var id int64

	// Get the id from the request parameter
	id, err = strconv.ParseInt(ctx.Params.ByName("id"), 10, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request, Please check the request parameter and try again.",
			"error":   err.Error(),
		})
		return
	}

	user, ok := ctx.Get("user")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	userId, _ := user.(jwt.MapClaims)["id"].(int64)

	// Get the event by id
	event, err = service.GetEventByUserId(id, userId)
	if err != nil { // If there is an error, return the error
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Could not get event, does not exist or you do not have permission to access it.",
			"error":   err.Error(),
		})
		return
	}

	// Bind the request body to the event
	err = ctx.ShouldBindJSON(event)
	if err != nil { // If there is an error, return the error
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request, Please check the body and try again.",
			"error":   err.Error(),
		})
		return
	}

	err = service.UpdateEventService(event)
	if err != nil { // If there is an error, return the error
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error, could not update event",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Event Updated",
		"event":   *event,
	})

}

func (e *EventHandler) DeleteEventById(ctx *gin.Context) {
	var err error
	var id int64

	id, err = strconv.ParseInt(ctx.Params.ByName("id"), 10, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request, Please check the request parameter and try again.",
			"error":   err.Error(),
		})
		return
	}

	user, ok := ctx.Get("user")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	userId, _ := user.(jwt.MapClaims)["id"].(int64)

	_, err = service.GetEventByUserId(id, userId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Could not get event, does not exist or you do not have permission to access it.",
			"error":   err.Error(),
		})
		return
	}

	err = service.DeleteEventByIdService(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error, could not delete event",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Event Deleted",
	})

}
