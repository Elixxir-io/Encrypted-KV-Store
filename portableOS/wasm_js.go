///////////////////////////////////////////////////////////////////////////////
// Copyright © 2020 xx network SEZC                                          //
//                                                                           //
// Use of this source code is governed by a license that can be found in the //
// LICENSE file                                                              //
///////////////////////////////////////////////////////////////////////////////

package portableOS

// This file is only compiled for WebAssembly.

import (
	"github.com/pkg/errors"
)

// Open opens the named file for reading. If successful, methods on the returned
// file can be used for reading.
var Open = func(name string) (File, error) {
	keyValue, err := storage.Get(name)
	if err != nil {
		return nil, errors.Errorf("could not open %q: %+v", name, err)
	}

	return initFile(name, keyValue, storage), nil
}

// Create creates or truncates the named file. If the file already exists, it is
// truncated. If the file does not exist, it is created. If successful, methods
// on the returned File can be used for I/O.
var Create = func(name string) (File, error) {
	storage.Set(name, "")
	return initFile(name, "", storage), nil
}

// Remove removes the named file or directory.
var Remove = func(name string) error {
	storage.Delete(name)
	return nil
}

// MkdirAll creates a directory named path, along with any necessary parents,
// and returns nil, or else returns an error. The permission bits perm (before
// umask) are used for all directories that MkdirAll creates. If path is already
// a directory, MkdirAll does nothing and returns nil.
var MkdirAll = func(path string, perm FileMode) error {
	return nil
}

// Stat returns a FileInfo describing the named file.
var Stat = func(name string) (FileInfo, error) {
	keyValue, err := storage.Get(name)
	if err != nil {
		return nil, errors.Errorf("could not open %q: %+v", name, err)
	}

	return &JsFileInfo{
		keyName: name,
		size:    int64(len(keyValue)),
	}, nil
}
