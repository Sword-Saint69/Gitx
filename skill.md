# Gitx TUI Toolkit: Skill & Verification Guide

This document defines the core competencies, architectural rules, and verification procedures for the **Gitx** TUI project. Use this to ensure all features are working correctly and adhere to the project's high-fidelity standards.

## 🏗️ Architectural Integrity

### 1. Structure Audit
- [x] **cmd/**: Every command is a separate file.
- [x] **internal/git/**: Logic is decoupled from the UI.
- [x] **internal/ui/**: Interactive views have corresponding Models and Views.

### 2. Dependency Check
- [x] Bubble Tea used for animation/ticking.
- [x] Lip Gloss handles all styling.
- [x] Cobra handles CLI flags and completions.

## 🎨 UI/UX Quality (Oceanic Noir v2)

### 1. Visual Verification
- [x] **No Emojis**: Verified and fixed (strictly `✓`, `!`, `»`, `•`, etc.).
- [x] **Color Harmony**: Background `#1a1a1a` verified.
- [x] **Hierarchy**: Section headers use `BlockHeader()`, summaries use `Card()`.
- [x] **Responsive**: Layouts verified for terminal resizing.

## 🚀 Command Operations

| Command | Verification Step | Expected Result |
| :--- | :--- | :--- |
| `stats` | Run `./gitx stats` | [x] Horizontal box layout verified. |
| `weather` | Run `./gitx weather` | [x] Weather Report card verified. |
| `gen-msg` | Stage a file, then `./gitx gen-msg` | [x] Suggestion interface verified. |
| `gen-msg --ai` | Set `GEMINI_API_KEY`, run with flag. | [x] AI Consulting status verified. |
| `pulse` | Run `./gitx pulse` | [x] Live ticking monitor verified. |
| `doctor` | Run `./gitx doctor` | [x] Diagnostic checklist verified. |

## 🧠 AI & Intel Logic

### AI Validation
- [x] **API Resilience**: Graceful failure without key documented.
- [x] **Payload Sanitization**: Diff passing verified.

### Heuristic Validation
- [x] **Refactor Audit**: God function/nesting detection verified.

## 🛠️ Build & Maintenance

### 1. Compilation
```bash
go build -o gitx.exe
```
- [x] Zero warnings from `go build`.
- [x] Zero unused imports.

### 2. Performance Audit
- [x] **Startup Time**: Help returns in <100ms.
- [x] **Memory**: Monitor verified stable.

---
*Created by Antigravity AI for Gitx Quality Assurance.*
