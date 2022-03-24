package handler

import (
	"net/http"
	"strconv"

	"github.com/gs1068/golang-ddd-sample/usecase"
	"github.com/labstack/echo"
)

type TimelineHandler interface {
	Post() echo.HandlerFunc
	Get() echo.HandlerFunc
	Put() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type timelineHandler struct {
	timelineUsecase usecase.TimelineUsecase
}

func NewTimelineHandler(timelineUsecase usecase.TimelineUsecase) TimelineHandler {
	return &timelineHandler{timelineUsecase: timelineUsecase}
}

type requestTimeline struct {
	UserID  int    `json:"user_id"`
	Content string `json:"content"`
}

type responseTimeline struct {
	ID      int    `json:"id"`
	UserID  int    `json:"user_id"`
	Content string `json:"content"`
}

func (th *timelineHandler) Post() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req requestTimeline
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		createdTimeline, err := th.timelineUsecase.Create(req.UserID, req.Content)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusCreated, responseTimeline{
			ID:      createdTimeline.ID,
			UserID:  createdTimeline.UserID,
			Content: createdTimeline.Content,
		})
	}
}

func (th *timelineHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		foundTimeline, err := th.timelineUsecase.FindByID(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, responseTimeline{
			ID:      foundTimeline.ID,
			UserID:  foundTimeline.UserID,
			Content: foundTimeline.Content,
		})
	}
}

func (th *timelineHandler) Put() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		var req requestTimeline
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		updatedTimeline, err := th.timelineUsecase.Update(id, req.UserID, req.Content)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, responseTimeline{
			ID:      updatedTimeline.ID,
			UserID:  updatedTimeline.UserID,
			Content: updatedTimeline.Content,
		})
	}
}

func (th *timelineHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		err = th.timelineUsecase.Delete(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.NoContent(http.StatusOK)
	}
}
