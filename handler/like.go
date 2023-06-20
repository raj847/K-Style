package handler

import (
	"errors"
	"kstyle-test/entity"
	"kstyle-test/handler/request"
	"kstyle-test/handler/response"
	"kstyle-test/service"
	"kstyle-test/validate"
	"net/http"

	"github.com/labstack/echo/v4"
)

type LikeHandler struct {
	likeService *service.LikeService
}

func NewLikeHandler(
	likeService *service.LikeService,
) *LikeHandler {
	return &LikeHandler{
		likeService: likeService,
	}
}

func (m *LikeHandler) Like(c echo.Context) error {

	var likeReq request.Like
	err := c.Bind(&likeReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Errors: []response.ErrorDetail{
				{
					Message: "failed to read json request",
					Code:    "BAD_REQUEST",
				},
			},
		})
	}
	err = validate.Validate(likeReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Errors: []response.ErrorDetail{
				{
					Message: err.Error(),
					Code:    "MEMBER_INVALID-(EMPTY)",
				},
			},
		})
	}
	like := entity.Like{
		IDReview: likeReq.IDReview,
		IDMember: likeReq.IDMember,
	}
	check, err := m.likeService.Like(c.Request().Context(), &like)
	if err != nil {
		if errors.Is(err, service.ErrMemberNotFound) {
			return c.JSON(http.StatusBadRequest, response.ErrorResponse{
				Errors: []response.ErrorDetail{
					{
						Message: "failed to create like",
						Code:    "MEMBER_NOT-FOUND-ERROR",
					},
				},
			})
		} else if errors.Is(err, service.ErrReviewNotFound) {
			return c.JSON(http.StatusBadRequest, response.ErrorResponse{
				Errors: []response.ErrorDetail{
					{
						Message: "failed to create like",
						Code:    "REVIEW_NOT-FOUND-ERROR",
					},
				},
			})
		}
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Errors: []response.ErrorDetail{
				{
					Message: "failed to create like",
					Code:    "LIKE_CREATE-ERROR",
				},
			},
		})
	}
	if check {
		return c.JSON(http.StatusOK, response.SuccessMessage{
			Message: "Like success",
		})
	}
	return c.JSON(http.StatusOK, response.SuccessMessage{
		Message: "Unlike success",
	})
}
