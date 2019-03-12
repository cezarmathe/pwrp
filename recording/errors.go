package recording

import (
	"errors"
	"pppi/git"
)

var (
	/*ErrNoRepositories indicates that no repositories were specified in the config*/
	ErrNoRepositories = errors.New("no repositories given")
)

/*ErrMissingBranch describes a missing branch*/
type ErrMissingBranch struct {
	Branch string
}

/*NewErrMissingBranch returns a new ErrMissingBranch*/
func NewErrMissingBranch(branch string) *ErrMissingBranch {
	return &ErrMissingBranch{branch}
}

func (err *ErrMissingBranch) Error() string {
	return "metadata branch " + err.Branch + " not found"
}

/*ErrNoPermissions describes a missing branch*/
type ErrNoPermissions struct {
	Path string
}

/*NewErrNoPermissions returns a new ErrNoPermissions*/
func NewErrNoPermissions(path string) *ErrNoPermissions {
	return &ErrNoPermissions{path}
}

func (err *ErrNoPermissions) Error() string {
	return "no permissions for writing data at " + err.Path
}

/*ErrBadURL describes a missing branch*/
type ErrBadURL struct {
	URL string
}

/*NewErrBadURL returns a new ErrBadURL*/
func NewErrBadURL(url string) *ErrBadURL {
	return &ErrBadURL{url}
}

func (err *ErrBadURL) Error() string {
	return "bad url: " + err.URL
}

/*ErrBadProtocol describes a missing branch*/
type ErrBadProtocol struct {
	Protocol git.Protocol
}

/*NewErrBadProtocol returns a new ErrBadProtocol*/
func NewErrBadProtocol(protocol git.Protocol) *ErrBadProtocol {
	return &ErrBadProtocol{protocol}
}

func (err *ErrBadProtocol) Error() string {
	return "bad protocol: " + string(err.Protocol)
}