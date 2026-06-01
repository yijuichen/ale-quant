// Package agent 承载 Agent 侧业务逻辑。
//
// 职责：接收 SaaS 下发的交易意图，调用交易所下单执行，并上报成交结果。
// 不含任何策略代码；交易所凭证仅读取自 config.agent.yaml。
package agent
