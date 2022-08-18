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

// Stat returns a FileInfo describing the named file.
var Stat = func(name string) (FileInfo, error) {
	return os.Stat(name)
}
