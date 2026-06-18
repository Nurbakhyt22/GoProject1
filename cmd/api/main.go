package main

import (
	"fmt"
	"net/http"

	"github.com/Nurbakhyt22/GoProject1/internal/handler"
	"github.com/Nurbakhyt22/GoProject1/internal/repository"
	"github.com/Nurbakhyt22/GoProject1/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	courseRepo := repository.NewCourseRepo()
	courseService := service.NewCourseService(courseRepo)
	courseHandler := handler.NewCourseHandler(courseService)

	router := gin.New()

	router.GET("/api/courses", courseHandler.GetAll)

	srv := &http.Server{Addr: ":8080", Handler: router}

	if err := srv.ListenAndServe(); err != nil {
		fmt.Println("Server error")
	}
}
