package instructor

import (
	"instructors-api/deliveries/controllers/common"
	"instructors-api/deliveries/middlewares"
	"instructors-api/deliveries/validators"
	_InstructorRepo "instructors-api/repositories/instructor"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type InstructorController struct {
	repo _InstructorRepo.Instructor
}

func NewInstructorController(repository _InstructorRepo.Instructor) *InstructorController {
	return &InstructorController{
		repo: repository,
	}
}

func (ctl *InstructorController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newInstructor RequestCreateInstructor

		if err := c.Bind(&newInstructor); err != nil || strings.TrimSpace(newInstructor.InstructorName) == "" || strings.TrimSpace(newInstructor.Email) == "" || strings.TrimSpace(newInstructor.Password) == "" {
			return c.JSON(http.StatusBadRequest, common.BadRequest("input dari user tidak sesuai, nama, email atau password tidak boleh kosong"))
		}

		if err := validators.ValidateCreateInstructor(newInstructor.InstructorName, newInstructor.Email, newInstructor.Password); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(err.Error()))
		}

		res, err := ctl.repo.Create(newInstructor.ToEntityInstructor())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "sukses menambahkan user baru", ToResponseCreateInstructor(res)))
	}
}

func (ctl *InstructorController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := middlewares.ExtractTokenUserID(c)
		var updateInstructor RequestUpdateInstructor

		if err := c.Bind(&updateInstructor); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest("terdapat kesalahan input dari client"))
		}

		if err := validators.ValidateUpdateInstructor(updateInstructor.InstructorName, updateInstructor.Email, updateInstructor.Password); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(err.Error()))
		}

		res, err := ctl.repo.Update(updateInstructor.ToEntityUser(uint(userID)))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses memperbarui data user", ToResponseUpdateInstructor(res)))
	}
}