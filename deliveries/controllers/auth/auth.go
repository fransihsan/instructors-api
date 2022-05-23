package auth

import (
	"instructors-api/deliveries/controllers/common"
	"instructors-api/deliveries/middlewares"
	_AuthRepo "instructors-api/repositories/auth"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	repo _AuthRepo.Auth
}

func NewAuthController(repository _AuthRepo.Auth) *AuthController {
	return &AuthController{
		repo: repository,
	}
}

func (ctl *AuthController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var loginFormat RequestLogin
		if err := c.Bind(&loginFormat); err != nil || strings.TrimSpace(loginFormat.Email) == "" || strings.TrimSpace(loginFormat.Password) == "" {
			return c.JSON(http.StatusBadRequest, common.BadRequest("input dari user tidak sesuai, email atau password tidak boleh kosong"))
		}

		checkedUser, err := ctl.repo.Login(loginFormat.Email, loginFormat.Password)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}

		tokenID, err := middlewares.GenerateToken(checkedUser.ID)
		if err != nil {
			return c.JSON(http.StatusNotAcceptable, common.NotAcceptable())
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "berhasil masuk, mendapatkan token baru", ToResponseLogin(checkedUser, tokenID)))
	}
}
