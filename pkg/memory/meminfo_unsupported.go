//go:build !linux && !windows
// +build !linux,!windows

package memory

import (
	errsys "github.com/bhojpur/errors/pkg/system"
)

// ReadMemInfo is not supported on platforms other than linux and windows.
func ReadMemInfo() (*MemInfo, error) {
	return nil, errsys.ErrNotSupportedPlatform
}
