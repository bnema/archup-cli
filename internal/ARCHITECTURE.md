# ArchUp CLI Architecture

This document describes the domain-driven architecture of ArchUp CLI.

## Overview

ArchUp follows a **pragmatic domain-driven design** approach with clear separation of concerns:

```
internal/
├── domain/        # Core business logic (framework-independent)
├── app/           # Application services (orchestration)
├── infra/         # Infrastructure implementations
└── tui/           # User interface (Bubbletea)
```

## Layers

### 1. Domain Layer (`internal/domain/`)

**Pure business logic with no external dependencies.**

- `hardware/` - Hardware detection and profiling
  - `Profile` - Value object representing system hardware
  - `Detector` - Interface for hardware detection
  - `Governor` - CPU governor domain logic

- `packages/` - Package selection and management
  - `Selection` - Aggregate for user's package choices
  - `Group` - Package grouping (terminals, browsers, etc.)
  - `Repository` - Interface for package queries

- `config/` - System configuration
  - `System` - Root aggregate combining all configs
  - `CPUConfig` - CPU governor and frequency settings
  - `PowerConfig` - TLP power management settings
  - `PacmanConfig` - Pacman configuration

- `preset/` - Themes and configuration presets
  - `Preset` - Root aggregate for themes (Omarchy, Bleu, Stock)
  - `Repository` - Interface for preset storage
  - `Renderer` - Template rendering for dotfiles

- `installation/` - Installation orchestration
  - `Session` - Root aggregate coordinating entire installation
  - `Plan` - Ordered list of installation steps
  - `Executor` - Interface for executing plans

### 2. Application Layer (`internal/app/`)

**Orchestrates domain objects and coordinates use cases.**

- `services/` - Application services
  - `InstallationService` - Main service coordinating installation flow
  - `HardwareService` - Hardware detection workflows (future)
  - `PackageService` - Package resolution workflows (future)

### 3. Infrastructure Layer (`internal/infra/`)

**Concrete implementations of domain interfaces.**

- `hwdetect/` - Hardware detection implementations
  - Uses: `lscpu`, `lspci`, `dmidecode`, `/proc/cpuinfo`

- `pacman/` - Pacman package manager operations
  - Executes: `pacman` commands
  - Parses: pacman output into domain objects

- `systemd/` - Systemd service management
  - Executes: `systemctl` commands

- `filesystem/` - File operations
  - Preset loading from disk/git
  - Dotfile deployment
  - Backup/restore operations

### 4. Interface Layer (`internal/tui/`)

**User interface - adapts domain for presentation.**

- `layouts/` - Wizard layouts (linear, tree, hybrid)
- `components/` - Reusable UI components

## Data Flow

```
User Input (TUI)
    ↓
Application Service
    ↓
Domain Logic (business rules)
    ↓
Infrastructure (system calls)
    ↓
System Changes
```

## Example: Installation Flow

1. **TUI** collects user selections
2. **Application Service** creates `Session` aggregate
3. **Domain** generates `Plan` with ordered steps
4. **Infrastructure** executes plan using `Executor`
5. **Domain Events** update progress
6. **TUI** displays progress to user

## Key Principles

### Domain Layer
- ✅ No dependencies on infrastructure or frameworks
- ✅ Contains all business rules
- ✅ Easily testable without mocks
- ✅ Defines interfaces that infrastructure implements

### Application Layer
- ✅ Coordinates domain objects
- ✅ Implements use cases
- ✅ Thin orchestration layer

### Infrastructure Layer
- ✅ Implements domain interfaces
- ✅ Handles external dependencies (filesystem, system calls)
- ✅ Can be swapped without changing domain

### Interface Layer
- ✅ Thin adapter to domain
- ✅ Minimal logic (UI state only)
- ✅ Could be replaced with CLI, web API, etc.

## Migration Strategy

Current implementation has domain logic embedded in TUI. Migration path:

1. ✅ Create domain packages (DONE - placeholders)
2. ⏳ Implement hardware detection in `infra/hwdetect`
3. ⏳ Move `HardwareInfo` from TUI to `domain/hardware`
4. ⏳ Implement `InstallationService` for real installation
5. ⏳ Gradually thin TUI to pure UI state

## Testing Strategy

```
Domain Layer:     Unit tests (no mocks needed)
Application Layer: Integration tests (mock infrastructure)
Infrastructure:   Integration tests (test doubles for system)
TUI:             E2E tests (optional)
```

## Status

**Current State:** Placeholder structure created with interfaces and stubs.

**Next Steps:**
1. Implement `hwdetect.SystemDetector` for real hardware detection
2. Move TUI's `HardwareInfo` to domain
3. Implement `InstallationService.Install()` for actual installation
4. Build `Executor` to execute installation plans

## Notes

- All files marked with `// PLACEHOLDER: Not implemented yet`
- Domain interfaces are stable, implementations will be added incrementally
- This is **pragmatic** DDD - not full tactical patterns (no repositories everywhere)
- Focus on separation of concerns and testability
