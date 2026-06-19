package repository

import (
	"errors"
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

func (cr *CourseRepo) GetById(id int) (model.Course, error) {
	course, ok := cr.coursesMap[id]
	if !ok {
		return model.Course{}, errors.New("course not found")
	}

	return course, nil
}

func (cr *CourseRepo) Delete(id int) error {
	if _, ok := cr.coursesMap[id]; !ok {
		return errors.New("course not found")
	}

	delete(cr.coursesMap, id)
	return nil
}

func (cr *CourseRepo) Create(course model.Course) (int, error) {
	course.ID = len(cr.coursesMap) + 1
	cr.coursesMap[course.ID] = course

	return course.ID, nil
}

func (cr *CourseRepo) Update(id int, input model.UpdateCourseInput) (model.Course, error) {
	course, ok := cr.coursesMap[id]
	if !ok {
		return model.Course{}, errors.New("Course not found")
	}

	if input.Title != nil {
		course.Title = *input.Title
	}
	if input.Price != nil {
		course.Price = *input.Price
	}
	if input.IsActive != nil {
		course.IsActive = *input.IsActive
	}
	course.UpdatedAt = time.Now()

	cr.coursesMap[id] = course
	return course, nil
}
