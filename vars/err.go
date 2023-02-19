package vars

import (
	"errors"
	"fmt"
	"os"
)

//import "errors"
//var ErrCantOpenListFile error =  errors.New("cannot open file")
//var ErrCantWriteFile error = errors.New("cannot write new text")

var (
	// ErrInvalid indicates an invalid argument.
	// Methods on File will return this error when the receiver is nil.
	ErrInvalid = os.ErrInvalid // "invalid argument"

	ErrPermission = os.ErrPermission // "permission denied"
	ErrExist      = os.ErrExist      // "file already exists"
	ErrNotExist   = os.ErrNotExist   // "file does not exist"
	ErrClosed     = os.ErrClosed     // "file already closed"
)

var ErrNotSymlink = errors.New("file is not symlink")
var ErrLinkToDiffFile = errors.New("link to different file")

// 参考: https://0e0.pw/5SyM
type fileErr struct {
	Err  error
	Path string
}

func (e *fileErr) Error() string {
	return fmt.Sprintf("%v: %v", e.Path, e.Err.Error())
}

func ErrNoSuchEntry(path string) *fileErr {
	err := fileErr{
		Err:  errors.New("no entry which has such path"),
		Path: path,
	}
	return &err

}
