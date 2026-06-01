---
name: quant-math-expert
description: 量化交易数学专家。当任务涉及 Sigmoid 动态天平、目标权重、信号合成、β/γ 参数、楔形区过滤、无量纲指标（对数收益率、MAV、VolatilityRatio），或遗传算法基因空间与适应度时使用。真源为 docs/量化交易平仓策略.md 与 docs/进化文档.md。金融知识优先查 alphagbm skill。
metadata:
  author: ale-quant
  version: "1.0"
---

# 量化交易数学专家

## 何时使用

- 实现或审查 Sigmoid 目标权重、信号合成、楔形区过滤、粉尘拦截
- 设计无量纲特征（X1/X2/X3...）与 GA 染色体、适应度函数
- 任何价格相关计算

## 唯一真源

- 策略数学引擎：[docs/量化交易平仓策略.md](../../../docs/量化交易平仓策略.md)
- 进化计算引擎：[docs/进化文档.md](../../../docs/进化文档.md)
- 金融学知识优先查 [alphagbm](../alphagbm) skill 是否支持。

## 核心公式（Sigmoid 动态天平）

```
CurrentWeight = FloatBTC × Price / TotalEquity
EffectiveBeta = max(0.01, β × MarketBetaMultiplier)
InventoryBias = clamp(CurrentWeight, 0, 1) − 0.5
Exponent      = EffectiveBeta × Signal + γ × InventoryBias
TargetWeight  = 1 / (1 + e^Exponent), clamp(0, 1)
```

- Signal 正值倾向减仓、负值倾向加仓；`Signal = a×X1 + b×X2 + c×X3 + ...`，系数是 GA 染色体。
- 楔形过滤：`|TheoreticalUSD| ≥ 阈值`直接下单；小额仅在非安静态且楔形突破时强制最小订单。
- `VolatilityRatio = clip(MAV短期 / MAV长期, 0.1, 3.0)`，MAV 为平均绝对涨跌（非 ATR）。

## 铁律

- **无量纲计算**：用对数收益率或比率，禁止跨标的比较绝对价格。
- **策略纯函数**：计算原语放 `internal/quant/`，无副作用、可单测，禁止 I/O。
- **复利前置**：策略必须满足复利前置条件。
- **防过拟合**：GA 回测需自行加入乱序回测 / 蒙特卡洛（性能影响大，本期 plan 未写入，实现时评估）。

## 工作方式

1. 先读对应文档章节，确认公式与边界。
2. 纯计算落 `internal/quant/`，先写测试再写实现。
3. 文档未定义的因子或机制不臆造。
