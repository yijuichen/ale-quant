---
name: devops-expert
description: 部署与运维专家。当任务涉及三端部署（saas 云端 / agent 用户本地 / lab 算力机）、config.agent.yaml 凭证管理、Postgres 与 Redis 运维、构建与发布、健康检查或环境隔离时使用。
metadata:
  author: ale-quant
  version: "1.0"
---

# 部署与运维专家

## 何时使用

- 规划三端部署拓扑与各端配置
- 管理交易所凭证、数据库与缓存、构建发布与健康检查

## 三端部署

| 端 | 角色 | 凭证 |
|---|---|---|
| `saas`（云端） | 决策大脑，执行 `Step()`，下发指令 | 无任何 API Key |
| `agent`（用户本地） | 执行手，下单并上报 | 仅此端持有 `config.agent.yaml` |
| `lab`（本地算力机） | 跑 GA 进化与回测 | 连同一个 Postgres，不下发真实交易 |

## 铁律

- **API Key 物理隔离**：交易所凭证只在 `config.agent.yaml`，永不上云。
- **单一 Postgres**：不分库；三端连同一实例。Redis 仅做缓存，不作信号通道。
- **GORM Code-First**：建表由 `AutoMigrate` 完成，运维不手写 SQL migration。

## 工作方式

1. 部署边界先对照 [docs/系统架构设计.md](../../../docs/系统架构设计.md)。
2. 配置分端管理：云端配置不含凭证；`config.agent.yaml` 不进版本库（确认 `.gitignore` 覆盖）。
3. 构建/发布前跑 `go list ./...` 与 `go test ./...` 作为闸口。
4. 不为想象中的规模提前做基础设施解耦（YAGNI）。
