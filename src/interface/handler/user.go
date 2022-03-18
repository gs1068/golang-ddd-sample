package handler

import (
	"net/http"
	"strconv"

	"github.com/gs1068/golang_ddd_sample/usecase"
	"github.com/labstack/echo"
)

type UserHandler interface {
	Post() echo.HandlerFunc
	Get() echo.HandlerFunc
	GetAll() echo.HandlerFunc
	Put() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type userHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) UserHandler {
	return &userHandler{userUsecase: userUsecase}
}

type requestUser struct {
	UserName string `json:"username"`
}

type responseUser struct {
	ID       int    `json:"id"`
	UserName string `json:"username"`
}

func (uh *userHandler) Post() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req requestUser
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		createdUser, err := uh.userUsecase.Create(req.UserName)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseUser{
			ID:       createdUser.ID,
			UserName: createdUser.UserName,
		}

		return c.JSON(http.StatusCreated, res)
	}
}

func (uh *userHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi((c.Param("id")))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		foundUser, err := uh.userUsecase.FindByID(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseUser{
			ID:       foundUser.ID,
			UserName: foundUser.UserName,
		}

		return c.JSON(http.StatusOK, res)
	}
}

func (uh *userHandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {

		foundUsers, err := uh.userUsecase.FindAll()
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		var res []responseUser

		for _, _user := range *foundUsers {
			user := _user
			new_user := responseUser{
				ID:       user.ID,
				UserName: user.UserName,
			}
			res = append(res, new_user)
		}

		return c.JSON(http.StatusOK, res)
	}
}

func (uh *userHandler) Put() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		var req requestUser
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		updatedUser, err := uh.userUsecase.Update(id, req.UserName)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseUser{
			ID:       updatedUser.ID,
			UserName: updatedUser.UserName,
		}

		return c.JSON(http.StatusOK, res)
	}
}

func (uh *userHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		err = uh.userUsecase.Delete(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.NoContent(http.StatusNoContent)
	}
}
