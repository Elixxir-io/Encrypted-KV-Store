///////////////////////////////////////////////////////////////////////////////
// Copyright © 2020 xx network SEZC                                          //
//                                                                           //
// Use of this source code is governed by a license that can be found in the //
// LICENSE file                                                              //
///////////////////////////////////////////////////////////////////////////////

package portableOS

import (
	"github.com/mattetti/filebuffer"
)

// JsFile represents a File for a Javascript value saved in local storage.
type JsFile struct {
	keyName string
	fb      *filebuffer.Buffer
	storage *jsStorage
}

// initFile creates a new in-memory file buffer of the key value.
func initFile(keyName, keyValue string, storage *jsStorage) *JsFile {
	f := &JsFile{
		keyName: keyName,
		fb:      filebuffer.New([]byte(keyValue)),
		storage: storage,
	}

	return f
}

// Close closes the File, rendering it unusable for I/O.
// On files that support SetDeadline, any pending I/O operations will
// be canceled and return immediately with an ErrClosed error.
// Close will return an error if it has already been called.
func (f *JsFile) Close() error {
	return f.fb.Close()
}

// Name returns the name of the file as presented to Open.
func (f *JsFile) Name() string {
	return f.keyName
}

// Read reads up to len(b) bytes from the File and stores them in b.
// It returns the number of bytes read and any error encountered.
// At end of file, Read returns 0, io.EOF.
func (f *JsFile) Read(b []byte) (n int, err error) {
	return f.fb.Read(b)
}

// ReadAt reads len(b) bytes from the File starting at byte offset off.
// It returns the number of bytes read and the error, if any.
// ReadAt always returns a non-nil error when n < len(b).
// At end of file, that error is io.EOF.
func (f *JsFile) ReadAt(b []byte, off int64) (n int, err error) {
	return f.fb.ReadAt(b, off)
}

// Seek sets the offset for the next Read or Write on file to offset,
// interpreted according to whence: 0 means relative to the origin of the
// file, 1 means relative to the current offset, and 2 means relative to the
// end. It returns the new offset and an error, if any. The behavior of Seek
// on a file opened with os.O_APPEND is not specified.
//
// If f is a directory, the behavior of Seek varies by operating system; you
// can seek to the beginning of the directory on Unix-like operating
// systems, but not on Windows.
func (f *JsFile) Seek(offset int64, whence int) (ret int64, err error) {
	return f.fb.Seek(offset, whence)
}

// Sync commits the current contents of the file to stable storage.
// Typically, this means flushing the file system's in-memory copy
// of recently written data to disk.
func (f *JsFile) Sync() error {
	f.storage.Set(f.keyName, f.fb.String())
	return nil
}

// Write writes len(b) bytes from b to the File.
// It returns the number of bytes written and an error, if any.
// Write returns a non-nil error when n != len(b).
func (f *JsFile) Write(b []byte) (n int, err error) {
	return f.fb.Write(b)
}

// JsFileInfo represents a FileInfo for a Javascript value saved in local
// storage.
type JsFileInfo struct {
	keyName string
	size    int64
}

// Name returns the base name of the file.
func (f *JsFileInfo) Name() string {
	return f.keyName
}

// Size returns the length in bytes.
func (f *JsFileInfo) Size() int64 {
	return f.size
}
