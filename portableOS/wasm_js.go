///////////////////////////////////////////////////////////////////////////////
// Copyright © 2020 xx network SEZC                                          //
//                                                                           //
// Use of this source code is governed by a license that can be found in the //
// LICENSE file                                                              //
///////////////////////////////////////////////////////////////////////////////

package portableOS

// This file is only compiled for WebAssembly.

import (
	"encoding/base64"
	"os"
	"strings"
	"syscall/js"
)

var storage = js.Global().Get("localStorage")

// directory is the contents of a value to signify that it represents a
// directory.
const directory = "\x1B[Directory]\x1B"

// Open opens the named file for reading. If successful, methods on the returned
// file can be used for reading.
var Open = func(name string) (File, error) {
	keyValue := storage.Call("getItem", name)
	if keyValue.IsNull() {
		return nil, os.ErrNotExist
	}

	s, err := base64.StdEncoding.DecodeString(keyValue.String())
	if err != nil {
		return nil, err
	}

	return open(name, string(s), storage), nil
}

// Create creates or truncates the named file. If the file already exists, it is
// truncated. If the file does not exist, it is created. If successful, methods
// on the returned File can be used for I/O.
var Create = func(name string) (File, error) {
	storage.Call("setItem", name, base64.StdEncoding.EncodeToString([]byte("")))

	return open(name, "", storage), nil
}

// Remove removes the named file or directory.
var Remove = func(name string) error {
	storage.Call("removeItem", name)
	return nil
}

// RemoveAll removes path and any children it contains.
// It removes everything it can but returns the first error
// it encounters. If the path does not exist, RemoveAll
// returns nil (no error).
// If there is an error, it will be of type *PathError.
var RemoveAll = func(path string) error {
	for i := 0; i < storage.Get("length").Int(); i++ {
		keyName := storage.Call("key", i).String()
		result := storage.Call("getItem", keyName)
		if result.IsNull() {
			return os.ErrNotExist
		}
		if strings.HasPrefix(keyName, path) {
			storage.Call("removeItem", keyName)
		}
	}

	return nil
}

// MkdirAll creates a directory named path, along with any necessary parents,
// and returns nil, or else returns an error. The permission bits perm (before
// umask) are used for all directories that MkdirAll creates. If path is already
// a directory, MkdirAll does nothing and returns nil.
var MkdirAll = func(path string, perm FileMode) error {
	storage.Call("setItem", path, base64.StdEncoding.EncodeToString(
		[]byte(directory)))
	open(path, "", storage)
	return nil
}

// Stat returns a FileInfo describing the named file.
var Stat = func(name string) (FileInfo, error) {
	keyValue := storage.Call("getItem", name)
	if keyValue.IsNull() {
		return nil, os.ErrNotExist
	}

	s, err := base64.StdEncoding.DecodeString(keyValue.String())
	if err != nil {
		return nil, err
	}

	return &jsFileInfo{
		keyName: name,
		size:    int64(len(s)),
		isDir:   string(s) == directory,
	}, nil
}
