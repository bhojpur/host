//go:build !linux
// +build !linux

package common

import (
	errsys "github.com/bhojpur/errors/pkg/system"
)

// Lgetxattr is not supported on platforms other than linux.
func Lgetxattr(path string, attr string) ([]byte, error) {
	return nil, errsys.ErrNotSupportedPlatform
}

// Lsetxattr is not supported on platforms other than linux.
func Lsetxattr(path string, attr string, data []byte, flags int) error {
	return errsys.ErrNotSupportedPlatform
}
