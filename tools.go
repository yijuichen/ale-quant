//go:build tools

// 依赖锚点：在各模块真正引用前，固定项目所需依赖的版本，使其留存于
// go.mod / go.sum 且不被 go mod tidy 删除。本文件带 tools 构建标签，
// 不进入任何二进制。某依赖被真实代码引用后，从此处移除对应行。
package tools

import (
	_ "github.com/gin-gonic/gin"
	_ "github.com/golang-jwt/jwt/v5"
	_ "github.com/gorilla/websocket"
	_ "github.com/redis/go-redis/v9"
	_ "github.com/robfig/cron/v3"
	_ "github.com/stretchr/testify"
	_ "go.uber.org/zap"
	_ "gorm.io/driver/postgres"
	_ "gorm.io/gorm"
)
