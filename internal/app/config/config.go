package config

var (
	// BuildTime is a time label of the moment when the binary was built
	BuildTime string = "unset"
	// Commit is a last commit hash at the moment when the binary was built
	Commit string = "unset"
	// Release is a semantic version of current build
	Release string = "unset"
)

func VersionStr() string {
	return "Version: " + Release + " / BuiidTime" + BuildTime + " / Commit: " + Commit
}
