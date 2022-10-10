package store

import (
	"fmt"
	"strconv"
	"strings"
)

// A Version represents a specific version of a Go module. The Version can be
// turned into a more granular SemVer with ParseSemVer.
type Version struct {
	ID       int32  `json:"id" db:"id"`
	ModuleID string `json:"module_id" db:"module_id"`
	Version  string `json:"version" db:"version"`
}

// A SemVer represents the components of a MAJOR.MINOR.PATCH-LABEL version,
// where label represents an optional pre-release tag. Build metadata is not supported.
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

// ParseSemVer parses a string into a SemVer struct.
func ParseSemVer(s string) (SemVer, error) {
	// Accommodate Go style semver prefixes, like v1.2.3-pre.
	s = strings.TrimPrefix(s, "v")

	// Ex. ["1", "2", "3-pre"]
	mmpl := strings.Split(s, ".")
	if len(mmpl) != 3 {
		return SemVer{}, fmt.Errorf("expected 3 parts, got %d", len(mmpl))
	}

	// Ex. "1"
	major, err := strconv.ParseInt(mmpl[0], 10, 32)
	if err != nil {
		return SemVer{}, fmt.Errorf("invalid major version: %s", mmpl[0])
	}

	// Ex. "2"
	minor, err := strconv.ParseInt(mmpl[1], 10, 32)
	if err != nil {
		return SemVer{}, fmt.Errorf("invalid minor version: %s", mmpl[1])
	}

	sv := SemVer{
		Major: int32(major),
		Minor: int32(minor),
	}

	// Handle the remaining <patch>-<label>.
	// Ex. "3-pre"
	pl := strings.Split(mmpl[2], "-")
	if len(pl) == 2 {
		sv.Label = pl[1]
	}

	patch, err := strconv.ParseInt(pl[0], 10, 32)
	if err != nil {
		return SemVer{}, fmt.Errorf("invalid patch version: %s", pl[0])

	}
	sv.Patch = int32(patch)

	return sv, nil
}
