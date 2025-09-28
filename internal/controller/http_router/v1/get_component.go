package v1

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/pc-configurator/components/internal/dto"
	"github.com/pc-configurator/components/pkg/base_errors"
	"github.com/pc-configurator/components/pkg/logger"
	"github.com/pc-configurator/components/pkg/validation"
)

func (h *Handlers) GetComponent(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	input := dto.GetComponentInput{
		ID: chi.URLParam(r, "id"),
	}

	invalidFields, err := validation.ValidateStruct(&input)
	if err != nil {
		logger.Error(err, "validation.ValidateStruct")
		errorResponse(w, http.StatusBadRequest, errorPayload{Type: base_errors.Validation, Details: invalidFields})
		return
	}

	output, err := h.usecase.GetComponent(ctx, input)
	if err != nil {
		logger.Error(err, "usecase.GetComponent")

		if errors.Is(err, base_errors.NotFound) {
			errorResponse(w, http.StatusBadRequest, errorPayload{Type: base_errors.NotFound})
			return
		}

		errorResponse(w, http.StatusInternalServerError, errorPayload{Type: base_errors.InternalServer})
		return

	}
	successResponse(w, http.StatusOK, output)
}
