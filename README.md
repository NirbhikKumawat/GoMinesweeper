# Minesweeper TUI
A high-performance, terminal-based Minesweeper game built with Go. This project demonstrates a professional CLI architecture using the "Golden Stack" of Go terminal development: __Cobra__, __Viper__, and __Bubble Tea__.

## Technical Stack
- __Language__: Go (1.21+)
- __CLI Framework__: Cobra (Standard for Kubernetes/Docker CLIs)
- __TUI Engine__: Bubble Tea (Functional state management)
- __Styling__: Lip Gloss & ANSI Escape Sequences
- __Configuration__: Viper 

## Architectural Overview
This project follows the Standard Go Project Layout to ensure a strict separation of concerns:

| Layer | Package            | Responsibility                                                                 |
|-------|--------------------|--------------------------------------------------------------------------------|
| CLI   | `cmd/`             | Handles flag parsing (--mines, --rows) and command execution.                  |
| Logic | `internal/engine/` | Pure game logic: mine placement, neighbor counting, and flood-fill algorithms. |
| UI    | `internal/ui/`     | The "View" layer. Implements the Bubble Tea Update/View lifecycle.             |


## Design Patterns Used

- __Model-View-Update (MVU)__: Used via Bubble Tea to handle UI states without side effects.
- __Dependency Injection__: The Game Engine is injected into the UI Model, making the core logic testable without a terminal.
- __Encapsulation__: All game internals are kept in `internal/` to prevent external package leakage.


## Key Features
- __In-Place Rendering__: Uses ANSI escape codes for smooth, flicker-free terminal updates.
- __APT-Style Progress Bar__: Real-time visual feedback on board clearance using the `bubbles/progress` component.
- __Dynamic Difficulty__: Fully configurable board dimensions and mine density via CLI flags.
- __Color-Coded Interface__: High-contrast ANSI colors for numerical proximity and hazard warnings.

## Getting Started
### Installation
```bash
go install
```
### Usage
```bash
# Start with default settings (10x10, 12 mines)
minesweeper
# Custom difficulty
minesweeper --rows 15 --cols 20 --mines 40
```
## Engineering Challenges Solved
1. __The Recursive Reveal (Flood Fill)__
Implemented an optimized flood-fill algorithm that automatically clears empty areas when a "0" cell is revealed. This ensures the game feels fluid and follows classic Minesweeper mechanics.
2. __State Management in TUI__
Managed the transition between "Input Mode" (typing coordinates) and "Action Mode" (revealing/flagging) within the Bubble Tea event loop, ensuring a 0-latency user experience.