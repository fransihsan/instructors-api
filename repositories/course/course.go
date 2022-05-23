package course

import (
	"errors"
	C "instructors-api/entities/course"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type CourseRepository struct {
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) *CourseRepository {
	return &CourseRepository{
		db: db,
	}
}

func (repo *CourseRepository) Create(newCourse C.Courses) (C.Courses, error) {
	if err := repo.db.Create(&newCourse).Error; err != nil {
		log.Warn(err)
		return C.Courses{}, errors.New("gagal membuat kursus baru")
	}
	return newCourse, nil
}

func (repo *CourseRepository) Get() ([]FormatGetCourse, error) {
	var courses []FormatGetCourse
	if rowsAffected := repo.db.Model(&C.Courses{}).Select("courses.id, courses.course_title, instructors.instructor_name, courses.course_price").Joins("inner join instructors on instructors.id = courses.istructor_id").Scan(&courses).RowsAffected; rowsAffected == 0 {
		return nil, errors.New("tidak terdapat kursus sama sekali")
	}
	return courses, nil
}

func (repo *CourseRepository) GetDetail(ID uint) (FormatGetDetailCourse, error) {
	var course FormatGetDetailCourse
	if rowsAffected := repo.db.Model(&C.Courses{}).Select("courses.id, courses.course_title, courses.course_benefit, courses.course_price, courses.course_desc, instructors.instructor_name").Joins("inner join instructors on instructors.id = courses.istructor_id").Where("courses.id = ?", ID).Scan(&course).RowsAffected; rowsAffected == 0 {
		return FormatGetDetailCourse{}, errors.New("tidak terdapat kursus sama sekali")
	}
	return course, nil
}

func (repo *CourseRepository) Update(courseUpdate C.Courses) (C.Courses, error) {
	if rowsAffected := repo.db.Model(&courseUpdate).Where("id = ? && istructor_id = ?", courseUpdate.ID, courseUpdate.IstructorID).Updates(courseUpdate).RowsAffected; rowsAffected == 0 {
		return C.Courses{}, errors.New("tidak ada data kursus yang diperbarui")
	}
	repo.db.First(&courseUpdate)
	return courseUpdate, nil
}

func (repo *CourseRepository) Delete(InstructorID, ID uint) error {
	if rowsAffected := repo.db.Delete(&C.Courses{}, "id = ? && istructor_id = ?", ID, InstructorID).RowsAffected; rowsAffected == 0 {
		return errors.New("tidak ada kursus yang dihapus")
	}
	return nil
}