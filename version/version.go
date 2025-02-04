package version

import (
	vgversion "code.vegaprotocol.io/shared/libs/version"
)

const (
	ReleasesAPI        = "https://api.github.com/repos/vegaprotocol/vegawallet/releases"
	ReleasesURL        = "https://github.com/vegaprotocol/vegawallet/releases"
	defaultVersionHash = "unknown"
	defaultVersion     = "v0.17.0+dev"
)

var (
	// VersionHash specifies the git commit used to build the application.
	// See VERSION_HASH in Makefile for details.
	VersionHash = defaultVersionHash

	// Version specifies the version used to build the application.
	// See VERSION in Makefile for details.
	Version = defaultVersion
)

func IsUnreleased() bool {
	return vgversion.IsUnreleased(Version)
}

type GetVersionResponse struct {
	Version string `json:"version"`
	GitHash string `json:"gitHash"`
}

func GetVersionInfo() *GetVersionResponse {
	return &GetVersionResponse{
		Version: Version,
		GitHash: VersionHash,
	}
}
