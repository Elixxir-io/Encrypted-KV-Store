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

// Open opens the named file for reading. If successful, methods on the returned
// file can be used for reading; the associated file descriptor has mode
// os.O_RDONLY.
var Open func(name string) (File, error) = func(name string) (File, error) {
	return os.Open(name)
}

// Create creates or truncates the named file. If the file already exists, it is
// truncated. If the file does not exist, it is created with mode 0666 (before
// umask). If successful, methods on the returned File can be used for I/O; the
// associated file descriptor has mode os.O_RDWR.
var Create func(name string) (File, error) = func(name string) (File, error) {
	return os.Create(name)
}

// Remove removes the named file or directory.
var Remove func(name string) error = os.Remove
