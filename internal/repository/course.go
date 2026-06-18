package repository

import (
	"time"

	"github.com/Nurbakhyt22/GoProject1/internal/model"
)

type CourseRepo struct {
	coursesMap map[int]model.Course
}

func NewCourseRepo() *CourseRepo {
	repo := &CourseRepo{
		coursesMap: make(map[int]model.Course),
	}

	repo.coursesMap[1] = model.Course{ID: 1, Title: "Golang", Price: 12000, IsActive: true, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	repo.coursesMap[2] = model.Course{ID: 2, Title: "Python", Price: 10000, IsActive: true, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	repo.coursesMap[3] = model.Course{ID: 3, Title: "C++", Price: 19000, IsActive: true, CreatedAt: time.Now(), UpdatedAt: time.Now()}

	return repo
}

func (cr *CourseRepo) GetAll() ([]model.Course, error) {
	courses := make([]model.Course, 0)

	for _, course := range cr.coursesMap {
		courses = append(courses, course)
	}

	return courses, nil
}
