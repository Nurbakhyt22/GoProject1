package handler

import (
	"net/http"
	"strconv"

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

func (ch *CourseHandler) GetById(c *gin.Context) {
	courseId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course id"})
	}

	course, err := ch.service.GetById(courseId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
	}

	c.JSON(http.StatusOK, course)
}

func (ch *CourseHandler) Delete(c *gin.Context) {
	courseId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID parameter"})
		return
	}

	err = ch.service.Delete(courseId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Course deleted"})
}
