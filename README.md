<p align="center">
  <img src="Assets/banner.png" height=300 width="auto" alt="GitX Banner" />
</p>

# GitX

<p align="center">
  <b>The modern, interactive toolkit for Git power users.</b>
</p>

<p align="center">
  An ultra-fast suite of terminal-based tools designed to streamline your daily Git workflow with beautiful TUIs and powerful insights.
</p>

<div align="center">

  <img src="https://img.shields.io/badge/Status-Beta-blueviolet?style=for-the-badge" alt="Status" />
  <img src="https://img.shields.io/badge/Written%20In-Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Language" />
  <img src="https://img.shields.io/badge/License-MIT-green?style=for-the-badge" alt="License" />

</div>

<br>

## 🚀 The Core Philosophy
Git is powerful, but its command-line interface often feels like looking at the world through a keyhole. **GitX** expands that view, bringing interactive visualizations and specialized workflows directly to your terminal.

<table width="100%">
  <tr>
    <td width="33%" valign="top">
      <h3 align="center">📸 Persistence</h3>
      <p align="center">Create named "snapshots" of your work-in-progress without clunky stash management.</p>
    </td>
    <td width="33%" valign="top">
      <h3 align="center">☕ Accountability</h3>
      <p align="center">Generate comprehensive "Standup" reports across all your projects in seconds.</p>
    </td>
    <td width="33%" valign="top">
      <h3 align="center">🔍 Insight</h3>
      <p align="center">Interactive ownership analysis and reflog exploration to understand how code evolves.</p>
    </td>
  </tr>
</table>

<br>

## ✨ Key Features

### 📋 Daily Standup Reporter
Stop hunting through `git log`. Get a consolidated view of everything you've committed in the last 24 hours across multiple repositories.
```bash
gitx standup
```

### 📸 Lightweight Snapshots
Save your current state with a name. Think of it as a labeled checkpoint that doesn't interfere with your commit history or stash.
```bash
gitx snapshot create "feature-x-pre-refactor"
gitx snapshot list
```

### 🔎 Interactive Ownership
Analyze who owns which parts of the codebase with a rich terminal interface. Perfect for large-scale refactors and onboarding.
```bash
gitx ownership
```

<br>

## 🛠️ Tech Stack
Built for speed and resilience using the modern Go ecosystem.

| **Component**          | **Technology**                                                                                                         | **Description**                                                                     |
| :--------------------- | :--------------------------------------------------------------------------------------------------------------------- | :---------------------------------------------------------------------------------- |
| **CLI Framework**      | <img src="https://img.shields.io/badge/Cobra-v1.10-blue?style=flat-square" valign="middle" />                           | Industry-standard CLI command architecture.                                         |
| **TUI Engine**         | <img src="https://img.shields.io/badge/BubbleTea-Charms-EB5745?style=flat-square" valign="middle" />                   | Functional terminal UI framework for Go.                                            |
| **Styling**            | <img src="https://img.shields.io/badge/Lipgloss-Charms-00ADD8?style=flat-square" valign="middle" />                    | CSS-like terminal string styling.                                                   |
| **Git Backend**        | <img src="https://img.shields.io/badge/go--git-v5-green?style=flat-square" valign="middle" />                          | Pure Go implementation of Git for deep integration.                                 |

<br>

## 📦 Installation

### Pre-built Binaries
Download the latest executable from the [Releases](https://github.com/user/gitx/releases) page.

### From Source
```bash
git clone https://github.com/user/gitx.git
cd gitx
go build -o gitx
```

<br>

## 🗺️ Roadmap
- [x] **Core TUI Engine:** High-performance rendering with BubbleTea.
- [x] **Snapshot System:** Named checkpoints for work-in-progress.
- [x] **Standup Command:** Automated commit aggregation.
- [ ] **Interactive Reflog:** A visual way to navigate and restore from the reflog.
- [ ] **Remote Management:** TUI for managing multiple remotes and branch tracking.
- [ ] **Custom Style Themes:** User-definable color schemes for the entire suite.

<br>

## 📄 License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

<p align="center">
  <sub>Built with 💙 by Git power users, for Git power users.</sub>
</p>