package recording

import (
	"errors"
	"pppi/git"
)

var (
	/*ErrNoRepositories indicates that no repositories were specified in the config*/
	ErrNoRepositories = errors.New("empty repository url list")
)

/*ErrMissingBranch describes a missing branch*/
type ErrMissingBranch struct {
	Branch     string
	Repository string
}

/*NewErrMissingBranch returns a new ErrMissingBranch*/
func NewErrMissingBranch(branch, repository string) *ErrMissingBranch {
	return &ErrMissingBranch{branch, repository}
}

func (err *ErrMissingBranch) Error() string {
	return "metadata branch " + err.Branch + " not found in " + err.Repository
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
	return "the protocol \"" + string(err.Protocol) + "\" specified in the configuration is bad"
}

/*ErrCreateStorageDir describes the creation of a certain storage directory*/
type ErrCreateStorageDir struct {
	Dirname string
}

/*NewErrCreateStorageDir returns a new ErrCreateStorageDir*/
func NewErrCreateStorageDir(dirname string) *ErrCreateStorageDir {
	return &ErrCreateStorageDir{dirname}
}

func (err *ErrCreateStorageDir) Error() string {
	return "encountered an error when creating the storage directory at \"" + err.Dirname + "\""
}
