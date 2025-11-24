package redis_adapter

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/pc-configurator/components/internal/dto"
	"github.com/pc-configurator/components/pkg/logger"
)

func (r *RedisAdapter) SetComponent(ctx context.Context, component dto.GetComponentIDOutput) {
	data, err := json.Marshal(component)
	if err != nil {
		logger.Info(logger.NewErrorWithPath("redis_adapter.json.Marshal", err).Error())
		return
	}

	err = r.client.Set(ctx, ComponentIDKey(strconv.Itoa(component.ID)), data, time.Hour).Err()
	if err != nil {
		logger.Info(logger.NewErrorWithPath("redis_adapter.client.Set", err).Error())
	}
}
