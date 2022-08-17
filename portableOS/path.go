///////////////////////////////////////////////////////////////////////////////
// Copyright © 2020 xx network SEZC                                          //
//                                                                           //
// Use of this source code is governed by a license that can be found in the //
// LICENSE file                                                              //
///////////////////////////////////////////////////////////////////////////////

package portableOS

import (
	"os"
)

// PathSeparator is the OS-specific path separator.
const PathSeparator rune = os.PathSeparator

// MkdirAll creates a directory named path, along with any necessary parents,
// and returns nil, or else returns an error. The permission bits perm (before
// umask) are used for all directories that MkdirAll creates. If path is already
// a directory, MkdirAll does nothing and returns nil.
var MkdirAll func(path string, perm FileMode) error = func(path string, perm FileMode) error {
	return os.MkdirAll(path, os.FileMode(5))
}
