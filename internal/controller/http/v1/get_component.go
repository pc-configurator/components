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

func (h Handlers) GetComponentID(ctx context.Context, request http_server.GetComponentIDRequestObject) (http_server.GetComponentIDResponseObject, error) {
	input := dto.GetComponentIDInput{
		ID: request.ID,
	}

	res, err := h.usecase.GetComponent(ctx, input)
	if err != nil {
		logger.SetCtxError(ctx, err)

		if errors.Is(err, domain.ErrComponentNotFound) {
			return http_server.GetComponentID404JSONResponse{Error: domain.NewErrorWithDetails(domain.ErrComponentNotFound)}, nil
		}

		var errorFields validation.ErrorFields
		if errors.As(err, &errorFields) {
			return http_server.GetComponentID400JSONResponse{Error: domain.NewValidationError(errorFields)}, nil
		}

		return http_server.GetComponentID500JSONResponse{Error: domain.NewErrorWithDetails(domain.ErrInternal)}, nil
	}

	return http_server.GetComponentID200JSONResponse{ID: &res.ID}, nil
}
