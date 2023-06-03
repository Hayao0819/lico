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

	errPermission = os.ErrPermission // "permission denied"
	errExist      = os.ErrExist      // "file already exists"
	errNotExist   = os.ErrNotExist   // "file does not exist"
	errClosed     = os.ErrClosed     // "file already closed"
)

var errNotSymlink = errors.New("file is not symlink")
var errLinkToDiffFile = errors.New("link to different file")
var ErrNoRepoDir = errors.New("repository has not been cloned. Please run init command")
var errNotManaged = errors.New("this file is not managed by lico")

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

func ErrNotSymlink(path string) *fileErr{
	return &fileErr{
		Err: errNotSymlink,
		Path: path,
	}
}

func ErrLinkToDiffFile(path string) *fileErr{
	return &fileErr{
		Err: errLinkToDiffFile,
		Path: path,
	}
}

func ErrPermission(path string) *fileErr{
	return &fileErr{
		Err: errPermission,
		Path: path,
	}
}

func ErrExist(path string) *fileErr{
	return &fileErr{
		Err: errExist,
		Path: path,
	}
}

func ErrNotExist(path string) *fileErr{
	return &fileErr{
		Err: errNotExist,
		Path: path,
	}
}

func ErrClosed(path string)*fileErr{
	return &fileErr{
		Err: errClosed,
		Path: path,
	}
}

func ErrNotManaged(path string)*fileErr{
	return &fileErr{
		Err: errNotManaged,
		Path: path,
	}
}
