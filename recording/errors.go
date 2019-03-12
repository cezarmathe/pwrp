package recording

import "errors"

var (
	/*ErrMissingBranch indicates that the metadata branch is missing.*/
	ErrMissingBranch = errors.New("metadata branch not found")

	/*ErrBadURL indicates that the metadata branch is missing.*/
	ErrBadURL = errors.New("bad url")

	/*ErrNoPermissions indicates that the metadata branch is missing.*/
	ErrNoPermissions = errors.New("no permissions for writing data")

	/*ErrBadProtocol indicates that the metadata branch is missing.*/
	ErrBadProtocol = errors.New("bad cloning protocol")
)
