package handler

import (
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

type ProductHandler struct {
	productService *service.ProductService
}

func NewProductHandler(
	productService *service.ProductService,
) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

func (m *ProductHandler) Create(c echo.Context) error {

	var productReq request.Product
	err := c.Bind(&productReq)
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
	err = validate.Validate(productReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Errors: []response.ErrorDetail{
				{
					Message: err.Error(),
					Code:    "PRODUCT_INVALID",
				},
			},
		})
	}
	product := entity.Product{
		NameProduct: productReq.NameProduct,
		Price:       productReq.Price,
	}
	err = m.productService.AddProduct(c.Request().Context(), &product)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Errors: []response.ErrorDetail{
				{
					Message: "failed to create product",
					Code:    "PRODUCT_CREATE-ERROR",
				},
			},
		})
	}
	return c.JSON(http.StatusCreated, product)
}

func (m *ProductHandler) GetAll(c echo.Context) error {
	products, err := m.productService.GetAllProduct(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Errors: []response.ErrorDetail{
				{
					Message: "failed to read all product",
					Code:    "PRODUCT_READ-ALL-ERROR",
				},
			},
		})
	}
	list, err := m.productService.GetAllLikes(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Errors: []response.ErrorDetail{
				{
					Message: "failed to read all product",
					Code:    "PRODUCT_READ-LIKES-ERROR",
				},
			},
		})
	}
	res := response.BuildProducts(products, list)

	return c.JSON(http.StatusOK, res)
}

func (m *ProductHandler) GetByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	product, err := m.productService.GetProductByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Errors: []response.ErrorDetail{
				{
					Message: "failed to read all product",
					Code:    "PRODUCT_READ-ALL-ERROR",
				},
			},
		})
	}
	list, err := m.productService.GetAllLikes(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Errors: []response.ErrorDetail{
				{
					Message: "failed to read all product",
					Code:    "PRODUCT_READ-LIKES-ERROR",
				},
			},
		})
	}
	res := response.BuildProduct(product, list)

	return c.JSON(http.StatusOK, res)
}

func (m *ProductHandler) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := m.productService.DeleteProduct(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Errors: []response.ErrorDetail{
				{
					Message: "failed to delete product",
					Code:    "PRODUCT_DELETE-ERROR",
				},
			},
		})
	}
	return c.JSON(http.StatusOK, response.SuccessMessage{
		Message: "Product has been deleted",
	})
}

func (m *ProductHandler) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var productReq request.Product
	err := c.Bind(&productReq)
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
	err = validate.Validate(productReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Errors: []response.ErrorDetail{
				{
					Message: err.Error(),
					Code:    "PRODUCT_INVALID",
				},
			},
		})
	}
	res, err := m.productService.UpdateProduct(c.Request().Context(), &entity.Product{
		Model: gorm.Model{
			ID: uint(id),
		},
		NameProduct: productReq.NameProduct,
		Price:       productReq.Price,
	})
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Errors: []response.ErrorDetail{
				{
					Message: "failed to update product",
					Code:    "PRODUCT_UPDATE-ERROR",
				},
			},
		})
	}
	return c.JSON(http.StatusOK, res)
}
