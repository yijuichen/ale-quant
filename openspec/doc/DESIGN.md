---
version: alpha
name: QuantSaaS Console
description: Dark trading-console UI for QuantSaaS web frontend (React, Tailwind v4). Tokens mirror web-frontend/src/index.css @theme and qs-* utility classes.

colors:
  primary: "#2dd4bf"
  on-primary: "#0f1115"
  secondary: "#64748b"
  tertiary: "#0ea5e9"
  neutral: "#0f1115"
  surface: "#020617"
  surface-elevated: "#0f1115"
  on-surface: "#e2e8f0"
  on-surface-muted: "#94a3b8"
  border-subtle: "#1e293b"
  accent-text: "#9ff7e0"
  selection: "#00a8ff"
  error: "#f87171"
  warn: "#fbbf24"
  safe: "#34d399"
  backdrop-warm: "#ff8c6b"
  scrollbar-thumb: "#334155"
  scrollbar-thumb-hover: "#475569"

typography:
  headline-display:
    fontFamily: Inter
    fontSize: 36px
    fontWeight: 600
    lineHeight: 1.15
    letterSpacing: -0.02em
  headline-lg:
    fontFamily: Inter
    fontSize: 24px
    fontWeight: 600
    lineHeight: 1.25
    letterSpacing: -0.01em
  headline-md:
    fontFamily: Inter
    fontSize: 18px
    fontWeight: 600
    lineHeight: 1.35
  body-lg:
    fontFamily: Inter
    fontSize: 16px
    fontWeight: 400
    lineHeight: 1.6
  body-md:
    fontFamily: Inter
    fontSize: 14px
    fontWeight: 400
    lineHeight: 1.55
  body-sm:
    fontFamily: Inter
    fontSize: 12px
    fontWeight: 400
    lineHeight: 1.5
  label-lg:
    fontFamily: JetBrains Mono
    fontSize: 12px
    fontWeight: 500
    lineHeight: 1.25
    letterSpacing: 0.06em
  label-md:
    fontFamily: Inter
    fontSize: 10px
    fontWeight: 700
    lineHeight: 1
    letterSpacing: 0.22em
  label-sm:
    fontFamily: JetBrains Mono
    fontSize: 11px
    fontWeight: 400
    lineHeight: 1.3
  label-cta:
    fontFamily: Inter
    fontSize: 12px
    fontWeight: 600
    lineHeight: 1
    letterSpacing: 0.16em

rounded:
  none: 0px
  sm: 4px
  md: 8px
  lg: 12px
  xl: 16px
  full: 9999px

spacing:
  xs: 4px
  sm: 8px
  md: 16px
  lg: 24px
  xl: 32px
  gutter: 24px
  panel-padding-x: 20px
  panel-header-y: 16px
  bento-gap: 24px

components:
  panel:
    backgroundColor: "{colors.surface}"
    rounded: "{rounded.lg}"
    typography: "{typography.body-md}"
  panel-elevated:
    backgroundColor: "{colors.surface}"
    rounded: "{rounded.lg}"
  card-header:
    padding: "{spacing.panel-padding-x}"
    typography: "{typography.label-md}"
  button-primary:
    backgroundColor: "{colors.primary}"
    textColor: "{colors.accent-text}"
    rounded: "{rounded.lg}"
    padding: 12px
    typography: "{typography.label-cta}"
  button-primary-hover:
    backgroundColor: "{colors.primary}"
  button-ghost:
    backgroundColor: "#000000"
    textColor: "{colors.on-surface-muted}"
    rounded: "{rounded.lg}"
    typography: "{typography.label-sm}"
  button-danger:
    backgroundColor: "{colors.error}"
    textColor: "{colors.error}"
    rounded: "{rounded.lg}"
    padding: 10px
  input:
    backgroundColor: "#000000"
    textColor: "{colors.on-surface}"
    rounded: "{rounded.lg}"
    padding: 10px
  notice-error:
    backgroundColor: "{colors.error}"
    textColor: "{colors.error}"
    rounded: "{rounded.lg}"
  notice-warn:
    backgroundColor: "{colors.warn}"
    textColor: "{colors.warn}"
    rounded: "{rounded.lg}"
  notice-safe:
    backgroundColor: "{colors.safe}"
    textColor: "{colors.safe}"
    rounded: "{rounded.lg}"
  notice-info:
    backgroundColor: "{colors.primary}"
    textColor: "{colors.accent-text}"
    rounded: "{rounded.lg}"
---

## Overview

QuantSaaS Console is a dark, data-dense operator interface for quantitative trading and lab workflows. The voice is institutional and calm: deep slate and ink backgrounds, one dominant teal accent for progression and commitment, minimal ornament, and monospace only where precision matters (IDs, telemetry, dense tables).

The shell stacks a fixed sidebar and top bar over a full-viewport atmospheric backdrop (soft coral, sky, and teal glows with subtle grain). Content uses a **bento-style grid** on large screens so dashboards read as modular instruments rather than long scrolling documents.

Agents implementing UI should preserve **WCAG-aware contrast** on body copy (`on-surface` on `surface`), reserve **primary teal** for the single strongest affordance per focused region, and use **error red** only for destructive actions and failure states.

## Colors

The palette is anchored on near-black foundations and cold slate text, with teal as the sole primary interaction hue. Supporting hues separate concerns without competing with teal: sky for informational variety in marketing-style backdrops and strategy tone `sky`, amber for warmth and warnings, emerald for positive execution health.

- **Primary (#2dd4bf):** Teal accent for CTAs, focus rings, panel glow accents, and success-path emphasis. Pair labels with **accent text (#9ff7e0)** on tinted fills for readability.
- **On primary (#0f1115):** Deep ink used when stacking opaque text or icons directly on saturated teal if ever required; default chrome uses translucent teal fills instead.
- **Secondary (#64748b):** Slate for secondary labels, sidebar metadata, and de-emphasized navigation.
- **Tertiary (#0ea5e9):** Sky accent for alternate strategy tone highlights and ambient backdrop gradients; not a default text color.
- **Neutral (#0f1115):** Canonical app background (`qs-bg` / body); slightly lifted from pure black for OLED-friendly depth.
- **Surface (#020617):** Panel and card interior base; aligns with `qs-panel` inner stack.
- **Surface elevated (#0f1115):** Outer frame when separating from backdrop.
- **On surface (#e2e8f0):** Primary readable text on dark surfaces (slate-200).
- **On surface muted (#94a3b8):** Captions and helper lines (slate-400).
- **Border subtle (#1e293b):** Dividers and hairlines; implementation often uses white at fractional opacity instead for blending.
- **Accent text (#9ff7e0):** High-legibility tint for uppercase CTAs on teal-tinted buttons.
- **Selection (#00a8ff):** Text selection highlight base (applied with transparency in CSS).
- **Error (#f87171):** Failures, validation, destructive emphasis.
- **Warn (#fbbf24):** Non-fatal alerts and reconcile prompts.
- **Safe (#34d399):** Healthy runtime and positive confirmations.
- **Backdrop warm (#ff8c6b):** Decorative ambient orb only; never for semantic status alone.
- **Scrollbar thumb (#334155 / #475569):** Scrollbars stay narrow and low-contrast so data panels remain focal.

Implementation note: `--color-qs-surface` in CSS uses `rgba(255,255,255,0.04)` for glass overlays; tokens above are **solid sRGB anchors** for agents and linting. Prefer semantic tokens and match opacity in code when reproducing glass.

### Design Tokens

Normative color tokens are defined in the YAML frontmatter under `colors`.

## Typography

**Inter** carries all UI chrome, headings, and body copy at weights 400, 500, and 600. **JetBrains Mono** appears for compact numeric columns, hashes, and technical annotations where alignment beats warmth.

Hierarchy rules:

- Page titles and hero metrics use **headline-display** or **headline-lg**.
- Section titles inside cards use **headline-md**.
- Dense tables and forms default to **body-md**; secondary lines use **body-sm** at muted color.
- Uppercase section rails (`.qs-section-label`) use **label-md**: 10px, bold, wide tracking.
- Primary and ghost buttons use uppercase **label-cta** or mono **label-sm** for ghost variants.

### Design Tokens

Normative typography tokens are defined in the YAML frontmatter under `typography`.

## Layout

Primary layout is **sidebar + main column** inside a full-height flex shell without document scroll on the outer frame; inner routes scroll independently.

Grid content uses **qs-bento-grid**: `gap` corresponds to spacing token **lg (24px)**. Desktop (min-width 1024px) uses four columns with selective spans; tablet collapses to two columns; mobile is a single column. Prefer consistent **gutter (24px)** between cards unless a tighter analytic stack explicitly calls for **md (16px)**.

Card chrome uses **panel-padding-x (20px)** on headers (`.qs-card-header` uses `px-5`). Vertical rhythm between stacked sections typically alternates **md** and **lg** spacing.

Touch targets for frequent actions should meet at least **44px** vertical extent where possible; dense toolbars may use **button-primary-sm** with smaller padding but should remain keyboard-focusable.

### Design Tokens

Spacing tokens are defined in the YAML frontmatter under `spacing`.

## Elevation & Depth

Depth is conveyed through **layered translucency**, not skeuomorphic elevation. Panels use `backdrop-blur`, inner teal-tinted inset glow (`qs-panel`), and optional outer shadow (`qs-panel-elevated`). Hover may deepen shadow slightly on cards.

Flat alternatives for hierarchy: **border-subtle** outlines, **stat cells** with darker inset fills (`#000000` at fractional opacity), and tonal strategy accents (`teal`, `amber`, `sky`, `slate` tone classes) rather than competing drop shadows.

Ambient background gradients remain **non-interactive** and stay behind `z-index` structured chrome.

## Shapes

Corners are predominantly **12px (`rounded.lg`)** on panels, cards, inputs, buttons, notices, and stat tiles. Outer card wrappers may use **16px (`rounded.xl`)** on the motion frame. **Full** radius is reserved for dots and pills. Scrollbar thumbs use **4px** half-cylindrical ends.

Avoid mixing ad-hoc radii within a single card; inner elements follow the same 12px language unless an icon badge explicitly uses **md (8px)** (`rounded-xl` on 32px badges is 12px in the component; token `lg` aligns with Tailwind `rounded-xl`).

### Design Tokens

Corner tokens are defined in the YAML frontmatter under `rounded`.

## Components

**Panels and cards:** `qs-panel` provides the frosted slab; `Card` wraps with optional `qs-panel-elevated` shadow. Headers use `qs-card-header` with bottom border at low white alpha.

**Buttons:** `qs-btn-primary` and variants are the default commitment control. `qs-btn-ghost` is for secondary navigation. `qs-btn-danger` is strictly destructive. Disabled states reduce opacity (`0.5`–`0.6`) and drop pointer cursor.

**Inputs:** `qs-input` uses dark fill, subtle border, teal focus ring via border tint.

**Notices:** `qs-notice-ok`, `qs-notice-err`, `qs-notice-warn`, `qs-notice-info` align border and fill opacity by semantic color.

**Strategy tones:** Strategy catalog maps optional **teal / amber / sky / slate** panel and badge treatments for template cards; default unknown strategies use slate.

Motion: entrance uses short fade-up with easing `[0.23, 1, 0.32, 1]` unless `prefers-reduced-motion` is set, in which case motion is suppressed.

### Design Tokens

Component-level defaults are defined in the YAML frontmatter under `components`. Properties follow the DESIGN.md spec (`backgroundColor`, `textColor`, `typography`, `rounded`, `padding`, etc.).

## Do's and Don'ts

- Do use **primary teal** for the primary action in a card or modal; avoid multiple competing teal-filled buttons in the same viewport region.
- Do use **error** only for irreversible or harmful actions and critical failures; prefer **warn** for recoverable trading halts or reconciliations.
- Do keep body text at **on-surface** or **on-surface-muted**; avoid placing long prose in **accent-text** alone.
- Do respect **reduced motion**: disable decorative pulse and shimmer when the user prefers reduced motion (already mirrored in CSS).
- Don't introduce ad-hoc neon accents outside teal, sky, amber, red, emerald semantic roles used in code.
- Don't use emoji or decorative unicode symbols in production UI strings (product convention).
- Don't rely on backdrop gradient colors alone to communicate engine state; always pair with explicit labels or badges.

---

Format compliance: This file follows the open **DESIGN.md** specification maintained by Google Labs (`google-labs-code/design.md`, `docs/spec.md`, version **alpha**). Validate or export with the reference tooling when available: `npx @google/design.md lint` and related commands from that project.
