package service

import (
	"errors"

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
	courses, err := cs.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return courses, nil
}

func (cs *CourseService) GetById(id int) (model.Course, error) {
	return cs.repo.GetById(id)
}

func (cs *CourseService) Delete(id int) error {
	if id <= 0 {
		return errors.New("Invalid ID")
	}
	return cs.repo.Delete(id)
}
