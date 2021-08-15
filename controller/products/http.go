package products

import (
	"consignku/bussiness/products"
	"consignku/controller"
	"consignku/controller/products/request"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type ProductsController struct {
	productsUseCase products.Usecase
}

func NewProductsController(productsUC products.Usecase) *ProductsController {
	return &ProductsController{
		productsUseCase: productsUC,
	}
}

func (ctrl *ProductsController) Store(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.Products{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	_, err := ctrl.productsUseCase.Store(ctx, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, "Successfully inserted")
}

func (ctrl *ProductsController) Update(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")

	if strings.TrimSpace(id) == "" {
		return controller.NewErrorResponse(c, http.StatusBadRequest, errors.New("missing required id"))
	}

	req := request.Products{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	domainReq := req.ToDomain()
	idInt, _ := strconv.Atoi(id)
	domainReq.Id = idInt

	_, err := ctrl.productsUseCase.Update(ctx, domainReq)

	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, "Successfully updated")
}

func (ctrl *ProductsController) GetByID(c echo.Context) error {
	ctx := c.Request().Context()

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	discounts, err := ctrl.productsUseCase.GetByID(ctx, id)

	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controller.NewSuccessResponse(c, discounts)
}

func (ctrl *ProductsController) Delete(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")

	if strings.TrimSpace(id) == "" {
		return controller.NewErrorResponse(c, http.StatusBadRequest, errors.New("missing required id"))
	}

	req := request.Products{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	domainReq := req.ToDomain()
	idInt, _ := strconv.Atoi(id)
	domainReq.Id = idInt

	_, err := ctrl.productsUseCase.Delete(ctx, domainReq)

	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, "Successfully deleted")
}
