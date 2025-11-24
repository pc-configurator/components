package redis_adapter

import "fmt"

func ComponentIDKey(componentID string) string {
	return fmt.Sprintf("component_%s", componentID)
}
