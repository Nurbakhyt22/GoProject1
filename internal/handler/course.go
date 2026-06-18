package handler

import (
	"net/http"

	"github.com/Nurbakhyt22/GoProject1/internal/service"
	"github.com/gin-gonic/gin"
)

type CourseHandler struct {
	service *service.CourseService
}

func NewCourseHandler(service *service.CourseService) *CourseHandler {
	return &CourseHandler{service: service}
}
func (ch *CourseHandler) GetAll(c *gin.Context) {
	courses, err := ch.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}
	c.JSON(http.StatusOK, courses)
}
