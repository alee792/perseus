package store

import "fmt"

// A Version represents a specific version of a Go module
type Version struct {
	ID       int32  `json:"id" db:"id"`
	ModuleID string `json:"module_id" db:"module_id"`
	SemVer   string `json:"semver" db:"version"`
}

// A SemVer represents the components of a MAJOR.MINOR.PATCH-LABEL version,
// where label represents optional pre-release or build metadata.
type SemVer struct {
	Major int32  `json:"major" db:"major"`
	Minor int32  `json:"minor" db:"minor"`
	Patch int32  `json:"patch" db:"patch"`
	Label string `json:"label" db:"label"`
}

// String compounds the version components into a standard semver string without
// a the Go conventional 'v' prefix.
func (s *SemVer) String() string {
	return fmt.Sprintf("%d.%d.%d.-%s", s.Major, s.Minor, s.Patch, s.Label)
}
