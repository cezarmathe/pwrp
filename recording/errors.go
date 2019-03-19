package recording

import (
	"errors"

	"github.com/cezarmathe/pwrp/gitops"
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
	Protocol gitops.Protocol
}

/*NewErrBadProtocol returns a new ErrBadProtocol*/
func NewErrBadProtocol(protocol gitops.Protocol) *ErrBadProtocol {
	return &ErrBadProtocol{protocol}
}

func (err *ErrBadProtocol) Error() string {
	return "the protocol \"" + string(err.Protocol) + "\" specified in the configuration is bad"
}
