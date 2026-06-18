package service

import (
	"github.com/Nurbakhyt22/GoProject1/internal/model"
	"github.com/Nurbakhyt22/GoProject1/internal/repository"
)

type CourseService struct {
	repo *repository.CourseRepo
}

func NewCourseService(courseRepo *repository.CourseRepo) *CourseService {
	service := CourseService{
		repo: courseRepo,
	}

	return &service
}

func (cs *CourseService) GetAll() ([]model.Course, error) {
	return cs.repo.GetAll()
}
