package version

// Version information, can be set at build time with ldflags
var (
	// Version is the current version of the application
	Version = "dev"

	// GitCommit is the git commit hash
	GitCommit = "unknown"

	// BuildDate is the date the binary was built
	BuildDate = "unknown"
)

// GetVersion returns the full version string
func GetVersion() string {
	return Version
}

// GetBuildInfo returns the build information
func GetBuildInfo() map[string]string {
	return map[string]string{
		"version":   Version,
		"commit":    GitCommit,
		"buildDate": BuildDate,
	}
}
