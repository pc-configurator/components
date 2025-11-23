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

func (h Handlers) CreateComponent(ctx context.Context, request http_server.CreateComponentRequestObject) (http_server.CreateComponentResponseObject, error) {
	input := dto.CreateComponentInput{
		Name:        request.Body.Name,
		Price:       request.Body.Price,
		CategoryID:  request.Body.CategoryID,
		Description: request.Body.Description,
	}

	res, err := h.usecase.CreateComponent(ctx, input)
	if err != nil {
		logger.SetCtxError(ctx, err)

		if errors.Is(err, domain.ErrComponentNameExists) {
			return http_server.CreateComponent400JSONResponse{Error: domain.NewErrorWithDetails(domain.ErrComponentNameExists)}, nil
		}

		if errors.Is(err, domain.ErrCategoryNotFound) {
			return http_server.CreateComponent400JSONResponse{Error: domain.NewErrorWithDetails(domain.ErrCategoryNotFound)}, nil
		}

		var errorFields validation.ErrorFields
		if errors.As(err, &errorFields) {
			return http_server.CreateComponent400JSONResponse{Error: domain.NewValidationError(errorFields)}, nil
		}

		return http_server.CreateComponent500JSONResponse{Error: domain.NewErrorWithDetails(domain.ErrInternal)}, nil
	}

	return http_server.CreateComponent201JSONResponse{ID: &res.ID}, nil
}
