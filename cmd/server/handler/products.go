package handler

import (
	"net/http"
	"strconv"

	"github.com/davimerotto/web-server/internal/products"
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

func (c *ProductHandler) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		p, err := c.service.GetAll()
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if len(p) == 0 {
			ctx.Status(http.StatusNoContent)
			return
		}

		ctx.JSON(http.StatusOK, p)
	}
}

func (c *ProductHandler) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req products.Product
		if err := ctx.Bind(&req); err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
				"error": err.Error(),
			})
			return
		}

		p, err := c.service.Create(req)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusCreated, p)
	}
}

func (c *ProductHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseUint(ctx.Param("id"), 10, 0)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err = c.service.Delete(uint(id))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusAccepted, gin.H{"message": "Produto deletado com sucesso"})
	}
}

func (c *ProductHandler) UpdateFull() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		p := products.Product{}
		if err := ctx.Bind(&p); err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}
		if p.Id == 0 {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "ID não informado"})
			return
		}
		if p.Name == "" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Nome não informado"})
			return
		}
		if p.Color == "" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Cor não informada"})
			return
		}
		if p.Price == 0 {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Preço não informado"})
			return
		}
		if p.Stock == 0 {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Estoque não informado"})
			return
		}
		if p.Code == "" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Código não informado"})
			return
		}
		if p.Creation_date == "" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Data de criação não informada"})
			return
		}

		prod, err := c.service.UpdateFull(p)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, prod)
	}
}

func (c *ProductHandler) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		p := products.Product{}
		id, err := strconv.ParseUint(ctx.Param("id"), 10, 0)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}
		if err := ctx.Bind(&p); err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}
		if id == 0 {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "ID não informado"})
			return
		}
		prod, err := c.service.Update(uint(id), p)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, prod)
	}
}
