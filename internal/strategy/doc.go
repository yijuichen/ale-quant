// Package strategy 定义策略框架与契约。
//
// 职责：声明 Step() 接口契约、策略注册机制、基因空间与评估接口。
// 回测与实盘共用同一 Step() 实现；Step() 内部禁止网络、数据库、文件 I/O。
package strategy
