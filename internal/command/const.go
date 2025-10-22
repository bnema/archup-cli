package command

// System commands
const (
	Go     = "go"
	Pacman = "pacman"
	Sudo   = "sudo"
	Lspci  = "lspci"
	Uname  = "uname"
)

// Go commands
const (
	GoInstall = "install"
)

// Pacman flags
const (
	PacmanSync      = "-S"
	PacmanNoConfirm = "--noconfirm"
	PacmanNeeded    = "--needed"
)

// Repository URLs
const (
	ArchupCLIRepo   = "github.com/bnema/archup-cli"
	ArchupCLILatest = ArchupCLIRepo + "@latest"
)
