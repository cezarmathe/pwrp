/*
	PWRP - Personal Work Recorder Processor
	Copyright (C) 2019  Cezar Mathe <cezarmathe@gmail.com> [https://cezarmathe.com]

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU Affero General Public License as published
	by the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU Affero General Public License for more details.

	You should have received a copy of the GNU Affero General Public License
	along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

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
