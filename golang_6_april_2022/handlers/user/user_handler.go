package user

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/tobolkin8/go_april_2022/database"
	"github.com/tobolkin8/go_april_2022/helpers"
	authModel "github.com/tobolkin8/go_april_2022/models/auth"
)

type UserService struct {
	Repo *database.Repo
}

func (svc UserService) Show(c *echo.Context) error {
	var user authModel.LoginInfo
	svc.Repo.First(&user, c.Param("id"))
	return helpers.EncodeJSON(c, http.StatusOK, &user)
}

func (svc UserService) List(c *echo.Context) error {
	var users []*authModel.LoginInfo
	svc.Repo.Find(&users)
	return helpers.EncodeJSON(c, http.StatusOK, users)
}

func (svc UserService) Create(c *echo.Context) error {
	var user authModel.LoginInfo
	if err := helpers.Decode(c.Request(), &user); err != nil {
		return err
	}

	if errs := user.Validate(); errs != nil {
		return errs[0]
	}
	svc.Repo.Create(&user)

	return c.JSON(http.StatusCreated, &user)
}

func (svc UserService) Update(c *echo.Context) error {
	return c.String(http.StatusOK, "hello, v1 world")
}

func (svc UserService) Delete(c *echo.Context) error {
	return c.String(http.StatusOK, "hello, v1 world")
}
