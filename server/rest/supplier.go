package rest

import (
	"net/http"

	"github.com/Tonmoy404/Assessment/logger"
	"github.com/Tonmoy404/Assessment/service"
	"github.com/Tonmoy404/Assessment/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (s *Server) createSupplier(ctx *gin.Context) {
	var req createSupplierReq
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		logger.Error(ctx, "cannot pass validation", err)
		ctx.JSON(http.StatusBadRequest, s.svc.Response(ctx, util.EN_API_PARAMETER_INVALID_ERROR, err))
		return
	}

	ID, err := uuid.NewUUID()
	if err != nil {
		logger.Error(ctx, "cannot create uuid", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Response(ctx, util.EN_INTERNAL_SERVER_ERROR, "Internal Server Error"))
		return
	}

	newSupplier := &service.Supplier{
		Id:                 ID.String(),
		Name:               req.Name,
		Email:              req.Email,
		Phone:              req.Phone,
		StatusId:           req.StatusId,
		IsVerifiedSupplier: req.IsVerifiedSupplier,
		CreatedAt:          util.GetCurrentTimestamp(),
	}

	supp, err := s.svc.CreateSupplier(ctx, newSupplier)
	if err != nil {
		logger.Error(ctx, "cannot create supplier", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Response(ctx, util.EN_INTERNAL_SERVER_ERROR, "Internal Server Error"))
		return
	}

	ctx.JSON(http.StatusOK, s.svc.Response(ctx, "Supplier Added Successfully", supp))
}

func (s *Server) deleteSupplier(ctx *gin.Context) {
	id := ctx.Param("id")

	existSupplier, err := s.svc.GetSupplier(ctx, id)
	if err != nil {
		logger.Error(ctx, "cannot get supplier", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Response(ctx, util.EN_INTERNAL_SERVER_ERROR, err))
		return
	}

	if existSupplier == nil {
		logger.Error(ctx, "supplier not found", nil)
		ctx.JSON(http.StatusNotFound, s.svc.Response(ctx, util.EN_NOT_FOUND, "Not found"))
		return
	}

	err = s.svc.DeleteSupplier(ctx, id)
	if err != nil {
		logger.Error(ctx, "cannot delete supplier", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Response(ctx, util.EN_INTERNAL_SERVER_ERROR, "Internal Server Error"))
		return
	}

	ctx.JSON(http.StatusOK, s.svc.Response(ctx, "Deleted supplier successfully", nil))
}

func (s *Server) getSupplier(ctx *gin.Context) {
	id := ctx.Param("id")

	supplier, err := s.svc.GetSupplier(ctx, id)
	if err != nil {
		logger.Error(ctx, "cannot get supplier", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Response(ctx, util.EN_INTERNAL_SERVER_ERROR, "Internal Server Error"))
		return
	}

	if supplier == nil {
		logger.Error(ctx, "supplier not exists", err)
		ctx.JSON(http.StatusNotFound, s.svc.Error(ctx, util.EN_NOT_FOUND, "Not Found"))
		return
	}

	ctx.JSON(http.StatusOK, s.svc.Response(ctx, "Successfully fetched", supplier))
}

func (s *Server) getSuppliers(ctx *gin.Context) {
	var req getSuppliersReq
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		logger.Error(ctx, "cannot pass validation", err)
		ctx.JSON(http.StatusBadRequest, s.svc.Response(ctx, util.EN_API_PARAMETER_INVALID_ERROR, "Bad Request"))
		return
	}

	suppliers, err := s.svc.GetSuppliers(ctx, req.Page, req.Limit)
	if err != nil {
		logger.Error(ctx, "cannot get suppliers", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Response(ctx, "Internal Server Error", err))
		return
	}

	if suppliers == nil {
		logger.Error(ctx, "No Supplier Found", err)
		ctx.JSON(http.StatusNotFound, s.svc.Error(ctx, util.EN_NOT_FOUND, "No Supplier Found"))
	}

	ctx.JSON(http.StatusOK, s.svc.Response(ctx, "Fetched suppliers", suppliers))
}

func (s *Server) updateSupplier(ctx *gin.Context) {
	var req updateSupplierReq
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		logger.Error(ctx, "cannot pass validation", err)
		ctx.JSON(http.StatusBadRequest, s.svc.Response(ctx, util.EN_API_PARAMETER_INVALID_ERROR, "Bad Request"))
		return
	}

	id := ctx.Param("id")

	supplier, err := s.svc.GetSupplier(ctx, id)
	if err != nil {
		logger.Error(ctx, "cannot get supplier", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Response(ctx, util.EN_INTERNAL_SERVER_ERROR, "Internal Server Error"))
		return
	}

	if supplier == nil {
		logger.Error(ctx, "supplier not found", nil)
		ctx.JSON(http.StatusNotFound, s.svc.Response(ctx, util.EN_NOT_FOUND, "Not found"))
		return
	}

	supplier.Name = req.Name
	supplier.Email = req.Email
	supplier.Phone = req.Phone
	supplier.StatusId = req.StatusID
	supplier.IsVerifiedSupplier = req.IsVerifiedSupplier

	err = s.svc.UpdateSupplier(ctx, supplier)
	if err != nil {
		logger.Error(ctx, "cannot update supplier", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Error(ctx, util.EN_INTERNAL_SERVER_ERROR, "Internal Server Error"))
		return
	}

	ctx.JSON(http.StatusOK, s.svc.Response(ctx, "Successfully updated", supplier))
}
