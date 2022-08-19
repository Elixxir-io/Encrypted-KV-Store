///////////////////////////////////////////////////////////////////////////////
// Copyright © 2020 xx network SEZC                                          //
//                                                                           //
// Use of this source code is governed by a license that can be found in the //
// LICENSE file                                                              //
///////////////////////////////////////////////////////////////////////////////

package portableOS

import (
	"os"
	"syscall/js"
)

// localStoragePrefix is prepended to every key stored in Javascript's
// localStorage to avoid collisions.
const localStoragePrefix = "xxdk" + string(os.PathSeparator)

var storage = newJsStorage("ekv", js.Global())

// jsStorage contains the Javascript localStorage object and a prefix that makes
// all keys saves unique to this instance.
type jsStorage struct {
	prefix string
	v      js.Value
}

// newJsStorage creates a new jsStorage and gets the localStorage.
func newJsStorage(prefix string, global js.Value) *jsStorage {
	return &jsStorage{
		prefix: prefix + string(os.PathSeparator),
		v:      global.Get("localStorage"),
	}
}

// Get gets the keyName's value from localStorage or returns os.ErrNotExist if
// it does not exist.
//
// This function wraps Javascript's [localStorage.getItem] method.
//
// [localStorage.getItem]: https://developer.mozilla.org/en-US/docs/Web/API/Storage/getItem
func (jss *jsStorage) Get(keyName string) (string, error) {
	result := jss.v.Call("getItem", jss.key(keyName))
	if result.IsNull() {
		return "", os.ErrNotExist
	}
	return result.String(), nil
}

// Set adds the keyValue to localStorage or update it if it already exists.
//
// This function wraps Javascript's [localStorage.setItem] method.
//
// [localStorage.setItem]: https://developer.mozilla.org/en-US/docs/Web/API/Storage/setItem
func (jss *jsStorage) Set(keyName, keyValue string) {
	jss.v.Call("setItem", jss.key(keyName), keyValue)
}

// Delete removes the keyName from the localStorage if it exists. If the key
// does not exist, this function does nothing.
//
// This function wraps Javascript's [localStorage.removeItem] method.
//
// [localStorage.removeItem]: https://developer.mozilla.org/en-US/docs/Web/API/Storage/removeItem
func (jss *jsStorage) Delete(keyName string) {
	jss.v.Call("removeItem", keyName)
}

// key prepends the necessary prefixes to the keyName.
func (jss *jsStorage) key(keyName string) string {
	return localStoragePrefix + jss.prefix + keyName
}
