package course

import (
	"instructors-api/deliveries/controllers/common"
	"instructors-api/deliveries/middlewares"
	"instructors-api/deliveries/validators"
	_CourseRepo "instructors-api/repositories/course"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type CourseController struct {
	repo        _CourseRepo.Course
}

func NewCourseController(repository _CourseRepo.Course) *CourseController {
	return &CourseController{
		repo:        repository,
	}
}

func (ctl *CourseController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newCourse RequestCreate
		if err := c.Bind(&newCourse); err != nil || strings.TrimSpace(newCourse.CourseTitle) == "" || strings.TrimSpace(newCourse.CourseDesc) == "" || newCourse.CoursePrice == 0 {
			return c.JSON(http.StatusBadRequest, common.BadRequest("input dari user tidak sesuai, title, description, atau price tidak boleh kosong"))
		}

		if err := validators.ValidateCreateCourse(newCourse.CourseTitle, newCourse.CourseDesc); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(err.Error()))
		}

		userID := middlewares.ExtractTokenUserID(c)
		res, err := ctl.repo.Create(newCourse.ToEntityService(userID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "sukses menambahkan Service baru", ToResponseCreate(res)))
	}
}

func (ctl *CourseController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := ctl.repo.Get()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses mendapatkan semua service", res))
	}
}

func (ctl *CourseController) GetDetail() echo.HandlerFunc {
	return func(c echo.Context) error {
		ID, _ := strconv.Atoi(c.Param("id"))

		res, err := ctl.repo.GetDetail(uint(ID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses mendapatkan detail service", res))
	}
}

func (ctl *CourseController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var updateService RequestUpdate
		ID, _ := strconv.Atoi(c.Param("id"))
		if err := c.Bind(&updateService); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest("input dari user tidak sesuai"))
		}

		if err := validators.ValidateUpdateCourse(updateService.CourseTitle, updateService.CourseDesc); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(err.Error()))
		}

		res, err := ctl.repo.Update(updateService.ToEntityService(uint(ID)))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses memperbarui data service", ToResponseUpdate(res)))
	}
}

func (ctl *CourseController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		InstructorID := middlewares.ExtractTokenUserID(c)

		ID, _ := strconv.Atoi(c.Param("id"))
		err := ctl.repo.Delete(InstructorID, uint(ID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses menghapus service", err))
	}
}
