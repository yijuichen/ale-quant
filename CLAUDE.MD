## 憲法
### 1.
文檔全部使用中文
### 2.
1. Think Before Coding
   Don't assume. Don't hide confusion. Surface tradeoffs.

Before implementing:

State your assumptions explicitly. If uncertain, ask.
If multiple interpretations exist, present them - don't pick silently.
If a simpler approach exists, say so. Push back when warranted.
If something is unclear, stop. Name what's confusing. Ask.
2. Simplicity First
   Minimum code that solves the problem. Nothing speculative.

No features beyond what was asked.
No abstractions for single-use code.
No "flexibility" or "configurability" that wasn't requested.
No error handling for impossible scenarios.
If you write 200 lines and it could be 50, rewrite it.
Ask yourself: "Would a senior engineer say this is overcomplicated?" If yes, simplify.

3. Surgical Changes
   Touch only what you must. Clean up only your own mess.

When editing existing code:

Don't "improve" adjacent code, comments, or formatting.
Don't refactor things that aren't broken.
Match existing style, even if you'd do it differently.
If you notice unrelated dead code, mention it - don't delete it.
When your changes create orphans:

Remove imports/variables/functions that YOUR changes made unused.
Don't remove pre-existing dead code unless asked.
The test: Every changed line should trace directly to the user's request.

4. Goal-Driven Execution
   Define success criteria. Loop until verified.

Transform tasks into verifiable goals:

"Add validation" → "Write tests for invalid inputs, then make them pass"
"Fix the bug" → "Write a test that reproduces it, then make it pass"
"Refactor X" → "Ensure tests pass before and after"
For multi-step tasks, state a brief plan:

1. [Step] → verify: [check]
2. [Step] → verify: [check]
3. [Step] → verify: [check]
   Strong success criteria let you loop independently. Weak criteria ("make it work") require constant clarification.
---
### 任何金融學知識都優先參考看看 alphagbm skill 有沒有支援 [alphagbm](.claude/skills/alphagbm)

---

## 一、系统定位与核心架构

### 系统是什么

一款面向量化交易投资者的**全天候智能量化管理工具**：通过华尔街级风控状态机、遗传算法参数寻优，将复杂的动态策略降维成普通人能一键托管的 SaaS 财富水库。

### 三端物理部署

- **`saas`（云端）**：决策大脑，执行 `Step()`，下发交易指令，不持有任何 API Key
- **`agent`（用户本地）**：执行手，只负责调用交易所下单并上报结果，不含任何策略代码
- **`lab`（本地算力机）**：实验室，专跑 GA 进化与回测，连同一个 Postgres 实例，不下发真实交易

### 全局技术决策（不可推翻的铁律）

**在开始前，把这六条铁律贴在显眼的地方。违反任何一条都要立即停下来。**

1. **策略同构**：回测与实盘必须调用同一个 `Step()` 实现，内部禁止出现 `if isBacktest` 分支
2. **策略纯函数**：`Step()` 内部禁止定时器、网络请求、数据库读写、任何文件 I/O
3. **API Key 物理隔离**：交易所凭证只存在于 `config.agent.yaml`，永不进入 SaaS 侧
4. **GORM Code-First**：数据库 schema 真源是 Go struct，只用 AutoMigrate，永不写 SQL migration 文件
5. **无量纲计算**：价格相关计算使用对数收益率或比率，禁止跨标的比较绝对价格
6. **单一 Postgres**：不分库，Redis 仅做缓存，不作信号传递通道

---

## 二、Sigmoid 动态天平 — 微观引擎设计哲学

> 这是系统的核心创新，其余策略机制需自行设计填充。

### 本质：仓位弹簧系统

Sigmoid 动态天平用 Sigmoid 函数实时计算**目标持仓权重**，通过买卖浮动仓使实际权重趋近目标权重。它是一个与信号来源无关的通用框架——你可以把任何归一化的市场信号接入它。

### 核心公式：Sigmoid 目标权重

```
CurrentWeight  = FloatBTC × Price / TotalEquity

Signal         = 【你的市场信号，任意归一化标量，正值倾向减仓，负值倾向加仓】

EffectiveBeta  = max(0.01,  β × MarketBetaMultiplier)
InventoryBias  = clamp(CurrentWeight, 0, 1) − 0.5
Exponent       = EffectiveBeta × Signal + γ × InventoryBias

TargetWeight   = 1 / (1 + e^Exponent),  clamp(0, 1)
```

### 公式解读

| 场景 | Signal | Exponent 方向 | TargetWeight | 动作 |
|---|---|---|---|---|
| 你的信号看空 | 正 | 增大 | < 0.5 | 减仓 |
| 你的信号看多 | 负 | 减小 | > 0.5 | 加仓 |
| 仓位 > 0.5，γ > 0 | 任意 | 额外增大 | 进一步压低 | 均值回归 |
| 仓位 < 0.5，γ > 0 | 任意 | 额外减小 | 进一步拉高 | 均值回归 |

- **Signal**：你的策略信号，可以是均值回归信号、动量信号、突破信号，甚至多信号的加权合成——只要归一化为一个标量即可接入
- **β（激进系数）**：越大调仓越频繁，市场状态感知层可在极端行情时动态放大 β
- **γ（仓位偏置系数）**：`γ=0` 时纯信号驱动，`γ>0` 时叠加均值回归力防止仓位极端漂移
- **MarketBetaMultiplier**：来自上层市场状态感知，行情极端时放大 β 自动加速响应

### GA 在这里做什么

Signal 通常由多个因子线性合成：

```
Signal = a × X1 + b × X2 + c × X3 + ...
```

其中 X1、X2、X3 是你选择的市场特征（无量纲化后的标量），a、b、c 是对应的权重系数。这些系数就是**染色体的一部分**，由遗传算法在历史数据上搜索最优值。

举例：如果你设计了三个无量纲特征 X1（价格偏离类）、X2（动量类）、X3（加速度/突破类），合成信号就是 `a×X1 + b×X2 + c×X3`，GA 搜索的就是让这个合成信号喂进 Sigmoid 后，在多时段历史回测中跑赢被动 DCA 基准的最优 (a, b, c)。

你选什么特征完全由你的策略逻辑决定——均值回归、动量、突破、成交量异动、链上数据……GA 的搜索机制是完全通用的，它不关心 X 是什么，只负责找到让系统在历史上表现最好的系数组合。连同 β、γ 等 Sigmoid 本身的参数一起，整个微观引擎的参数空间就是一个高维搜索问题，GA 是求解它的工具。所以为了防止过拟合，你需要在回测中自行加入乱序回测或者蒙特卡洛算法等内容。因此类内容对性能的影响过于大，本次 plan 没有写入。

### 理论订单与楔形区过滤

```
DeltaWeight    = TargetWeight − CurrentWeight
TheoreticalUSD = DeltaWeight × TotalEquity
```

粉尘拦截规则（防止无效小额交易）：
- `|TheoreticalUSD| ≥ 最小阈值`：直接下单
- `|TheoreticalUSD| ∈ (0, 最小阈值)`：仅在非安静态 **且** 满足楔形突破条件时强制最小订单，否则归零
- 楔形突破条件：`|DeltaWeight| ≥ 阈值` 或 `VolatilityRatio ≥ 阈值`
- `VolatilityRatio = clip(MAV短期 / MAV长期, 0.1, 3.0)`，MAV 为平均绝对涨跌（非 ATR）