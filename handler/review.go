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

type ReviewHandler struct {
	reviewService *service.ReviewService
}

func NewReviewHandler(
	reviewService *service.ReviewService,
) *ReviewHandler {
	return &ReviewHandler{
		reviewService: reviewService,
	}
}

func (m *ReviewHandler) Create(c echo.Context) error {

	var reviewReq request.Review
	err := c.Bind(&reviewReq)
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
	err = validate.Validate(reviewReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Errors: []response.ErrorDetail{
				{
					Message: err.Error(),
					Code:    "MEMBER_INVALID",
				},
			},
		})
	}
	review := entity.Review{
		IDProduct:  reviewReq.IDProduct,
		IDMember:   reviewReq.IDMember,
		DescReview: reviewReq.DescReview,
	}
	err = m.reviewService.AddReview(c.Request().Context(), &review)
	if err != nil {
		if errors.Is(err, service.ErrMemberNotFound) {
			return c.JSON(http.StatusBadRequest, response.ErrorResponse{
				Errors: []response.ErrorDetail{
					{
						Message: "failed to create review",
						Code:    "MEMBER_NOT-FOUND-ERROR",
					},
				},
			})
		} else if errors.Is(err, service.ErrProductNotFound) {
			return c.JSON(http.StatusBadRequest, response.ErrorResponse{
				Errors: []response.ErrorDetail{
					{
						Message: "failed to create review",
						Code:    "PRODUCT_NOT-FOUND-ERROR",
					},
				},
			})
		}
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Errors: []response.ErrorDetail{
				{
					Message: "failed to create review",
					Code:    "REVIEW_CREATE-ERROR",
				},
			},
		})
	}
	res := response.Review{
		IDProduct:  review.IDProduct,
		IDMember:   review.IDMember,
		DescReview: review.DescReview,
	}
	return c.JSON(http.StatusCreated, res)
}
