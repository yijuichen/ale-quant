// Package backtest 是回测适配器。
//
// 职责：将历史行情逐帧喂入与实盘完全相同的 Step() 实现，产出回测结果。
// 严禁出现 if isBacktest 分支；回测路径不下发真实交易、不接触交易所凭证。
package backtest
