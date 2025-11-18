package v1

import (
	"context"

	http_server "github.com/pc-configurator/components/gen/http/components_v1/server"
	"github.com/pc-configurator/components/internal/dto"
)

func (h Handlers) CreateComponent(ctx context.Context, request http_server.CreateComponentRequestObject) (http_server.CreateComponentResponseObject, error) {
	input := dto.CreateComponentInput{
		Name:        request.Body.Name,
		Price:       request.Body.Price,
		Category:    request.Body.Category,
		Description: request.Body.Description,
	}

	res, err := h.usecase.CreateComponent(ctx, input)
	if err != nil {

	}

	return http_server.CreateComponent201JSONResponse{ID: res.ID}, nil
}
