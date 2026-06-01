// Package saas 承载 SaaS 侧业务逻辑。
//
// 职责：决策编排、HTTP/WebSocket API、定时调度。Step() 只在本侧执行。
// 不持有交易所 API Key。
package saas
