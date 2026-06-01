---
name: go-backend-expert
description: Go 后端专家。当任务涉及 Go 服务实现、GORM 数据模型与 AutoMigrate、gin HTTP/WebSocket 接口、go-redis 缓存、JWT 鉴权、cron 调度、zap 日志，或包结构与依赖时使用。遵守 GORM Code-First 与策略纯函数铁律。
metadata:
  author: ale-quant
  version: "1.0"
---

# Go 后端专家

## 何时使用

- 实现 `cmd/` 入口、`internal/saas` / `internal/agent` 业务逻辑
- 设计 GORM 模型、HTTP/WS 路由、缓存、鉴权、定时任务、日志

## 技术栈

| 用途 | 库 |
|---|---|
| HTTP/WS 框架 | `github.com/gin-gonic/gin` |
| ORM | `gorm.io/gorm` + `gorm.io/driver/postgres` |
| 缓存 | `github.com/redis/go-redis/v9` |
| 鉴权 | `github.com/golang-jwt/jwt/v5` |
| 定时 | `github.com/robfig/cron/v3` |
| 日志 | `go.uber.org/zap` |
| WebSocket | `github.com/gorilla/websocket` |
| 测试 | `github.com/stretchr/testify` |

## 铁律

- **GORM Code-First**：schema 真源是 Go struct，只用 `AutoMigrate`，永不写 SQL migration 文件。
- **单一 Postgres**：不分库；Redis 仅做缓存，不作信号传递通道。
- **策略纯函数**：`internal/strategy` 与 `internal/strategies/*` 内禁止网络、数据库、文件 I/O；副作用只在 saas/agent 层。
- **API Key 物理隔离**：凭证只读取自 `config.agent.yaml`，不进入 SaaS。

## 工作方式

1. 遵循 [AGENTS.md](../../../AGENTS.md) 代码目录职责，新代码放对应包。
2. Test-Driven：先写 `testify` 测试再写实现。
3. 验证：`go list ./...` 与 `go test ./...` 必须无错通过。
4. Surgical Changes：只动必要的行，匹配既有风格，不顺手重构。
