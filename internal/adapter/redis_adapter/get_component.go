package redis_adapter

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/pc-configurator/components/internal/dto"
	"github.com/pc-configurator/components/pkg/logger"
)

func (r *RedisAdapter) GetComponent(ctx context.Context, componentID string) (dto.GetComponentIDOutput, error) {
	val, err := r.client.Get(ctx, ComponentIDKey(componentID)).Bytes()
	if err != nil {
		return dto.GetComponentIDOutput{}, err
	}

	var output dto.GetComponentIDOutput
	if err = json.Unmarshal(val, &output); err != nil {
		logger.Info(logger.NewErrorWithPath("json.Unmarshal", err).Error())
		return dto.GetComponentIDOutput{}, err
	}

	logger.Info(fmt.Sprintf("redis.GetComponent: component %s exists in cache", componentID))
	return output, nil
}
