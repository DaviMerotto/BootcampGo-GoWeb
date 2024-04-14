package handler

import (
	"net/http"
	"strconv"

	"github.com/davimerotto/web-server/internal/products"
	"github.com/davimerotto/web-server/pkg/web"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service products.Service
}

func NewProduct(p products.Service) *ProductHandler {
	return &ProductHandler{
		service: p,
	}
}

// ListProducts godoc
// @Summary List products
// @Tags Product
// @Description get products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.HttpError
// @Router /products [GET]
func (c *ProductHandler) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		p, err := c.service.GetAll()
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest,
				web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}

		if len(p) == 0 {
			ctx.Status(http.StatusNoContent)
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, p, ""))
	}
}

// CreateProduct godoc
// @Summary Create product
// @Tags Product
// @Description create product
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param product body products.CreateProduct true "Add product"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.HttpError
// @Router /products [POST]
func (c *ProductHandler) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req products.Product
		if err := ctx.Bind(&req); err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity,
				web.NewResponse(http.StatusUnprocessableEntity, nil, err.Error()))
			return
		}

		p, err := c.service.Create(req)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest,
				web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusCreated, web.NewResponse(http.StatusCreated, p, ""))
	}
}

// DeleteProduct godoc
// @Summary Delete product
// @Tags Product
// @Description delete product
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param id path int true "Product ID"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.HttpError
// @Router /products/{id} [DELETE]
func (c *ProductHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseUint(ctx.Param("id"), 10, 0)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest,
				web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}
		err = c.service.Delete(uint(id))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest,
				web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusAccepted,
			web.NewResponse(http.StatusAccepted, "produto deletado com sucesso!", ""))
	}
}

// UpdateAllFieldsProduct godoc
// @Summary Update all fields product
// @Tags Product
// @Description update all fields of product
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param product body products.Product true "Update all fields of product"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.HttpError
// @Router /products [PATCH]
func (c *ProductHandler) UpdateFull() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		p := products.Product{}
		if err := ctx.Bind(&p); err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity,
				web.NewResponse(http.StatusUnprocessableEntity, nil, err.Error()))
			return
		}
		if p.Id == 0 {
			ctx.AbortWithStatusJSON(http.StatusBadRequest,
				web.NewResponse(http.StatusBadRequest, nil, "ID não informado"))
			return
		}
		if p.Name == "" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest,
				web.NewResponse(http.StatusBadRequest, nil, "Nome não informado"))
			return
		}
		if p.Color == "" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest,
				web.NewResponse(http.StatusBadRequest, nil, "Cor não informada"))
			return
		}
		if p.Price == 0 {
			ctx.AbortWithStatusJSON(http.StatusBadRequest,
				web.NewResponse(http.StatusBadRequest, nil, "Preço não informado"))
			return
		}
		if p.Stock == 0 {
			ctx.AbortWithStatusJSON(http.StatusBadRequest,
				web.NewResponse(http.StatusBadRequest, nil, "Estoque não informado"))
			return
		}
		if p.Code == "" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest,
				web.NewResponse(http.StatusBadRequest, nil, "Código não informado"))
			return
		}
		if p.Creation_date == "" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest,
				web.NewResponse(http.StatusBadRequest, nil, "Data de criação não informada"))
			return
		}

		prod, err := c.service.UpdateFull(p)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest,
				web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, prod, ""))
	}
}

// UpdateProduct godoc
// @Summary Update product
// @Tags Product
// @Description uptade any field of product
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param id path int true "Product ID"
// @Param product body products.UpdateProduct true "Update product"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.HttpError
// @Router /products/{id} [PUT]
func (c *ProductHandler) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		p := products.Product{}
		id, err := strconv.ParseUint(ctx.Param("id"), 10, 0)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity,
				web.NewResponse(http.StatusUnprocessableEntity, nil, err.Error()))
			return
		}
		if err := ctx.Bind(&p); err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity,
				web.NewResponse(http.StatusUnprocessableEntity, nil, err.Error()))
			return
		}
		if id == 0 {
			ctx.AbortWithStatusJSON(http.StatusBadRequest,
				web.NewResponse(http.StatusBadRequest, nil, "ID não informado"))
			return
		}
		prod, err := c.service.Update(uint(id), p)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest,
				web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, prod, ""))
	}
}
