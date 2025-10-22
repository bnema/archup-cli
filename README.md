# archup-cli

Desktop wizard for ArchUp - Install Tier 2 (Desktop Foundation) + Tier 3 (Compositor)

## Usage

```bash
archup wizard          # Run interactive setup
archup version         # Show version
archup --debug wizard  # Debug mode
```

## What It Does

After ArchUp Tier 1 (barebone) install, this wizard adds:

- **Tier 2**: Desktop foundation (~50 packages)
  - Graphics, audio, Wayland, Bluetooth, printing
- **Tier 3**: Compositor + apps (30-50 packages)
  - Niri, Hyprland, Sway, or River

## Development

```bash
# Build
go build -o archup .

# Run
./archup wizard

# Extract packages from Omarchy
./scripts/extract-foundation.sh
```

## License

MIT
