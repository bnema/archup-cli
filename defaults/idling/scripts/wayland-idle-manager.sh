#!/bin/bash
# Wayland Idle Manager - WM-agnostic idle management script
# Supports: niri, hyprland, sway
# Usage: wayland-idle-manager.sh <action> [options]
#   Actions: lock-screen, power-off-monitors, power-on-monitors, restore-brightness

set -euo pipefail

# Configuration
readonly BRIGHTNESS_FILE="${XDG_RUNTIME_DIR:-/tmp}/wayland_brightness_before_idle"
readonly SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
readonly LOG_FILE="${XDG_RUNTIME_DIR:-/tmp}/wayland-idle-manager.log"

# Logging
log() {
    echo "[$(date '+%Y-%m-%d %H:%M:%S')] $*" >> "$LOG_FILE"
}

error() {
    log "ERROR: $*"
    echo "ERROR: $*" >&2
}

# Detect running Wayland WM
detect_wm() {
    if pgrep -x niri > /dev/null 2>&1; then
        echo "niri"
    elif pgrep -x hyprland > /dev/null 2>&1; then
        echo "hyprland"
    elif pgrep -x sway > /dev/null 2>&1; then
        echo "sway"
    else
        error "No supported Wayland WM detected"
        return 1
    fi
}

# Prevent multiple hyprlock instances
lock_screen() {
    log "Action: lock_screen"
    if ! pgrep -x hyprlock > /dev/null 2>&1; then
        if command -v hyprlock &> /dev/null; then
            hyprlock -v
        else
            error "hyprlock not found"
            return 1
        fi
    else
        log "hyprlock already running, skipping"
    fi
}

# Power off monitors and save brightness
power_off_monitors() {
    log "Action: power_off_monitors"
    local wm
    wm=$(detect_wm) || return 1

    # Save current brightness before turning off
    if command -v brightnessctl &> /dev/null; then
        if brightnessctl g > "$BRIGHTNESS_FILE" 2>/dev/null; then
            log "Brightness saved: $(cat "$BRIGHTNESS_FILE")"
        else
            echo "100" > "$BRIGHTNESS_FILE"
            log "Brightness file created with default value"
        fi
    fi

    case "$wm" in
        niri)
            niri msg action power-off-monitors || error "Failed to power off monitors on niri"
            ;;
        hyprland)
            hyprctl dispatch dpms off || error "Failed to power off monitors on hyprland"
            ;;
        sway)
            swaymsg "output * dpms off" || error "Failed to power off monitors on sway"
            ;;
        *)
            error "Unknown WM: $wm"
            return 1
            ;;
    esac
    log "Monitors powered off"
}

# Power on monitors
power_on_monitors() {
    log "Action: power_on_monitors"
    local wm
    wm=$(detect_wm) || return 1

    case "$wm" in
        niri)
            niri msg action power-on-monitors || error "Failed to power on monitors on niri"
            ;;
        hyprland)
            hyprctl dispatch dpms on || error "Failed to power on monitors on hyprland"
            ;;
        sway)
            swaymsg "output * dpms on" || error "Failed to power on monitors on sway"
            ;;
        *)
            error "Unknown WM: $wm"
            return 1
            ;;
    esac
    log "Monitors powered on"
}

# Restore brightness from saved value
restore_brightness() {
    log "Action: restore_brightness"
    if ! command -v brightnessctl &> /dev/null; then
        log "brightnessctl not found, skipping brightness restore"
        return 0
    fi

    if [ -f "$BRIGHTNESS_FILE" ]; then
        local saved_brightness
        saved_brightness=$(cat "$BRIGHTNESS_FILE" 2>/dev/null || echo "")
        if [ -n "$saved_brightness" ] && [ "$saved_brightness" -gt 0 ]; then
            if brightnessctl s "$saved_brightness" 2>/dev/null; then
                log "Brightness restored to: $saved_brightness"
            else
                error "Failed to restore brightness"
                return 1
            fi
        fi
    else
        log "No saved brightness file found"
    fi
}

# Main entry point
main() {
    local action="${1:-}"

    if [ -z "$action" ]; then
        error "No action specified"
        echo "Usage: wayland-idle-manager.sh <action>"
        echo "Actions: lock-screen, power-off-monitors, power-on-monitors, restore-brightness"
        return 1
    fi

    log "Starting action: $action"

    case "$action" in
        lock-screen)
            lock_screen
            ;;
        power-off-monitors)
            power_off_monitors
            ;;
        power-on-monitors)
            power_on_monitors
            ;;
        restore-brightness)
            restore_brightness
            ;;
        *)
            error "Unknown action: $action"
            return 1
            ;;
    esac

    log "Action completed: $action"
}

main "$@"
