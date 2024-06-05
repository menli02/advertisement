package handler

import (
	"Advertisement/internal/domain"
	"Advertisement/internal/handler/metrics"
	"context"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"net/http"
	"strconv"
)

var tracer = otel.GetTracerProvider().Tracer("Advertisement")

func (h *Handler) create(c *gin.Context) {
	ctx := c.Request.Context()
	ctx, span := tracer.Start(ctx, "create-handler")
	defer span.End()
	var ad *domain.Advertisement
	if err := c.BindJSON(&ad); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.services.CreateAdvertisement(context.Background(), ad)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	metrics.IncrementHTTPRequestTotal(http.MethodPost, strconv.Itoa(http.StatusOK))
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Advertisement created successfully",
	})
}

func (h *Handler) getAll(c *gin.Context) {
	ctx := c.Request.Context()
	ctx, span := tracer.Start(ctx, "getAll-handler")
	defer span.End()
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	offset := (page - 1) * pageSize
	limit := pageSize

	sortBy := c.DefaultQuery("sortBy", "createdAt")
	sortOrder := c.DefaultQuery("sortOrder", "desc")

	advertisements, err := h.services.GetAllSortedAndPaginated(context.Background(), sortBy, sortOrder, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	metrics.IncrementHTTPRequestTotal(http.MethodGet, strconv.Itoa(http.StatusOK))

	c.JSON(http.StatusOK, advertisements)
}

func (h *Handler) getById(c *gin.Context) {
	ctx := c.Request.Context()
	ctx, span := tracer.Start(ctx, "getById-handler")
	defer span.End()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid advertisement ID"})
		return
	}

	ad, err := h.services.GetAdvertisementByID(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch advertisement"})
		return
	}

	if ad == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Advertisement not found"})
		return
	}
	metrics.IncrementHTTPRequestTotal(http.MethodGet, strconv.Itoa(http.StatusOK))

	c.JSON(http.StatusOK, ad)
}

func (h *Handler) update(c *gin.Context) {
	ctx := c.Request.Context()
	ctx, span := tracer.Start(ctx, "update-handler")
	defer span.End()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid advertisement ID"})
		return
	}
	var input *domain.UpdateAdvertisementInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	if err := input.ValidateUpdateInput(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedAd := domain.UpdateAdvertisementInput{
		Title:       input.Title,
		Description: input.Description,
		Price:       input.Price,
		IsActive:    input.IsActive,
	}
	if err := h.services.UpdateAdvertisement(context.Background(), id, &updatedAd); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update advertisement"})
		return
	}
	metrics.IncrementHTTPRequestTotal(http.MethodPut, strconv.Itoa(http.StatusOK))

	c.JSON(http.StatusOK, gin.H{"message": "Advertisement updated successfully"})
}

func (h *Handler) delete(c *gin.Context) {
	ctx := c.Request.Context()
	ctx, span := tracer.Start(ctx, "delete-handler")
	defer span.End()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid advertisement ID"})
		return
	}

	if err := h.services.DeleteAdvertisement(context.Background(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete advertisement"})
		return
	}
	metrics.IncrementHTTPRequestTotal(http.MethodDelete, strconv.Itoa(http.StatusOK))

	c.JSON(http.StatusOK, gin.H{"message": "Advertisement deleted successfully"})
}
