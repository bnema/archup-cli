package hardware

// NOTE: This is a placeholder for CPU governor domain logic.
// Will contain business rules for selecting appropriate CPU governors.

// Governor represents a CPU frequency scaling governor
type Governor string

const (
	GovPerformance Governor = "performance"
	GovPowersave   Governor = "powersave"
	GovSchedutil   Governor = "schedutil"
	GovOndemand    Governor = "ondemand"
	GovConservative Governor = "conservative"
)

// GovernorInfo contains metadata about a CPU governor
type GovernorInfo struct {
	Name        Governor
	Description string
	PowerUsage  string // "high", "medium", "low"
	Performance string // "high", "medium", "low"
}

// AvailableGovernors returns all known governors with metadata
// PLACEHOLDER: Not implemented yet
func AvailableGovernors() []GovernorInfo {
	return []GovernorInfo{
		{
			Name:        GovPerformance,
			Description: "Maximum CPU performance (higher power consumption)",
			PowerUsage:  "high",
			Performance: "high",
		},
		{
			Name:        GovPowersave,
			Description: "Better battery life (lower performance)",
			PowerUsage:  "low",
			Performance: "low",
		},
		{
			Name:        GovSchedutil,
			Description: "Balanced - CPU scheduler-driven (recommended)",
			PowerUsage:  "medium",
			Performance: "medium",
		},
		{
			Name:        GovOndemand,
			Description: "Dynamic scaling based on load",
			PowerUsage:  "medium",
			Performance: "medium",
		},
		{
			Name:        GovConservative,
			Description: "Gradual scaling for smooth performance",
			PowerUsage:  "medium",
			Performance: "medium",
		},
	}
}

// RecommendGovernor suggests a governor based on hardware profile
// PLACEHOLDER: Not implemented yet
func RecommendGovernor(profile *Profile) Governor {
	if profile.IsLaptop {
		return GovPowersave
	}
	return GovSchedutil
}
