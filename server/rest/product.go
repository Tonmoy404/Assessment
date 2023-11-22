package rest

import (
	"net/http"

	"github.com/Tonmoy404/Assessment/logger"
	"github.com/Tonmoy404/Assessment/service"
	"github.com/Tonmoy404/Assessment/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (s *Server) createProduct(ctx *gin.Context) {
	var req createProductReq
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		logger.Error(ctx, "cannot pass validation", err)
		ctx.JSON(http.StatusBadRequest, s.svc.Error(ctx, util.EN_API_PARAMETER_INVALID_ERROR, "Bad request"))
		return
	}

	existBrand, err := s.svc.GetBrand(ctx, req.BrandId)
	if err != nil {
		logger.Error(ctx, "cannot find brand", err)
		ctx.JSON(http.StatusNotFound, s.svc.Error(ctx, util.EN_NOT_FOUND, "Not Found"))
		return
	}

	if existBrand == nil {
		logger.Error(ctx, "Brand not found", nil)
		ctx.JSON(http.StatusBadRequest, s.svc.Response(ctx, "Brand not found", err))
		return
	}

	existSupplier, err := s.svc.GetSupplier(ctx, req.SupplierId)
	if err != nil {
		logger.Error(ctx, "cannot find supplier", err)
		ctx.JSON(http.StatusNotFound, s.svc.Error(ctx, util.EN_NOT_FOUND, "Not Found"))
		return
	}

	if existSupplier == nil {
		logger.Error(ctx, "Supplier not found", nil)
		ctx.JSON(http.StatusBadRequest, s.svc.Response(ctx, "Supplier not found", err))
		return
	}

	existCategory, err := s.svc.GetCategory(ctx, req.CategoryId)
	if err != nil {
		logger.Error(ctx, "cannot find category", err)
		ctx.JSON(http.StatusNotFound, s.svc.Error(ctx, util.EN_NOT_FOUND, "Not Found"))
		return
	}

	if existCategory == nil {
		logger.Error(ctx, "Brand not found", nil)
		ctx.JSON(http.StatusBadRequest, s.svc.Response(ctx, "Brand not found", err))
		return
	}

	existProduct, err := s.svc.GetProduct(ctx, req.Name)
	if err != nil {
		logger.Error(ctx, "cannot get product", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Error(ctx, util.EN_INTERNAL_SERVER_ERROR, "Internal Server Error"))
		return
	}

	if existProduct != nil {
		logger.Error(ctx, "already exists", err)
		ctx.JSON(http.StatusBadRequest, s.svc.Error(ctx, util.EN_API_PARAMETER_INVALID_ERROR, "Bad Request"))
		return
	}

	id, err := uuid.NewUUID()
	if err != nil {
		logger.Error(ctx, "cannot get uniqueId", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Error(ctx, util.EN_INTERNAL_SERVER_ERROR, "Internal Server Error"))
		return
	}

	newProduct := &service.Product{
		Id:            id.String(),
		Name:          req.Name,
		Description:   req.Description,
		Specification: req.Specification,
		BrandId:       req.BrandId,
		CategoryId:    req.CategoryId,
		SupplierId:    req.SupplierId,
		UnitPrice:     req.UnitPrice,
		DiscountPrice: req.DiscountPrice,
		Tags:          req.Tags,
		StatusId:      req.StatusId,
		CreatedAt:     util.GetCurrentTimestamp(),
	}

	newPro, err := s.svc.CreateProduct(ctx, newProduct)
	if err != nil {
		logger.Error(ctx, "cannot create product", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Error(ctx, util.EN_INTERNAL_SERVER_ERROR, "Internal server error"))
		return
	}

	ctx.JSON(http.StatusCreated, s.svc.Response(ctx, "Success", newPro))
}

func (s *Server) getProducts(ctx *gin.Context) {
	var req filterProductsReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		logger.Error(ctx, "cannot pass validation", err)
		ctx.JSON(http.StatusBadRequest, s.svc.Error(ctx, util.EN_API_PARAMETER_INVALID_ERROR, "Bad request"))
		return
	}

	filter := &service.FilterProducts{
		Name:       req.Name,
		MaxPrice:   req.MaxPrice,
		MinPrice:   req.MinPrice,
		BrandId:    req.BrandId,
		CategoryId: req.CategoryId,
		SupplierId: req.SupplierId,
		Limit:      req.Limit,
	}

	products, err := s.svc.GetProducts(ctx, filter)
	if err != nil {
		logger.Error(ctx, "cannot get products", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Error(ctx, util.EN_INTERNAL_SERVER_ERROR, "Internal Server Error"))
		return
	}

	ctx.JSON(http.StatusOK, s.svc.Response(ctx, "Fetced Products Successfully", products))
}

func (s *Server) getProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	product, err := s.svc.GetProduct(ctx, id)
	if err != nil {
		logger.Error(ctx, "cannot get product", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Error(ctx, util.EN_INTERNAL_SERVER_ERROR, "Internal Server Error"))
		return
	}

	if product == nil {
		logger.Error(ctx, "cannot find product", err)
		ctx.JSON(http.StatusNotFound, s.svc.Response(ctx, util.EN_NOT_FOUND, "Not Found"))
		return
	}
}

func (s *Server) updateProduct(ctx *gin.Context) {
	var req updateProductReq
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		logger.Error(ctx, "cannot pass validation", err)
		ctx.JSON(http.StatusBadRequest, s.svc.Response(ctx, util.EN_API_PARAMETER_INVALID_ERROR, err))
		return
	}

	productId := ctx.Param("id")

	product, err := s.svc.GetProduct(ctx, productId)
	if err != nil {
		logger.Error(ctx, "cannot get product", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Response(ctx, util.EN_INTERNAL_SERVER_ERROR, err))
		return
	}

	if product == nil {
		logger.Error(ctx, "product not found", nil)
		ctx.JSON(http.StatusNotFound, s.svc.Response(ctx, util.EN_NOT_FOUND, "Not found"))
		return
	}

	product.Name = req.Name
	product.Description = req.Description
	product.Specification = req.Specifications
	product.BrandId = req.BrandID
	product.CategoryId = req.CategoryID
	product.SupplierId = req.SupplierID
	product.Stock = req.Stock
	product.UnitPrice = req.UnitPrice
	product.DiscountPrice = req.DiscountPrice
	product.Tags = req.Tags
	product.StatusId = req.StatusID

	err = s.svc.UpdateProduct(ctx, product)
	if err != nil {
		logger.Error(ctx, "cannot update product", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Response(ctx, util.EN_INTERNAL_SERVER_ERROR, err))
		return
	}

	ctx.JSON(http.StatusOK, s.svc.Response(ctx, "Upadted Successfully", product))
}

func (s *Server) deleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	product, err := s.svc.GetProduct(ctx, id)
	if err != nil {
		logger.Error(ctx, "cannot get product", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Response(ctx, util.EN_INTERNAL_SERVER_ERROR, err))
		return
	}

	if product == nil {
		logger.Error(ctx, "product not found", nil)
		ctx.JSON(http.StatusNotFound, s.svc.Response(ctx, util.EN_NOT_FOUND, "Not found"))
		return
	}

	err = s.svc.DeleteProduct(ctx, id)
	if err != nil {
		logger.Error(ctx, "cannot delete product", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Response(ctx, util.EN_INTERNAL_SERVER_ERROR, err))
		return
	}

	ctx.JSON(http.StatusOK, s.svc.Response(ctx, "Deleted Successfully", product))
}
