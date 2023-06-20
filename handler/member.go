package handler

import (
	"errors"
	"kstyle-test/entity"
	"kstyle-test/handler/request"
	"kstyle-test/handler/response"
	"kstyle-test/service"
	"kstyle-test/validate"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type MemberHandler struct {
	memberService *service.MemberService
}

func NewMemberHandler(
	memberService *service.MemberService,
) *MemberHandler {
	return &MemberHandler{
		memberService: memberService,
	}
}

func (m *MemberHandler) Create(c echo.Context) error {

	var memberReq request.Member
	err := c.Bind(&memberReq)
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
	err = validate.Validate(memberReq)
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
	member := entity.Member{
		Username:  memberReq.Username,
		Gender:    memberReq.Gender,
		Skintype:  memberReq.Skintype,
		Skincolor: memberReq.Skincolor,
	}
	err = m.memberService.AddMember(c.Request().Context(), &member)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Errors: []response.ErrorDetail{
				{
					Message: "failed to create member",
					Code:    "MEMBER_CREATE-ERROR",
				},
			},
		})
	}
	res := response.BuildMember(member)
	return c.JSON(http.StatusCreated, res)
}

func (m *MemberHandler) GetAll(c echo.Context) error {
	members, err := m.memberService.GetAllMember(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Errors: []response.ErrorDetail{
				{
					Message: "failed to read all member",
					Code:    "MEMBER_READ-ALL-ERROR",
				},
			},
		})
	}
	res := response.BuildMembers(members)
	return c.JSON(http.StatusOK, res)
}

func (m *MemberHandler) GetByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	members, err := m.memberService.GetMemberByID(c.Request().Context(), id)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			return c.JSON(http.StatusBadRequest, response.ErrorResponse{
				Errors: []response.ErrorDetail{
					{
						Message: "failed to read member",
						Code:    "MEMBER_NOT-FOUND-ERROR",
					},
				},
			})
		}
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Errors: []response.ErrorDetail{
				{
					Message: "failed to read member by id",
					Code:    "MEMBER_READ-BY-ID-ERROR",
				},
			},
		})
	}
	res := response.BuildMember(members)
	return c.JSON(http.StatusOK, res)
}

func (m *MemberHandler) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := m.memberService.DeleteMember(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Errors: []response.ErrorDetail{
				{
					Message: "failed to delete member",
					Code:    "MEMBER_DELETE-ERROR",
				},
			},
		})
	}
	return c.JSON(http.StatusOK, response.SuccessMessage{
		Message: "Member has been deleted",
	})
}

func (m *MemberHandler) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var memberReq request.Member
	err := c.Bind(&memberReq)
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
	err = validate.Validate(memberReq)
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
	res, err := m.memberService.UpdateMember(c.Request().Context(), &entity.Member{
		Model: gorm.Model{
			ID: uint(id),
		},
		Username:  memberReq.Username,
		Gender:    memberReq.Gender,
		Skintype:  memberReq.Skintype,
		Skincolor: memberReq.Skincolor,
	})
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Errors: []response.ErrorDetail{
				{
					Message: "failed to update member",
					Code:    "MEMBER_UPDATE-ERROR",
				},
			},
		})
	}
	result := response.BuildMember(res)
	return c.JSON(http.StatusOK, result)
}
