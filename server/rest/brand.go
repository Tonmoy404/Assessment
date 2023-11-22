package rest

import (
	"net/http"

	"github.com/Tonmoy404/Assessment/logger"
	"github.com/Tonmoy404/Assessment/service"
	"github.com/Tonmoy404/Assessment/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (s *Server) createBrand(ctx *gin.Context) {
	var req createBrandReq
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		logger.Error(ctx, "cannot pass validation", err)
		ctx.JSON(http.StatusBadRequest, s.svc.Response(ctx, util.EN_API_PARAMETER_INVALID_ERROR, err))
		return
	}

	id, err := uuid.NewUUID()
	if err != nil {
		logger.Error(ctx, "cannot create uuid", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Error(ctx, util.EN_INTERNAL_SERVER_ERROR, "Internal Server Error"))
		return 
	}

	newBrand := &service.Brand{
		Id: id.String(),
		Name:      req.Name,
		StatusId:  req.StatusId,
		CreatedAt: util.GetCurrentTimestamp(),
	}

	brand, err := s.svc.CreateBrand(ctx, newBrand)
	if err != nil {
		logger.Error(ctx, "cannot create brand", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Response(ctx, util.EN_INTERNAL_SERVER_ERROR, err))
		return
	}

	ctx.JSON(http.StatusOK, s.svc.Response(ctx, "Successfully created", brand))
}

func (s *Server) deleteBrand(ctx *gin.Context) {
	id := ctx.Param("id")
	brand, err := s.svc.GetBrand(ctx, id)
	if err != nil {
		logger.Error(ctx, "cannot get brand", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Response(ctx, util.EN_INTERNAL_SERVER_ERROR, err))
		return
	}

	if brand == nil {
		logger.Error(ctx, "brand not found", nil)
		ctx.JSON(http.StatusNotFound, s.svc.Response(ctx, util.EN_NOT_FOUND, "Not found"))
		return
	}

	err = s.svc.DeleteBrand(ctx, id)
	if err != nil {
		logger.Error(ctx, "cannot delete brand", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Response(ctx, "Internal Server Error", err))
		return
	}

	ctx.JSON(http.StatusOK, s.svc.Response(ctx, "Deleted Successfully", nil))
}

func (s *Server) getBrand(ctx *gin.Context) {
	id := ctx.Param("id")
	brand, err := s.svc.GetBrand(ctx, id)
	if err != nil {
		logger.Error(ctx, "cannot get brand", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Response(ctx, "Internal Server Error", err))
		return
	}

	if brand == nil {
		logger.Error(ctx, "brand not found", err)
		ctx.JSON(http.StatusNotFound, s.svc.Response(ctx, util.EN_NOT_FOUND, err))
		return
	}

	ctx.JSON(http.StatusOK, s.svc.Response(ctx, "Successfully fetched", brand))
}

func (s *Server) getBrands(ctx *gin.Context) {
	var req getBrandsReq
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		logger.Error(ctx, "cannot pass validation", err)
		ctx.JSON(http.StatusBadRequest, s.svc.Response(ctx, util.EN_API_PARAMETER_INVALID_ERROR, err))
		return
	}

	result, err := s.svc.GetBrands(ctx, req.Page, req.Limit)
	if err != nil {
		logger.Error(ctx, "cannot get brands", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Response(ctx, "Internal Server Error", err))
		return
	}

	ctx.JSON(http.StatusOK, s.svc.Response(ctx, "Fetched brands", result))
}

func (s *Server) updateBrand(ctx *gin.Context) {
	id := ctx.Param("id")

	var req updateBrandReq
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		logger.Error(ctx, "cannot pass validation", err)
		ctx.JSON(http.StatusBadRequest, s.svc.Response(ctx, util.EN_API_PARAMETER_INVALID_ERROR, err))
		return
	}

	existBrand, err := s.svc.GetBrand(ctx, id)
	if err != nil {
		logger.Error(ctx, "cannot get brand", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Response(ctx, util.EN_INTERNAL_SERVER_ERROR, err))
		return
	}

	if existBrand == nil {
		logger.Error(ctx, "brand not found", nil)
		ctx.JSON(http.StatusNotFound, s.svc.Response(ctx, util.EN_NOT_FOUND, "Not found"))
		return
	}

	existBrand.Name = req.Name
	existBrand.StatusId = req.StatusID

	err = s.svc.UpdateBrand(ctx, existBrand)
	if err != nil {
		logger.Error(ctx, "cannot update brand", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Response(ctx, "Internal Server Error", err))
		return
	}

	ctx.JSON(http.StatusOK, s.svc.Response(ctx, "Successfully updated", nil))
}
