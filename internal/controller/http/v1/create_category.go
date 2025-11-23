package v1

import (
	"context"
	"errors"

	http_server "github.com/pc-configurator/components/gen/http/components_v1/server"
	"github.com/pc-configurator/components/internal/domain"
	"github.com/pc-configurator/components/internal/dto"
	"github.com/pc-configurator/components/pkg/logger"
	"github.com/pc-configurator/components/pkg/validation"
)

func (h Handlers) CreateCategory(ctx context.Context, request http_server.CreateCategoryRequestObject) (http_server.CreateCategoryResponseObject, error) {
	input := dto.CreateCategoryInput{
		Name: request.Body.Name,
	}

	res, err := h.usecase.CreateCategory(ctx, input)
	if err != nil {
		logger.SetCtxError(ctx, err)

		if errors.Is(err, domain.ErrCategoryNameExists) {
			return http_server.CreateCategory400JSONResponse{Error: domain.NewErrorWithDetails(domain.ErrCategoryNameExists)}, nil
		}

		var errorFields validation.ErrorFields
		if errors.As(err, &errorFields) {
			return http_server.CreateCategory400JSONResponse{Error: domain.NewValidationError(errorFields)}, nil
		}

		return http_server.CreateCategory500JSONResponse{Error: domain.NewErrorWithDetails(domain.ErrInternal)}, nil
	}

	return http_server.CreateCategory201JSONResponse{ID: res.ID}, nil
}
