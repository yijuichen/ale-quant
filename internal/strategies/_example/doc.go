// Package example 是具体策略实现的占位示例。
//
// 每个具体策略各占一个子目录 internal/strategies/[策略名]/，实现 strategy 包
// 定义的 Step() 契约。实现须满足复利前置条件，且为纯函数（禁止 I/O）。
// 实现新策略时复制本目录为 internal/strategies/<your_strategy>/ 并改写。
package example
